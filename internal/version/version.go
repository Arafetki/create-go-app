package version

import (
	"runtime"

	"github.com/arafetki/create-go-app/internal/git"
)

type Info struct {
	CLIVersion     string
	RuntimeVersion string
	CommitHash     string
}

func New(cliVersion string) *Info {
	return &Info{
		CLIVersion:     cliVersion,
		RuntimeVersion: runtime.Version(),
		CommitHash:     git.Revision(),
	}
}
