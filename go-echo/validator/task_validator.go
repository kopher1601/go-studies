package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-echo/model"
)

type TaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct {
}

func NewTaskValidator() TaskValidator {
	return &taskValidator{}
}

func (t *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 characters"),
		),
	)
}
