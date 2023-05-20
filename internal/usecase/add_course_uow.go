package usecase

import (
	"context"

	"github.com/felipedias-dev/fullcycle-go-expert-uow/internal/entity"
	"github.com/felipedias-dev/fullcycle-go-expert-uow/internal/repository"
	"github.com/felipedias-dev/fullcycle-go-expert-uow/pkg/uow"
)

type InputUowUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUowUseCase struct {
	Uow uow.UowInterface
}

func NewAddCourseUowUseCase(uow uow.UowInterface) *AddCourseUowUseCase {
	return &AddCourseUowUseCase{
		Uow: uow,
	}
}

func (a *AddCourseUowUseCase) Execute(ctx context.Context, input InputUowUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		category := entity.Category{
			Name: input.CategoryName,
		}
		repoCategory := a.getCategoryRepository(ctx)
		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
		repoCourse := a.getCourseRepository(ctx)
		err = repoCourse.Insert(ctx, course)
		if err != nil {
			return err
		}
		return nil
	})
}

func (a *AddCourseUowUseCase) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUowUseCase) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}
