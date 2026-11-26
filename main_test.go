package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"нулевой размер", 0, 0},
		{"отрицательный размер", -5, 0},
		{"один элемент", 1, 1},
		{"обычный размер", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Fatalf("длина слайса = %d, ожидалось %d", len(got), tt.want)
			}
			for _, value := range got {
				if value <= 0 {
					t.Fatalf("получено неположительное число: %d", value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		want     int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{7}, 7},
		{"максимум в начале", []int{9, 2, 3}, 9},
		{"максимум в середине", []int{1, 9, 3}, 9},
		{"максимум в конце", []int{1, 2, 9}, 9},
		{"одинаковые элементы", []int{4, 4, 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.elements); got != tt.want {
				t.Errorf("maximum() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	sizes := []int{0, 1, 7, 8, 9, 15, 16, 100, 1001}

	for _, size := range sizes {
		elements := generateRandomElements(size)
		want := maximum(elements)
		if got := maxChunks(elements); got != want {
			t.Errorf("размер %d: maxChunks() = %d, ожидалось %d", size, got, want)
		}
	}
}

func TestMaxChunksMaxInLastElement(t *testing.T) {
	// Максимум в хвосте, который не делится на части поровну.
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 100}
	if got := maxChunks(elements); got != 100 {
		t.Errorf("maxChunks() = %d, ожидалось 100", got)
	}
}
