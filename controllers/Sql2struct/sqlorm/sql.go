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
	"go/ast"
	"strconv"
	"strings"
	"xorm.io/core"

	"gormat/controllers/Sql2struct/sqlorm/util"

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
	uniqIndces := map[string][]string{}
	for _, field := range ms.getStructFieds(ms.modelType) {
		switch t := field.Type.(type) {
		case *ast.Ident:
			tag, err := generateSqlTag(field)
			if err != nil {
				log.Printf("generateSqlTag [%s] failed:%v", t.Name, err)
			} else {
				tags = append(tags, fmt.Sprintf("%s %s", getColumnName(field), tag))
			}
		case *ast.SelectorExpr:
			tag, err := generateSqlTag(field)
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
		if idxName, ok := sqlSettings["UNIQUE_INDEX"]; ok {
			keys := uniqIndces[idxName]
			keys = append(keys, columnName)
			uniqIndces["uIdx_"+t.Name+"_"+sqlSettings["COLUMN"]] = keys
		}
	}

	var primaryKeyStr string
	if len(primaryKeys) > 0 {
		primaryKeyStr = fmt.Sprintf("PRIMARY KEY (%v)", strings.Join(primaryKeys, ", "))
	}

	indicesStrs := []string{}
	for idxName, keys := range indices {
		for _, v := range keys {
			indicesStrs = append(indicesStrs, fmt.Sprintf("KEY `%s` (%s)", idxName, v))
		}
	}

	uniqIndicesStrs := []string{}
	for idxName, keys := range uniqIndces {
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
	return fmt.Sprintf(`CREATE TABLE %v 
(
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

	if _, found := sqlSettings["NOT NULL"]; !found { // default: not null
		sqlSettings["NOT NULL"] = ""
	}

	additionalType := sqlSettings["NOT NULL"] + " " + sqlSettings["UNIQUE"]
	if value, ok := sqlSettings["DEFAULT"]; ok {
		additionalType = additionalType + " DEFAULT " + value
	}

	if value, ok := sqlSettings["COMMENT"]; ok {
		additionalType += " COMMENT " + value
	}

	if value, ok := sqlSettings["COLLATE"]; ok {
		additionalType += " COLLATE " + value
	}

	if value, ok := sqlSettings["TYPE"]; ok {
		sqlType = value
	}

	if sqlType == "" {
		var size = 128

		if value, ok := sqlSettings["SIZE"]; ok {
			size, _ = strconv.Atoi(value)
		}

		_, autoIncrease := sqlSettings["AUTO_INCREMENT"]
		if isPrimaryKey(field) {
			autoIncrease = true
		}

		sqlType, err = mysqlTag(field, size, autoIncrease)
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
	gormSettings := ParseTagSetting(tagStr)
	if _, ok := gormSettings["PRIMARY_KEY"]; ok {
		return true
	}

	if len(field.Names) > 0 && strings.ToUpper(field.Names[0].Name) == "ID" {
		return true
	}

	return false
}

func mysqlTag(field *ast.Field, size int, autoIncrease bool) (string, error) {
	typeName := ""
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
		if autoIncrease {
			return "int unsigned AUTO_INCREMENT", nil
		}
		return "int unsigned", nil
	case "int", "int8", "int16", "int32", "uint8", "uint16", "uint32", "uintptr":
		if autoIncrease {
			return "int AUTO_INCREMENT", nil
		}
		return "int", nil
	case "int64":
		if autoIncrease {
			return "bigint AUTO_INCREMENT", nil
		}
		return "bigint", nil
	case "uint64":
		if autoIncrease {
			return "bigint unsigned AUTO_INCREMENT", nil
		}
		return "bigint unsigned", nil
	case "float32", "float64":
		return "double", nil
	case "string", "NullString":
		if size > 0 && size < 65532 {
			return fmt.Sprintf("varchar(%d)", size), nil
		}
		return "longtext", nil
	case "Time":
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
		k := strings.TrimSpace(strings.ToUpper(v[0]))
		if len(v) == 2 {
			setting[k] = v[1]
		} else {
			setting[k] = k
		}
	}
	return setting
}
