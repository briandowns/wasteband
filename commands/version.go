package commands

import "github.com/mitchellh/cli"

type Version struct {
	version string
}

func NewVersion(version string) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Version{
			version: version,
		}, nil
	}
}

func (v *Version) Run(args []string) int { return 0 }
func (v *Version) Help() string          { return "" }
func (v *Version) Synopsis() string      { return "" }
