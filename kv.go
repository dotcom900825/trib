package trib

import (
	"strings"
)

type KeyValue struct {
	Key   string
	Value string
}

type Pattern struct {
	Prefix string
	Suffix string
}

func (p *Pattern) Match(k string) bool {
	ret := strings.HasPrefix(k, p.Prefix)
	ret = ret && strings.HasSuffix(k, p.Suffix)
	return ret
}

type List struct {
	L []string
}

func KV(k, v string) *KeyValue { return &KeyValue{k, v} }

// Key-value pair interfaces
// Default value for all keys is empty string
type KeyString interface {
	// Gets a value. Empty string by default.
	Get(key string, value *string) error

	// Set kv.Key to kv.Value. Set succ to true when no error.
	Set(kv *KeyValue, succ *bool) error

	// List all the keys of non-empty pairs where the key matches
	// the given pattern.
	Keys(p *Pattern, list *List) error
}

// Key-list interfaces.
// Default value for all lists is an empty list.
// After the call, list.L should never by nil.
type KeyList interface {
	// Get the list.
	ListGet(key string, list *List) error

	// Append a string to the list. Set succ to true when no error.
	ListAppend(kv *KeyValue, succ *bool) error

	// Removes all elements that equals to kv.Value in list kv.Key
	// n is set to the number of elements removed.
	ListRemove(kv *KeyValue, n *int) error

	// List all the keys of non-empty lists, where the key matches
	// the given pattern.
	ListKeys(p *Pattern, list *List) error
}

type Storage interface {
	// Returns an auto-incrementing clock, the returned value
	// will be no smaller than atLeast, and it will be
	// strictly larger than the value returned last time,
	// unless it was math.MaxUint64.
	Clock(atLeast uint64, ret *uint64) error

	KeyString
	KeyList
}
