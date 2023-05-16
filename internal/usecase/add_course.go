package usecase

import (
	"context"

	"github.com/felipedias-dev/fullcycle-go-expert-uow/internal/entity"
	"github.com/felipedias-dev/fullcycle-go-expert-uow/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCase struct {
	CategoryRepository repository.CategoryRepositoryInterface
	CourseRepository   repository.CourseRepositoryInterface
}

func NewAddCourseUseCase(categoryRepository repository.CategoryRepositoryInterface, courseRepository repository.CourseRepositoryInterface) *AddCourseUseCase {
	return &AddCourseUseCase{
		CategoryRepository: categoryRepository,
		CourseRepository:   courseRepository,
	}
}

func (a *AddCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := a.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}

	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
