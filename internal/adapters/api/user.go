package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zuri-stage-one/internal/core/helpers"
)

func (u *HTTPHandler) Get(c *gin.Context) {
	userLists, err := u.UserService.Get()
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"An internal server error"})
		return
	}
	helpers.JSON(c, "", http.StatusOK, userLists, nil)
}
