package telegram

import (
	"exp/blackpoly"
	
	"log"

	"github.com/NicoNex/echotron/v3"
)

var PlayerCache = make(map[int]*blackpoly.Player)



const token = "5517968679:AAH4mhYlU_hlx-4XK1crVXlr2MlzysB_FsA"

func newBot(chatID int64) echotron.Bot {

	return &bot{
		chatID,
		echotron.NewAPI(token),
	}
}

func (b *bot) Update(update *echotron.Update) {
	// fmt.Println(update.CallbackQuery)
	Call := update.CallbackQuery != nil
	Msg := update.Message != nil
	// edMsg := update.EditedMessage != nil

	if Call {
		call := update.CallbackQuery

		if call.Data == "join" {
			join(b, update)
		} else if call.Data == "continue" {
			b.DeleteMessage(b.chatID, call.Message.ID)
		} else if call.Data == "random" {
			randomCall(b, call)
		} else if call.Data == "game_id" {
			joinById(b, call)
		}else if call.Data == "create"{
			createCall(b, call)
	 	}else if call.Data == "game_start"{
			startGame(b, call)
		}

	} else if Msg {
		msg := update.Message.Text
		if msg == "/start" {

			start(b, update)
		} else if msg == "/help" {
			help(b, update)
		} else if msg == "/play" {
			play(b, update)
		}else{
			randomText(b, update)
		}
	}

}

func TeleRun() {
	dsp := echotron.NewDispatcher(token, newBot)
	log.Println(dsp.Poll())
}
