package orderedmap

type OrderedMap struct {
	kv map[interface{}]*Element
	ll list
}

func NewOrderedMap() *OrderedMap { _ = "STUB: not implemented"; return nil }

// NewOrderedMapWithCapacity creates a map with enough pre-allocated space to
// hold the specified number of elements.
func NewOrderedMapWithCapacity(capacity int) *OrderedMap { _ = "STUB: not implemented"; return nil }

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (m *OrderedMap) Get(key interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (m *OrderedMap) Set(key, value interface{}) bool { _ = "STUB: not implemented"; return false }

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (m *OrderedMap) GetOrDefault(key, defaultValue interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (m *OrderedMap) GetElement(key interface{}) *Element { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in the map.
func (m *OrderedMap) Len() int {
	_ = "STUB: not implemented"

	// Keys returns all of the keys in the order they were inserted. If a key was
	// replaced it will retain the same position. To ensure most recently set keys
	// are always at the end you must always Delete before Set.
	return 0
}

func (m *OrderedMap) Keys() (keys []interface{}) { _ = "STUB: not implemented"; return nil }

// Delete will remove a key from the map. It will return true if the key was
// removed (the key did exist).
func (m *OrderedMap) Delete(key interface{}) (didDelete bool) {
	_ = "STUB: not implemented"
	return false
}

// Front will return the element that is the first (oldest Set element). If
// there are no elements this will return nil.
func (m *OrderedMap) Front() *Element { _ = "STUB: not implemented"; return nil }

// Back will return the element that is the last (most recent Set element). If
// there are no elements this will return nil.
func (m *OrderedMap) Back() *Element {
	_ = "STUB: not implemented"

	// Copy returns a new OrderedMap with the same elements.
	// Using Copy while there are concurrent writes may mangle the result.
	return nil
}

func (m *OrderedMap) Copy() *OrderedMap { _ = "STUB: not implemented"; return nil }

// Has checks if a key exists in the map.
func (m *OrderedMap) Has(key interface{}) bool { _ = "STUB: not implemented"; return false }
