package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ZeroTheorem/my_wife_job_bot/db"
	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("sqlite", "file:mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	q := db.New(conn)
	err = q.CreateRow(ctx, db.CreateRowParams{
		Name:  "Alena",
		Val:   48880,
		Month: int64(time.Now().Month()),
		Year:  int64(time.Now().Year()),
	})

	if err != nil {
		log.Fatal(err)
	}

	val, err := q.GetAvg(ctx, db.GetAvgParams{
		Name:  "Alena",
		Month: int64(time.Now().Month()),
		Year:  int64(time.Now().Year()),
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}
