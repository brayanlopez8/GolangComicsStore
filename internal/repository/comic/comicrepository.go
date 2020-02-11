package comicRepository

import comicModel "github.com/brayanlopez8/GolangComicsStore/internal/model"

type Comic interface {
	DeleteByID(id int) error
	FindAll() ([]*comicModel.Comic, error)
	FindByID(id int) (*comicModel.Comic, error)
	Insert(*comicModel.Comic) (id int, err error)
	Update(*comicModel.Comic) error
}
