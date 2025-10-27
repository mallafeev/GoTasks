package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "пустой слайс",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "один элемент",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "уже отсортирован",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "обратный порядок",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "случайный порядок",
			input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			want:  []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name:  "с отрицательными числами",
			input: []int{-3, -1, -4, -1, -5},
			want:  []int{-5, -4, -3, -1, -1},
		},
		{
			name:  "первые 10 из маасива задания 1",
			input: []int{542, -565, 531, -294, -56, 14, 270, -51, -914, 605},
			want:  []int{-914, -565, -294, -56, -51, 14, 270, 531, 542, 605},
		},
		{
			name:  "все одинаковые",
			input: []int{7, 7, 7, 7},
			want:  []int{7, 7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			BubbleSort(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("BubbleSort() = %v, ждали %v", arr, tt.want)
			}
		})
	}
}
