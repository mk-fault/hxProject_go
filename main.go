package main

import (
	"hx/dao"
	"hx/routers"
)

func init() {
	dao.InitDB()
}

func main() {
	r := routers.SetupRouter()
	r.Run()
}
