package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youssefsiam38/youselect/db"
	"github.com/youssefsiam38/youselect/framework"
	"github.com/youssefsiam38/youselect/models"
)

func Login(c *gin.Context) {
	var admin models.Admin

	c.BindJSON(&admin)

	verified, tokenPtr, err := db.CheckAdmin(&admin)

	if err != nil {
		framework.Log(err)
		c.JSON(http.StatusNotFound, ErrorRes{
			Error: err.Error(),
		})
		return
	}

	if verified {
		err := framework.SendTokenEmail(tokenPtr)
		if err != nil {
			framework.Log(err)
			c.JSON(http.StatusNotFound, ErrorRes{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, MessageRes{
			Message: "Check your Email",
		})
		return
	}
	c.JSON(http.StatusOK, MessageRes{
		Message: "Wrong credentials",
	})

}
