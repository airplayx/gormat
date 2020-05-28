package sql2struct

import (
	"fmt"
	"github.com/xormplus/xorm/schemas"
)

//GetJSONTag ...
func GetJSONTag(column *schemas.Column, isOmitempty bool) string {
	if !isOmitempty {
		return fmt.Sprintf(`json:"%s"`, column.Name)
	}
	return fmt.Sprintf(`json:"%s,omitempty"`, column.Name)
}
