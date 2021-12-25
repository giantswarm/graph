package main

import (
	"github.com/giantswarm/graph/api"
)

func main() {
	router := api.CreateRouter()
	router.Run(":3000")
}
