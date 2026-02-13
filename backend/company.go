package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyModel struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name   string    `json:"name"`
	NodeID string    `json:"node_id"`
}

func (CompanyModel) TableName() string {
	return "companies"
}

type Company struct {
	version string
}

func (c *Company) createHandler(ctx *gin.Context) {
	db, found := ctx.Get("db")

	if !found {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	orm := db.(*gorm.DB)

	company := new(CompanyModel)
	if err := json.NewDecoder(ctx.Request.Body).Decode(company); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad Request",
		})
		return
	}

	orm.Create(company)

	ctx.JSON(http.StatusCreated, company)
}

func (c *Company) Router(router *gin.Engine) {
	group := router.Group(c.version)
	companies := group.Group("/companies")
	companies.POST("", c.createHandler)
}

func NewCompanyController(version string) *Company {
	return &Company{
		version,
	}
}
