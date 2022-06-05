package pkg

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	var engine *gin.Engine = gin.Default()

	return engine
}

func RunApi(uri string, engine *gin.Engine) {
	engine.Run(uri)
}
