/*Package configs ...
@Time : 2019/12/25 16:05
@Software: GoLand
@File : init
@Author : Bingo <airplayx@gmail.com>
*/
package configs

import (
	"testing"
)

func TestText(t *testing.T) {
	var test = map[string]interface{}{
		"language": "language",
	}
	for k, v := range test {
		lang := Text(k)
		if lang == "" || lang != v {
			t.Errorf("#%s: %s; want %s", k, lang, v)
		}
	}
}
