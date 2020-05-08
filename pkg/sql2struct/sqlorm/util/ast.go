package util

import (
	log "github.com/liudanking/goutil/logutil"
	"go/ast"
	"strconv"
)

//GetFieldTag ...
func GetFieldTag(field *ast.Field, key string) *Tag {
	if field.Tag == nil {
		return &Tag{}
	}

	s, _ := strconv.Unquote(field.Tag.Value)
	tags, err := Parse(s)
	if err != nil {
		log.Warning("parse tag string:%s failed:%v", field.Tag.Value, err)
		return &Tag{}
	}
	tag, err := tags.Get(key)
	if err != nil {
		return &Tag{}
	}

	return tag
}

//GetFieldName ...
func GetFieldName(field *ast.Field) string {
	if len(field.Names) > 0 {
		return field.Names[0].Name
	}

	return ""
}
