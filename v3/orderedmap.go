package orderedmap

import "iter"

type OrderedMap[K comparable, V any] struct {
	kv map[K]*Element[K, V]
	ll list[K, V]
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] { _ = "STUB: not implemented"; return nil }

// NewOrderedMapWithCapacity creates a map with enough pre-allocated space to
// hold the specified number of elements.
func NewOrderedMapWithCapacity[K comparable, V any](capacity int) *OrderedMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func NewOrderedMapWithElements[K comparable, V any](els ...*Element[K, V]) *OrderedMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (m *OrderedMap[K, V]) Get(key K) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (m *OrderedMap[K, V]) Set(key K, value V) bool { _ = "STUB: not implemented"; return false }

// ReplaceKey replaces an existing key with a new key while preserving order of
// the value. This function will return true if the operation was successful, or
// false if 'originalKey' is not found OR 'newKey' already exists (which would be an overwrite).
func (m *OrderedMap[K, V]) ReplaceKey(originalKey, newKey K) bool {
	_ = "STUB: not implemented"
	return false
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (m *OrderedMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	_ = "STUB: not implemented"
	return *new(V)
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (m *OrderedMap[K, V]) GetElement(key K) *Element[K, V] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in the map.
func (m *OrderedMap[K, V]) Len() int {
	_ = "STUB: not implemented"

	// AllFromFront returns an iterator that yields all elements in the map starting
	// at the front (oldest Set element).
	return 0
}

func (m *OrderedMap[K, V]) AllFromFront() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }

// AllFromBack returns an iterator that yields all elements in the map starting
// at the back (most recent Set element).
func (m *OrderedMap[K, V]) AllFromBack() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }

// Keys returns an iterator that yields all the keys in the map starting at the
// front (oldest Set element). To create a slice containing all the map keys,
// use the slices.Collect function on the returned iterator.
func (m *OrderedMap[K, V]) Keys() iter.Seq[K] { _ = "STUB: not implemented"; return nil }

// Values returns an iterator that yields all the values in the map starting at
// the front (oldest Set element). To create a slice containing all the map
// values, use the slices.Collect function on the returned iterator.
func (m *OrderedMap[K, V]) Values() iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// Delete will remove a key from the map. It will return true if the key was
// removed (the key did exist).
func (m *OrderedMap[K, V]) Delete(key K) (didDelete bool) { _ = "STUB: not implemented"; return false }

// Front will return the element that is the first (oldest Set element). If
// there are no elements this will return nil.
func (m *OrderedMap[K, V]) Front() *Element[K, V] { _ = "STUB: not implemented"; return nil }

// Back will return the element that is the last (most recent Set element). If
// there are no elements this will return nil.
func (m *OrderedMap[K, V]) Back() *Element[K, V] {
	_ = "STUB: not implemented"

	// Copy returns a new OrderedMap with the same elements.
	// Using Copy while there are concurrent writes may mangle the result.
	return nil
}

func (m *OrderedMap[K, V]) Copy() *OrderedMap[K, V] { _ = "STUB: not implemented"; return nil }

// Has checks if a key exists in the map.
func (m *OrderedMap[K, V]) Has(key K) bool { _ = "STUB: not implemented"; return false }
