package answers

import (
	"Online-Test-GO/internal/entity"
	basic_repo "Online-Test-GO/internal/repository/_basic_repo"
	"golang.org/x/net/context"
)

type Repository interface {
	GetAll(ctx context.Context, filter users.Filter) ([]users.List, int, error)
	GetById(ctx context.Context, id int) (users.Detail, error)
	Create(ctx context.Context, data users.Create) (entity.Answers, error)
	Update(ctx context.Context, data users.Update) (entity.Answers, error)
	Delete(ctx context.Context, data basic_repo.Delete) error
}
