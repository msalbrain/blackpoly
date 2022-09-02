package telegram

import (
	"exp/blackpoly/game"

	"fmt"
	"github.com/NicoNex/echotron/v3"
	"time"
	// "log"
	// "fmt"
	// "strconv"
	// "strings"
	// "github.com/NicoNex/echotron/v3"
)

func updatePropertyStatId(G *game.GameSession, b *bot) {
	for _, i := range G.PlayerList {
		f1, _ := b.SendMessage("there are:\n32 houses\n12 hotels", i.Id, nil)
		i.Msgid.PropertyStat = f1.Result.ID

	}
}

func updateBoardId(G *game.GameSession, b *bot) {
	for _, i := range G.PlayerList {
		// E := echotron.NewInputFileID("https://m.media-amazon.com/images/I/916dKhEsGyL._AC_SL1500_.jpg")
		E1 := echotron.NewInputFilePath("monopoly2.jpg")
		f2, _ := b.SendPhoto(E1, i.Id, nil)
		i.Msgid.Board = f2.Result.ID
	}
}

/*
	G *game.GameSession the usual
	b *bot the usual
	re bool is use to detemine if

	return
	in holds the player index with the hightest throw
	val holds the value thrown
*/
func allrolldice(G *game.GameSession, b *bot, re bool) (int, int) {
	var s []int
	for _, i := range G.PlayerList {
		if re == true {
			b.DeleteMessage(i.Id, i.Msgid.Dice1)
			b.DeleteMessage(i.Id, i.Msgid.Dice2)
		}
		a1, _ := b.SendDice(i.Id, echotron.Die, nil)
		i.Msgid.Dice1 = a1.Result.ID
		a2, _ := b.SendDice(i.Id, echotron.Die, nil)

		i.Msgid.Dice2 = a2.Result.ID

		fmt.Print("this is before s_append")
		s = append(s, a1.Result.Dice.Value+a2.Result.Dice.Value)
		fmt.Print("this is after s_append")
	}
	var in int
	val := 0
	for i, j := range s {
		if val < j {
			val = j
			in = i
		} else if val == j {
			in, val = allrolldice(G, b, true) // recursion baby
			break
		}
	}
	return in, val
}

func GameInit(G *game.GameSession, b *bot) {
	updatePropertyStatId(G, b)
	updateBoardId(G, b)

	for _, i := range G.PlayerList {
		action, _ := b.SendMessage("۰۰۰", i.Id, nil)
		play, _ := b.SendMessage("۰۰۰", i.Id, nil)
		info, _ := b.SendMessage("۰۰۰", i.Id, nil)
		// E := echotron.NewInputFileID("https://m.media-amazon.com/images/I/916dKhEsGyL._AC_SL1500_.jpg")

		i.Msgid.Action = action.Result.ID
		i.Msgid.Play = play.Result.ID
		i.Msgid.Info = info.Result.ID
	}

	in, val := allrolldice(G, b, false)

	time.Sleep(2200 * time.Millisecond)
	// set turn in game session
	G.Turn = G.PlayerList[in]

	for _, i := range G.PlayerList {
		fmt.Println(G.PlayerList[in].Id, i.Id)
		if int64(G.PlayerList[in].Id) == i.Id {
			E := echotron.NewMessageID(i.Id, i.Msgid.Action)
			b.EditMessageText(fmt.Sprintf("you roled a %d so you start first", val), E, nil)
			continue
		} else {
			E := echotron.NewMessageID(i.Id, i.Msgid.Action)
			b.EditMessageText(fmt.Sprintf("%s roled a %d so he/she/they start first", G.PlayerList[in].UserName, val), E, nil)
		}
	}

	for _, i := range G.PlayerList {
		if i == G.Turn {
			
			E := echotron.NewMessageID(i.Id, i.Msgid.Play)
			inline := InlineKeyboardCreate([]string{"roll dice 🎲", "make deal", "buy", "auction", "players 🫂", "property", "done", "😜"},
				[]string{"roll", "deal", "buy", "auction", "player", "property", "done", "😜"}, false)
			E1 := echotron.MessageTextOptions{ReplyMarkup: inline}
			play_id, err := b.EditMessageText("you own $1500", E, &E1)
			if err != nil{
				fmt.Println(err)
			}
			// fmt.Println(play_id.Result.ID)
			i.Msgid.Play = play_id.Result.ID

		} else {
			E := echotron.NewMessageID(i.Id, i.Msgid.Play)
			inline := InlineKeyboardCreate([]string{"players 🫂", "property"}, []string{"player", "property"}, false)
			E1 := echotron.MessageTextOptions{ReplyMarkup: inline}
			play_id, _ := b.EditMessageText("you own $1500", E, &E1)
			i.Msgid.Play = play_id.Result.ID

		}
	}
}
