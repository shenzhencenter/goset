package goset

import "encoding/json"

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable](v ...T) *Set[T] {
	s := &Set[T]{m: make(map[T]struct{})}
	for _, v := range v {
		s.Add(v)
	}
	return s
}

func (s *Set[T]) Add(v ...T) {
	for _, v := range v {
		s.m[v] = struct{}{}
	}
}

func (s *Set[T]) Remove(v ...T) {
	for _, v := range v {
		delete(s.m, v)
	}
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set[T]) Clone() *Set[T] {
	res := New[T]()
	for v := range s.m {
		res.Add(v)
	}
	return res
}

func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for v := range s.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	if s.Size() > other.Size() {
		return false
	}
	for v := range s.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSupersetOf(other *Set[T]) bool {
	return other.IsSubsetOf(s)
}

func (s *Set[T]) ToSlice() []T {
	var res []T
	for v := range s.m {
		res = append(res, v)
	}
	return res
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	res := New[T]()
	for v := range s.m {
		res.Add(v)
	}
	for v := range other.m {
		res.Add(v)
	}
	return res
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	res := New[T]()
	for v := range s.m {
		if other.Contains(v) {
			res.Add(v)
		}
	}
	return res
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	res := New[T]()
	for v := range s.m {
		if !other.Contains(v) {
			res.Add(v)
		}
	}
	return res
}

func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	res := New[T]()
	for v := range s.m {
		if !other.Contains(v) {
			res.Add(v)
		}
	}
	for v := range other.m {
		if !s.Contains(v) {
			res.Add(v)
		}
	}
	return res
}

func (s *Set[T]) JsonEncode() ([]byte, error) {
	slice := s.ToSlice()
	return json.Marshal(slice)
}

func (s *Set[T]) JsonDecode(data []byte) error {
	var slice []T
	err := json.Unmarshal(data, &slice)
	if err != nil {
		return err
	}
	s.Clear()
	s.Add(slice...)
	return nil
}

func (s *Set[T]) SearchOne(f func(v T) bool) (bool, T) {
	var null T
	for v := range s.m {
		if f(v) {
			return true, v
		}
	}
	return false, null
}

func (s *Set[T]) SearchAll(f func(v T) bool) []T {
	var res []T
	for v := range s.m {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
