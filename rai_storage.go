package main

import (
	"errors"
	"github.com/google/uuid"
)

var errNotFound = errors.New("not found")

type Storage interface {
	CreateTask(task Task) (string, error)
	UpdateTask(id string, task Task) error
	GetTask(id string) (*Task, error)
	DeleteTask(id string) error
}

type RaiStorage struct {
	data map[string]Task
}

func NewRaiStorage() *RaiStorage {
	return &RaiStorage{
		data: make(map[string]Task),
	}
}

func (rs *RaiStorage) GetTask(id string) (*Task, error) {
	value, exists := rs.data[id]
	if !exists {
		return nil, errNotFound
	}
	return &value, nil
}

func (rs *RaiStorage) UpdateTask(id string, task Task) error {
	if _, exists := rs.data[id]; !exists {
		return errNotFound
	}
	rs.data[id] = task
	return nil
}

func (rs *RaiStorage) CreateTask(task Task) (string, error) {
	id := uuid.New().String()
	task.Status = "in_progress"
	rs.data[id] = task
	return id, nil
}

func (rs *RaiStorage) DeleteTask(id string) error {
	if _, exists := rs.data[id]; !exists {
		return errNotFound
	}
	delete(rs.data, id)
	return nil
}
