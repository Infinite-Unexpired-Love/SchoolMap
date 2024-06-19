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
		Mobile    string `json:"mobile" binding:"required"`
		Password  string `json:"password" binding:"required"`
		CaptchaID string `json:"captchaId" binding:"required"`
		Captcha   string `json:"captcha" binding:"required,min=6,max=6"`
	}
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, handler.StatusBad("登录凭证错误", nil))
		return
	}

	if !store.Verify(loginDetails.CaptchaID, loginDetails.Captcha, true) {
		c.JSON(http.StatusBadRequest, &gin.H{
			"msg": "验证码错误",
		})
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
