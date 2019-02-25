package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "test"
	app.Usage = "cli test program"
	//默认执行action
	app.Action = func(c *cli.Context) error {
		println("app test node")
		return nil
	}
	//可以通过命令执行下面字命令函数，此时跳过主action
	app.Commands = []cli.Command{{
		Name:               "sub",
		ShortName:          "s",
		Aliases:            nil,
		Usage:              "./node sub",
		UsageText:          "",
		Description:        "",
		ArgsUsage:          "",
		Category:           "",
		BashComplete:       nil,
		Before:             nil,
		After:              nil,
		Action: func(c *cli.Context) {
			println("sub action")
		},
		OnUsageError:       nil,
		Subcommands:        nil,
		Flags:              nil,
		SkipFlagParsing:    false,
		SkipArgReorder:     false,
		HideHelp:           false,
		Hidden:             false,
		HelpName:           "",
		CustomHelpTemplate: "",
	},}

	_ = app.Run(os.Args)
}
