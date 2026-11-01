package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	pathName := CASPathTransformFunc(key)
	expectedPathName := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"

	if pathName != expectedPathName {
		t.Errorf("Expected path name %s, but got %s", expectedPathName, pathName)
	}

	fmt.Println(pathName)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("mypictures", data); err != nil {
		t.Error(err)
	}
}
