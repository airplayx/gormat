package common

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/core"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func InitDb() (err error) {
	engine, err = xorm.NewEngine(Configs().Driver, Configs().Source)
	if err != nil {
		return
	}
	engine.SetLogLevel(core.LOG_WARNING)

	if err = engine.Ping(); err != nil {
		return fmt.Errorf("try to connect db faild: %s", err)
	}

	return nil
}

func DB() *xorm.Engine {
	return engine
}

func DBMetas(et []string, tryComplete bool) (tables []*core.Table, err error) {
	tmpTables, err := DB().Dialect().GetTables()
	if err != nil {
		return nil, fmt.Errorf("get tables faild: %s", err)
	}
	for _, v := range tmpTables {
		if len(et) > 0 {
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
				return fmt.Errorf("Unknown col %s in index %v of table %v, columns %v", name, index.Name, table.Name, table.ColumnsSeq())
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
