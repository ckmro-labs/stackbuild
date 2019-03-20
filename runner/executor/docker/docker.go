package docker

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/network"
	"docker.io/go-docker/api/types/volume"
	"github.com/laidingqing/stackbuild/runner/executor"
)

type dockerExecutor struct {
	client docker.APIClient
}

// NewEnv 创建一个新的docker环境引擎.
func NewEnv() (executor.Executor, error) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return New(cli), nil
}

// New use docker api create a executor..
func New(client docker.APIClient) executor.Executor {
	return &dockerExecutor{
		client: client,
	}
}

func (e *dockerExecutor) Setup(ctx context.Context, spec *executor.Spec) error {
	if spec.Docker != nil {
		for _, vol := range spec.Docker.Volumes {
			if vol.EmptyDir == nil {
				continue
			}

			_, err := e.client.VolumeCreate(ctx, volume.VolumesCreateBody{
				Name:   vol.Metadata.UID,
				Driver: "local",
				Labels: spec.Metadata.Labels,
			})
			if err != nil {
				return err
			}
		}
	}
	driver := "bridge"
	if spec.Platform.OS == "windows" {
		driver = "nat"
	}
	_, err := e.client.NetworkCreate(ctx, spec.Metadata.UID, types.NetworkCreate{
		Driver: driver,
		Labels: spec.Metadata.Labels,
	})

	return err
}

func (e *dockerExecutor) Create(ctx context.Context, spec *executor.Spec, step *executor.Step) error {
	if step.Docker == nil {
		return errors.New("executor: missing docker configuration")
	}
	_, domain, latest, err := parseImage(step.Docker.Image)
	if err != nil {
		return err
	}
	pullopts := types.ImagePullOptions{}
	auths, ok := executor.LookupAuth(spec, domain)
	if ok {
		pullopts.RegistryAuth = Encode(auths.Username, auths.Password)
	}
	if step.Docker.PullPolicy == executor.PullAlways ||
		(step.Docker.PullPolicy == executor.PullDefault && latest) {
		rc, perr := e.client.ImagePull(ctx, step.Docker.Image, pullopts)
		if perr == nil {
			io.Copy(ioutil.Discard, rc)
			rc.Close()
		}
		if perr != nil {
			return perr
		}
	}
	_, err = e.client.ContainerCreate(ctx,
		toConfig(spec, step),
		toHostConfig(spec, step),
		toNetConfig(spec, step),
		step.Metadata.UID,
	)
	if docker.IsErrImageNotFound(err) && step.Docker.PullPolicy != executor.PullNever {
		rc, perr := e.client.ImagePull(ctx, step.Docker.Image, pullopts)
		if perr != nil {
			return perr
		}
		io.Copy(ioutil.Discard, rc)
		rc.Close()
		_, err = e.client.ContainerCreate(ctx,
			toConfig(spec, step),
			toHostConfig(spec, step),
			toNetConfig(spec, step),
			step.Metadata.UID,
		)
	}
	if err != nil {
		return err
	}

	copyOpts := types.CopyToContainerOptions{}
	copyOpts.AllowOverwriteDirWithFile = false
	for _, mount := range step.Files {
		file, ok := executor.LookupFile(spec, mount.Name)
		if !ok {
			continue
		}
		tar := createTarfile(file, mount)
		err := e.client.CopyToContainer(ctx, step.Metadata.UID, "/", bytes.NewReader(tar), copyOpts)
		if err != nil {
			return err
		}
	}

	for _, net := range step.Docker.Networks {
		err = e.client.NetworkConnect(ctx, net, step.Metadata.UID, &network.EndpointSettings{
			Aliases: []string{net},
		})
		if err != nil {
			return nil
		}
	}

	return nil
}

func (e *dockerExecutor) Start(ctx context.Context, spec *executor.Spec, step *executor.Step) error {
	return e.client.ContainerStart(ctx, step.Metadata.UID, types.ContainerStartOptions{})
}

func (e *dockerExecutor) Wait(ctx context.Context, spec *executor.Spec, step *executor.Step) (*executor.State, error) {
	wait, errc := e.client.ContainerWait(ctx, step.Metadata.UID, "")
	select {
	case <-wait:
	case <-errc:
	}

	info, err := e.client.ContainerInspect(ctx, step.Metadata.UID)
	if err != nil {
		return nil, err
	}
	if info.State.Running {
		//TODO
	}

	return &executor.State{
		Exited:    true,
		ExitCode:  info.State.ExitCode,
		OOMKilled: info.State.OOMKilled,
	}, nil
}

func (e *dockerExecutor) Tail(ctx context.Context, spec *executor.Spec, step *executor.Step) (io.ReadCloser, error) {
	opts := types.ContainerLogsOptions{
		Follow:     true,
		ShowStdout: true,
		ShowStderr: true,
		Details:    false,
		Timestamps: false,
	}

	logs, err := e.client.ContainerLogs(ctx, step.Metadata.UID, opts)
	if err != nil {
		return nil, err
	}
	rc, wc := io.Pipe()

	go func() {
		StdCopy(wc, wc, logs)
		logs.Close()
		wc.Close()
		rc.Close()
	}()
	return rc, nil
}

func (e *dockerExecutor) Destroy(ctx context.Context, spec *executor.Spec) error {
	removeOpts := types.ContainerRemoveOptions{
		Force:         true,
		RemoveLinks:   false,
		RemoveVolumes: true,
	}

	// stop all containers
	for _, step := range spec.Steps {
		e.client.ContainerKill(ctx, step.Metadata.UID, "9")
	}

	// cleanup all containers
	for _, step := range spec.Steps {
		e.client.ContainerRemove(ctx, step.Metadata.UID, removeOpts)
	}

	// cleanup all volumes
	if spec.Docker != nil {
		for _, vol := range spec.Docker.Volumes {
			if vol.EmptyDir == nil {
				continue
			}
			if vol.EmptyDir.Medium == "memory" {
				continue
			}
			e.client.VolumeRemove(ctx, vol.Metadata.UID, true)
		}
	}

	// cleanup the network
	e.client.NetworkRemove(ctx, spec.Metadata.UID)
	return nil
}
