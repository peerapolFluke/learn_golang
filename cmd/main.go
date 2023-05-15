package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"

	"ginent/ent"
	"ginent/ent/migrate"
	"ginent/feature/resolvers"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db *ent.Client
}

var svr Server

func initDatabase() {
	client, err := ent.Open(dialect.Postgres, "host=localhost port=5433 user=root password=password dbname=simple_bank sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	svr.db = client
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(resolvers.NewSchema(svr.db))

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
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}

func main() {
	initDatabase()
	runHttpServer()
}

// package main

// import (
// 	"context"
// 	"log"

// 	"entgo.io/ent/dialect"
// 	"github.com/gin-gonic/gin"
// 	_ "github.com/lib/pq"
// 	"ginent/ent"
// 	"ginent/feature/todo"
// )

// type Server struct {
// 	db   *ent.Client
// 	http *gin.Engine
// }

// var svr Server

// func initDatabase() {
// 	client, err := ent.Open(dialect.Postgres, "host=localhost port=5433 user=root password=password dbname=simple_bank sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	svr.db = client

// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}
// }

// func runHttpServer() {
// 	r := gin.Default()
// 	svr.http = r
// 	r.GET("/get-todo", func(c *gin.Context) {
// 		todo.GetTodo(c, svr.db)
// 	})
// 	r.Run()
// }

// func main() {
// 	//setup server
// 	initDatabase()
// 	runHttpServer()
// }
