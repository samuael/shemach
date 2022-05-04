package main

// import (
// 	"database/sql"
// 	"os"
// 	"sync"

// 	// "github.com/qustavo/dotsql"
// 	"github.com/gchaincl/dotsql"
// 	"github.com/samuael/agri-net/agri-net-backend/pkg/storage/sql_db"
// 	"github.com/subosito/gotenv"
// )

// func init() {
// 	gotenv.Load("../.env")
// }

// var once sync.Once

// var conn *sql.DB
// var connError error

// func main() {
// 	once.Do(func() {
// 		conn, connError = sql_db.NewStorage(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
// 		if connError != nil {
// 			println(connError.Error())
// 			os.Exit(1)
// 		}
// 	})

// 	if dot, err := dotsql.LoadFromFile("../../../pkg/constants/query/ntables.sql"); err == nil {
// 		if _, err = dot.Exec(conn, "create-admin"); err != nil {
// 			println(err.Error(), os.Getenv("ERROR_CREATING_TABLE_ADMIN"))
// 			os.Exit(1)
// 		}
// 		if _, err = dot.Exec(conn, "create-categories"); err != nil {
// 			println(err.Error())
// 			println(os.Getenv("ERROR_CREATING_TABLE_CATEGORIES"))
// 			os.Exit(1)
// 		}
// 		if _, err = dot.Exec(conn, "create-round"); err != nil {
// 			println(err.Error())
// 			println(os.Getenv("ERROR_CREATING_TABLE_ROUND"))
// 			os.Exit(1)
// 		}
// 		if _, err = dot.Exec(conn, "birth_date"); err != nil {
// 			println(err.Error())
// 			println(os.Getenv("ERROR_CREATING_TABLE_ROUND"))
// 			os.Exit(1)
// 		}
// 		if _, err = dot.Exec(conn, "create-special-case"); err != nil {
// 			println(err.Error())
// 			println(os.Getenv("ERROR_CREATING_TABLE_Special_case"))
// 			os.Exit(1)
// 		}
// 		if _, err = dot.Exec(conn, "create-addresses"); err != nil {
// 			println(err.Error())
// 			println(os.Getenv("ERROR_CREATING_TABLE_ADDRESSES"))
// 			os.Exit(1)
// 		}
// 		if _, err = dot.Exec(conn, "create-student"); err != nil {
// 			println(os.Getenv("ERROR_CREATING_TABLE_STUDENT"))
// 			println(err.Error())
// 			os.Exit(1)
// 		}
// 		if _, err = dot.Exec(conn, "pay-in"); err != nil {
// 			println(err.Error())
// 			println(os.Getenv("ERROR_CREATING_TABLE_PAYIN"))
// 			os.Exit(1)
// 		}
// 		if _, err = dot.Exec(conn, "create-payout"); err != nil {
// 			println(err.Error())
// 			println(os.Getenv("ERROR_CREATING_TABLE_PAYOUT"))
// 			os.Exit(1)
// 		}
// 	}
// 	println("\nDatabase Tables succesfuly Initialized ... \n")
// 	defer conn.Close()
// }
