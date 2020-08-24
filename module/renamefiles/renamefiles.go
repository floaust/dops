package renamefiles

import (
	"github.com/dops-cli/dops/cli"

	"github.com/dops-cli/dops/categories"
)

// Module returns the created module
type Module struct{}

// GetCommands returns the commands of the module
func (Module) GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "rename-files",
			Aliases: []string{"rf"},
			Usage:   "Renames all selected files to a specific pattern",
			Description: `This module can be used to rename multiple files according to a specified pattern.
The pattern could be a timestamp, or the hashcode of the file, among others.`,
			Category: categories.IO,
			Action: func(context *cli.Context) error {

				return nil
			},
			Flags: []cli.Flag{},
		},
	}
}