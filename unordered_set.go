package set

// Set представляет обычное множество, реализованное с помощью map.
type Set struct {
	elements map[interface{}]struct{}
}

// NewSet создает новое множество.
func NewSet() *Set {
	return &Set{
		elements: make(map[interface{}]struct{}),
	}
}

// Add добавляет элемент в множество.
func (s *Set) Add(value interface{}) {
	s.elements[value] = struct{}{}
}

// Remove удаляет элемент из множества.
func (s *Set) Remove(value interface{}) {
	delete(s.elements, value)
}

// Contains проверяет, содержится ли элемент в множестве.
func (s *Set) Contains(value interface{}) bool {
	_, exists := s.elements[value]
	return exists
}

// Size возвращает количество элементов в множестве.
func (s *Set) Size() int {
	return len(s.elements)
}

// Clear очищает множество.
func (s *Set) Clear() {
	s.elements = make(map[interface{}]struct{})
}

// Elements возвращает все элементы множества.
func (s *Set) Elements() []interface{} {
	elements := make([]interface{}, 0, len(s.elements))
	for element := range s.elements {
		elements = append(elements, element)
	}
	return elements
}