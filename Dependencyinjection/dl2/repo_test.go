package main

import (
	"errors"
	"testing"
)

type MockDB struct {
	GetFunc func(id string) (string, error)
}

func (m *MockDB) Get(id string) (string, error) {
	return m.GetFunc(id)
}

func TestRepository_Get(t *testing.T) {
	mockDB := &MockDB{
		GetFunc: func(id string) (string, error) {
			if id != "1" {
				return "", errors.New("not found")
			}
			return "Data", nil
		},
	}
	repo := NewRepo(mockDB)
	data, err := repo.Get("1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if data != "Data" {
		t.Errorf("expected data to be 'Data', got '%v'", data)
	}
}
