package ctrl

import (
	"day1/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindSoldierByRaUnCvApi(c *gin.Context) {
	ra := c.Query("Rarity")
	un := c.Query("UnlockArena")
	cv := c.Query("Cvc")
	c.String(http.StatusOK, handler.FindSoldierByRaUnCv(ra, un, cv))
}

func FindSoldierRaByIdApi(c *gin.Context) {
	id := c.Query("Id")
	c.String(http.StatusOK, "稀有度:"+handler.FindSoldierRaById(id))
}

func FindSoldierCoByIdApi(c *gin.Context) {
	id := c.Query("Id")
	c.String(http.StatusOK, "战力："+handler.FindSoldierCoById(id))
}

func FindSoldierByCvApi(c *gin.Context) {
	cvc := c.Query("Cvc")
	c.String(http.StatusOK, handler.FindSoldierByCv(cvc))
}

func FindSoldierByUnApi(c *gin.Context) {
	c.String(http.StatusOK, handler.FindSoldierByUn())
}
