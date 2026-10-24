package main

import (
	"context"
	"fmt"
	"pg_practise/feature_postgres/simple_connection"
	"pg_practise/feature_postgres/simple_sql"
	"time"
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

	fmt.Println("Таблица создана")

	if err := simple_sql.InsertRow(ctx, conn, "test11", "test2", false, time.Now()); err != nil {
		panic(err)
	}

	fmt.Println("Строка добавлена")

	if err := simple_sql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}

	fmt.Println("Строка обновлена")
}
