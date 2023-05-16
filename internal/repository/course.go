package repository

import (
	"context"
	"database/sql"

	"github.com/felipedias-dev/fullcycle-go-expert-uow/internal/db"
	"github.com/felipedias-dev/fullcycle-go-expert-uow/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(dtb *sql.DB) *CourseRepository {
	return &CourseRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name: course.Name, CategoryID: int32(course.CategoryID),
	})
}
