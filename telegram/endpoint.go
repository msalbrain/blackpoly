package telegram

import (
	"exp/blackpoly/game"
	"fmt"
	"strconv"
	"strings"

	"github.com/NicoNex/echotron/v3"
)


type bot struct {
	chatID int64
	echotron.API
}

func start(b *bot, update *echotron.Update) {
	k := []echotron.KeyboardButton{{Text: "/play"}}
	Rm := echotron.ReplyKeyboardMarkup{Keyboard: [][]echotron.KeyboardButton{k}, ResizeKeyboard: true}
	E := echotron.MessageOptions{ReplyMarkup: Rm}
	b.SendMessage("welcome to blackpoly an online platform for multiplayer monopoly.", b.chatID, &E)
}

func help(b *bot, update *echotron.Update) {
	b.SendMessage("It is a very simple game. The rules are", b.chatID, nil)
}

func play(b *bot, update *echotron.Update) {
	b.DeleteMessage(b.chatID, update.Message.ID)
	I := InlineKeyboardCreate([]string{"🗣create", "👥join", "👀spectate"}, []string{"create", "join", "spectate"},
		true)
	P := echotron.MessageOptions{ReplyMarkup: I}
	b.SendMessage("It is a very simple game. The rules are simple. What do you want to do?", b.chatID, &P)
}

func randomText(b *bot, update *echotron.Update) {

	msg := update.Message.Text

	// a := strings.Split(msg, "gm-")
	if strings.Contains(msg, "gm-") {
		b.DeleteMessage(b.chatID, update.Message.ID)
		var gm_id bool
		num, err := strconv.Atoi(string(msg[3:]))
		if err != nil {
			gm_id = false
			b.SendMessage(fmt.Sprintf("the f*ck is %s🧐", msg), b.chatID, nil)
		} else {
			gm_id = true
		}

		if gm_id == true {
			//check if num is available to join
			// can := game.CheckgameId(num)
			gan, t := game.GetgameId(num)
			if t {
				joinIdLogic(b, gan)
			} else {
				b.SendMessage(fmt.Sprintf("seemed like game id %s is invalid or game session is full ", msg), b.chatID, nil)
			}
		}
	}
}

//--------------------------the callback$------------------------

func join(b *bot, update *echotron.Update) {
	call := update.CallbackQuery
	b.DeleteMessage(b.chatID, call.Message.ID)
	p := game.Playing(b.chatID)
	if p {
		I := InlineKeyboardCreate([]string{"continue"},
			[]string{"continue"}, true)

		p := echotron.MessageOptions{ReplyMarkup: I}
		b.SendMessage("you are curently playing in a game. would you like to continue", b.chatID, &p)
	} else {

		I := InlineKeyboardCreate([]string{"join random game", "join using id"},
			[]string{"random", "game_id"}, true)

		P := echotron.MessageOptions{ReplyMarkup: I}

		b.SendMessage("join By ...?", b.chatID, &P)
	}
}

func alreadyPlaying(b *bot) {
	I := InlineKeyboardCreate([]string{"Yes", "No"},
		[]string{"continue", "home"}, true)

	key := echotron.MessageOptions{ReplyMarkup: I}
	b.SendMessage("seems like you have a game in progress, do you want to continue?", b.chatID, &key)
}

func joinIdLogic(b *bot, ret_game *game.GameSession) {

	player_p := game.AddPlayer(b.chatID, ret_game)

	b.SendMessage(fmt.Sprintf("<b><pre>this is game session gm-%d</pre></b>", ret_game.Id), b.chatID, &echotron.MessageOptions{ParseMode: echotron.HTML})
	b.SendMessage(fmt.Sprintf("You are %s", player_p.UserName), player_p.Id, nil)

	for _, i := range ret_game.PlayerList { // broadcast a new user added to every body except player_p
		if i.Id == player_p.Id {
			continue
		}
		b.SendMessage(fmt.Sprintf("%s has joined session", player_p.UserName), i.Id, nil)
	}

	for _, i := range ret_game.PlayerList { // broadcast to player_p all previous players
		if i.Id == player_p.Id {
			continue
		}
		b.SendMessage(fmt.Sprintf("%s has joined session before", i.UserName), player_p.Id, nil)
	}

	if len(ret_game.PlayerList) == 6 { // used to set limit on the number of people that can join = 6
		ret_game.SetComplete()
		for _, i := range ret_game.PlayerList { // used to send full session to all players
			if i.Id == player_p.Id {
				continue
			}
			b.SendMessage("session already full ", i.Id, nil)
		}
	}

	// E := echotron.NewMessageID(ret_game.PlayerList[0].Id, ret_game.PlayerList[0].Last)
	// b.EditMessageText(fmt.Sprintf("%s has joined session", player_p.UserName), E,nil)
	// ret_game.PlayerList[0].Last += 1

}

func randomCall(b *bot, call *echotron.CallbackQuery) {
	b.DeleteMessage(b.chatID, call.Message.ID)
	p := game.Playing(b.chatID)
	switch p {
	case true:
		alreadyPlaying(b)
	case false:
		ret_game, t := game.FindrandomGame()
		if t {
			joinIdLogic(b, ret_game)
		} else if t == false {
			I := InlineKeyboardCreate([]string{"🗣create", "👀spectate"},
				[]string{"create", "spectate"}, true)

			key := echotron.MessageOptions{ReplyMarkup: I}
			b.SendMessage("could't find any open game. But you could create one or spectate one", b.chatID, &key)
		}
	}

}

func joinById(b *bot, call *echotron.CallbackQuery) {
	b.DeleteMessage(b.chatID, call.Message.ID)

	b.SendMessage("send in the game Id", b.chatID, nil)

	//implement a function to find a in wait mode
	// and create a player obj based of the game found and add the player to that
	//send back "type in the game id"
}

func startGame(b *bot, call *echotron.CallbackQuery) {
	b.DeleteMessage(b.chatID, call.Message.ID)
	P := game.Findplayer(b.chatID)
	if P == (&game.Player{}) {
		b.SendMessage("seem like you do not have any ongoing game", b.chatID, nil)
	} else {
		P.Game.SetComplete()
		P.Game.SetActive()
		// fmt.Println(P.UserName)
		
		// game.SetComp(P.Game)
		// game.SetAct(P.Game)

		GameInit(P.Game, b)
	}
}

func createCall(b *bot, call *echotron.CallbackQuery) {
	b.DeleteMessage(b.chatID, call.Message.ID)

	seep := game.Playing(b.chatID)
	if seep {
		print("i entere seep")
		alreadyPlaying(b)
	} else if seep == false {
		G, t := game.NewGame()
		if !t {
			f3, _ := b.SendMessage("Wow don't really know what happened, couldn't create a new game", b.chatID, nil)
			fmt.Printf("this is msgid after sending couldn't create  %d\n", f3.Result.ID)
		} else {
			val := strconv.Itoa(int(b.chatID))
			var gm_str string
			for i := 5; i < len(val); i++ {
				gm_str += string(val[i])
			}
			game_id_in := gm_str + string(strconv.Itoa(int(G.Start)))

			conv_int, _ := strconv.Atoi(game_id_in)
			G.Id = conv_int
			game.GameDb[conv_int] = G
			player_p := game.AddPlayer(b.chatID, G)

			b.SendMessage("your game id is:", b.chatID, nil)
			b.SendMessage(fmt.Sprintf("<b><pre>gm-%s</pre></b>", game_id_in), b.chatID, &echotron.MessageOptions{ParseMode: echotron.HTML})
			b.SendMessage(fmt.Sprintf("You are %s", player_p.UserName), b.chatID, nil)


			E1 := InlineKeyboardCreate([]string{"✅START✅"}, []string{"game_start"}, true)
			f4, _ := b.SendMessage("Start Game", b.chatID, &echotron.MessageOptions{ReplyMarkup: E1})
			fmt.Print("after start game", f4.Result.ID)

		}
	}
}
