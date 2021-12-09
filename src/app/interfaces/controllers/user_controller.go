package controller

import (
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/taise-hub/shellgame/src/app/interfaces/database"
)


type BattleController interface {
	Index(Context)
	Battle(Context)
	New(Context)
	Start(Context)
	Wait(Context)
	WsBattle(Context, Conn)

	Error500(Context)
}

type battleController struct {
	//uc  HOGEHOGEUsecase
	//svc HOGEHOGEservice
}

func NewBattleController(sqlHandler database.SqlHandler) BattleController {
	return &battleController {
	}
}

// GET /
func (con *battleController) Index(c Context) {
	c.HTML(200, "index.html", nil)
}

// GET /battle
func (con *battleController) Battle(c Context) {
	c.HTML(200, "new.html", nil)
}

// POST /battle
func (con *battleController) New(c Context) {
	// Same as sessions.Defalut()
	session := c.MustGet(sessions.DefaultKey).(sessions.Session)

	session.Set("name", c.PostForm("name"))
	session.Set("room", c.PostForm("room"))
	if err := session.Save(); err != nil {
		log.Printf("failed at PostJoinBattle(): %s\n", err.Error())
		con.Error500(c)
	}
	c.Redirect(302, "/battle/wait")
}

// GET /battle/start
func (con *battleController) Start(c Context) {
	c.HTML(200, "battle.html", nil)
}

/// GET /battle/wait
func (con *battleController) Wait(c Context) {
	c.HTML(200, "wait.html", nil)
}

func (con *battleController) WsBattle(c Context, conn Conn) {
	// session := c.MustGet(sessions.DefaultKey).(sessions.Session)
	// id := session.Get("room").(string) + "_" + session.Get("name").(string)
	
	// コンテナの生成。
	// containerUsecase.Run(id)
	// roomとplayerインスタンスの生成
	// PlayerUsecase.Join(player, room)
	// PlayerServiceの中で、
	// player.Join(room)
}

func (con *battleController) Error500(c Context) {
	c.HTML(500, "500.html", nil)
}

func (con *battleController) Error400(c Context) {
	c.HTML(400, "400.html", nil)
}
