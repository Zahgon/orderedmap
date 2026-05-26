package orderedmap

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The key that corresponds to this element in the ordered map.
	Key interface{}

	// The value stored with this element.
	Value interface{}
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	_ = "STUB: not implemented"

	// Prev returns the previous list element or nil.
	return nil
}

func (e *Element) Prev() *Element {
	_ = "STUB: not implemented"

	// list represents a null terminated (non circular) intrusive doubly linked list.
	// The list is immediately usable after instantiation without the need of a dedicated initialization.
	return nil
}

type list struct {
	root Element // list head and tail
}

func (l *list) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Front returns the first element of list l or nil if the list is empty.
func (l *list) Front() *Element {
	_ = "STUB: not implemented"

	// Back returns the last element of list l or nil if the list is empty.
	return nil
}

func (l *list) Back() *Element {
	_ = "STUB: not implemented"

	// Remove removes e from its list
	return nil
}

func (l *list) Remove(e *Element) { _ = "STUB: not implemented"; return }

// avoid memory leaks
// avoid memory leaks

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *list) PushFront(key interface{}, value interface{}) *Element {
	_ = "STUB: not implemented"
	return nil
}

// It's the first element

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *list) PushBack(key interface{}, value interface{}) *Element {
	_ = "STUB: not implemented"
	return nil
}

// It's the first element
