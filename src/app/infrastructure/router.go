package infrastructure

import (
	"log"
	"net/http"
	"golang.org/x/exp/utf8string"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"github.com/taise-hub/shellgame/src/app/interfaces/controllers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Router() {
	r := gin.Default()
	r.Static("/static/assets/", "../static/assets")
	r.LoadHTMLGlob("../templates/*")
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	controller := controller.NewBattleController(NewSqlHandler())

	r.GET("/", func(c *gin.Context) { controller.Index(c) })
	r.GET("/index", func(c *gin.Context) { controller.Index(c) })
	battle := r.Group("/battle")
	{
		battle.GET("", func(c *gin.Context) { controller.Battle(c) })
		battle.Use(validate())
		{
			battle.POST("", func(c *gin.Context) { controller.New(c) })
		}
		battle.Use(hasSession())
		{
			battle.GET("/start", func(c *gin.Context) { controller.Start(c) })
			battle.GET("/wait", func(c *gin.Context) { controller.Wait(c) })
			battle.GET("/ws", func(c *gin.Context) {
				conn, err :=  upgrader.Upgrade(c.Writer, c.Request, nil)
				if err != nil {
					log.Printf("failed at WsBattle(): %s\n", err.Error())
					controller.Error500(c)
					return
				}
				controller.WsBattle(c, conn)
			 })
			battle.GET("/wswait", func(c *gin.Context) { })
		}
	}

	// r.GET("/jointimeattack", api_v1.GetJoinTimeAttack)
	// r.POST("/jointimeattack", api_v1.PostJoinTimeAttack)
	// r.GET("/timeattack", api_v1.GetTimeAttack)

	// r.Use(sessionCheck())
	// {
	// 	r.GET("/startbattle", api_v1.GetStartBattle)
	// 	r.GET("/wstimeattack", api_v1.WsTimeAttack)
	// }
	r.Run(":8080")
}

func hasSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessions.Default(c).Get("name") == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "nosession"})
			return
		} 
		c.Next()
	}
}

func validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		room := c.PostForm("room")
		if name == "" || room == "" {
			c.HTML(400, "400.html", "empty requests are not allowed")
			return
		}
		if (!utf8string.NewString(name).IsASCII() || !utf8string.NewString(room).IsASCII()) {
			c.HTML(400, "400.html", "only ASCII characters are allowed in tne name and room")
			return
		}
		c.Next()
	}
}