package goset

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()
	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}

	s = New(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("expected size 3, got %d", s.Size())
	}
}

func TestAdd(t *testing.T) {
	s := New[int]()
	s.Add(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("expected size 3, got %d", s.Size())
	}
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("set does not contain expected elements")
	}
}

func TestRemove(t *testing.T) {
	s := New(1, 2, 3)
	s.Remove(2)
	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}
	if s.Contains(2) {
		t.Errorf("set contains removed element 2")
	}
}

func TestContains(t *testing.T) {
	s := New(1, 2, 3)
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("set does not contain expected elements")
	}
	if s.Contains(4) {
		t.Errorf("set contains unexpected element 4")
	}
}

func TestSize(t *testing.T) {
	s := New(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("expected size 3, got %d", s.Size())
	}
}

func TestClear(t *testing.T) {
	s := New(1, 2, 3)
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Errorf("expected set to be empty")
	}

	s.Add(1)
	if s.IsEmpty() {
		t.Errorf("expected set to not be empty")
	}
}

func TestClone(t *testing.T) {
	s := New(1, 2, 3)
	clone := s.Clone()
	if !s.Equal(clone) {
		t.Errorf("expected clone to be equal to original set")
	}
}

func TestEqual(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(1, 2, 3)
	if !s1.Equal(s2) {
		t.Errorf("expected sets to be equal")
	}

	s2.Add(4)
	if s1.Equal(s2) {
		t.Errorf("expected sets to not be equal")
	}
}

func TestIsSubsetOf(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(1, 2, 3)
	if !s1.IsSubsetOf(s2) {
		t.Errorf("expected s1 to be a subset of s2")
	}

	if s2.IsSubsetOf(s1) {
		t.Errorf("expected s2 to not be a subset of s1")
	}
}

func TestIsSupersetOf(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(1, 2)
	if !s1.IsSupersetOf(s2) {
		t.Errorf("expected s1 to be a superset of s2")
	}

	if s2.IsSupersetOf(s1) {
		t.Errorf("expected s2 to not be a superset of s1")
	}
}

func TestToSlice(t *testing.T) {
	s := New(1, 2, 3)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("expected slice length 3, got %d", len(slice))
	}
}

func TestUnion(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(2, 3)
	union := s1.Union(s2)
	expected := New(1, 2, 3)
	if !union.Equal(expected) {
		t.Errorf("expected union to be %v, got %v", expected, union)
	}
}

func TestIntersection(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(2, 3)
	intersection := s1.Intersection(s2)
	expected := New(2)
	if !intersection.Equal(expected) {
		t.Errorf("expected intersection to be %v, got %v", expected, intersection)
	}
}

func TestDifference(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(2, 3)
	difference := s1.Difference(s2)
	expected := New(1)
	if !difference.Equal(expected) {
		t.Errorf("expected difference to be %v, got %v", expected, difference)
	}
}

func TestSymmetricDifference(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(2, 3)
	symmetricDifference := s1.SymmetricDifference(s2)
	expected := New(1, 3)
	if !symmetricDifference.Equal(expected) {
		t.Errorf("expected symmetric difference to be %v, got %v", expected, symmetricDifference)
	}
}

func TestJsonEncodeDecode(t *testing.T) {
	s := New(1, 2, 3)
	jsonData, err := s.JsonEncode()
	if err != nil {
		t.Errorf("JsonEncode error: %v", err)
	}

	var decodedSet Set[int]
	err = decodedSet.JsonDecode(jsonData)
	if err != nil {
		t.Errorf("JsonDecode error: %v", err)
	}

	if !s.Equal(&decodedSet) {
		t.Errorf("expected decoded set to be equal to original set")
	}
}

func TestSearchOne(t *testing.T) {
	s := New(1, 2, 3, 4, 5)

	found, value := s.SearchOne(func(v int) bool {
		return v%2 == 0
	})

	if !found || (value != 2 && value != 4) {
		t.Errorf("Expected one even number, got %v", value)
	}

	found, value = s.SearchOne(func(v int) bool {
		return v > 5
	})

	if found {
		t.Errorf("Expected no number greater than 5, but found %v", value)
	}
}

func TestSearchAll(t *testing.T) {
	s := New(1, 2, 3, 4, 5)

	results := s.SearchAll(func(v int) bool {
		return v%2 == 0
	})

	expected := []int{2, 4}
	if len(results) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, results)
	}

	for _, v := range results {
		found := false
		for _, e := range expected {
			if v == e {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected result %v", v)
		}
	}

	results = s.SearchAll(func(v int) bool {
		return v > 5
	})

	if len(results) != 0 {
		t.Errorf("Expected no numbers greater than 5, but found %v", results)
	}
}
