package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //_ na frente é um import q será usado de forma implicita
)

func main() {

	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta")

	rows, err := db.Query("SELECT * FROM usuarios")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println(rows)
}
