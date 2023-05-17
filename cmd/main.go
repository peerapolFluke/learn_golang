package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	config "seamoor/config"
	"seamoor/cors/auth"
	gincontext "seamoor/cors/gin-context"
	"seamoor/cors/injection"
	"seamoor/resolvers"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(resolvers.NewSchema())

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func runHttpServer() {
	r := gin.Default()
	r.Use(gincontext.GinContextToContextMiddleware())
	r.Use(auth.Middleware())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(config.Env.ServerAddress)
}

func main() {
	var err error
	err = config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	injection.InitDatabase()
	runHttpServer()
}
