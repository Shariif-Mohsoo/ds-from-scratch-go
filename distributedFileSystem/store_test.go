package main

import (
	"bytes"
	"testing"
)

func TestPathTransformationFunc(t *testing.T) {
	key := "mineBestPicture"
	PathKey := CASPathTransformFunc(key)
	// fmt.Println(pathname)
	expectedOriginalKey := "7b31db6647a3b5581b88bbe5b9355eabf98796e0"
	expectedPathName := "7b31d/b6647/a3b55/81b88/bbe5b/9355e/abf98/796e0"

	if PathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", PathKey.PathName, expectedPathName)
	}

	if PathKey.Original != expectedOriginalKey {
		t.Errorf("have %s want %s", PathKey.Original, expectedOriginalKey)
	}

}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("Some jpg bytes"))
	if err := s.writeStream("mySpecialPicture", data); err != nil {
		t.Error(err)
	}

}
