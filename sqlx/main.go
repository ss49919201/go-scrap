package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func dsn() string {
	return (&mysql.Config{
		User:      "user",
		Passwd:    "password",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "mydb-for-sqlx",
		ParseTime: true,
	}).FormatDSN()
}

func a(v ...any) {
	// fmt.Println(v == nil)
	// fmt.Println(v)
	// fmt.Println(reflect.ValueOf(v).IsNil())
	// fmt.Println(v == ([]any)(nil))
}

func main() {
	var v []any
	a(v...)
	v = []any{}
	a(v...)
	db := check(sql.Open("mysql", dsn()))
	// initiate zerolog
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zlogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	// prepare logger
	loggerOptions := []sqldblogger.Option{
		sqldblogger.WithSQLQueryFieldname("sql"),
		sqldblogger.WithWrapResult(false),
		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
	}
	// wrap *sql.DB to transparent logger
	db = sqldblogger.OpenDriver(dsn(), db.Driver(), zerologadapter.New(zlogger), loggerOptions...)
	// pass it sqlx
	sqlxDB := sqlx.NewDb(db, "mysql")

	query := "SELECT ''"
	var args []any
	// Query()の結果のsql.Rowをsqlx.Rowsのフィールドに内包して返す
	check(sqlxDB.Queryx(query, args...))
	fmt.Println(check3(sqlx.Named("SELECT '::' FROM table WHERE tbl = 'tbl'", map[string]any{"bl": "table"})))

	rows := check(sqlxDB.NamedQuery("SELECT ':tbl' FROM tbl WHERE id IN (:ids)", map[string]any{"tbl": "table", "ids": []int{1, 2}}))
	// iterate over each row
	var itrErr error
	for rows.Next() {
		var v string
		itrErr = rows.Scan(&v)
		// pp.Println(v)
	}
	// check the error from rows
	check(0, itrErr)
	check(0, rows.Err())
}

func check[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func check3[T, T2 any](v T, v2 T2, err error) (T, T2) {
	if err != nil {
		panic(err)
	}
	return v, v2
}
