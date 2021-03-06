package demo

import (
	"github.com/dops-cli/dops/cli"

	"github.com/dops-cli/dops/categories"
)

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "demo",
			Usage:       "Demo module of dops",
			Description: `NOTICE: This module does nothing, except showing all possible flags for an interactive demo.`,
			Category:    categories.Dops,
			Action: func(c *cli.Context) error {
				return nil
			},
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "Boolean"},
				&cli.DurationFlag{Name: "Duration"},
				&cli.Float64Flag{Name: "Float64"},
				&cli.Float64SliceFlag{Name: "Float64List"},
				&cli.IntFlag{Name: "Int"},
				&cli.IntSliceFlag{Name: "IntList"},
				&cli.PathFlag{Name: "Path"},
				&cli.StringFlag{Name: "String"},
				&cli.StringSliceFlag{Name: "StringList"},
				&cli.TimestampFlag{Name: "Timestamp"},
				&cli.OptionFlag{Name: "Options"},
			},
		},
	}
}
