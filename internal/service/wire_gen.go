// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/yuweiweiouo/coding-exercise/internal/dao"
)

// Injectors from wire.go:

func CreateTaskService(taskDao dao.TaskDao) (TaskService, error) {
	serviceTaskService := NewTaskService(taskDao)
	return serviceTaskService, nil
}