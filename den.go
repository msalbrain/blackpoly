package blackpoly

var No_houses = 32
var No_hotels = 12

var actioncard = `
	34 ActionCard

	2 deal breaker
	3 just say no
	3 sly deal
	4 force deal
	3 debt collector
	3 it my birthday
	10 pass Go
	3 house
	3 hotel 
`

type Card interface{
	conn()
}

type Player struct{
	Name string
	UnDev []UnDevPropertyCard
	Dev []DevPropertyCard
}



type Mortgage struct{
	Price int
}

type UnDevPropertyCard struct {
	Name string
	PurchasePrice int
	Mortgage	

}

func (pc *UnDevPropertyCard) conn(){}

type TitleDeedCard struct{
	Id int
	Name string
	color string
	PurchasePrice int
	Rent int
	HouseRent1 int
	HouseRent2 int
	HouseRent3 int
	HouseRent4 int
	HotelRent int
	MortgagePrice int
	HouseCost int
	HotelCost int
}

type UtilityCard struct{
	Id int
	Name string
	PurchasePrice int
	Rent int
	MortgagePrice int
}

type ParkCard struct{
	Id int
	Name string
	PurchasePrice int
	Rent int
	MortgagePrice int
}

type DevPropertyCard struct {
	Name string
	Color string
	PurchasePrice int
	Mortgage	
}

func (pc *DevPropertyCard) conn(){}