package create

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	f, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("could not create test file: %s", err)
	}
	osStat = func(name string) (os.FileInfo, error) {
		return nil, os.ErrNotExist
	}
	osCreate = func(name string) (*os.File, error) {
		return f, nil
	}
	err = CreateFile("test", "content")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestCreateFile_StatError(t *testing.T) {
	osStat = func(name string) (os.FileInfo, error) {
		return nil, fmt.Errorf("testing error")
	}
	osCreate = func(name string) (*os.File, error) {
		return nil, nil
	}
	err := CreateFile("test", "content")
	if err == nil {
		t.Fatalf("expected an error, saw none")
	}
	if err.Error() != "file system error: testing error" {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestCreateFile_CreationError(t *testing.T) {
	osStat = func(name string) (os.FileInfo, error) {
		return nil, os.ErrNotExist
	}
	osCreate = func(name string) (*os.File, error) {
		return nil, fmt.Errorf("os.Create error")
	}
	err := CreateFile("test", "content")
	if err == nil {
		t.Fatalf("expected an error, saw none")
	}
	if err.Error() != "could not create file: os.Create error" {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestCreateFile_FileExistsError(t *testing.T) {
	osStat = func(name string) (os.FileInfo, error) {
		return nil, nil
	}
	osCreate = func(name string) (*os.File, error) {
		return nil, nil
	}
	err := CreateFile("test", "content")
	if err == nil {
		t.Fatalf("expected an error, saw none")
	}
	if err.Error() != "file already exists" {
		t.Fatalf("unexpected error: %s", err)
	}
}
