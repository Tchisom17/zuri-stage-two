package api

import (
	"github.com/gin-gonic/gin"
	"zuri-stage-one/internal/core/dto"
	"zuri-stage-one/internal/core/models"
)

func (u *HTTPHandler) Calculate(c *gin.Context) {
	var body dto.OperationRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, err.Error())
		return
	}

	user, err := u.UserService.Get()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	operation := &models.Operation{
		OperationType: body.OperationType,
		X:             body.X,
		Y:             body.Y,
	}

	response, err := u.OperationService.Calculate(operation)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"slackUsername":  user.SlackUsername,
		"result":         response.Result,
		"operation_type": response.OperationType,
	})
}
