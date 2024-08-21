package slice

import (
	"reflect"
	"testing"
)

func TestDeleteIndexes(t *testing.T) {
	type args struct {
		s       []int
		indexes []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0",args{[]int{0,1,2,3,4},[]int{3,1}},[]int{0,2,4}},
		{"0",args{[]int{0,1,2,3,4},[]int{-3,-1}},[]int{0,1,2,3,4}},
		{"0",args{[]int{0,1,2,3,4},[]int{0,-1}},[]int{1,2,3,4}},
		{"0",args{[]int{0,1,2,3,4},[]int{99,100}},[]int{0,1,2,3,4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteIndexes(tt.args.s, tt.args.indexes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}
