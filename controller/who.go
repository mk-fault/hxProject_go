package controller

import (
	"hx/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetWHOList(c *gin.Context) {
	code, ok := c.GetQuery("CODE")
	if !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "code参数必填",
		})
		return
	}
	year := c.Query("YEAR")
	spatialDimType := c.Query("SPAT")
	spatial := c.Query("SPA")
	dim1 := c.Query("D1")
	dim2 := c.Query("D2")
	dim3 := c.Query("D3")

	param := &models.WHO{
		IndicatorCode: code,
	}
	if year != "" {
		numYear, err := strconv.ParseFloat(year, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"msg": "year参数必须是数字",
			})
		}
		param.TimeDim = &numYear
	}
	if spatialDimType != "" {
		param.SpatialDimType = spatialDimType
	}
	if spatial != "" {
		param.SpatialDim = spatial
	}
	if dim1 != "" {
		param.Dim1 = &dim1
	}
	if dim2 != "" {
		param.Dim2 = &dim2
	}
	if dim3 != "" {
		param.Dim3 = &dim3
	}

	results := models.GetWHO(param)
	c.JSON(http.StatusOK, Response[models.WHO]{
		Count:   len(results),
		Results: results,
	})
}

func GetWHOIndex(c *gin.Context) {
	code, ok := c.GetQuery("CODE")
	if !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "code参数必填",
		})
		return
	}

	// 制作空间维度类型字典
	spatialDimType := models.GetWHOSingleColumnDistinct[string](&models.WHO{IndicatorCode: code}, "SpatialDimType")

	// 制作时间维度字典
	lst := models.GetWHOSingleColumnDistinct[string](&models.WHO{IndicatorCode: code}, "TimeDimType")
	year := make(map[string][]int)
	for _, v := range lst {
		if v == "" {
			continue
		}
		year[v] = models.GetWHOSingleColumnDistinct[int](&models.WHO{IndicatorCode: code, TimeDimType: &v}, "TimeDim")
	}

	// 制作第一维度字典
	lst = models.GetWHOSingleColumnDistinct[string](&models.WHO{IndicatorCode: code}, "Dim1Type")
	Dim1Dic := make(map[string][]string)
	for _, v := range lst {
		if v == "" {
			continue
		}
		Dim1Dic[v] = models.GetWHOSingleColumnDistinct[string](&models.WHO{IndicatorCode: code, Dim1Type: &v}, "Dim1")
	}

	// 制作第二维度字典
	lst = models.GetWHOSingleColumnDistinct[string](&models.WHO{IndicatorCode: code}, "Dim2Type")
	Dim2Dic := make(map[string][]string)
	for _, v := range lst {
		if v == "" {
			continue
		}
		Dim2Dic[v] = models.GetWHOSingleColumnDistinct[string](&models.WHO{IndicatorCode: code, Dim2Type: &v}, "Dim2")
	}

	// 制作第三维度字典
	lst = models.GetWHOSingleColumnDistinct[string](&models.WHO{IndicatorCode: code}, "Dim3Type")
	Dim3Dic := make(map[string][]string)
	for _, v := range lst {
		if v == "" {
			continue
		}
		Dim3Dic[v] = models.GetWHOSingleColumnDistinct[string](&models.WHO{IndicatorCode: code, Dim3Type: &v}, "Dim3")
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"SpatialDimTypeDic": spatialDimType,
		"TimeDimDic":        year,
		"Dim1Dic":           Dim1Dic,
		"Dim2Dic":           Dim2Dic,
		"Dim3Dic":           Dim3Dic,
	})

	// lst := models.GetWHOSingleColumn(code, "Dim1Type")
	// Dim1Dic := make(map[string][]string)
	// for _, v := range lst {
	// 	Dim1Dic[v] = models.GetWHO(map[string]interface{}{"IndicatorCode": code, Dim1Type: v})
	// }

}
