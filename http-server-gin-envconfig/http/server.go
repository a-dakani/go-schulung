package http

import (
	"context"
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin-envconfig/ginserver"
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

func handleAddAudi(repository ginserver.AutoRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var audi ginserver.Audi
		err := c.BindJSON(&audi)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err = repository.AddAuto(c.Request.Context(), audi)

	}
}
func handleGetAllAutos(repository ginserver.AutoRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		autos, err := repository.GetAllAutos(c.Request.Context())
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

func StartServer(config Config, repository ginserver.AutoRepository) error {
	engine := gin.Default()

	authorized := engine.Group("/api", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"bar": "foo",
	}), XtraceIdMiddleware)

	authorized.GET("autos", handleGetAllAutos(repository))
	authorized.PUT("autos/audi", handleAddAudi(repository))

	engine.GET("/ping", handlePing)

	err := engine.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		return err
	}
	return nil
}
