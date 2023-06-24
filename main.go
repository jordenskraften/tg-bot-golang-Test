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
	if len(os.Args) >= 2 {

		token := os.Args[1] // Получаем токен из аргумента командной строки

		token_string := string(token)

		bot, err := tgbotapi.NewBotAPI(token_string)
		if err != nil {
			log.Fatal(err)
		}

		// Установка режима отладки
		bot.Debug = true

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
		final_message := "Обновленная цена: " + fmt.Sprintf("%.2f", number)
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
