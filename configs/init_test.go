package configs

import (
	"github.com/buger/jsonparser"
	"testing"
)

func TestText(t *testing.T) {
	JSON, _ = jsonparser.Set(JSON, []byte("\"cn\""), "const", "language")
	type args struct {
		key     string
		replace []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1 language",
			args: args{
				key: "light",
			},
			want: "白色",
		},
		{
			name: "test2 language",
			args: args{
				key: "ddd",
			},
			want: "ddd",
		},
		{
			name: "test3 language",
			args: args{
				key: "test %s and more %d",
				replace: []interface{}{
					"one",
					3,
				},
			},
			want: "test one and more 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Text(tt.args.key, tt.args.replace...); got != tt.want {
				t.Errorf("Text() = %v, want %v", got, tt.want)
			}
		})
	}
}
