package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	login    = "postgres"
	password = "1"
	dbname   = "postgres"
)

func AddUserTodb(user_id, balance, dep, concl int, banned bool, refer string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into us(user_id, balance, banned,referer,deposit,conclusion) values ($1, $2, $3,$4,$5,$6)", 3, 123, false, "asd", 0, 0)
	if err != nil {
		panic(err)
	}

}
