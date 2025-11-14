package main

import (
	"context"
	"os"
	"rhel-drivers/internal/api"
	"rhel-drivers/internal/cli"
	"rhel-drivers/internal/dnf"
	"rhel-drivers/internal/provider/amd"
	"rhel-drivers/internal/provider/nvidia"
	"rhel-drivers/internal/rhsm"
	"rhel-drivers/internal/sysinfo"
)

// set at build time via -ldflags, eg: go build -ldflags="-X main.version=1.0.0" ./cmd/rhel-drivers
var version = "dev"

func main() {
	ctx := context.Background()
	sysinfo := sysinfo.DetectSysInfo()

	pm := dnf.New()
	repoVerifier := rhsm.NewVerifier(sysinfo)
	providers := []api.Provider{nvidia.NewProvider(pm), amd.NewProvider(pm)}
	deps := api.CoreDeps{
		PM:           pm,
		RepoVerifier: repoVerifier,
		Providers:    providers,
	}

	root := cli.NewRootCmd(deps, version)

	if err := root.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
