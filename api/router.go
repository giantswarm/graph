package api

import (
	"flag"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"

	"github.com/giantswarm/graph/api/github"
	"github.com/giantswarm/graph/ent"
)

func CreateRouter() *gin.Engine {
	addr := flag.String("gremlin-server", os.Getenv("GREMLIN_SERVER"), "gremlin server address")
	flag.Parse()
	if *addr == "" {
		log.Fatal("missing gremlin server address")
	}
	client, err := ent.Open(dialect.Gremlin, *addr)
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}
	defer client.Close()

	r := gin.Default()

	// Setup routes for GitHub entities
	err = github.SetupRoutes(r, client)
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}

	return r
}
