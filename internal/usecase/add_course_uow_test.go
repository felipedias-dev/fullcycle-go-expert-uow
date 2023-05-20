package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/felipedias-dev/fullcycle-go-expert-uow/internal/db"
	"github.com/felipedias-dev/fullcycle-go-expert-uow/internal/repository"
	"github.com/felipedias-dev/fullcycle-go-expert-uow/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestAddCourseUow(t *testing.T) {
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dtb.Exec("DROP TABLE IF EXISTS `courses`;")
	dtb.Exec("DROP TABLE IF EXISTS `categories`;")

	dtb.Exec("CREATE TABLE IF NOT EXISTS `categories` (id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255) NOT NULL);")
	dtb.Exec("CREATE TABLE IF NOT EXISTS `categories` (id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255) NOT NULL, category_id INT NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	ctx := context.Background()
	uow := uow.NewUow(ctx, dtb)

	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dtb)
		repo.Queries = db.New(tx)
		return repo
	})
	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dtb)
		repo.Queries = db.New(tx)
		return repo
	})

	input := InputUowUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	useCase := NewAddCourseUowUseCase(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
