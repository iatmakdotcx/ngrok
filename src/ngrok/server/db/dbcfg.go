package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db, _ = sql.Open("mysql", "root:@/ngrokdb")
