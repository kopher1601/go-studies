package repository

import (
	"fmt"
	"go-echo/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func (t *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	// select * from task t join user u on t.user_id = u.id
	// select * from task t where t.user_id = ?
	if err := t.db.
		Joins("User").
		Where("user_id = ?", userId).
		Order("created_at").Find(tasks).
		Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := t.db.
		Joins("User").
		Where("user_id = ? and task_id =?", userId, taskId).
		First(task).
		Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) CreateTask(task *model.Task) error {
	if err := t.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	tx := t.db.
		Model(task).
		Clauses(clause.Returning{}).
		Where("id = ? AND user_id = ?", taskId, userId).
		Update("title", task.Title)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (t *taskRepository) DeleteTask(userId uint, taskId uint) error {
	tx := t.db.Where("id=? and user_id = ?", taskId, userId).Delete(&model.Task{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}
