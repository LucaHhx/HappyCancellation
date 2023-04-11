package serve

import (
	"github.com/LucaHhx/nano/component"
	"github.com/LucaHhx/nano/session"
	"serve/game"
	"serve/usergame"
)

type Serve struct {
	component.Base
}

func NewServe() *Serve {
	return &Serve{}
}

var UserGames = make(map[int64]*usergame.UserGame)

type UserRequest struct {
	Name    string `json:"name" from:"name"`
	Id      int64  `json:"id" from:"id"`
	Refresh bool   `json:"refresh" from:"refresh"`
}

func (g *Serve) NewGame(s *session.Session, msg *UserRequest) error {
	if UserGames[msg.Id] == nil || msg.Refresh {
		UserGames[msg.Id] = &usergame.UserGame{
			Game: game.NewGame(),
			User: s,
		}
	}
	UserGames[msg.Id].Ergodic(nil, false)
	s.Push("onNewGame", UserGames[msg.Id].Game)
	return nil
}

type TagRequest struct {
	User UserRequest `json:"user" from:"user"`
	Tag1 game.Tag    `json:"tag1" from:"tag1"`
	Tag2 game.Tag    `json:"tag2" from:"tag2"`
}

func (g *Serve) PositionExchange(s *session.Session, msg *TagRequest) error {
	ug := UserGames[msg.User.Id]
	ug.User = s
	err := ug.ExchangePosition(msg.Tag1, msg.Tag2)
	if err != nil {
		s.Push("onError", err.Error())
	}
	return nil
}
