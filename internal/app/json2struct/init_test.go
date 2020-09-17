/*Package json2struct ...
@Time : 2020/1/3 14:28
@Software: GoLand
@File : init
@Author : Bingo <airplayx@gmail.com>
*/
package json2struct

import (
	"reflect"
	"testing"

	"fyne.io/fyne"
)

func TestScreen(t *testing.T) {
	tests := []struct {
		name string
		want *fyne.Container
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Screen(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Screen() = %v, want %v", got, tt.want)
			}
		})
	}
}
