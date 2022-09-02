package game

import (
	"fmt"
	"math/rand"
)

var (
	GameDb    = make(map[int]*GameSession)
	AllPlayer = make([]*Player, 0)
)




func AddPlayer(id int64, G *GameSession) *Player {
	p := &Player{}
	p.Id = id
	p.worth = 0
	p.Position = 0
	no := len(G.PlayerList) + 1
	
	
	p.UserName = fmt.Sprintf("player%v", TheTokens[(no-1)])
	// println("this is ",p.UserName, no)
	p.Token = TheTokens[(no - 1)]
	p.Game = G

	AllPlayer = append(AllPlayer, p)
	G.PlayerList = append(G.PlayerList, &p.PlayerBase)
	// G.Complete = fa

	return p
}

func Playing(id int64) bool {
	if len(AllPlayer) < 1 {
		return false
	}
	for _, player := range AllPlayer {
		if player.Id == id && player.Game.Id == 0 {
			return false
		} else if player.Id == id && player.Game.Id != 0 {
			return true
		}
	}

	return false
}

func Findplayer(id int64) *Player{

	for _, player := range AllPlayer {
		if player.Id == id{
			return player
		} 
	}
	return &Player{}
}

func FindrandomGame() (*GameSession, bool) {
	var gamelist []*GameSession
	for _, value := range GameDb {
		print(value)
		if value.Complete == false {
			print("it entered here and ")
			gamelist = append(gamelist, value)
		}
	}
	if len(gamelist) == 0 {
		return &GameSession{}, false
	}

	return gamelist[rand.Intn(len(gamelist))], true

}

func GetgameId(num int) (*GameSession, bool) {
	for index, val := range GameDb {
		if index == num {
			if val.Complete == true {
				return &GameSession{}, false
			} else {
				return val, true
			}
		} else if index != num {
			continue
		}
	}
	return &GameSession{}, false
}

func CheckgameId(num int) bool {

	for val := range GameDb {
		if int(val) == num {
			return true
		}
	}
	return false
}
