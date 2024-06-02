package player

import (
	"github.com/gin-gonic/gin"
	"github.com/sebaperi/go-skeleton/internal/domain"
	"net/http"
)

func (h Handler) CreatePlayer(c *gin.Context) {
	// Traducir el request
	// Validacion a nivel http (ej json incorrecto)
	// Consumir un servicio
	// Traducir el response
	var playerCreateParams domain.Player
	if err := c.BindJSON(&playerCreateParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedId, err := h.PlayerService.Create(playerCreateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"player_id": insertedId})
}
