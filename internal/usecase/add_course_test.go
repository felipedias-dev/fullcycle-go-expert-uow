package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/felipedias-dev/fullcycle-go-expert-uow/internal/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dtb.Exec("DROP TABLE IF EXISTS `courses`;")
	dtb.Exec("DROP TABLE IF EXISTS `categories`;")

	dtb.Exec("CREATE TABLE IF NOT EXISTS `categories` (id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255) NOT NULL);")
	dtb.Exec("CREATE TABLE IF NOT EXISTS `categories` (id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255) NOT NULL, category_id INT NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	input := InputUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	ctx := context.Background()
	useCase := NewAddCourseUseCase(repository.NewCategoryRepository(dtb), repository.NewCourseRepository(dtb))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
