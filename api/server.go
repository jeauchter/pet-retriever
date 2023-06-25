package api

import (
	"github.com/jeremyauchter/pet-retriever/api/controllers"
)

var server = controllers.Server{}

func Run() {
	server.Initialize()
	server.Run(":8080")
}
