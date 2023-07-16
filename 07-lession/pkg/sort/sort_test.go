package sort

import (
	"reflect"
	"sort"
	"testing"
)

func Test_Strings(t *testing.T) {
	tests := []struct {
		name string
		got  []string
		want []string
	}{
		{
			name: "Random order",
			got:  []string{"banana", "apple", "kiwi", "orange"},
			want: []string{"apple", "banana", "kiwi", "orange"},
		},
		{
			name: "Order by asc",
			got:  []string{"apple", "banana", "kiwi", "orange"},
			want: []string{"apple", "banana", "kiwi", "orange"},
		},
		{
			name: "Order by desc",
			got:  []string{"orange", "kiwi", "banana", "apple"},
			want: []string{"apple", "banana", "kiwi", "orange"},
		},
		{
			name: "Order by cyrilic chars",
			got:  []string{"яблоко", "банан", "киви", "апельсин"},
			want: []string{"апельсин", "банан", "киви", "яблоко"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.got)

			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Expected: %v, but got: %v", tt.want, tt.got)
			}
		})
	}
}

func Test_Ints(t *testing.T) {
	got := []int{5, 2, 6, 3, 1, 4}
	want := []int{1, 2, 3, 4, 5, 6}
	sort.Ints(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected: %v, but got: %v", want, got)
	}
}

func Benchmark_Ints(b *testing.B) {
	s := []int{5, 2, 6, 3, 1, 4}
	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}

func Benchmark_Float64s(b *testing.B) {
	f := []float64{0.7, 0.2, 0.8, 0.1, 0.9, 0.5}
	for i := 0; i < b.N; i++ {
		sort.Float64s(f)
	}
}
