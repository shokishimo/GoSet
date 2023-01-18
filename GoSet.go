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
