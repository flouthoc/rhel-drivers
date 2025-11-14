package api

import (
	"context"
	"fmt"
)

type PackageInfo struct {
	Name       string
	Epoch      string
	Version    string
	Release    string
	Arch       string
	SourceName string
	Repo       string
}

func (p PackageInfo) NEVRA() string {
	if p.Epoch != "" && p.Epoch != "0" {
		return fmt.Sprintf("%s-%s:%s-%s.%s", p.Name, p.Epoch, p.Version, p.Release, p.Arch)
	}
	return fmt.Sprintf("%s-%s-%s.%s", p.Name, p.Version, p.Release, p.Arch)
}

type PackageManager interface {
	ListAvailablePackages(ctx context.Context) ([]PackageInfo, error)
	ListInstalledPackages(ctx context.Context) ([]PackageInfo, error)
	Install(ctx context.Context, packages []string, opts InstallOptions) error
	Remove(ctx context.Context, packages []string, opts RemoveOptions) error
}
