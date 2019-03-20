package docker

import (
	"archive/tar"
	"bytes"
	"strings"
	"time"

	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/mount"
	"docker.io/go-docker/api/types/network"
	"github.com/laidingqing/stackbuild/runner/executor"
)

// returns a container configuration.
func toConfig(spec *executor.Spec, step *executor.Step) *container.Config {
	config := &container.Config{
		Image:        step.Docker.Image,
		Labels:       step.Metadata.Labels,
		WorkingDir:   step.WorkingDir,
		AttachStdin:  false,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
		OpenStdin:    false,
		StdinOnce:    false,
		ArgsEscaped:  false,
	}

	if len(step.Envs) != 0 {
		config.Env = toEnv(step.Envs)
	}
	for _, sec := range step.Secrets {
		secret, ok := executor.LookupSecret(spec, sec)
		if ok {
			config.Env = append(config.Env, sec.Env+"="+secret.Data)
		}
	}
	if len(step.Docker.Args) != 0 {
		config.Cmd = step.Docker.Args
	}
	if len(step.Docker.Command) != 0 {
		config.Entrypoint = step.Docker.Command
	}

	// NOTE it appears this is no longer required,
	// however this could cause incompatibility with
	// certain docker versions.
	//
	//   if len(step.Volumes) != 0 {
	// 	    config.Volumes = toVolumeSet(spec, step)
	//   }
	return config
}

// returns a container host configuration.
func toHostConfig(spec *executor.Spec, step *executor.Step) *container.HostConfig {
	config := &container.HostConfig{
		LogConfig: container.LogConfig{
			Type: "json-file",
		},
		Privileged: step.Docker.Privileged,
		// TODO(bradrydzewski) set ShmSize
	}
	// windows does not support privileged so we hard-code
	// this value to false.
	if spec.Platform.OS == "windows" {
		config.Privileged = false
	}
	if len(step.Docker.DNS) > 0 {
		config.DNS = step.Docker.DNS
	}
	if len(step.Docker.DNSSearch) > 0 {
		config.DNSSearch = step.Docker.DNSSearch
	}
	if len(step.Docker.ExtraHosts) > 0 {
		config.ExtraHosts = step.Docker.ExtraHosts
	}
	if step.Resources != nil {
		config.Resources = container.Resources{}
		if limits := step.Resources.Limits; limits != nil {
			config.Resources.Memory = limits.Memory
		}
	}

	if len(step.Volumes) != 0 {
		config.Devices = toDeviceSlice(spec, step)
		config.Binds = toVolumeSlice(spec, step)
		config.Mounts = toVolumeMounts(spec, step)
	}
	return config
}

// helper function returns the container network configuration.
func toNetConfig(spec *executor.Spec, proc *executor.Step) *network.NetworkingConfig {
	endpoints := map[string]*network.EndpointSettings{}
	endpoints[spec.Metadata.UID] = &network.EndpointSettings{
		NetworkID: spec.Metadata.UID,
		Aliases:   []string{proc.Metadata.Name},
	}
	return &network.NetworkingConfig{
		EndpointsConfig: endpoints,
	}
}

// helper function that converts a slice of device paths to a slice of
// container.DeviceMapping.
func toDeviceSlice(spec *executor.Spec, step *executor.Step) []container.DeviceMapping {
	var to []container.DeviceMapping
	for _, mount := range step.Devices {
		device, ok := executor.LookupVolume(spec, mount.Name)
		if !ok {
			continue
		}
		if isDevice(device) == false {
			continue
		}
		to = append(to, container.DeviceMapping{
			PathOnHost:        device.HostPath.Path,
			PathInContainer:   mount.DevicePath,
			CgroupPermissions: "rwm",
		})
	}
	if len(to) == 0 {
		return nil
	}
	return to
}

// helper function returns a slice of volume mounts.
func toVolumeSlice(spec *executor.Spec, step *executor.Step) []string {
	var to []string
	for _, mount := range step.Volumes {
		volume, ok := executor.LookupVolume(spec, mount.Name)
		if !ok {
			continue
		}
		if isDevice(volume) {
			continue
		}
		if isDataVolume(volume) == false {
			continue
		}
		path := volume.Metadata.UID + ":" + mount.Path
		to = append(to, path)
	}
	return to
}

func toVolumeMounts(spec *executor.Spec, step *executor.Step) []mount.Mount {
	var mounts []mount.Mount
	for _, target := range step.Volumes {
		source, ok := executor.LookupVolume(spec, target.Name)
		if !ok {
			continue
		}
		if isDataVolume(source) {
			continue
		}
		mounts = append(mounts, toMount(source, target))
	}
	if len(mounts) == 0 {
		return nil
	}
	return mounts
}

func toMount(source *executor.Volume, target *executor.VolumeMount) mount.Mount {
	to := mount.Mount{
		Target: target.Path,
		Type:   toVolumeType(source),
	}
	if isBindMount(source) || isNamedPipe(source) {
		to.Source = source.HostPath.Path
	}
	if isTempfs(source) {
		to.TmpfsOptions = &mount.TmpfsOptions{
			SizeBytes: source.EmptyDir.SizeLimit,
			Mode:      0700,
		}
	}
	return to
}

// return given volume enum.
func toVolumeType(from *executor.Volume) mount.Type {
	switch {
	case isDataVolume(from):
		return mount.TypeVolume
	case isTempfs(from):
		return mount.TypeTmpfs
	case isNamedPipe(from):
		return mount.TypeNamedPipe
	default:
		return mount.TypeBind
	}
}

// converts a key value map of environment variables to a string slice
func toEnv(env map[string]string) []string {
	var envs []string
	for k, v := range env {
		envs = append(envs, k+"="+v)
	}
	return envs
}

// returns true if the volume is a bind mount.
func isBindMount(volume *executor.Volume) bool {
	return volume.HostPath != nil
}

// returns true if the volume is in-memory.
func isTempfs(volume *executor.Volume) bool {
	return volume.EmptyDir != nil && volume.EmptyDir.Medium == "memory"
}

// returns true if the volume is a data-volume.
func isDataVolume(volume *executor.Volume) bool {
	return volume.EmptyDir != nil && volume.EmptyDir.Medium != "memory"
}

// returns true if the volume is a device
func isDevice(volume *executor.Volume) bool {
	return volume.HostPath != nil && strings.HasPrefix(volume.HostPath.Path, "/dev/")
}

// returns true if the volume is a named pipe.
func isNamedPipe(volume *executor.Volume) bool {
	return volume.HostPath != nil &&
		strings.HasPrefix(volume.HostPath.Path, `\\.\pipe\`)
}

func createTarfile(file *executor.File, mount *executor.FileMount) []byte {
	w := new(bytes.Buffer)
	t := tar.NewWriter(w)
	h := &tar.Header{
		Name:    mount.Path,
		Mode:    mount.Mode,
		Size:    int64(len(file.Data)),
		ModTime: time.Now(),
	}
	t.WriteHeader(h)
	t.Write(file.Data)
	return w.Bytes()
}
