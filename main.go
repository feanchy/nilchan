package main

import (
	"context"
	"fmt"
	"postgres/feature/postgres/simple_connection"
	"postgres/feature/postgres/simple_sql"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 1003 {

			task.Title = "покормить кошку"
			task.Description = "отсыпать кошке 30 грамм корма"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			err := simple_sql.UpdateTask(ctx, conn, task)
			if err != nil {
				panic(err)
			}
			break
		}

	}
	pp.Println(tasks)

	fmt.Println("succes!")
}
