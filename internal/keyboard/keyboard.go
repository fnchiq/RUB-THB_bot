package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	buttonRow := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Binance"),
		tgbotapi.NewKeyboardButton("Bybit"),
	}

	keyboard := tgbotapi.NewReplyKeyboard(
		buttonRow,
	)

	return keyboard
}

func GetDefaultKeyboard() tgbotapi.ReplyKeyboardMarkup {
	buttonRow := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("RUB"),
		tgbotapi.NewKeyboardButton("THB"),
	}

	buttonRow1 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Узнать курс THB"),
		tgbotapi.NewKeyboardButton("Узнать курс KZT"),
	}

	buttonRow2 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Главное меню"),
	}

	keyboard := tgbotapi.NewReplyKeyboard(
		buttonRow,
		buttonRow1,
		buttonRow2,
	)

	return keyboard
}

func GetBankKeyboard() tgbotapi.ReplyKeyboardMarkup {
	buttonRow := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Сбербанк"),
		tgbotapi.NewKeyboardButton("Тинькофф"),
	}

	buttonRow1 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Главное меню"),
	}

	keyboard := tgbotapi.NewReplyKeyboard(
		buttonRow,
		buttonRow1,
	)

	return keyboard
}

func GetAmountKeyboard() tgbotapi.ReplyKeyboardMarkup {
	buttonRow := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Главное меню"),
	}

	keyboard := tgbotapi.NewReplyKeyboard(buttonRow)
	return keyboard
}
