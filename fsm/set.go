package fsm

import (
	"fmt"
	"sort"
	"strings"
)

// Set is a type where we can store a list that has no duplicates. Using map[T]struct{} as a way to store a set
type Set[T comparable] map[T]struct{}

// NewSet creates a new set and optionally initializes it with elements
func NewSet[T comparable](elements ...T) Set[T] {
	s := make(Set[T])
	for _, e := range elements {
		s[e] = struct{}{}
	}
	return s
}

// Add inserts an element into the set
func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

// Remove deletes an element from the set
func (s Set[T]) Remove(val T) {
	delete(s, val)
}

// Contains checks if an element exists in the set
func (s Set[T]) Contains(val T) bool {
	_, exists := s[val]
	return exists
}

// Size returns the number of elements in the set
func (s Set[T]) Size() int {
	return len(s)
}

// String returns a string representation of the set
func (s Set[T]) String() string {
	var elements []string
	for k := range s {
		elements = append(elements, fmt.Sprintf("%v", k))
	}
	// Sort for deterministic order
	sort.Strings(elements)
	return "(" + strings.Join(elements, ", ") + ")"
}

// Parse adds elements from input like "(a,b,c)" into the set.
// A converter function is required to transform each token into T.
func (s Set[T]) Parse(input string, convert func(string) (T, error)) error {
	// remove surrounding parentheses
	input = strings.TrimSpace(input)
	if strings.HasPrefix(input, "(") && strings.HasSuffix(input, ")") {
		input = input[1 : len(input)-1]
	} else {
		fmt.Printf("Don't have parantheses")
		return fmt.Errorf("input string doesn't have surrounding parentheses")
	}

	// split and convert
	for _, tok := range strings.Split(input, ",") {
		trimmed := strings.TrimSpace(tok)
		if trimmed == "" {
			continue
		}
		val, err := convert(trimmed)
		if err != nil {
			return err
		}
		s.Add(val)
	}
	return nil
}

// DeepCopy returns a new Set[T] containing all elements of s
func (s Set[T]) DeepCopy() Set[T] {
	copy := make(Set[T], len(s))
	for k := range s {
		copy[k] = struct{}{}
	}
	return copy
}

// IsSubset checks if the set s is a subset of other
func (s Set[T]) IsSubset(other Set[T]) bool {
	for k := range s {
		if _, exists := other[k]; !exists {
			return false
		}
	}
	return true
}
