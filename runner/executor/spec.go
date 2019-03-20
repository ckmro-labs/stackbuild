package executor

type (
	//Metadata defines executor's meta
	Metadata struct {
		UID       string            `json:"uid,omitempty"`
		Namespace string            `json:"namespace,omitempty"`
		Name      string            `json:"name,omitempty"`
		Labels    map[string]string `json:"labels,omitempty"`
	}

	//Spec ..
	Spec struct {
		Metadata Metadata      `json:"metadata,omitempty"`
		Platform Platform      `json:"platform,omitempty"`
		Secrets  []*Secret     `json:"secrets,omitempty"`
		Steps    []*Step       `json:"steps,omitempty"`
		Files    []*File       `json:"files,omitempty"`
		Docker   *DockerConfig `json:"docker,omitempty"`
	}

	// Step defines a pipeline step.
	Step struct {
		Metadata     Metadata          `json:"metadata,omitempty"`
		Detach       bool              `json:"detach,omitempty"`
		DependsOn    []string          `json:"depends_on,omitempty"`
		Devices      []*VolumeDevice   `json:"devices,omitempty"`
		Envs         map[string]string `json:"environment,omitempty"`
		Files        []*FileMount      `json:"files,omitempty"`
		IgnoreErr    bool              `json:"ignore_err,omitempty"`
		IgnoreStdout bool              `json:"ignore_stderr,omitempty"`
		IgnoreStderr bool              `json:"ignore_stdout,omitempty"`
		Resources    *Resources        `json:"resources,omitempty"`
		RunPolicy    RunPolicy         `json:"run_policy,omitempty"`
		Secrets      []*SecretVar      `json:"secrets,omitempty"`
		Volumes      []*VolumeMount    `json:"volumes,omitempty"`
		WorkingDir   string            `json:"working_dir,omitempty"`
		Docker       *DockerStep       `json:"docker,omitempty"`
	}

	// SecretVar represents an environment variable
	SecretVar struct {
		Name string `json:"name,omitempty"`
		Env  string `json:"env,omitempty"`
	}

	// Resources describes the compute resource
	Resources struct {
		// Limits maximum
		Limits *ResourceObject `json:"limits,omitempty"`
		// Requests minimum
		Requests *ResourceObject `json:"requests,omitempty"`
	}

	// ResourceObject describes compute resource
	ResourceObject struct {
		CPU    int64 `json:"cpu,omitempty"`
		Memory int64 `json:"memory,omitempty"`
	}

	// FileMount defines how a file resource
	FileMount struct {
		Name string `json:"name,omitempty"`
		Path string `json:"path,omitempty"`
		Mode int64  `json:"mode,omitempty"`
	}
	//File ..
	File struct {
		Metadata Metadata `json:"metadata,omitempty"`
		Data     []byte   `json:"data,omitempty"`
	}

	// DockerStep configures a docker step.
	DockerStep struct {
		Args       []string   `json:"args,omitempty"`
		Command    []string   `json:"command,omitempty"`
		DNS        []string   `json:"dns,omitempty"`
		DNSSearch  []string   `json:"dns_search,omitempty"`
		ExtraHosts []string   `json:"extra_hosts,omitempty"`
		Image      string     `json:"image,omitempty"`
		Networks   []string   `json:"networks,omitempty"`
		Ports      []*Port    `json:"ports,omitempty"`
		Privileged bool       `json:"privileged,omitempty"`
		PullPolicy PullPolicy `json:"pull_policy,omitempty"`
	}

	// Platform defines the target platform.
	Platform struct {
		OS      string `json:"os,omitempty"`
		Arch    string `json:"arch,omitempty"`
		Variant string `json:"variant,omitempty"`
		Version string `json:"version,omitempty"`
	}

	// Secret represents a secret variable.
	Secret struct {
		Metadata Metadata `json:"metadata,omitempty"`
		Data     string   `json:"data,omitempty"`
	}

	// DockerConfig configures a Docker-based pipeline.
	DockerConfig struct {
		Auths   []*DockerAuth `json:"auths,omitempty"`
		Volumes []*Volume     `json:"volumes,omitempty"`
	}

	// DockerAuth defines dockerhub authentication credentials.
	DockerAuth struct {
		Address  string `json:"address,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}
	// Port represents a network port in a single container.
	Port struct {
		Port     int    `json:"port,omitempty"`
		Host     int    `json:"host,omitempty"`
		Protocol string `json:"protocol,omitempty"`
	}

	// Volume that can be mounted by containers.
	Volume struct {
		Metadata Metadata        `json:"metadata,omitempty"`
		EmptyDir *VolumeEmptyDir `json:"temp,omitempty"`
		HostPath *VolumeHostPath `json:"host,omitempty"`
	}

	// VolumeEmptyDir mounts a temporary directory
	VolumeEmptyDir struct {
		Medium    string `json:"medium,omitempty"`
		SizeLimit int64  `json:"size_limit,omitempty"`
	}

	// VolumeHostPath mounts a file or directory
	VolumeHostPath struct {
		Path string `json:"path,omitempty"`
	}

	// State represents the container state.
	State struct {
		ExitCode  int  // Container exit code
		Exited    bool // Container exited
		OOMKilled bool // Container is oom killed
	}

	// VolumeDevice describes a mapping of a raw block
	// device within a container.
	VolumeDevice struct {
		Name       string `json:"name,omitempty"`
		DevicePath string `json:"path,omitempty"`
	}

	// VolumeMount describes a mounting of a Volume
	// within a container.
	VolumeMount struct {
		Name string `json:"name,omitempty"`
		Path string `json:"path,omitempty"`
	}
)
