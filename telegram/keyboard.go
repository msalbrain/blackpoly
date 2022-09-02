package telegram

import (
	"fmt"

	"github.com/NicoNex/echotron/v3"
)

var (
	print = fmt.Println
)


func InlineKeyboardCreate(text []string, callbacktext []string, step bool) echotron.InlineKeyboardMarkup {
	l := len(text)

	if step == true {
		playerKeyboard := make([][]echotron.InlineKeyboardButton, 0)
		// oddkeyboard := make([]echotron.InlineKeyboardButton,1)
		for i := 0; i < l; i++ {
			
			// oddkeyboard[0] = echotron.InlineKeyboardButton{Text: text[i], CallbackData: callbacktext[i]}

			playerKeyboard = append(playerKeyboard, []echotron.InlineKeyboardButton{{Text: text[i], CallbackData: callbacktext[i]}})

		}
		// print(playerKeyboard)
		I := echotron.InlineKeyboardMarkup{InlineKeyboard: playerKeyboard}
		return I
	}

	if step == false {
		playerKeyboard := make([][]echotron.InlineKeyboardButton, 0)

		if l%2 == 0 {
			v := 0
			for i := 0; i < l; i++ {
				if v == 0{
					fmt.Println(i, []echotron.InlineKeyboardButton{{Text: text[i], CallbackData: callbacktext[i]}})
					playerKeyboard = append(playerKeyboard, []echotron.InlineKeyboardButton{{Text: text[v], CallbackData: callbacktext[v]}, {Text: text[v+1], CallbackData: callbacktext[v+1]}})
					v += 1
				}else if v == 1{
					len_play := len(playerKeyboard)
					playerKeyboard[len_play - 1] = append(playerKeyboard[len_play-1], echotron.InlineKeyboardButton{Text: text[i], CallbackData: callbacktext[i]})
					v -= 1
				}else{
					break
				}
				
			}
			fmt.Println(playerKeyboard)
			In := echotron.InlineKeyboardMarkup{InlineKeyboard: playerKeyboard}
			return In
		}

	}
	return echotron.InlineKeyboardMarkup{}
}
