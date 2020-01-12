package util

import (
	"go/ast"
	"strconv"

	log "github.com/liudanking/goutil/logutil"
)

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

func GetFieldName(field *ast.Field) string {
	if len(field.Names) > 0 {
		return field.Names[0].Name
	}

	return ""
}
