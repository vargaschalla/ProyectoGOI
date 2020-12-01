package app

import (
	"PROYECTintegrador/ProyectoGOI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IntemsIndex(c *gin.Context) {
	s := models.Item{Title: "Wagner", Notes: "Hola mundo"}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome " + s.Title,
	})
}
