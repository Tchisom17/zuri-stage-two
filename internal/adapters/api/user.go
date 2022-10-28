package api

import (
	"github.com/gin-gonic/gin"
)

func (u *HTTPHandler) Get(c *gin.Context) {
	userLists, err := u.UserService.Get()
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, userLists)
}
