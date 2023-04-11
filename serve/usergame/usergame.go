package usergame

import (
	"github.com/LucaHhx/nano/session"
	"serve/game"
	"time"
)

type UserGame struct {
	User *session.Session
	Game *game.RmGame
}

func (ug *UserGame) ExchangePosition(tag1 game.Tag, tag2 game.Tag) error {
	err := ug.Game.ExchangePosition(tag1, tag2)
	if err != nil {
		return err
	}
	ug.User.Push("onNewGame", ug.Game)
	tags := ug.SameExtraction([]game.Tag{tag1, tag2}, true)
	if len(tags) == 0 {
		ug.Game.ExchangePosition(tag2, tag1)
		ug.User.Push("onNewGame", ug.Game)
	}
	return nil
}
func (ug *UserGame) SameExtraction(tags []game.Tag, isPush bool) []game.Tag {
	Tags := make([]game.Tag, 0)
	for _, tag := range tags {
		Tags = append(Tags, ug.Game.SameExtraction(tag, isPush)...)
	}
	if isPush {
		time.Sleep(time.Millisecond * 250)
		ug.User.Push("onNewGame", ug.Game)
	}
	if len(Tags) > 0 {
		ug.Drop(Tags, isPush)
	}
	return Tags
}
func (ug *UserGame) Drop(tags []game.Tag, isPush bool) {
	ug.Game.Drop(tags)
	if isPush {
		time.Sleep(time.Millisecond * 250)
		ug.User.Push("onNewGame", ug.Game)
	}
	time.Sleep(time.Millisecond * 250)
	ug.Game.AddTag(tags)
	if isPush {
		time.Sleep(time.Millisecond * 250)
		ug.User.Push("onNewGame", ug.Game)
	}
	ug.Ergodic(nil, isPush)
}
func (ug *UserGame) Ergodic(tags []game.Tag, isPush bool) {
	time.Sleep(time.Millisecond * 250)
	if len(tags) == 0 {
		for _, v := range ug.Game.Table {
			tags = append(tags, v...)
		}
	}
	ug.SameExtraction(tags, isPush)
}
