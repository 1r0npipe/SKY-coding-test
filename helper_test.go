package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_stringToSlice(t *testing.T) {
	tests := []struct {
		input   string
		want    []int
		wantErr error
	}{
		{input: "1,2,3,4,5,6", want: []int{1, 2, 3, 4, 5, 6}, wantErr: nil},
		{input: "15,25,45,78,567", want: []int{15, 25, 45, 78, 567}, wantErr: nil},
	}
	for _, tt := range tests {
		want, err := stringToSlice(tt.input)
		fmt.Println(reflect.DeepEqual(want, tt.want))
		if !reflect.DeepEqual(want, tt.want) {
			t.Errorf("stringToSlice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}

}

func Test_contains(t *testing.T) {
	tests := []struct {
		in   []int
		s    int
		want bool
	}{
		{in: []int{1, 2, 3, 4, 5}, s: 5, want: true},
		{in: []int{1, 2, 3, 4, 5}, s: 55, want: false},
	}
	for _, tt := range tests {
		if got := contains(tt.in, tt.s); got != tt.want {
			t.Errorf("contains() = %v, want %v", got, tt.want)
		}
	}
}

func Test_sumSlice(t *testing.T) {
	tests := []struct {
		in  []int
		want int
	}{
		{ in: []int{1,2,3,4,5}, want: 15},
		{ in: []int{11,222,3345,4567,58}, want: 8203},
	}
	for _, tt := range tests {
		if got := sumSlice(tt.in); got != tt.want {
			t.Errorf("sumSlice() = %v, want %v", got, tt.want)
		}
	}
}

func Test_findMaxMap(t *testing.T) {
	tests := []struct {
		in map[int]int
		max  int
		ind int
	}{
		{in: map[int]int{1:5,2:10,4:55,55:123},max:123, ind: 55},
		{in: map[int]int{11:55,29:111,44:456,5:99},max:456, ind: 44},
	}
	for _, tt := range tests {
		max, ind := findMaxMap(tt.in)
		if max != tt.max {
			t.Errorf("findMaxMap() max = %v, want %v", max, tt.max)
		}
		if ind != tt.ind {
			t.Errorf("findMaxMap() ind = %v, want %v", ind, tt.ind)
		}
	}
}