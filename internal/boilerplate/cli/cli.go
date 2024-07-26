package cli

import (
	"context"
	"sort"
	"strings"

	"github.com/rraymondgh/arr-interface/internal/version"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	Args       []string `name:"cli_args"`
	Lifecycle  fx.Lifecycle
	Shutdowner fx.Shutdowner
	Commands   []*cli.Command `group:"commands"`
	Logger     *zap.SugaredLogger
}

type Result struct {
	fx.Out
	App *cli.App
}

func New(p Params) (Result, error) {
	commands := p.Commands
	sort.Slice(commands, func(i, j int) bool {
		return strings.Compare(commands[i].Name, commands[j].Name) < 0
	})
	name := "arr-interface"
	if version.GitTag != "" {
		name += " " + version.GitTag
	}
	app := &cli.App{
		Name:     name,
		Commands: commands,
		CommandNotFound: func(_ *cli.Context, command string) {
			p.Logger.Errorw("command not found", "command", command)
		},
		After: func(ctx *cli.Context) error {
			return p.Shutdowner.Shutdown()
		},
		Version: version.GitTag,
		// disabling the version flag as for some reason the After hook never gets called
		HideVersion: true,
	}
	app.Setup()
	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go (func() {
				if err := app.RunContext(context.Background(), p.Args); err != nil {
					panic(err)
				}
			})()
			return nil
		},
	})
	return Result{
		App: app,
	}, nil
}
