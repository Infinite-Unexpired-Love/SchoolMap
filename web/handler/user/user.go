package user

import (
	"TGU-MAP/service"
	"TGU-MAP/utils"
	"TGU-MAP/web/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleLogin(c *gin.Context) {
	var loginDetails struct {
		Mobile   string `json:"mobile" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, handler.StatusBad("登录凭证错误", nil))
		return
	}

	user, err := service.UserClient.FindElemByMobile(loginDetails.Mobile)
	if err != nil || loginDetails.Password != user.Password {
		c.JSON(http.StatusUnauthorized, handler.StatusBad("登录凭证错误", nil))
		return
	}

	token, err1 := utils.GenerateToken(user.ID)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, handler.StatusBad("出错了……", nil))
		return
	}
	c.JSON(http.StatusOK, handler.StatusOK("success", gin.H{
		"token":    token,
		"username": user.Username,
		"mobile":   user.Mobile,
	}))
}
