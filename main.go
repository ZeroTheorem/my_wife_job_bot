package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ZeroTheorem/my_wife_job_bot/db"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
	_ "modernc.org/sqlite"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	b, err := tele.NewBot(tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open("sqlite", "file:mydb.db")

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	q := db.New(conn)

	b.Handle("/add", func(c tele.Context) error {
		vals := strings.Split(c.Message().Text, " ")
		if len(vals) != 3 {
			return c.Send("Необходимо ввести все значение в формате\n\n/add <Имя> <Значение>")
		}
		nameLower := strings.ToLower(vals[1])
		if nameLower != "даша" && nameLower != "алена" {
			return c.Send(
				"Допустимые имена:\n\nДаша\nАлена\n\nможешь писать их с маленькой или большой буквы - это не важно, но другие имена не допустимы!")
		}

		intValue, err := strconv.ParseInt(vals[2], 10, 64)
		if err != nil {
			return c.Send(
				fmt.Sprintf("%v -- второе значение после /add должно быть числом", vals[2]))
		}
		err = q.CreateRow(ctx, db.CreateRowParams{
			Name:  strings.ToLower(vals[1]),
			Val:   intValue,
			Month: int64(time.Now().Month()),
			Year:  int64(time.Now().Year()),
		})
		if err != nil {
			return c.Send(
				fmt.Sprintf("Ууупс... что-то пошло не так: %v", err))

		}
		return c.Send("Запись была успешно добавлена 😉")
	})
	b.Handle("/deletelast", func(c tele.Context) error {
		lastVal, err := q.DeleteLastRow(ctx)
		if err != nil {
			return c.Send(
				fmt.Sprintf("Ууупс... что-то пошло не так: %v", err))
		}
		return c.Send(
			fmt.Sprintf("Запись:\n\n%v: %v\n\nбыла успешно удалена 😉", lastVal.Name, lastVal.Val))
	})
	b.Start()

}
