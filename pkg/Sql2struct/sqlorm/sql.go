/*
@Time : 2020/1/11 12:20
@Software: GoLand
@File : sql
@Author : https://github.com/liudanking/gorm2sql
*/
package sqlorm

import (
	"errors"
	"fmt"
	"github.com/xormplus/core"
	"go/ast"
	"strconv"
	"strings"

	"gormat/pkg/Sql2struct/sqlorm/util"

	"github.com/pinzolo/casee"
	"log"
)

type SqlGenerator struct {
	structName string
	modelType  *ast.StructType
}

func NewSqlGenerator(typeSpec *ast.TypeSpec) (*SqlGenerator, error) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil, errors.New("typeSpec is not struct type")
	}

	return &SqlGenerator{
		structName: typeSpec.Name.Name,
		modelType:  structType,
	}, nil
}

func (ms *SqlGenerator) GetCreateTableSql(t *core.Table) (string, error) {
	var tags []string
	var primaryKeys []string
	indices := map[string][]string{}
	uniqInd := map[string][]string{}
	for _, field := range ms.getStructFieds(ms.modelType) {
		tag, err := generateSqlTag(field)
		switch t := field.Type.(type) {
		case *ast.Ident:
			if err != nil {
				log.Printf("generateSqlTag [%s] failed:%v", t.Name, err)
			} else {
				tags = append(tags, fmt.Sprintf("%s %s", getColumnName(field), tag))
			}
		case *ast.SelectorExpr:
			if err != nil {
				log.Printf("generateSqlTag [%s] failed:%v", t.Sel.Name, err)
			} else {
				tags = append(tags, fmt.Sprintf("%s %s", getColumnName(field), tag))
			}
		default:
			log.Printf("field %s not supported, ignore", util.GetFieldName(field))
		}

		columnName := getColumnName(field)
		if isPrimaryKey(field) {
			primaryKeys = append(primaryKeys, columnName)
		}

		sqlSettings := ParseTagSetting(util.GetFieldTag(field, "gorm").Name)
		if idxName, ok := sqlSettings["INDEX"]; ok {
			keys := indices[idxName]
			keys = append(keys, columnName)
			indices["idx_"+t.Name+"_"+sqlSettings["COLUMN"]] = keys
		}
		if idxName, ok := sqlSettings["UNIQUE"]; ok {
			keys := uniqInd[idxName]
			keys = append(keys, columnName)
			uniqInd["uIdx_"+t.Name+"_"+sqlSettings["COLUMN"]] = keys
		}
		if idxName, ok := sqlSettings["UNIQUE_INDEX"]; ok {
			keys := uniqInd[idxName]
			keys = append(keys, columnName)
			uniqInd["uIdx_"+t.Name+"_"+sqlSettings["COLUMN"]] = keys
		}
	}

	var primaryKeyStr string
	if len(primaryKeys) > 0 {
		primaryKeyStr = fmt.Sprintf("PRIMARY KEY (%v)", strings.Join(primaryKeys, ", "))
	}

	var indicesStrs []string
	for idxName, keys := range indices {
		for _, v := range keys {
			indicesStrs = append(indicesStrs, fmt.Sprintf("KEY `%s` (%s)", idxName, v))
		}
	}

	var uniqIndicesStrs []string
	for idxName, keys := range uniqInd {
		for _, v := range keys {
			uniqIndicesStrs = append(uniqIndicesStrs, fmt.Sprintf("UNIQUE KEY `%s` (%s)", idxName, v))
		}
	}
	if t.Charset == "" {
		t.Charset = "utf8mb4"
	}
	options := []string{
		"engine=" + t.StoreEngine,
		"DEFAULT charset=" + t.Charset,
		"COMMENT='" + t.Comment + "'",
	}
	return fmt.Sprintf(`CREATE TABLE %v (
  %v,
  %v
) %v;`,
		"`"+t.Name+"`",
		strings.Join(append(tags, append(indicesStrs, uniqIndicesStrs...)...), ",\n  "),
		primaryKeyStr,
		strings.Join(options, " ")), nil
}

func (ms *SqlGenerator) getStructFieds(node ast.Node) []*ast.Field {
	var fields []*ast.Field
	nodeType, ok := node.(*ast.StructType)
	if !ok {
		return nil
	}
	for _, field := range nodeType.Fields.List {
		if util.GetFieldTag(field, "gorm").Name == "-" {
			continue
		}

		switch t := field.Type.(type) {
		case *ast.Ident:
			if t.Obj != nil && t.Obj.Kind == ast.Typ {
				if typeSpec, ok := t.Obj.Decl.(*ast.TypeSpec); ok {
					fields = append(fields, ms.getStructFieds(typeSpec.Type)...)
				}
			} else {
				fields = append(fields, field)
			}
		case *ast.SelectorExpr:
			fields = append(fields, field)
		default:
			log.Printf("filed %s not supported, ignore", util.GetFieldName(field))
		}
	}

	return fields
}

func generateSqlTag(field *ast.Field) (string, error) {
	var sqlType string
	var err error

	sqlSettings := ParseTagSetting(util.GetFieldTag(field, "gorm").Name)
	additionalType := sqlSettings["NOT NULL"]

	if value, ok := sqlSettings["DEFAULT"]; ok {
		additionalType += " DEFAULT " + value
	}

	if value, ok := sqlSettings["COMMENT"]; ok {
		additionalType += " COMMENT " + value
	}

	if value, ok := sqlSettings["COLLATE"]; ok {
		additionalType += " COLLATE " + value
	}

	if value, ok := sqlSettings["TYPE"]; ok {
		sqlType = value
		if value == "timestamp" && sqlSettings["NOT NULL"] == "" {
			sqlType += " NULL"
		}
	}

	if sqlType == "" || sqlSettings["PRIMARY_KEY"] != "" {
		var size = 255

		if value, ok := sqlSettings["SIZE"]; ok {
			size, _ = strconv.Atoi(value)
		}

		_, autoIncrease := sqlSettings["AUTO_INCREMENT"]
		sqlType, err = mysqlTag(field, sqlType, size, autoIncrease)
		if err != nil {
			log.Printf("get mysql field tag failed:%v", err)
			return "", err
		}
	}

	if strings.TrimSpace(additionalType) == "" {
		return sqlType, nil
	} else {
		return fmt.Sprintf("%v %v", sqlType, additionalType), nil
	}

}

func getColumnName(field *ast.Field) string {
	tagStr := util.GetFieldTag(field, "gorm").Name
	gormSettings := ParseTagSetting(tagStr)
	if columnName, ok := gormSettings["COLUMN"]; ok {
		return "`" + columnName + "`"
	}

	if len(field.Names) > 0 {
		return fmt.Sprintf("`%s`", casee.ToSnakeCase(field.Names[0].Name))
	}

	return ""
}

func isPrimaryKey(field *ast.Field) bool {
	tagStr := util.GetFieldTag(field, "gorm").Name
	if _, ok := ParseTagSetting(tagStr)["PRIMARY_KEY"]; ok {
		return true
	}
	return false
}

func mysqlTag(field *ast.Field, sqlType string, size int, autoIncrease bool) (s string, err error) {
	defer func() {
		if sqlType != "" {
			s = sqlType + " " + strings.Join(strings.Split(s, " ")[1:], " ")
		}
		if autoIncrease {
			s += " AUTO_INCREMENT"
		}
	}()
	var typeName string
	switch t := field.Type.(type) {
	case *ast.Ident:
		typeName = t.Name
	case *ast.SelectorExpr:
		typeName = t.Sel.Name
	default:
		return "", errors.New(fmt.Sprintf("field %s not supported", util.GetFieldName(field)))
	}
	switch typeName {
	case "bool":
		return "boolean", nil
	case "uint":
		return "int unsigned", nil
	case "int", "int8", "int16", "int32", "uint8", "uint16", "uint32", "uintptr":
		return "int", nil
	case "int64":
		return "bigint", nil
	case "uint64":
		return "bigint unsigned", nil
	case "float32", "float64":
		return "double", nil
	case "string", "NullString":
		if size > 0 && size < 65532 {
			return fmt.Sprintf("varchar(%d)", size), nil
		}
		return "longtext", nil
	case "time":
		return "datetime", nil
	default:
		return "", errors.New(fmt.Sprintf("type %s not supported", typeName))

	}
}

func ParseTagSetting(str string) map[string]string {
	tags := strings.Split(str, ";")
	setting := map[string]string{}
	for _, value := range tags {
		v := strings.Split(value, ":")
		if len(v) == 0 {
			continue
		}
		k := strings.TrimSpace(strings.ToUpper(v[0]))
		switch len(v) {
		case 1:
			setting[k] = k
		default:
			setting[k] = strings.Join(v[1:], ":")
		}
	}
	return setting
}
