package executor

import (
	"net/url"
	"strings"
)

// LookupVolume lookup the named volume.
func LookupVolume(spec *Spec, name string) (*Volume, bool) {
	if spec.Docker == nil {
		return nil, false
	}
	for _, vol := range spec.Docker.Volumes {
		if vol.Metadata.Name == name {
			return vol, true
		}
	}
	return nil, false
}

// LookupSecret is a helper
func LookupSecret(spec *Spec, secret *SecretVar) (*Secret, bool) {
	for _, sec := range spec.Secrets {
		if sec.Metadata.Name == secret.Name {
			return sec, true
		}
	}
	return nil, false
}

// LookupFile is a helper function
func LookupFile(spec *Spec, name string) (*File, bool) {
	for _, file := range spec.Files {
		if file.Metadata.Name == name {
			return file, true
		}
	}
	return nil, false
}

// LookupAuth is a helper function
func LookupAuth(spec *Spec, domain string) (*DockerAuth, bool) {
	if spec.Docker == nil {
		return nil, false
	}
	for _, auth := range spec.Docker.Auths {
		host := auth.Address
		if strings.HasPrefix(host, "http://") ||
			strings.HasPrefix(host, "https://") {
			uri, err := url.Parse(auth.Address)
			if err != nil {
				continue
			}
			host = uri.Host
		}
		if host == "index.docker.io" {
			host = "docker.io"
		}

		if host == domain {
			return auth, true
		}
	}
	return nil, false
}
