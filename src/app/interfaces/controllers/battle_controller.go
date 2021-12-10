package controller

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"log"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/taise-hub/shellgame/src/app/interfaces/database"
	"github.com/taise-hub/shellgame/src/app/interfaces/websocket"
	"github.com/taise-hub/shellgame/src/app/interfaces/container"
	"github.com/taise-hub/shellgame/src/app/usecase"
	"github.com/taise-hub/shellgame/src/app/domain/service"
)


type BattleController interface {
	Index(Context)
	Battle(Context)
	New(Context)
	Start(Context)
	Wait(Context)
	WsBattle(Context, model.Connection)

	Error500(Context)
}

type battleController struct {
	battleService	 service.BattleService
	questionUsecase  usecase.QuestionUsecase
}

func NewBattleController(sqlHandler database.SqlHandler, containerHandler container.ContainerHandler, webSocketHandler websocket.WebSocketHandler) BattleController {
	return &battleController {
		battleService: service.NewBattleService(
			websocket.NewWebSocketRepository(webSocketHandler),
			service.NewContainerService(container.NewContainerRepository(containerHandler)),
		),
		questionUsecase: usecase.NewQuestionUsecase(
			database.NewQuestionRepository(sqlHandler)),
	}
}

// GET /
func (ctrl *battleController) Index(c Context) {
	c.HTML(200, "index.html", nil)
}

// GET /battle
func (ctrl *battleController) Battle(c Context) {
	c.HTML(200, "new.html", nil)
}

// POST /battle
func (ctrl *battleController) New(c Context) {
	// Same as sessions.Defalut()
	session := c.MustGet(sessions.DefaultKey).(sessions.Session)

	session.Set("player", c.PostForm("name"))
	session.Set("room", c.PostForm("room"))
	if err := session.Save(); err != nil {
		log.Printf("failed at PostJoinBattle(): %s\n", err.Error())
		ctrl.Error500(c)
	}
	c.Redirect(302, "/battle/wait")
}

// GET /battle/start
func (ctrl *battleController) Start(c Context) {
	c.HTML(200, "battle.html", nil)
}

/// GET /battle/wait
func (ctrl *battleController) Wait(c Context) {
	c.HTML(200, "wait.html", nil)
}

func (ctrl *battleController) WsBattle(c Context, conn model.Connection) {
	// Same as sessions.Defalut()
	session := c.MustGet(sessions.DefaultKey).(sessions.Session)
	roomName := session.Get("room").(string)
	playerName := session.Get("player").(string)
	player := model.NewPlayer(fmt.Sprintf("%s_%s",roomName, playerName), conn)
	ctrl.battleService.Start(player.ID)
	ctrl.battleService.ParticipateIn(player, roomName)
	go ctrl.battleService.Receiver(player)
}

func (ctrl *battleController) Error500(c Context) {
	c.HTML(500, "500.html", nil)
}

func (ctrl *battleController) Error400(c Context) {
	c.HTML(400, "400.html", nil)
}
