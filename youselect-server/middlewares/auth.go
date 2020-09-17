package middlewares

import (
	"github.com/youssefsiam38/youselect/framework"
	"github.com/youssefsiam38/youselect/handlers"
	"github.com/youssefsiam38/youselect/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Auth to make authentication required
func Auth(c *gin.Context) {
	// Get Authorization token
	token := c.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)

	// retrievedUser, err := utils.VerifyJWT(auth)

	err := db.CheckAdminByToken(&token)

	if err != nil {
		framework.Log(err)
		c.JSON(http.StatusUnauthorized, handlers.ErrorRes{
			Error: "Something went wrong with authorization please login",
		})
		c.Abort()
	}
	c.Next()
}

// (c, http.StatusUnauthorized, "Something went wrong with authorization please login")