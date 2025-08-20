package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	if PathKey.Filename != expectedOriginalKey {
		t.Errorf("have %s want %s", PathKey.Filename, expectedOriginalKey)
	}

}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "mySpecialPics"
	data := []byte("Some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
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
}
