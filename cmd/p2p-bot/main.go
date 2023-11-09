package main

import (
	"fmt"
	"log"
	"strconv"

	"p2p_bot/internal/binance"
	"p2p_bot/internal/bybit"
	"p2p_bot/internal/keyboard"
	"p2p_bot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const botToken = "токен"

func main() {
	var binanceAPI binance.Binance
	var bybitAPI bybit.Bybit
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	storage := make(map[int64]map[string]int)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			if _, ok := storage[update.Message.Chat.ID]; !ok {
				storage[update.Message.Chat.ID] = make(map[string]int)
			}
			switch update.Message.Text {
			case "Binance":
				storage[update.Message.Chat.ID]["Binance"] = 1
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите пункт:")
				msg.ReplyMarkup = keyboard.GetDefaultKeyboard()
				_, err := bot.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			case "Bybit":
				storage[update.Message.Chat.ID]["Bybit"] = 1
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите пункт:")
				msg.ReplyMarkup = keyboard.GetDefaultKeyboard()
				_, err := bot.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			case "RUB":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите банк:")
				msg.ReplyMarkup = keyboard.GetBankKeyboard()
				_, err := bot.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			case "Сбербанк":
				storage[update.Message.Chat.ID]["Sber"] = 1
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введети суммму от 1000 до 750000:")
				msg.ReplyMarkup = keyboard.GetAmountKeyboard()
				_, err := bot.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			case "Тинькофф":
				storage[update.Message.Chat.ID]["Tinkoff"] = 1
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введети суммму от 1000 до 750000:")
				msg.ReplyMarkup = keyboard.GetAmountKeyboard()
				_, err := bot.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			case "THB":
				storage[update.Message.Chat.ID]["THB"] = 1
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введети суммму от 1000 до 750000:")
				msg.ReplyMarkup = keyboard.GetAmountKeyboard()
				_, err := bot.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			case "Узнать курс THB":
				if storage[update.Message.Chat.ID]["Binance"] == 1 {
					var n1 []float64 = []float64{
						binanceAPI.GetPrice("RussianStandardBank", "RUB", "BUY", "10000", "null"),
						binanceAPI.GetPrice("RaiffeisenBank", "RUB", "BUY", "10000", "null"),
					}

					var n2 float64 = binanceAPI.GetPrice("BANK", "THB", "SELL", "4000", `"merchant"`)

					result := utils.ExchangeRateTHB(n1, n2)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
					_, err := bot.Send(msg)
					if err != nil {
						fmt.Println(err)
					}
				} else if storage[update.Message.Chat.ID]["Bybit"] == 1 {
					var n1 []float64 = []float64{
						bybitAPI.GetPrice("582", "RUB", "1", "10000", "false"),
						bybitAPI.GetPrice("581", "RUB", "1", "10000", "false"),
					}

					var n2 float64 = bybitAPI.GetPrice("14", "THB", "0", "4000", "false")

					result := utils.ExchangeRateTHB(n1, n2)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
					_, err := bot.Send(msg)
					if err != nil {
						fmt.Println(err)
					}
				}
			case "Узнать курс KZT":
				if storage[update.Message.Chat.ID]["Binance"] == 1 {
					var n1 []float64 = []float64{
						binanceAPI.GetPrice("RussianStandardBank", "RUB", "BUY", "10000", "null"),
						binanceAPI.GetPrice("RaiffeisenBank", "RUB", "BUY", "10000", "null"),
					}

					var n2 float64 = binanceAPI.GetPrice("KaspiBank", "KZT", "SELL", "45000", "null")

					result := utils.ExchangeRateKZT(n1, n2)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
					_, err := bot.Send(msg)
					if err != nil {
						fmt.Println(err)
					}
				} else if storage[update.Message.Chat.ID]["Bybit"] == 1 {
					var n1 []float64 = []float64{
						bybitAPI.GetPrice("582", "RUB", "1", "10000", "false"),
						bybitAPI.GetPrice("581", "RUB", "1", "10000", "false"),
					}

					var n2 float64 = bybitAPI.GetPrice("150", "KZT", "0", "45000", "false")

					result := utils.ExchangeRateKZT(n1, n2)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
					_, err := bot.Send(msg)
					if err != nil {
						fmt.Println(err)
					}
				}
			case "Главное меню":
				storage[update.Message.Chat.ID]["Binance"] = 0
				storage[update.Message.Chat.ID]["Bybit"] = 0
				storage[update.Message.Chat.ID]["Sber"] = 0
				storage[update.Message.Chat.ID]["Tinkoff"] = 0
				storage[update.Message.Chat.ID]["THB"] = 0
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите биржу")
				msg.ReplyMarkup = keyboard.GetMainMenuKeyboard()
				_, err := bot.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			default:
				if storage[update.Message.Chat.ID]["Sber"] == 0 && storage[update.Message.Chat.ID]["Tinkoff"] == 0 && storage[update.Message.Chat.ID]["THB"] == 0 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите пункт")
					msg.ReplyMarkup = keyboard.GetMainMenuKeyboard()
					_, err := bot.Send(msg)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					amount, err := strconv.Atoi(update.Message.Text)
					if err != nil || amount < 1000 || amount > 750000 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите сумму от 1000 до 750000")
						_, err := bot.Send(msg)
						if err != nil {
							fmt.Println(err)
						}
					} else if storage[update.Message.Chat.ID]["Binance"] == 1 && storage[update.Message.Chat.ID]["Sber"] == 1 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", binanceAPI.GetPrice("RussianStandardBank", "RUB", "BUY", update.Message.Text, "null")))
						_, err := bot.Send(msg)
						if err != nil {
							fmt.Println(err)
						}
					} else if storage[update.Message.Chat.ID]["Binance"] == 1 && storage[update.Message.Chat.ID]["Tinkoff"] == 1 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", binanceAPI.GetPrice("RaiffeisenBank", "RUB", "BUY", update.Message.Text, "null")))
						_, err := bot.Send(msg)
						if err != nil {
							fmt.Println(err)
						}
					} else if storage[update.Message.Chat.ID]["Binance"] == 1 && storage[update.Message.Chat.ID]["THB"] == 1 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", binanceAPI.GetPrice("BANK", "THB", "SELL", update.Message.Text, `"merchant"`)))
						_, err := bot.Send(msg)
						if err != nil {
							fmt.Println(err)
						}
					} else if storage[update.Message.Chat.ID]["Bybit"] == 1 && storage[update.Message.Chat.ID]["Sber"] == 1 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", bybitAPI.GetPrice("533", "RUB", "1", update.Message.Text, "false")))
						_, err := bot.Send(msg)
						if err != nil {
							fmt.Println(err)
						}

					} else if storage[update.Message.Chat.ID]["Bybit"] == 1 && storage[update.Message.Chat.ID]["Tinkoff"] == 1 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", bybitAPI.GetPrice("64", "RUB", "1", update.Message.Text, "false")))
						_, err := bot.Send(msg)
						if err != nil {
							fmt.Println(err)
						}
					} else if storage[update.Message.Chat.ID]["Bybit"] == 1 && storage[update.Message.Chat.ID]["THB"] == 1 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", bybitAPI.GetPrice("14", "THB", "0", update.Message.Text, "true")))
						_, err := bot.Send(msg)
						if err != nil {
							fmt.Println(err)
						}
					}
				}
			}
		}
	}
}
