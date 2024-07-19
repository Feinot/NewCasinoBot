package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Users struct {
	user_id int
	balance int
	dep     int
	concl   int
	banned  bool
	refer   string
}

const (
	host     = "localhost"
	port     = 5432
	login    = "postgres"
	password = "1"
	dbname   = "postgres"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

}
func AddUserTodb(user_id, balance, dep, concl int, banned bool, refer string) error {
	fmt.Println(user_id)
	_, err = DB.Exec("insert into us(user_id, balance, banned,referer,deposit,conclusion) values ($1, $2, $3,$4,$5,$6)", user_id, balance, false, refer, 0, 0)
	if err != nil {
		fmt.Println(err)
		return (err)
	}
	return nil
}
func Balance(user_id int) int {

	rows, err := DB.Query("select balance from us where user_id=$1", user_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	p := Users{}
	for rows.Next() {

		err := rows.Scan(&p.balance)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return p.balance

	}
	return p.balance
}
func BalanceAdding(id int, dep int) {
	_, err = DB.Exec("update us set balance = balance + $1 where user_id = $2;", dep, id)
	if err != nil {
		fmt.Println(err)

	}

}
