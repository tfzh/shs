package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var SqlDB *sql.DB

func init() {
   var err error
   SqlDB, err = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/t_admin_db?parseTime=true")
   if err != nil {
      log.Fatal(err.Error())
  }
  err = SqlDB.Ping()
  if err != nil {
      log.Fatal(err.Error())
  }
}