package comicUseCase

import (
	comicModel "github.com/brayanlopez8/GolangComicsStore/internal/model"
	"github.com/brayanlopez8/GolangComicsStore/internal/repository"
)

//ListComicUseCase ...
type ListComicUseCase struct {
	ComicRepositoryInterface repository.ComicRepositoryInterface
}

//ListComics ...
func (luc *ListComicUseCase) ListComics() ([]comicModel.Comic, error) {
	return luc.ComicRepositoryInterface.FindAll()
}
