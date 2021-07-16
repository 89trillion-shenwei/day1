package router

import (
	"day1/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	group1 := r.Group("")
	{
		group1.GET("/FindSoldierByRaUnCv", ctrl.FindSoldierByCvApi)
	}

	group2 := r.Group("")
	{
		group2.GET("/FindSoldierRaById", ctrl.FindSoldierRaByIdApi)
	}

	group3 := r.Group("")
	{
		group3.GET("/FindSoldierCoById", ctrl.FindSoldierCoByIdApi)
	}

	group4 := r.Group("")
	{
		group4.GET("/FindSoldierByCv", ctrl.FindSoldierByCvApi)
	}

	group5 := r.Group("")
	{
		group5.GET("/FindSoldierByUn", ctrl.FindSoldierByUnApi)
	}

	return r
}