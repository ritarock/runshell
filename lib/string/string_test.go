package string

import (
	"reflect"
	"testing"
)

func TestCreateCmd(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "createCmd",
			args: args{str: "[pwd, ls]"},
			want: [][]string{{"pwd, ls"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCmd(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
