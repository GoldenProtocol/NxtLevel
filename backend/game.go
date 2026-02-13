package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameModel struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name   string    `json:"name"`
	NodeID string    `json:"node_id"`
}

func (GameModel) TableName() string {
	return "games"
}

type Game struct {
	version string
}

func (g *Game) createHandler(ctx *gin.Context) {
	db, found := ctx.Get("db")

	if !found {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	orm := db.(*gorm.DB)

	game := new(GameModel)
	if err := json.NewDecoder(ctx.Request.Body).Decode(game); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad Request",
		})
		return
	}

	orm.Create(game)

	ctx.JSON(http.StatusCreated, game)
}

func (g *Game) Router(router *gin.Engine) {
	group := router.Group(g.version)
	games := group.Group("/games")
	games.POST("", g.createHandler)
}

func NewGameController(version string) *Game {
	return &Game{
		version,
	}
}
