package main

import (
	"fmt"

	"github.com/urfave/cli"
	pb "k8s.io/kubernetes/pkg/kubelet/api/v1alpha1/runtime"
)

// TODO: --security-opt or equiv for apparmor, SELinux, seccomp
// TODO: resource limit flags (cpu-period & etc)

type createConfig struct {
	image      string
	command    string
	args       []string
	pod        string
	privileged bool
	rm         bool
	hostNet    bool
	hostPID    bool
	hostIPC    bool
	name       string
	labels     map[string]string
	podConfig  *newPodConfig
	workDir    string
	env        map[string]string
	detach     bool
	stdin      bool
	tty        bool
	devices    []*pb.Device
	mounts     []*pb.Mount
	capAdd     []string
	capDrop    []string
}

type newPodConfig struct {
	dnsServers       []string
	dnsSearch        []string
	dnsOpt           []string
	ports            []*pb.PortMapping
	hostname         string
	cgroupParent     string
	sysctl           string
	user             int64
	additionalGroups []int64
	readOnlyRootfs   bool
}

var createFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "pod",
		Usage: "Run container in an existing pod",
	},
	cli.BoolFlag{
		Name:  "privileged",
		Usage: "Run a privileged container",
	},
	cli.BoolFlag{
		Name:  "rm",
		Usage: "Remove container (and pod if created) after exit",
	},
	cli.StringFlag{
		Name:  "network",
		Usage: "Use `host` network namespace",
	},
	cli.StringFlag{
		Name:  "pid",
		Usage: "Use `host` PID namespace",
	},
	cli.StringFlag{
		Name:  "ipc",
		Usage: "Use `host` IPC namespace",
	},
	cli.StringFlag{
		Name:  "name",
		Usage: "Assign a name to the container",
	},
	cli.StringSliceFlag{
		Name:  "label",
		Usage: "Set label metadata on container",
	},
	cli.StringFlag{
		Name:  "workdir, w",
		Usage: "Set working `directory` of container",
		Value: "/",
	},
	cli.StringSliceFlag{
		Name:  "env, e",
		Usage: "Set environment variables in container",
	},
	cli.BoolFlag{
		Name:  "detach, d",
		Usage: "Start container detached",
	},
	cli.BoolFlag{
		Name:  "interactive, i",
		Usage: "Keep STDIN open even if deatched",
	},
	cli.BoolFlag{
		Name:  "tty, t",
		Usage: "Allocate a TTY for container",
	},
	cli.StringSliceFlag{
		Name:  "device",
		Usage: "Mount devices into the container",
	},
	cli.StringSliceFlag{
		Name:  "volume, v",
		Usage: "Mount volumes into the container",
	},
	cli.StringSliceFlag{
		Name:  "cap-add",
		Usage: "Add capabilities to the container",
	},
	cli.StringSliceFlag{
		Name:  "cap-drop",
		Usage: "Drop capabilities from the container",
	},
}

// Flags that can only be used when creating a new pod
var podOnlyFlags = []cli.Flag{
	cli.StringSliceFlag{
		Name:  "dns",
		Usage: "Set custom DNS servers. Conflicts with --pod",
	},
	cli.StringSliceFlag{
		Name:  "dns-search",
		Usage: "Set custom DNS search domains. Conflicts with --pod",
	},
	cli.StringSliceFlag{
		Name:  "dns-opt",
		Usage: "Set custom DNS options. Conflicts with --pod",
	},
	cli.StringSliceFlag{
		Name:  "expose",
		Usage: "Expose a port. Conflicts with --pod",
	},
	cli.StringFlag{
		Name:  "hostname, h",
		Usage: "Set hostname. Conflicts with --pod",
		Value: defaultHostname,
	},
	cli.StringFlag{
		Name:  "cgroup-parent",
		Usage: "Set CGroup parent. Conflicts with --pod",
		Value: defaultCgroupParent,
	},
	cli.StringFlag{
		Name:  "sysctl",
		Usage: "Set namespaces SYSCTL. Conflicts with --pod",
	},
	cli.StringFlag{
		Name:  "user, u",
		Usage: "Specify user to run as. Conflicts with --pod",
	},
	cli.StringFlag{
		Name:  "group-add",
		Usage: "Specify additional groups to run as. Conflicts with --pod",
	},
	cli.BoolFlag{
		Name:  "read-only",
		Usage: "Make root filesystem read-only. Conflicts with --pod",
	},
}

var kpodRunFlags = append(append(append([]cli.Flag{}, podOnlyFlags...),
	createFlags...), commonFlags...)

var runCommand = cli.Command{
	Name:   "run",
	Usage:  "launch a container",
	Flags:  kpodRunFlags,
	Action: kpodRun,
}

// Parse CLI arguments of 'kpod run'
// Returns image to run, command, arguments for command, and error
func parseRunArguments(args []string) (string, string, []string, error) {
	if len(args) == 0 {
		return "", "", []string{}, fmt.Errorf("must provide as least one argument: image to run")
	}

	image := args[0]
	command := ""
	finalArgs := []string{}

	if len(args) > 1 {
		command = args[1]

		if len(args) > 2 {
			finalArgs = append(finalArgs, args[2:]...)
		}
	}

	return image, command, finalArgs, nil
}

func parseRunFlags(c *cli.Context) (*createConfig, error) {
	return fmt.Errorf("TODO")
}

// Implementation of 'kpod run' command
func kpodRun(c *cli.Context) error {
	return fmt.Errorf("not yet implemented")
}
