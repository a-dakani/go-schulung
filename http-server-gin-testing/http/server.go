package http

import (
	"github.com/a-dakani/go-schulung/http-server-gin-testing/ginserver"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handlePing(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func handleAddAudi(repository ginserver.AutoRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		var audi ginserver.Audi
		err := context.BindJSON(&audi)
		if err != nil {
			log.Println(err)
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err = repository.AddAuto(audi)

	}
}
func handleGetAllAutos(repository ginserver.AutoRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		autos, err := repository.GetAllAutos()
		if err != nil {
			log.Println(err)
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		context.Negotiate(http.StatusOK, gin.Negotiate{
			Offered: []string{gin.MIMEJSON, gin.MIMEXML},
			Data:    autos,
		})
	}
}

func StartServer(repository ginserver.AutoRepository) error {
	engine := gin.Default()

	authorized := engine.Group("/api", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"bar": "foo",
	}))

	authorized.GET("autos", handleGetAllAutos(repository))
	authorized.PUT("autos/audi", handleAddAudi(repository))

	engine.GET("/ping", handlePing)

	err := engine.Run("localhost:8080")
	if err != nil {
		return err
	}
	return nil
}
