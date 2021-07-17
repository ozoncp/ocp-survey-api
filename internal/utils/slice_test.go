package utils

import (
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	type args struct {
		s         []int
		batchSize uint
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Nil slice",
			args: args{
				s:         nil,
				batchSize: 3,
			},
			want: nil,
		},
		{
			name: "Empty slice",
			args: args{
				s:         []int{},
				batchSize: 3,
			},
			want: nil,
		},
		{
			name: "Invalid batch size",
			args: args{
				s:         []int{0, 1, 2},
				batchSize: 0,
			},
			want: nil,
		},
		{
			name: "Smaller than batch",
			args: args{
				s:         []int{0, 1, 2},
				batchSize: 5,
			},
			want: [][]int{{0, 1, 2}},
		},
		{
			name: "Same size as batch",
			args: args{
				s:         []int{0, 1, 2, 3},
				batchSize: 4,
			},
			want: [][]int{{0, 1, 2, 3}},
		},
		{
			name: "Larger than batch",
			args: args{
				s:         []int{0, 1, 2, 3, 4, 5, 6, 7},
				batchSize: 3,
			},
			want: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitSlice(tt.args.s, tt.args.batchSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseMap(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[int]string
	}{
		{
			name: "Nil map",
			args: args{m: nil},
			want: nil,
		},
		{
			name: "Empty map",
			args: args{m: map[string]int{}},
			want: nil,
		},
		{
			name: "Normal map",
			args: args{m: map[string]int{
				"admin": 0,
				"user":  1,
				"guest": 2,
			}},
			want: map[int]string{
				0: "admin",
				1: "user",
				2: "guest",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseMap(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseMap_Duplicates(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []map[int]string // due to random map iteration order, we have multiple valid results
	}{
		{
			name: "Duplicate values",
			args: args{m: map[string]int{
				"admin": 0,
				"root":  0,
				"user":  1,
				"guest": 2,
			}},
			want: []map[int]string{
				{
					0: "root",
					1: "user",
					2: "guest",
				},
				{
					0: "admin",
					1: "user",
					2: "guest",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseMap(tt.args.m)
			ok := false
			for _, want := range tt.want {
				if reflect.DeepEqual(got, want) {
					ok = true
					break
				}
			}
			if !ok {
				t.Errorf("ReverseMap() = %v, want one of %v", got, tt.want)
			}
		})
	}
}

func TestFilterSlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Nil slice",
			args: args{s: nil},
			want: nil,
		},
		{
			name: "Empty slice",
			args: args{s: []int{}},
			want: nil,
		},
		{
			name: "Nothing to remove",
			args: args{s: []int{2, 3, 4, 5}},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "Remove some items",
			args: args{s: []int{0, 1, 2, 3, 4, 5, -1}},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "Remove all items",
			args: args{s: []int{-1, 0, 1, 0}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
