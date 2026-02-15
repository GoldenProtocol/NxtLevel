package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NodeModel struct {
	gorm.Model
	CompanyNode
	ID string `gorm:"primaryKey" json:"id"`
}

func (NodeModel) TableName() string {
	return "nodes"
}

type CompanyNode struct {
	Type              string  `json:"type"`
	PositionAbsoluteX float64 `json:"positionAbsoluteX"`
	PositionAbsoluteY float64 `json:"positionAbsoluteY"`
	Selected          bool    `json:"selected"`
	Selectable        bool    `json:"selectable"`
	Draggable         bool    `json:"draggable"`
	Deletable         bool    `json:"deletable"`
	IsConnectable     bool    `json:"isConnectable"`
	Dragging          bool    `json:"dragging"`
	ZIndex            int     `json:"zIndex"`
	Width             float64 `json:"width"`
	Height            float64 `json:"height"`
}

type NodeData struct {
	Label string `json:"label"`
}

type Node struct {
	version string
}

func (n *Node) readAll(ctx *gin.Context) {
	db, found := ctx.Get("db")

	if !found {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	g := db.(*gorm.DB)

	nodes := []*NodeModel{}

	g.Find(&nodes)

	companyNodes := ExtractNodeId(FilterNodesByType(nodes, "CompanyNode"))
	gameNodes := ExtractNodeId(FilterNodesByType(nodes, "GameNode"))

	companyRecords := make([]*CompanyModel, 0, len(companyNodes))
	gameRecords := make([]*GameModel, 0, len(gameNodes))
	g.Find(&companyRecords).Where("node_id IN ?", companyNodes)
	g.Find(&gameRecords).Where("node_id IN ?", gameNodes)

}

func (n *Node) createHandler(ctx *gin.Context) {
	db, found := ctx.Get("db")

	if !found {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	g := db.(*gorm.DB)

	node := new(NodeModel)
	if err := json.NewDecoder(ctx.Request.Body).Decode(node); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad Request",
		})
		return
	}

	g.Create(node)

	ctx.JSON(http.StatusCreated, node)
}

func (n *Node) Router(router *gin.Engine) {
	group := router.Group(n.version)
	nodes := group.Group("/nodes")
	nodes.POST("", n.createHandler)
	nodes.GET("", n.readAll)
}

func NewNodeController(version string) *Node {
	return &Node{
		version,
	}
}
