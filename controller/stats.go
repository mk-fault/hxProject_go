package controller

import (
	"hx/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func GetStatsQGList(c *gin.Context) {
	code, ok := c.GetQuery("CODE")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "code参数必填",
		})
		return
	}
	year := c.Query("YEAR")

	results := models.GetStatsQG(code, year)
	c.JSON(http.StatusOK, Response[models.StatsQG]{
		Count: len(results),
		Results: results,
	})
}

func GetStatsGSList(c *gin.Context) {
	code, ok := c.GetQuery("CODE")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "code参数必填",
		})
		return
	}
	year := c.Query("YEAR")
	prov := c.Query("PROV")

	results := models.GetStatsGS(code, year, prov)
	c.JSON(http.StatusOK, Response[models.StatsGS]{
		Count: len(results),
		Results: results,
	})
}