package handlers

import (
	"net/http"
	"github.com/youssefsiam38/youselect/db"
	"github.com/youssefsiam38/youselect/framework"
	"github.com/gin-gonic/gin"
)

func GetStoresNames(c *gin.Context) {
	type storeName struct{
		ID uint `json:"id"`
		Name string `json:"name"`
	}

	storesPtr, err := db.SelectAllStores()
	if err != nil {
		framework.Log(err)
		c.JSON(http.StatusNotFound, ErrorRes{
			Error: err.Error(),
		})
		return
	}
	stores := *storesPtr

	var storesNames []storeName

	for i := 0; i < len(stores); i++ {
		storesNames = append(storesNames, storeName{
			ID: stores[i].ID,
			Name: stores[i].Name,
		})
	}

	c.JSON(http.StatusOK, storesNames)
}