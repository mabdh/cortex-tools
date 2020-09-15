package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/grafana/cortex-tools/pkg/commands"
	"github.com/grafana/cortex-tools/pkg/version"
)

var (
	ruleCommand         commands.RuleCommand
	alertCommand        commands.AlertCommand
	alertmanagerCommand commands.AlertmanagerCommand
	logConfig           commands.LoggerConfig
	pushGateway         commands.PushGatewayConfig
	loadgenCommand      commands.LoadgenCommand
)

func main() {
	app := kingpin.New("cortextool", "A command-line tool to manage cortex.")
	logConfig.Register(app)
	alertCommand.Register(app)
	alertmanagerCommand.Register(app)
	ruleCommand.Register(app)
	pushGateway.Register(app)
	loadgenCommand.Register(app)

	app.Command("version", "Get the version of the cortextool CLI").Action(func(k *kingpin.ParseContext) error {
		fmt.Print(version.Template)
		version.CheckLatest()

		return nil
	})

	kingpin.MustParse(app.Parse(os.Args[1:]))

	pushGateway.Stop()
}
