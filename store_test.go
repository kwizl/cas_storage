package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"
	expectedPathName := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"

	if pathKey.PathName != expectedPathName {
		t.Errorf("Expected path name %s, but got %s", expectedPathName, pathKey.PathName)
	}

	if pathKey.Original != expectedOriginalKey {
		t.Errorf("Expected original key %s, but got %s", expectedOriginalKey, pathKey.Original)
	}

	fmt.Println(pathKey)
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
