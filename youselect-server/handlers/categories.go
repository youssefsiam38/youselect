package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/youssefsiam38/youselect/db"
	"github.com/youssefsiam38/youselect/framework"
)

func GetCategories(c *gin.Context) {
	categs, err := db.SelectAllCategories()
	if err != nil {
		framework.Log(err)
		c.JSON(http.StatusNotFound, ErrorRes{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, categs)
}
