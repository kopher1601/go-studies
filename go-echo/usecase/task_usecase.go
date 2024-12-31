package usecase

import (
	"go-echo/model"
	"go-echo/repository"
)

type TaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId, taskId uint) error
}

type taskUsecase struct {
	tr repository.TaskRepository
}

func (t *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	var tasks []model.Task
	if err := t.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	var responses []model.TaskResponse
	for _, task := range tasks {
		responses = append(responses, model.TaskResponse{
			ID:        task.Id,
			Title:     task.Title,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		})
	}
	return responses, nil
}

func (t *taskUsecase) GetTaskById(userId, taskId uint) (model.TaskResponse, error) {
	var task model.Task
	if err := t.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	response := model.TaskResponse{
		ID:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return response, nil
}

func (t *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := t.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	response := model.TaskResponse{
		ID:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return response, nil
}

func (t *taskUsecase) UpdateTask(task model.Task, userId, taskId uint) (model.TaskResponse, error) {
	if err := t.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	response := model.TaskResponse{
		ID:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return response, nil
}

func (t *taskUsecase) DeleteTask(userId, taskId uint) error {
	if err := t.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}

func NewTaskUsecase(tr repository.TaskRepository) TaskUsecase {
	return &taskUsecase{tr: tr}
}
