package http

import (
	"github.com/a-dakani/go-schulung/http-server-gin/ginserver"
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
		context.JSON(http.StatusOK, autos)
	}
}

func StartServer(repository ginserver.AutoRepository) error {
	engine := gin.Default()
	engine.GET("/ping", handlePing)
	engine.GET("api/autos", handleGetAllAutos(repository))
	engine.PUT("api/autos/audi", handleAddAudi(repository))
	err := engine.Run("localhost:8080")
	if err != nil {
		return err
	}
	return nil
}
