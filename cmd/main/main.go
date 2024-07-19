package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/Feinot/NewCasinoBot/cmd/main/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Users struct {
	user_id int
	balance int
	dep     int
	concl   int
	banned  bool
	refer   string
}
type MyPhotoData struct {
	FileID string
}

func (p MyPhotoData) NeedsUpload() bool {
	return false
}

func (p MyPhotoData) UploadData() (string, io.Reader, error) {
	return "", nil, nil
}

func (p MyPhotoData) SendData() string {
	return p.FileID
}

var (
	adminChat   = -4226936363
	CollinkChat = -1002217455965
)

type ChatConfigWithUser struct {
	ChatID             int64
	SuperGroupUsername string
	UserID             int
}

var (
	telegramBotToken string
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {

	bot, err := tgbotapi.NewBotAPI("7446661096:AAFn11mrTdkpfJqQeciyl5c97LjfRg35HBg")
	if err != nil {
		log.Panic(err)
	}
	var member tgbotapi.ChatConfigWithUser
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates := bot.GetUpdatesChan(u)
	update, err := bot.GetUpdates(u)
	fmt.Printf("%+v\n", update)

	for update := range updates {

		reply := ""
		if update.Message == nil {
			continue
		}
		if update.Message.Photo != nil {

			if update.Message.Photo != nil {
				if update.Message.Photo != nil {
					// Получаем последнюю фотографию в списке
					photo := (update.Message.Photo)[len(update.Message.Photo)-1]

					// Создаем объект MyPhotoData с информацией о файле фотографии
					photoData := MyPhotoData{
						FileID: photo.FileID,
					}

					// Создаем объект tgbotapi.PhotoConfig с объектом MyPhotoData
					msg := tgbotapi.NewPhoto(int64(adminChat), photoData)
					msg.Caption = strconv.FormatInt(update.Message.From.ID, 10)

					// Отправляем изображение пользователю
					_, err = bot.Send(msg)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "start":
			member.ChatID = int64(CollinkChat)
			member.UserID = update.Message.From.ID

			_, err := bot.GetChatMember(tgbotapi.GetChatMemberConfig{member})
			if err != nil {
				reply = "sub pls"
				fmt.Println(err)
			} else {
				fmt.Println(int(update.Message.From.ID))
				addUserTODb(int(update.Message.From.ID), 0, 0, 0, false, "")
				reply = Menu()

			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

			bot.Send(msg)
		case "subs":
			member.ChatID = int64(CollinkChat)
			member.UserID = update.Message.From.ID
			//user_id int, balance int, dep int, concl int, banned bool, refer string
			_, err := bot.GetChatMember(tgbotapi.GetChatMemberConfig{member})
			if err != nil {
				reply = "sub pls"
				fmt.Println(err)
			} else {
				fmt.Println(int(update.Message.From.ID), "asd")
				addUserTODb(int(update.Message.From.ID), 0, 0, 0, false, "")

				reply = Menu()

			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

			bot.Send(msg)
		case "menu":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, Menu())

			bot.Send(msg)
		case "games":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, Games())

			bot.Send(msg)
		case "deposit":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, replenishmentBalans())
			bot.Send(msg)
		case "dart":
			bot.Send(tgbotapi.NewDiceWithEmoji(update.Message.Chat.ID, "🎰"))

		case "Balance":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, strconv.Itoa(Balns(int(update.Message.Chat.ID)))))
		case "adding":
			if update.Message.Chat.ID == int64(adminChat) {
				result := strings.TrimPrefix(update.Message.Text, "/adding ")
				b, err := strconv.Atoi(result)
				//fmt.Println(update.Message.From.ID, update.Message.Chat.ID)
				if err != nil {
					fmt.Println(err)
				}
				BalanceAdding(int(update.Message.From.ID), b)
				//int(update.Message.Chat.ID)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, strconv.Itoa(Balns(int(update.Message.From.ID)))))

			}

		}

		//st := tgbotapi.NewStickerSetConfig(update.Message.Chat.ID, sticer.Emoji)

		//tgbotapi.NewMessage(update.Message.Chat.ID, "🎯")
		//tgbotapi.NewDiceWithEmoji(update.Message.Chat.ID, "🎯")

	}

}

func Menu() string {
	//(tg chanel)https://t.me/+a6l0_viPDpIzZGI6
	return "/games /deposit"

}
func Games() string {
	return "/cube /dart /futbal /bascketbal"

}

func addUserTODb(user_id int, balance int, dep int, concl int, banned bool, refer string) error {
	return database.AddUserTodb(user_id, balance, dep, concl, banned, refer)
}
func Balns(id int) int {
	return database.Balance(id)
}
func replenishmentBalans() string {
	return "replenishment CARDNUMBER and attach a screenshot of the deposit, if whona you cancel dep write /menu"
}
func Referals() {}
func URLGenerate(id uint64) string {
	url := fmt.Sprint("https://t.me/FJcasino_bot?start=", id)
	return url
}
func TwentiWan() {
	return
}
func BalanceAdding(id int, dep int) {
	database.BalanceAdding(id, dep)

}
