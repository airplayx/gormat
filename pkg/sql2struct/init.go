package sql2struct

import (
	"fmt"
	"strings"
	//load some db pkg
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xormplus/xorm"
	l "github.com/xormplus/xorm/log"
)

//Engine ...
var Engine *xorm.Engine

//Init ...
func Init(db *SourceMap) (err error) {
	switch strings.Title(db.Driver) {
	case "PostgreSQL":
		Engine, err = xorm.NewPostgreSQL(fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", xorm.POSTGRESQL_DRIVER, db.Password, db.User, db.Host, db.Port, strings.Join(db.Db, "")))
	case "Sqlite3":
		Engine, err = xorm.NewSqlite3(strings.Join(db.Db, ""))
	case "Mssql":
		Engine, err = xorm.NewMSSQL(xorm.MSSQL_DRIVER, fmt.Sprintf("%s://%s:%s@%s/instance?database=%s", xorm.MSSQL_DRIVER, db.User, db.Password, db.Host, strings.Join(db.Db, "")))
	default:
		Engine, err = xorm.NewMySQL(xorm.MYSQL_DRIVER, fmt.Sprintf("%s:%s@(%s:%s)/%s", db.User, db.Password, db.Host, db.Port, strings.Join(db.Db, "")))
	}
	if err != nil {
		return
	}
	Engine.SetLogLevel(l.LOG_WARNING)
	return Engine.Ping()
}
