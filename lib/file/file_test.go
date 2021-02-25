package file

import (
	"testing"
)

func TestRead(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "read line",
			args: args{path: "../../test_data/test.txt"},
			want: "[sleep 1, sleep 1, echo 'sleep 2']\n[sleep 1, sleep 1, sleep 1, echo 'sleep 3']\n[echo 'hello']\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Read(tt.args.path); got != tt.want {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
