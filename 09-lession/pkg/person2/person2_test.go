package person2

import (
	"reflect"
	"testing"
)

func TestGetOlder(t *testing.T) {
	type args struct {
		p []Person
	}
	tests := []struct {
		name string
		args args
		want Person
	}{
		{
			name: "With Customer types",
			args: args{p: []Person{Customer{ages: 22}, Customer{ages: 59}, Customer{ages: 45}}},
			want: Customer{ages: 59},
		},
		{
			name: "With Emploer types",
			args: args{p: []Person{Emploer{ages: 33}, Emploer{ages: 12}, Emploer{ages: 11}}},
			want: Emploer{ages: 33},
		},
		{
			name: "With both types",
			args: args{p: []Person{Emploer{ages: 33}, Customer{ages: 45}, Emploer{ages: 11}}},
			want: Customer{ages: 45},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOlder(tt.args.p...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOlder() = %v, want %v", got, tt.want)
			}
		})
	}
}
