package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Создание экземпляра бота с использованием токена

	token := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	// Установка режима отладки
	bot.Debug = false

	// Создание канала для получения обновлений
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Обработка входящих сообщений
	for update := range updates {
		// Проверка типа обновления (сообщение, команда и т.д.)
		if update.Message == nil {
			continue
		}

		// Получение текста сообщения
		messageText := update.Message.Text

		if update.Message.IsCommand() { // проверяем, является ли сообщение командой
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, напишите стоимость заказа в российских рублях.")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			// Обработка полученного сообщения
			response := processMessage(messageText)

			// Отправка ответа пользователю
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			_, err := bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// Обработка входящего сообщения
func processMessage(message string) string {
	//обработка сообщения пользователя и передача ему ответа
	number, err := findNumberInString(message)
	//найдем число в сообщении пользователя
	if err != nil {
		return "Введите число в вашем сообщении"
	} else {
		//если число найдено прогоним его по формулам
		number = SelectPriceFormula(number)
		final_message := "Обновленная цена: " + fmt.Sprintf("%.2f", number) + string('\u20BD')
		return final_message
	}
}

func findNumberInString(str string) (float64, error) {
	//находим число в сообщении пользователя
	re := regexp.MustCompile(`(\d+(\.\d+)?)`)

	match := re.FindStringSubmatch(str)
	if len(match) > 1 {
		numberStr := match[1]
		number, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			return 0, err
		}
		return number, nil
	}

	return 0, fmt.Errorf("число не найдено")
}
