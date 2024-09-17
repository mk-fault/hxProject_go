package controller

import (
	"hx/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGBDList(c *gin.Context) {
	var results []map[string]interface{}

	title := c.Query("TITLE")
	if title == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "title参数必填",
		})
		return
	}

	year := c.Query("YEAR")
	location := c.Query("LOCATION")
	page := c.DefaultQuery("PAGE", "1")
	pageSize := c.DefaultQuery("PAGE_SIZE", "10000")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "page参数必须是数字",
		})
		return
	}

	intPageSize, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "pageSize参数必须是数字",
		})
		return
	}

	query := dao.DBHXGBD.Table(title)
	if year != "" {
		query = query.Where("year = ?", year).Order("year")
	}
	if location != "" {
		query = query.Where("location = ?", location)
	}
	var totalCount int64
	query.Count(&totalCount)
	query.Offset((intPage - 1) * intPageSize).Limit(intPageSize).Find(&results)

	c.JSON(http.StatusOK, gin.H{
		"total_count":       totalCount,
		"total_page":        int(totalCount)/intPageSize + 1,
		"current_page":      intPage,
		"current_page_size": intPageSize,
		"data":              results,
	})
}
