package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPathTransformationFunc(t *testing.T) {
	key := "mineBestPicture"
	pathKey := CASPathTransformFunc(key)
	// fmt.Println(pathname)
	expectedFilename := "7b31db6647a3b5581b88bbe5b9355eabf98796e0"
	expectedPathName := "7b31d/b6647/a3b55/81b88/bbe5b/9355e/abf98/796e0"

	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Filename != expectedFilename {
		t.Errorf("have %s want %s", pathKey.Filename, expectedFilename)
	}

}

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("Some jpg bytes")
		if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Errorf("Expected to have key %s", key)
		}

		r, err := s.readStream(key)
		if err != nil {
			t.Error(err)
		}

		b, _ := ioutil.ReadAll(r)
		fmt.Println(string(b))
		if string(b) != string(data) {
			t.Errorf("want %s but get %s", data, b)
		}

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); ok {
			t.Errorf("Expected to not have key %s", key)
		}

	}

}

// helper funcs
func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
