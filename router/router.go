package routers

import (
  "github.com/gin-gonic/gin"
  . "../controllers" //constroller部分
 )

func InitRouter() *gin.Engine{
    router := gin.Default()
    router.GET("/", Test)
    return router
}

