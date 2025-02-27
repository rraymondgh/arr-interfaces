package healthcheck

import (
	"github.com/hellofresh/health-go/v5"
	"github.com/rraymondgh/arr-interfaces/internal/version"
	"go.uber.org/fx"
)

type Result struct {
	fx.Out
	HealthcheckOption health.Option `group:"healthcheck_options"`
}

func New() Result {
	return Result{
		HealthcheckOption: health.WithComponent(health.Component{
			Name:    "arr-interfaces",
			Version: version.GitTag,
		}),
	}
}
