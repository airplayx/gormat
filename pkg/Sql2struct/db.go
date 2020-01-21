package Sql2struct

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func InitDb(db *SourceMap) (err error) {
	switch strings.Title(db.Driver) {
	case "PostgreSQL":
		engine, err = xorm.NewPostgreSQL(fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", xorm.POSTGRESQL_DRIVER, db.Password, db.User, db.Host, db.Port, strings.Join(db.Db, "")))
	case "Sqlite3":
		engine, err = xorm.NewSqlite3(strings.Join(db.Db, ""))
	case "Mssql":
		engine, err = xorm.NewMSSQL(xorm.MSSQL_DRIVER, fmt.Sprintf("%s://%s:%s@%s/instance?database=%s", xorm.MSSQL_DRIVER, db.User, db.Password, db.Host, strings.Join(db.Db, "")))
	default:
		engine, err = xorm.NewMySQL(xorm.MYSQL_DRIVER, fmt.Sprintf("%s:%s@(%s:%s)/%s", db.User, db.Password, db.Host, db.Port, strings.Join(db.Db, "")))
	}
	if err != nil {
		return
	}
	engine.SetLogLevel(core.LOG_WARNING)
	return engine.Ping()
}

func DB() *xorm.Engine {
	return engine
}

func DBMetas(t []string, et []string, tryComplete bool) (tables []*core.Table, err error) {
	tmpTables, err := DB().Dialect().GetTables()
	if err != nil {
		return nil, fmt.Errorf("get tables faild: %s", err)
	}
	for _, v := range tmpTables {
		if len(t) > 0 {
			if !InStringSlice(v.Name, t) {
				continue
			}
		} else if len(et) > 0 {
			if InStringSlice(v.Name, et) {
				continue
			}
		}
		if err = loadTableInfo(v); err != nil {
			if tryComplete {
				log.Printf("load table:%s info faild: %s, strip", v.Name, err)
				continue
			}
			return nil, fmt.Errorf("load table:%s info faild: %s, please add it into exclude_tables, or set try_complete=true", v.Name, err)
		}
		tables = append(tables, v)
	}
	return
}

func loadTableInfo(table *core.Table) error {
	colSeq, cols, err := DB().Dialect().GetColumns(table.Name)
	if err != nil {
		return err
	}
	for _, name := range colSeq {
		table.AddColumn(cols[name])
	}
	indexes, err := DB().Dialect().GetIndexes(table.Name)
	if err != nil {
		return err
	}
	table.Indexes = indexes

	for _, index := range indexes {
		for _, name := range index.Cols {
			if col := table.GetColumn(name); col != nil {
				col.Indexes[index.Name] = index.Type
			} else {
				return fmt.Errorf("Unknown col %s in index %v of table %v, columns %v ", name, index.Name, table.Name, table.ColumnsSeq())
			}
		}
	}
	return nil
}

func sqlType2TypeString(st core.SQLType) string {
	t := core.SQLType2Type(st)
	s := t.String()
	if s == "[]uint8" {
		return "[]byte"
	}
	return s
}

func getGoImports(column *core.Column) map[string]string {
	imports := make(map[string]string)

	if sqlType2TypeString(column.SQLType) == "time.Time" {
		imports["time"] = "time"
	}

	return imports
}
