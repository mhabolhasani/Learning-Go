package main

import (
	"fmt"
	"sync"
)

type Player struct {
	name    string
	area    Map
	channel chan string
}

type Map struct {
	players    []Player
	Id         int
	mapChannel chan string
}

type Game struct {
	players []Player
	areas   []Map
}

func NewGame(mapIds []int) (*Game, error) {
	game := new(Game)
	Err := error(nil)
	for _, id := range mapIds {
		for _, m := range game.areas {
			if m.Id == id {
				Err = nil
				continue
			} else {
				_ = append(game.areas, Map{Id: id, mapChannel: make(chan string)})
			}
		}
	}
	return game, Err
}

func (g *Game) ConnectPlayer(name string) error {
	err := error(nil)
	for _, player := range g.players {
		if player.name == name {

		}
	}
	return err
}

func (g *Game) SwitchPlayerMap(name string, mapId int) error {
	return nil
}

func (g *Game) GetPlayer(name string) (*Player, error) {
	return nil, nil
}

func (g *Game) GetMap(mapId int) (*Map, error) {
	return nil, nil
}

func (m *Map) FanOutMessages() {
}

func (p *Player) GetChannel() <-chan string {
	return nil
}

func (p *Player) SendMessage(msg string) error {
	return nil
}

func (p *Player) GetName() string {
	return ""
}
