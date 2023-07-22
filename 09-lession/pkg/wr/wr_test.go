package wr

import (
	"bytes"
	"testing"
)

func TestCustomWriter(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		args   args
		wantWr string
	}{
		{
			name:   "Strings as input",
			args:   args{args: []interface{}{"Hello", "World"}},
			wantWr: "HelloWorld",
		},
		{
			name:   "Mixed types as input",
			args:   args{args: []interface{}{"Name", 25, true}},
			wantWr: "Name",
		},
		{
			name:   "Empty input",
			args:   args{args: []interface{}{}},
			wantWr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := &bytes.Buffer{}
			CustomWriter(wr, tt.args.args...)
			if gotWr := wr.String(); gotWr != tt.wantWr {
				t.Errorf("CustomWriter() = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}
