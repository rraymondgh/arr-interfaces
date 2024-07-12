package httpserver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/httpserver"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	Schema graphql.ExecutableSchema
	Logger *zap.SugaredLogger
}

type Result struct {
	fx.Out
	Option httpserver.Option `group:"http_server_options"`
}

func New(p Params) Result {
	return Result{
		Option: &builder{
			schema: p.Schema,
		},
	}
}

type builder struct {
	schema graphql.ExecutableSchema
}

func (builder) Key() string {
	return "graphql"
}

func (b builder) Apply(e *gin.Engine) error {
	gql := handler.NewDefaultServer(b.schema)
	e.POST("/graphql", func(c *gin.Context) {
		gql.ServeHTTP(c.Writer, c.Request)
	})
	pg := playground.Handler("GraphQL playground", "/graphql")
	e.GET("/graphql", func(c *gin.Context) {
		pg.ServeHTTP(c.Writer, c.Request)
	})
	return nil
}
