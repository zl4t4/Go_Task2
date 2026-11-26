package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE           = 100_000_000
	CHUNKS         = 8
	maxRandomValue = 1_000_000
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Intn(maxRandomValue) + 1
	}
	return elements
}

func maximum(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	maxValue := elements[0]
	for _, value := range elements[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func maxChunks(elements []int) int {
	if len(elements) == 0 {
		return 0
	}
	// Слайс слишком мал, чтобы делить его на части.
	if len(elements) < CHUNKS {
		return maximum(elements)
	}

	chunkSize := len(elements) / CHUNKS
	chunkMaximums := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		// Последняя часть забирает остаток, если длина не делится на CHUNKS.
		if i == CHUNKS-1 {
			end = len(elements)
		}

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaximums[index] = maximum(chunk)
		}(i, elements[start:end])
	}
	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(elements)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nЭто заняло %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(elements)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nЭто заняло %d мкс\n", max, elapsed)
}
