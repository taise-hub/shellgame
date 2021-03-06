package controller

import (
	"github.com/taise-hub/shellgame/src/app/usecase"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/exp/utf8string"
	"github.com/gin-contrib/sessions"
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"github.com/taise-hub/shellgame/src/app/domain/service"
	"github.com/taise-hub/shellgame/src/app/interfaces/container"
	"github.com/taise-hub/shellgame/src/app/interfaces/database"
	"github.com/taise-hub/shellgame/src/app/interfaces/websocket"
	"log"
)

type BattleController interface {
	Index(Context)
	NewGame(Context, string)
	Register(Context)
	Start(Context)
	Wait(Context)
	WsWait(Context, model.Connection)
	WsBattle(Context, model.Connection)
	Error500(Context)
	Error400(Context, string)
}

type battleController struct {
	uc usecase.BattleUsecase
}

func NewBattleController(sqlHandler database.SqlHandler, containerHandler container.ContainerHandler, webSocketHandler websocket.WebSocketHandler) BattleController {
	return &battleController{
		uc: usecase.NewBattleUsecase(
			service.NewBattleService(
			database.NewQuestionRepository(sqlHandler),
			websocket.NewWebSocketRepository(webSocketHandler),
			service.NewContainerService(container.NewContainerRepository(containerHandler))),
		),
	}
}

// GET /
func (ctrl *battleController) Index(c Context) {
	c.HTML(200, "index.html", nil)
}

// GET /standard
func (ctrl *battleController) NewGame(c Context, mode string) {
	c.HTML(200, "new.html", mode)
}

// POST /standard
func (ctrl *battleController) Register(c Context) {
	ctrl.register(c)
	c.Redirect(302, "/battle/wait")
}

func (ctrl *battleController) register(c Context) {
	// Same as sessions.Defalut()
	session := c.MustGet(sessions.DefaultKey).(sessions.Session)

	roomName := c.PostForm("room")
	mode := c.PostForm("mode")
	if roomName == "" {
		ctrl.Error400(c, "あいことばを入力してください。")
		return
	}
	if !utf8string.NewString(roomName).IsASCII() {
		ctrl.Error400(c, "入力値は英数字でお願いします。")
		return
	}
	if !ctrl.uc.CanCreateRoom(roomName) {
		ctrl.Error400(c, "指定したあいことばが、現在既に利用されています。")
		return
	}
	session.Set("room", roomName)
	session.Set("mode", mode)
	if err := session.Save(); err != nil {
		log.Printf("failed at PostJoinBattle(): %s\n", err.Error())
		ctrl.Error500(c)
	}
}

// GET /standard/start
func (ctrl *battleController) Start(c Context) {
	c.HTML(200, "battle.html", nil)
}

/// GET /standard/wait
func (ctrl *battleController) Wait(c Context) {
	c.HTML(200, "wait.html", nil)
}

func (ctrl *battleController) WsBattle(c Context, conn model.Connection) {
	// Same as sessions.Defalut()
	session := c.MustGet(sessions.DefaultKey).(sessions.Session)
	roomName := session.Get("room").(string)
	u, err := uuid.NewRandom()
	if err != nil {
		c.HTML(500, "500.html", nil)
	}
	mode := session.Get("mode").(string)
	player := model.NewPlayer(fmt.Sprintf("%s_%s", roomName, u.String()), conn)
	if err := ctrl.uc.Start(ctrl.uc.SelectMode(mode), player.ID); err != nil {
		log.Println("invalid request!!")
		return
	}
	ctrl.uc.ParticipateIn(player, roomName)
	go ctrl.uc.Receiver(player)
	go ctrl.uc.Sender(player)
}

func (ctrl *battleController) Error500(c Context) {
	c.HTML(500, "500.html", nil)
}

func (ctrl *battleController) Error400(c Context, err string) {
	c.HTML(400, "400.html", err)
}

func (ctrl *battleController) WsWait(c Context, conn model.Connection) {
	// Same as sessions.Defalut()
	session := c.MustGet(sessions.DefaultKey).(sessions.Session)
	roomName := session.Get("room").(string)
	u, err := uuid.NewRandom()
	if err != nil {
		c.HTML(500, "500.html", nil)
	}
	player := model.NewPlayer(u.String(), conn)
	ctrl.uc.StartSignalSender(player, roomName)
}
