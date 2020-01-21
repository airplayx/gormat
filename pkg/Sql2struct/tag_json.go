package Sql2struct

import (
	"fmt"

	"github.com/xormplus/core"
)

func GetJsonTag(column *core.Column, isOmitempty bool) string {
	if !isOmitempty {
		return fmt.Sprintf(`json:"%s"`, column.Name)
	}
	return fmt.Sprintf(`json:"%s,omitempty"`, column.Name)
}
