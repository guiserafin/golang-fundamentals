package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o mysql
)

// Conectar abre a conexão com o db
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		fmt.Println(erro)
		return nil, erro
	}

	if erro := db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil

}
