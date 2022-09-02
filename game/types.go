package game

import (
	// "encoding/json"
	"fmt"
	"time"
	// "log"
)

type PlayerToken string

var TheTokens = []string{"🔋", "🏮", "⚾️", "🦕", "⚖️", "🛵", "🐈", "🔗"}

type GameSession struct {
	Id         int
	Start      int64
	End        int
	Active     bool
	Complete   bool
	PlayerList []*PlayerBase
	Turn       *PlayerBase
	Winner     *PlayerBase
	Looser     *PlayerBase
	Asset      map[int]*BoardCard
}

// type Property struct {
// 	Property interface{}
// }

func NewGame() (*GameSession, bool) {
	G := &GameSession{Start: time.Now().Unix()}
	return G, true
}

func (G *GameSession) SetActive() {
	G.Active = true
}

func (G *GameSession) SetComplete() {
	fmt.Println(G.Complete)
	G.Complete = true
}

// func SetComp(G *GameSession){
// 	fmt.Println("this is the part that has to do with completing")
// 	G.Complete = true
// }

// func SetAct(G *GameSession){
// 	fmt.Println("this is the part that has to do with completing")
// 	G.Active = true
// }

type PlayerBase struct {
	Id       int64
	UserName string
	UniqueId string
	Last     int

	/*this struct hold ths msgid in a telegram chat,...
	PropertStat hold the id for displaying number of houses and hotel,...
	Board holds th id for dislayin board image,...
	Action holds id for thing  happening in game e.g player 1 rolled a 10 and landed on wishing well,...
	Play contains id for inline keyboard actions of the player e.g rolldice, view cards,  make deal,...
	Info this is used for any extra info like display of cards
	Dice1 and Dice2 are self explanatory like the rest,...*/
	Msgid MsgIds
}

type Player struct {
	PlayerBase
	Position int
	worth    int
	Token    string `default:""`
	Game     *GameSession
	Asset    map[int]interface{}
}

type CardBase struct {
	Id int `json:"id"`
}

type TitleDeedCard struct {
	CardBase
	Name  string `json:"name"`
	Color string `json:"color"`

	PurchasePrice int `json:"price"`
	Rent          int `json:"rent"`
	RentCom       int `json:"rent_with_set"`
	HouseRent1    int `json:"h1rent"`
	HouseRent2    int `json:"h2rent"`
	HouseRent3    int `json:"h3rent"`
	HouseRent4    int `json:"h4rent"`
	HotelRent     int `json:"hotel_rent"`
	MortgagePrice int `json:"mortgage_price"`

	HouseCost int `json:"house_cost"`
	HotelCost int `json:"hotel_cost"`
}

type AllCards struct {
	*TitleDeedCard
	*UtilityCard
	*TrainCard
	*CommunityCard
	*ChanceCard
	*Jailcard
	*FreePark
	*GoToJail
}

type UtilityCard struct {
	CardBase
	Utility       string
	PurchasePrice int `json:"price"`
	MortgagePrice int `json:"mortgage_price"`
}

type TrainCard struct {
	CardBase

	Train         string `json:"train"`
	PurchasePrice int    `json:"price"`
	Rent          int
	MortgagePrice int `json:"mortgage_price"`
}

type CommunityCard struct {
	CardBase
	Community bool
}

type ChanceCard struct {
	CardBase
	Chance bool
}

type Jailcard struct {
	CardBase
	Jail bool
}

type FreePark struct {
	CardBase
	park bool
}

type GoToJail struct {
	CardBase
}

type TaxCard struct {
	Train   TrainCard   `json:"train"`
	Utility UtilityCard `json:"utility"`
	CardBase
	Name string
	Tax  int
}

type BoardCard struct {
	Position  int
	Earned    int
	Landed    int
	No_houses int
	No_hotels int
	asset     interface{}
	owned     PlayerBase
}

type MsgIds struct {
	PropertyStat int
	Board        int
	Play         int
	Action       int
	Info         int
	Dice1        int
	Dice2        int
}
