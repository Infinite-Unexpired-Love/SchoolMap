package user

import (
	"TGU-MAP/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

// 使用了最基础的存储方式，验证码存储在内存中，重启应用后数据丢失
var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, ans, err := cp.Generate()
	utils.Debug("生成验证码：", ans)
	if err != nil {
		utils.Error("生成验证码错误：", err)
		c.JSON(http.StatusInternalServerError, &gin.H{
			"msg": "暂时无法登录",
		})
		return
	}
	c.JSON(http.StatusOK, &gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
