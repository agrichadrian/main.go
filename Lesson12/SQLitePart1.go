package main

import(
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" 
)

func main(){
	db, err := sql.Open("sqlite", "animals.db")
	if err != nil {
		log.Fatal("Ошибка при подключении к базе:", err)
	}
	defer db.Close()

	fmt.Println("Подключиться к базе данных!")
} 