package web

import (
	"TGU-MAP/web/items"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	r := gin.Default()
	r.GET("/data", items.HandleGetData)

	r.POST("/data", items.HandleInsertData)

	r.PUT("/node/:id", items.HandleUpdateNode)

	r.DELETE("/node/:id", items.HandleDeleteNode)

	return r.Run(":8080")
}
