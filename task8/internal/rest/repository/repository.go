package repository

import (
	"GoTasks/task8/internal/rest/model"
	"context"
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee model.Employee) error
	Read(ctx context.Context, id int64) (model.Employee, error)
	Update(ctx context.Context, employee model.Employee) error
	Delete(ctx context.Context, id int64) error

	List(ctx context.Context) ([]model.Employee, error)
}
