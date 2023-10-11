package http

import (
	"context"
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin-kafka/ginserver"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

func XtraceIdMiddleware(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "x-trace-id", c.GetHeader("x-trace-id"))
	c.Request = c.Request.WithContext(ctx)
}

func handlePing(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func handleAddAudi(as *ginserver.AutoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var audi ginserver.Audi
		err := c.BindJSON(&audi)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err = as.AddAuto(c.Request.Context(), audi)

	}
}
func handleGetAllAutos(as *ginserver.AutoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		autos, err := as.GetAllAutos(c.Request.Context())
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Negotiate(http.StatusOK, gin.Negotiate{
			Offered: []string{gin.MIMEJSON, gin.MIMEXML},
			Data:    autos,
		})
	}
}

func StartServer(binding string, as *ginserver.AutoService) error {
	engine := gin.Default()

	authorized := engine.Group("/api", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"bar": "foo",
	}), XtraceIdMiddleware)

	authorized.GET("autos", handleGetAllAutos(as))
	authorized.PUT("autos/audi", handleAddAudi(as))

	engine.GET("/ping", handlePing)

	err := engine.Run(fmt.Sprintf(binding))
	if err != nil {
		return err
	}
	return nil
}
