package cmd

import (
	"testing"
	"errors"
)

func Test_runCmd_instant(t *testing.T) {
	wait = 0
	tests := []struct {
		name string
		cf cmdFunc
		want int
	}{
		{
			name: "true,nil",
			cf: func () (bool, error) {
				return true, nil
			},
			want: 0,
		},
		{
			name: "false,nil",
			cf: func () (bool, error) {
				return false, nil
			},
			want: 1,
		},
		{
			name: "true,err",
			cf: func () (bool, error) {
				return true, errors.New("test")
			},
			want: 1,
		},
		{
			name: "false,err",
			cf: func () (bool, error) {
				return false, errors.New("test")
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runCmd(tt.cf); got != tt.want {
				t.Errorf("runCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
