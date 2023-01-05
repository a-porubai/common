package http_handler

import (
	"github.com/gin-gonic/gin"

	"github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/ports"
)

type GrpcHandler struct {
	gamesService ports.GamesService
}

func NewGrpcHandler(gamesService ports.GamesService) *GrpcHandler {
	return &GrpcHandler{
		gamesService: gamesService,
	}
}

func (hdl *GrpcHandler) Get(c *gin.Context) {
	game, err := hdl.gamesService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, game)
}

func (hdl *GrpcHandler) Create(c *gin.Context) {
	body := BodyCreate{}
	c.BindJSON(&body)

	game, err := hdl.gamesService.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, BuildResponseCreate(game))
}

func (hdl *GrpcHandler) RevealCell(c *gin.Context) {
	body := BodyRevealCell{}
	c.BindJSON(&body)

	game, err := hdl.gamesService.Reveal(c.Param("id"), body.Row, body.Col)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, BuildResponseRevealCell(game))
}
