package docker

import (
	"context"
	"encoding/json"
	"errors"
	"os/exec"
	"strings"
)

type ContainerInfo struct {
	ID   string
	Name string

	Image    string
	Platform string

	Cmd  []string
	Env  map[string]string
	Args []string

	Dir string

	Ports map[int]int

	IPAddress string
}

func Info(ctx context.Context, container string) (*ContainerInfo, error) {
	tool, _, err := Tool(ctx)

	if err != nil {
		return nil, err
	}

	type networkHostPort struct {
		HostIP   string `json:"HostIp"`
		HostPort string `json:"HostPort"`
	}

	type inspectType struct {
		ID string `json:"Id"`
		// Created time.Time `json:"Created"`
		// Path string   `json:"Path"`
		// Args []string `json:"Args"`
		// State   struct {
		// 	Status     string    `json:"Status"`
		// 	Running    bool      `json:"Running"`
		// 	Paused     bool      `json:"Paused"`
		// 	Restarting bool      `json:"Restarting"`
		// 	OOMKilled  bool      `json:"OOMKilled"`
		// 	Dead       bool      `json:"Dead"`
		// 	Pid        int       `json:"Pid"`
		// 	ExitCode   int       `json:"ExitCode"`
		// 	Error      string    `json:"Error"`
		// 	StartedAt  time.Time `json:"StartedAt"`
		// 	FinishedAt time.Time `json:"FinishedAt"`
		// } `json:"State"`
		Image string `json:"Image"`
		// ResolvConfPath  string      `json:"ResolvConfPath"`
		// HostnamePath    string      `json:"HostnamePath"`
		// HostsPath       string      `json:"HostsPath"`
		// LogPath         string      `json:"LogPath"`
		Name string `json:"Name"`
		// RestartCount    int         `json:"RestartCount"`
		// Driver          string      `json:"Driver"`
		Platform string `json:"Platform"`
		// MountLabel      string      `json:"MountLabel"`
		// ProcessLabel    string      `json:"ProcessLabel"`
		// AppArmorProfile string      `json:"AppArmorProfile"`
		// ExecIDs         interface{} `json:"ExecIDs"`
		// HostConfig      struct {
		// 	Binds           interface{} `json:"Binds"`
		// 	ContainerIDFile string      `json:"ContainerIDFile"`
		// 	LogConfig       struct {
		// 		Type   string `json:"Type"`
		// 		Config struct {
		// 		} `json:"Config"`
		// 	} `json:"LogConfig"`
		// 	NetworkMode  string `json:"NetworkMode"`
		// 	PortBindings struct {
		// 		Nine090TCP []struct {
		// 			HostIP   string `json:"HostIp"`
		// 			HostPort string `json:"HostPort"`
		// 		} `json:"9090/tcp"`
		// 	} `json:"PortBindings"`
		// 	RestartPolicy struct {
		// 		Name              string `json:"Name"`
		// 		MaximumRetryCount int    `json:"MaximumRetryCount"`
		// 	} `json:"RestartPolicy"`
		// 	AutoRemove           bool          `json:"AutoRemove"`
		// 	VolumeDriver         string        `json:"VolumeDriver"`
		// 	VolumesFrom          interface{}   `json:"VolumesFrom"`
		// 	CapAdd               interface{}   `json:"CapAdd"`
		// 	CapDrop              interface{}   `json:"CapDrop"`
		// 	CgroupnsMode         string        `json:"CgroupnsMode"`
		// 	DNS                  []interface{} `json:"Dns"`
		// 	DNSOptions           []interface{} `json:"DnsOptions"`
		// 	DNSSearch            []interface{} `json:"DnsSearch"`
		// 	ExtraHosts           interface{}   `json:"ExtraHosts"`
		// 	GroupAdd             interface{}   `json:"GroupAdd"`
		// 	IpcMode              string        `json:"IpcMode"`
		// 	Cgroup               string        `json:"Cgroup"`
		// 	Links                interface{}   `json:"Links"`
		// 	OomScoreAdj          int           `json:"OomScoreAdj"`
		// 	PidMode              string        `json:"PidMode"`
		// 	Privileged           bool          `json:"Privileged"`
		// 	PublishAllPorts      bool          `json:"PublishAllPorts"`
		// 	ReadonlyRootfs       bool          `json:"ReadonlyRootfs"`
		// 	SecurityOpt          interface{}   `json:"SecurityOpt"`
		// 	UTSMode              string        `json:"UTSMode"`
		// 	UsernsMode           string        `json:"UsernsMode"`
		// 	ShmSize              int           `json:"ShmSize"`
		// 	Runtime              string        `json:"Runtime"`
		// 	ConsoleSize          []int         `json:"ConsoleSize"`
		// 	Isolation            string        `json:"Isolation"`
		// 	CPUShares            int           `json:"CpuShares"`
		// 	Memory               int           `json:"Memory"`
		// 	NanoCpus             int           `json:"NanoCpus"`
		// 	CgroupParent         string        `json:"CgroupParent"`
		// 	BlkioWeight          int           `json:"BlkioWeight"`
		// 	BlkioWeightDevice    []interface{} `json:"BlkioWeightDevice"`
		// 	BlkioDeviceReadBps   interface{}   `json:"BlkioDeviceReadBps"`
		// 	BlkioDeviceWriteBps  interface{}   `json:"BlkioDeviceWriteBps"`
		// 	BlkioDeviceReadIOps  interface{}   `json:"BlkioDeviceReadIOps"`
		// 	BlkioDeviceWriteIOps interface{}   `json:"BlkioDeviceWriteIOps"`
		// 	CPUPeriod            int           `json:"CpuPeriod"`
		// 	CPUQuota             int           `json:"CpuQuota"`
		// 	CPURealtimePeriod    int           `json:"CpuRealtimePeriod"`
		// 	CPURealtimeRuntime   int           `json:"CpuRealtimeRuntime"`
		// 	CpusetCpus           string        `json:"CpusetCpus"`
		// 	CpusetMems           string        `json:"CpusetMems"`
		// 	Devices              []interface{} `json:"Devices"`
		// 	DeviceCgroupRules    interface{}   `json:"DeviceCgroupRules"`
		// 	DeviceRequests       interface{}   `json:"DeviceRequests"`
		// 	KernelMemory         int           `json:"KernelMemory"`
		// 	KernelMemoryTCP      int           `json:"KernelMemoryTCP"`
		// 	MemoryReservation    int           `json:"MemoryReservation"`
		// 	MemorySwap           int           `json:"MemorySwap"`
		// 	MemorySwappiness     interface{}   `json:"MemorySwappiness"`
		// 	OomKillDisable       interface{}   `json:"OomKillDisable"`
		// 	PidsLimit            interface{}   `json:"PidsLimit"`
		// 	Ulimits              interface{}   `json:"Ulimits"`
		// 	CPUCount             int           `json:"CpuCount"`
		// 	CPUPercent           int           `json:"CpuPercent"`
		// 	IOMaximumIOps        int           `json:"IOMaximumIOps"`
		// 	IOMaximumBandwidth   int           `json:"IOMaximumBandwidth"`
		// 	MaskedPaths          []string      `json:"MaskedPaths"`
		// 	ReadonlyPaths        []string      `json:"ReadonlyPaths"`
		// } `json:"HostConfig"`
		// GraphDriver struct {
		// 	Data struct {
		// 		LowerDir  string `json:"LowerDir"`
		// 		MergedDir string `json:"MergedDir"`
		// 		UpperDir  string `json:"UpperDir"`
		// 		WorkDir   string `json:"WorkDir"`
		// 	} `json:"Data"`
		// 	Name string `json:"Name"`
		// } `json:"GraphDriver"`
		// Mounts []interface{} `json:"Mounts"`
		Config struct {
			// 	Hostname     string `json:"Hostname"`
			// 	Domainname   string `json:"Domainname"`
			// 	User         string `json:"User"`
			// 	AttachStdin  bool   `json:"AttachStdin"`
			// 	AttachStdout bool   `json:"AttachStdout"`
			// 	AttachStderr bool   `json:"AttachStderr"`
			// 	ExposedPorts struct {
			// 		Nine090TCP struct {
			// 		} `json:"9090/tcp"`
			// 	} `json:"ExposedPorts"`
			// 	Tty        bool        `json:"Tty"`
			// 	OpenStdin  bool        `json:"OpenStdin"`
			// 	StdinOnce  bool        `json:"StdinOnce"`
			Env []string `json:"Env"`
			Cmd []string `json:"Cmd"`
			// 	Image      string      `json:"Image"`
			// 	Volumes    interface{} `json:"Volumes"`
			WorkingDir string   `json:"WorkingDir"`
			Entrypoint []string `json:"Entrypoint"`
			// 	OnBuild    interface{} `json:"OnBuild"`
			// 	Labels     struct {
			// 	} `json:"Labels"`
		} `json:"Config"`
		NetworkSettings struct {
			// 	Bridge                 string `json:"Bridge"`
			// 	SandboxID              string `json:"SandboxID"`
			// 	HairpinMode            bool   `json:"HairpinMode"`
			// 	LinkLocalIPv6Address   string `json:"LinkLocalIPv6Address"`
			// 	LinkLocalIPv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen"`
			Ports map[string][]struct {
				HostIP   string `json:"HostIp"`
				HostPort string `json:"HostPort"`
			} `json:"Ports"`
			// 	SandboxKey             string      `json:"SandboxKey"`
			// 	SecondaryIPAddresses   interface{} `json:"SecondaryIPAddresses"`
			// 	SecondaryIPv6Addresses interface{} `json:"SecondaryIPv6Addresses"`
			// 	EndpointID             string      `json:"EndpointID"`
			// 	Gateway                string      `json:"Gateway"`
			// 	GlobalIPv6Address      string      `json:"GlobalIPv6Address"`
			// 	GlobalIPv6PrefixLen    int         `json:"GlobalIPv6PrefixLen"`
			IPAddress string `json:"IPAddress"`
			// 	IPPrefixLen            int         `json:"IPPrefixLen"`
			// 	IPv6Gateway            string      `json:"IPv6Gateway"`
			// 	MacAddress             string      `json:"MacAddress"`
			// 	Networks               struct {
			// 		Bridge struct {
			// 			IPAMConfig          interface{} `json:"IPAMConfig"`
			// 			Links               interface{} `json:"Links"`
			// 			Aliases             interface{} `json:"Aliases"`
			// 			NetworkID           string      `json:"NetworkID"`
			// 			EndpointID          string      `json:"EndpointID"`
			// 			Gateway             string      `json:"Gateway"`
			// 			IPAddress           string      `json:"IPAddress"`
			// 			IPPrefixLen         int         `json:"IPPrefixLen"`
			// 			IPv6Gateway         string      `json:"IPv6Gateway"`
			// 			GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
			// 			GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
			// 			MacAddress          string      `json:"MacAddress"`
			// 			DriverOpts          interface{} `json:"DriverOpts"`
			// 		} `json:"bridge"`
			// 	} `json:"Networks"`
		} `json:"NetworkSettings"`
	}

	run := exec.CommandContext(ctx, tool, "inspect", container)

	data, err := run.Output()

	if err != nil {
		return nil, err
	}

	var inspectResult []inspectType

	if err := json.Unmarshal(data, &inspectResult); err != nil {
		return nil, err
	}

	env := map[string]string{}

	if len(inspectResult) != 1 {
		return nil, errors.New("unexpected inspect data")
	}

	inspectData := inspectResult[0]

	for _, e := range inspectData.Config.Env {
		pair := strings.SplitN(e, "=", 2)

		if len(pair) < 2 {
			continue
		}

		k := pair[0]
		v := pair[1]

		env[k] = v
	}

	// for c, h := range inspectData.NetworkSettings.Ports {
	// }

	return &ContainerInfo{
		ID:   inspectData.ID,
		Name: strings.TrimLeft(inspectData.Name, "/"),

		Image: inspectData.Image,

		Cmd:  inspectData.Config.Entrypoint,
		Env:  env,
		Args: inspectData.Config.Cmd,

		Dir: inspectData.Config.WorkingDir,

		IPAddress: inspectData.NetworkSettings.IPAddress,
	}, nil
}
