package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.Default()
	g = CollectRouter(g)
	panic(g.Run())
}
