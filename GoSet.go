package GoSet

type Set[T comparable] struct {
	m map[T]bool
}

// NewSet Constructor
func (set *Set[T]) NewSet() {
	set.m = make(map[T]bool)
}

// Add method adds the key t. Return false if there is already the key in the map
func (set *Set[T]) Add(t T) bool {
	_, ok := set.m[t]
	if ok {
		return false
	}
	set.m[t] = true
	return true
}

// Delete method deletes the key if it exists in the map. Return false if there is no the key in the map
func (set *Set[T]) Delete(t T) bool {
	for k, _ := range set.m {
		if t == k {
			delete(set.m, k)
			return true
		}
	}
	return false
}

// DeleteAll method deletes all the elements in the set
func (set *Set[T]) DeleteAll() {
	set.m = make(map[T]bool)
}

// Contain check if the input key exists in the map. Return true if it exists, false otherwise.
func (set *Set[T]) Contain(t T) bool {
	if _, ok := set.m[t]; ok {
		return true
	}
	return false
}

// IsEmpty method return true if the set is empty. Return false otherwise
func (set *Set[T]) IsEmpty() bool {
	if l := len(set.m); l == 0 {
		return true
	}
	return false
}

// Size method return the size of the set
func (set *Set[T]) Size() int {
	return len(set.m)
}
