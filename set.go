package set

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// SortedSet представляет сортированное множество.
type SortedSet[T constraints.Ordered] struct {
	elements map[T]struct{}
	sorted   []T
}

// NewSortedSet создает новое пустое сортированное множество.
func NewSortedSet[T constraints.Ordered]() *SortedSet[T] {
	return &SortedSet[T]{
		elements: make(map[T]struct{}),
		sorted:   []T{},
	}
}

// Add добавляет элемент в множество, поддерживая сортировку.
func (s *SortedSet[T]) Add(value T) {
	if _, exists := s.elements[value]; !exists {
		s.elements[value] = struct{}{}
		s.sorted = append(s.sorted, value)
		sort.Slice(s.sorted, func(i, j int) bool {
			return s.sorted[i] < s.sorted[j] // Сравнение для стандартных типов
		})
	}
}

// Remove удаляет элемент из множества.
func (s *SortedSet[T]) Remove(value T) {
	if _, exists := s.elements[value]; exists {
		delete(s.elements, value)

		// Удаляем элемент из отсортированного среза
		index := sort.Search(len(s.sorted), func(i int) bool {
			return s.sorted[i] >= value
		})
		if index < len(s.sorted) && s.sorted[index] == value {
			s.sorted = append(s.sorted[:index], s.sorted[index+1:]...)
		}
	}
}

// Contains проверяет, содержится ли элемент в множестве.
func (s *SortedSet[T]) Contains(value T) bool {
	_, exists := s.elements[value]
	return exists
}

// Size возвращает количество элементов в множестве.
func (s *SortedSet[T]) Size() int {
	return len(s.elements)
}

// Elements возвращает отсортированные элементы множества.
func (s *SortedSet[T]) Elements() []T {
	return s.sorted
}