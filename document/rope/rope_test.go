package rope

import "testing"

func TestSimpleInsertion(t *testing.T) {
	store := NewRope("abcdef")
	store.Insert(3, "aaa")
	str := store.ToString()
	if str != "abcaaadef" {
		t.Error("Expected abcaaadef, got ", str)
	}
}

func TestSimpleDeletion(t *testing.T) {
	store := NewRope("abcdef")
	store.Delete(1, 1)
	str := store.ToString()
	if str != "acdef" {
		t.Error("Expected acdef, got ", str)
	}
}

func TestSameness(t *testing.T) {
	store := NewRope("abcdef")
	for i := 0; i < 100; i++ {
		store.Insert(2, "a")
		store.Delete(2, 1)
	}
	str := store.ToString()
	if str != "abcdef" {
		t.Error("Expected abcdef, got ", str)
	}
}

func TestSeriesInsertion(t *testing.T) {
	store := NewRope("abcdef")
	store.Insert(3, "aaa")
	str := store.ToString()
	if str != "abcaaadef" {
		t.Error("Expected abcaaadef, got ", str)
	}
	store.Insert(0, "!")
	str = store.ToString()
	if str != "!abcaaadef" {
		t.Error("Expected !abcaaadef, got ", str)
	}
	store.Insert(10, "!")
	str = store.ToString()
	if str != "!abcaaadef!" {
		t.Error("Expected !abcaaadef!, got ", str)
	}
	value := store.Insert(13, "aaa")
	if value {
		t.Error("Expected false, got ", value)
	}
	value = store.Insert(3, "aaa")
	if !value {
		t.Error("Expected true, got ", value)
	}
}

func TestSeriesDeletion(t *testing.T) {
	store := NewRope("abcdef")
	store.Delete(1, 1)
	str := store.ToString()
	if str != "acdef" {
		t.Error("Expected abcaaadef, got ", str)
	}
	store.Delete(0, 1)
	str = store.ToString()
	if str != "cdef" {
		t.Error("Expected cdef, got ", str)
	}
	store.Delete(3, 1)
	str = store.ToString()
	if str != "cde" {
		t.Error("Expected cde, got ", str)
	}
	value := store.Delete(0, 45)
	if value {
		t.Error("Expected false, got ", value)
	}
	value = store.Delete(0, 1)
	if !value {
		t.Error("Expected true, got ", value)
	}
}

func TestLargeInsertion(t *testing.T) {
	store := NewRope("aaaaaaaaaa")
	for i := 0; i < 10000; i++ {
		store.Insert(9, "aaaaaaaaaa")
	}
	str := store.ToString()
	length := len(str)
	if length != 100010 {
		t.Error("Expected length 100010, got ", length)
	}
}

func TestLargeDeletion(t *testing.T) {
	store := NewRope("aaaaaaaaaa")
	for i := 0; i < 10000; i++ {
		store.Insert(9, "aaaaaaaaaa")
	}

	for i := 0; i < 100; i++ {
		store.Delete(42, 30)
	}
	str := store.ToString()
	length := len(str)
	if length != 97010 {
		t.Error("Expected length 97010, got ", length)
	}
}
