package configfx

import (
	"github.com/adrg/xdg"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/config"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/config/configresolver"
	"go.uber.org/fx"
)

func New() fx.Option {
	options := []fx.Option{
		fx.Provide(config.New),
		fx.Provide(fx.Annotated{
			Group: "config_resolvers",
			Target: func() (configresolver.Resolver, error) {
				return configresolver.NewFromOsEnv(
					configresolver.WithPriority(-10),
				), nil
			},
		}),
		fx.Provide(
			fx.Annotated{
				Group: "config_resolvers",
				Target: func() (configresolver.Resolver, error) {
					return configresolver.NewFromYamlFile(
						"./config.yml",
						true,
						configresolver.WithPriority(10),
					)
				},
			},
		),
	}
	if configFilePath, err := xdg.ConfigFile("arr-interface/config.yml"); err == nil {
		options = append(options,
			fx.Provide(
				fx.Annotated{
					Group: "config_resolvers",
					Target: func() (configresolver.Resolver, error) {
						return configresolver.NewFromYamlFile(
							configFilePath,
							true,
							configresolver.WithPriority(20),
						)
					},
				},
			),
		)
	}
	return fx.Module(
		"config",
		fx.Options(options...),
	)
}
