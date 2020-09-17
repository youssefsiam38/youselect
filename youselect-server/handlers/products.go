package handlers

import (
	"github.com/youssefsiam38/youselect/framework"
	"net/http"
	"github.com/youssefsiam38/youselect/db"
	"github.com/gin-gonic/gin"
)

func AllProducts(c *gin.Context) {
	products, err := db.SelectAllProducts()
	if err != nil {
		framework.Log(err)
		c.JSON(http.StatusNotFound, ErrorRes{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func Products(c *gin.Context) {
	page, min, max, store := c.Query("page"), c.Query("min"), c.Query("max"), c.Query("store")
	products, err := db.SelectProducts(false, false, &page, &min, &max, &store, &max)
	if err != nil {
		framework.Log(err)
		c.JSON(http.StatusNotFound, ErrorRes{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func Search(c *gin.Context) {
	page, min, max, store, search := c.Query("page"), c.Query("min"), c.Query("max"), c.Query("store"), c.Query("s")
	products, err := db.SelectProducts(true, false, &page, &min, &max, &store, &search)
	if err != nil {
		if err.Error() == "Not Found" {
			framework.Log(err, search)
		} else {
			framework.Log(err)
		}
		c.JSON(http.StatusNotFound, ErrorRes{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}