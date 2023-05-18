package usecase

import (
	"context"

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
	return a.Uow.Do(ctx, func(uow uow.UowInterface) error {

		return nil
	})

	// category := entity.Category{
	// 	Name: input.CategoryName,
	// }

	// err := a.CategoryRepository.Insert(ctx, category)
	// if err != nil {
	// 	return err
	// }

	// course := entity.Course{
	// 	Name:       input.CourseName,
	// 	CategoryID: input.CourseCategoryID,
	// }

	// err = a.CourseRepository.Insert(ctx, course)
	// if err != nil {
	// 	return err
	// }

	// return nil
}
