package file

import (
	"reflect"
	"testing"
)

func TestReadLines(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "read line",
			args: args{path: "../../test_data/test.txt"},
			want: []string{"[pwd, ls]", "[echo 'hello']"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadLines(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLines() = %v, want %v", got, tt.want)
			}
		})
	}
}