package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zazin/go-country/src/database"
	"github.com/zazin/go-country/src/models"
	"gorm.io/gorm"
	"net/http"
)

type CountryRepository struct {
	Db *gorm.DB
}

func New() *CountryRepository {
	db := database.InitDb()
	err := db.AutoMigrate(&models.Country{})
	if err != nil {
		return nil
	}
	return &CountryRepository{Db: db}
}

//get countries
func (repository *CountryRepository) GetCountries(c *gin.Context) {
	var countries []models.Country
	err := models.ListCountry(repository.Db, &countries)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, countries)
}
