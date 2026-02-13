package main

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
	ID     string `json:"id"`
}

type EdgeModel struct {
	gorm.Model
	Edge
	ID string `json:"id" gorm:"primaryKey"`
}

func (EdgeModel) TableName() string {
	return "connections"
}

type EdgeController struct {
	version string
}

func (e *EdgeController) createHandler(ctx *gin.Context) {
	db, found := ctx.Get("db")

	if !found {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	g := db.(*gorm.DB)

	edges := []*EdgeModel{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&edges); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad Request",
		})
		return
	}

	ids := []string{}

	for _, edge := range edges {
		ids = append(ids, edge.ID)
	}

	existing := []*EdgeModel{}

	g.Find(&existing, ids)

	toStore := []*EdgeModel{}

	for _, edge := range edges {
		index := slices.IndexFunc(existing, func(e *EdgeModel) bool {
			return edge.ID == e.ID
		})

		if index == -1 {
			toStore = append(toStore, edge)
		}
	}

	g.Create(toStore)

	ctx.JSON(http.StatusCreated, toStore)
}

func (e *EdgeController) Router(router *gin.Engine) {
	group := router.Group(e.version)
	nodes := group.Group("/connections")
	nodes.POST("", e.createHandler)
}

func NewEdgeController(version string) *EdgeController {
	return &EdgeController{
		version,
	}
}
