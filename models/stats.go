package models

import (
	"hx/dao"
)

type StatsQG struct {
	Id      int     `gorm:"column:id" json:"-"`
	Measure string  `gorm:"column:measure" json:"measure"`
	Code    string  `gorm:"column:code" json:"-"`
	Year    string     `gorm:"column:year" json:"year"`
	Value   float64 `gorm:"column:value" json:"value"`
}

type StatsGS struct {
	Id       int     `gorm:"column:id" json:"-"`
	Measure  string  `gorm:"column:measure" json:"measure"`
	MeaCode  string  `gorm:"column:mea_code" json:"-"`
	Province string  `gorm:"column:province" json:"province"`
	ProvCode string  `gorm:"column:prov_code" json:"-"`
	Year     string     `gorm:"column:year" json:"year"`
	Value    float64 `gorm:"column:value" json:"value"`
}

func (StatsQG) TableName() string {
	return "stats_qg"
}

func (StatsGS) TableName() string {
	return "stats_gs"
}

func GetStatsQG(code string, year string) (results []StatsQG) {
	query := dao.DBHX.Where("code = ?", code)
	if year != "" {
		query = query.Where("year = ?", year).Order("year")
	}
	query.Find(&results)
	return
}

func GetStatsGS(code string, year string, prov string) (results []StatsGS) {
	query := dao.DBHX.Where("mea_code = ?", code)
	if year != "" {
		query = query.Where("year = ?", year).Order("year")
	}
	if prov != "" {
		query = query.Where("prov_code = ?", prov)
	}
	query.Find(&results)
	return
}
