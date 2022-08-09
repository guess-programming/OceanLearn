package main

import (
	"github.com/gin-gonic/gin"
	"github.com/myusername/OceanLearn/controller"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auto/register", controller.Register)
	return r
}
