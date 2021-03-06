// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"
	dao "github.com/yuweiweiouo/coding-exercise/internal/dao"
	model "github.com/yuweiweiouo/coding-exercise/internal/model"
)

// DetailsRepository is an autogenerated mock type for the DetailsRepository type
type TaskDao struct {
	mock.Mock
}

var _ dao.TaskDao = (*TaskDao)(nil)

func (dao *TaskDao) GetAll() []model.Task {
	args := dao.Called()
	return args.Get(0).([]model.Task)
}
func (dao *TaskDao) Create(task model.Task) (model.Task, error){
	args := dao.Called(task)
	return args.Get(0).(model.Task), args.Error(1)
}
func (dao *TaskDao) Update(task model.Task) (model.Task, error){
	args := dao.Called(task)
	return args.Get(0).(model.Task), args.Error(1)
}
func (dao *TaskDao) Delete(id int) error{
	args := dao.Called(id)
	return args.Error(0)
}