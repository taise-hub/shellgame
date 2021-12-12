package infrastructure

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	controller "github.com/taise-hub/shellgame/src/app/interfaces/controllers"
)

func Router() {
	r := gin.Default()
	r.Static("/static/assets/", "../static/assets")
	r.LoadHTMLGlob("../templates/*")
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	controller := controller.NewBattleController(NewSqlHandler(), NewContainerHandler(), NewWebSocketHandler())

	r.GET("/", func(c *gin.Context) { controller.Index(c) })
	r.GET("/index", func(c *gin.Context) { controller.Index(c) })
	battle := r.Group("/battle")
	{
		battle.GET("", func(c *gin.Context) { controller.Battle(c) })
		battle.POST("", func(c *gin.Context) { controller.New(c) })
		battle.Use(hasSession())
		{
			battle.GET("/start", func(c *gin.Context) { controller.Start(c) })
			battle.GET("/wait", func(c *gin.Context) { controller.Wait(c) })
			battle.GET("/wswait", func(c *gin.Context) { 
				conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
				if err != nil {
					controller.Error500(c, err.Error())
				}
				controller.WsWait(c, conn)
			})
			battle.GET("/ws", func(c *gin.Context) {
				conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
				if err != nil {
					controller.Error500(c, err.Error())
				}
				controller.WsBattle(c, conn)
			 })
		}
	}
	r.Run(":8080")
}

func hasSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessions.Default(c).Get("player") == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "nosession"})
			return
		} 
		c.Next()
	}
}