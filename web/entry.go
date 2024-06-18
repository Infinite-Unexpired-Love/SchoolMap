package web

import (
	"TGU-MAP/service"
	"TGU-MAP/utils"
	"TGU-MAP/web/handler"
	"TGU-MAP/web/handler/aliasItem"
	"TGU-MAP/web/handler/feedback"
	"TGU-MAP/web/handler/listItem"
	"TGU-MAP/web/handler/noticeItem"
	"TGU-MAP/web/handler/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"net/http"
)

func StartServer() error {
	r := gin.Default()
	r.Use(httpsHandler(), CorsMiddleware(), AuthMiddleware())

	r.POST("/login", user.HandleLogin)

	li := r.Group("/li")
	{
		li.GET("/", listItem.HandleGetData)
		//li.POST("/", handler.HandleInsertData)
		li.POST("/item/:id", listItem.HandleInsertNode)
		li.PUT("/item/:id", listItem.HandleUpdateNode)
		li.DELETE("/item/:id", listItem.HandleDeleteNode)
	}

	al := r.Group("/al")
	{
		al.GET("/", aliasItem.HandleGetData)
		al.POST("/item", aliasItem.HandleInsertNode)
		al.DELETE("/item/:id", aliasItem.HandleDeleteNode)
	}

	no := r.Group("/no")
	{
		no.GET("/", noticeItem.HandleGetData)
		no.POST("/item", noticeItem.HandleInsertNode)
		no.PUT("/item/:id", noticeItem.HandleUpdateNode)
		no.DELETE("/item/:id", noticeItem.HandleDeleteNode)
	}

	fe := r.Group("/fe")
	{
		fe.GET("/", feedback.HandleGetData)
		fe.POST("/item", feedback.HandleInsertNode)
		fe.DELETE("/item/:id", feedback.HandleDeleteNode)
	}

	return r.RunTLS(fmt.Sprintf(":%d", service.GlobalConfig.Web.Port), service.GlobalConfig.Web.Cert, service.GlobalConfig.Web.Key)
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusNoContent, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

var whiteList = []string{"/login", "/li/", "/al/", "/no/", "/fe/item"}

// AuthMiddleware 验证JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, uri := range whiteList {
			if c.Request.RequestURI == uri {
				c.Next()
				return
			}
		}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, handler.StatusBad("无效token", nil))
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, handler.StatusBad("无效token", nil))
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()

	}
}

func httpsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		secureMiddle := secure.New(secure.Options{
			SSLRedirect: true, //只允许https请求
			//SSLHost:"" //http到https的重定向
			STSSeconds:           1536000, //Strict-Transport-Security header的时效:1年
			STSIncludeSubdomains: true,    //includeSubdomains will be appended to the Strict-Transport-Security header
			STSPreload:           true,    //STS Preload(预加载)
			FrameDeny:            true,    //X-Frame-Options 有三个值:DENY（表示该页面不允许在 frame 中展示，即便是在相同域名的页面中嵌套也不允许）、SAMEORIGIN、ALLOW-FROM uri
			ContentTypeNosniff:   true,    //禁用浏览器的类型猜测行为,防止基于 MIME 类型混淆的攻击
			BrowserXssFilter:     true,    //启用XSS保护,并在检查到XSS攻击时，停止渲染页面
			//IsDevelopment:true,  //开发模式
		})
		err := secureMiddle.Process(context.Writer, context.Request)
		// 如果不安全，终止.
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "数据不安全")
			return
		}
		// 如果是重定向，终止
		if status := context.Writer.Status(); status > 300 && status < 399 {
			context.Abort()
			return
		}
		context.Next()
	}
}
