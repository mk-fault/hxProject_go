package models

import (
	"hx/dao"
)

type WHO struct {
	Id               int     `gorm:"column:id"`
	IndicatorCode    string  `gorm:"column:IndicatorCode"`
	SpatialDimType   string  `gorm:"column:SpatialDimType"`
	SpatialDim       string  `gorm:"column:SpatialDim"`
	TimeDimType      *string  `gorm:"column:TimeDimType"`
	TimeDim          *float64  `gorm:"column:TimeDim"`
	Dim1Type         *string `gorm:"column:Dim1Type"`
	Dim1             *string `gorm:"column:Dim1"`
	Dim2Type         *string `gorm:"column:Dim2Type"`
	Dim2             *string `gorm:"column:Dim2"`
	Dim3Type         *string `gorm:"column:Dim3Type"`
	Dim3             *string `gorm:"column:Dim3"`
	Value            string  `gorm:"column:Value"`
	NumericValue     float64 `gorm:"column:NumericValue"`
	SpatialDimDetail *string `gorm:"column:SpatialDimDetail"`
}

func (WHO) TableName() string {
	return "who"
}

func GetWHO(param *WHO) (results []WHO) {
	dao.DBHX.Where(param).Order("TimeDim").Find(&results)
	return
}

func GetWHOSingleColumnDistinct[T string|int](param *WHO, column string) (results []T) {
	dao.DBHX.Table("who").Select("COALESCE(" + column + ", '')").Where(param).Pluck("COALESCE(" + column + ", '')", &results)
	results = removeRepeatedElement(results)
	return
}

// 通过map键的唯一性去重
func removeRepeatedElement[T string|int](s []T) []T {
	result := make([]T, 0)
	m := make(map[T]interface{}) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = nil
		}
	}
	return result
}
