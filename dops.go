package main

import (
	"github.com/dops-cli/dops/module"
	"github.com/dops-cli/dops/module/modules"
	"github.com/dops-cli/dops/say"
	"github.com/dops-cli/dops/say/color"
	"github.com/dops-cli/dops/template"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"sort"
)

var (
	CliFlags    []cli.Flag
	CliCommands []*cli.Command
)

func init() {
	cli.AppHelpTemplate = template.AppHelp
	cli.CommandHelpTemplate = template.CommandHelp
	cli.SubcommandHelpTemplate = template.SubcommandHelp
	cli.MarkdownDocTemplate = template.MarkdownDoc
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		cli.HelpPrinterCustom(color.Output, templ, data, nil)
	}
	cli.VersionPrinter = func(c *cli.Context) {
		say.Info("dops is currently on version " + c.App.Version + "!")
	}
}

func main() {

	for _, f := range module.ActiveGlobalFlags {
		CliFlags = append(CliFlags, f.GetFlags()...)
	}

	for _, m := range module.ActiveModules {
		CliCommands = append(CliCommands, m.GetCommands()...)
	}

	CliCommands = append(CliCommands, modules.Module{}.GetCommands()...)

	app := &cli.App{
		Name:     "dops",
		HelpName: "dops",
		Usage:    "CLI DevOps Toolkit",
		Version:  "v1.8.1", // <---VERSION---> This comment is used for CI, do NOT modify it!
		Flags:    CliFlags,
		Commands: CliCommands,
		Authors: []*cli.Author{
			{
				Name:  "Marvin Wendt",
				Email: "dops@marvinjwendt.com",
			},
		},
		Copyright:              "(c) 2020 Marvin Wendt",
		Writer:                 color.Output,
		UseShortOptionHandling: true,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
