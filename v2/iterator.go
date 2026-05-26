//go:build go1.23
// +build go1.23

package orderedmap

import "iter"

func (m *OrderedMap[K, V]) Iterator() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }

func (m *OrderedMap[K, V]) ReverseIterator() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }
