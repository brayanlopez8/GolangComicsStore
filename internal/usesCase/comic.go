package comic

import (
	comicModel "github.com/brayanlopez8/GolangComicsStore/internal/model"
	comicRepository "github.com/brayanlopez8/GolangComicsStore/internal/repository/comic"
)

//ComicUseCase ...
type ComicUseCase struct {
	ComicRepositoryInterface comicRepository.Comic
}

//FindAll  handles GET /games requests and returns all Games from database.
func (ucc *ComicUseCase) FindAll() ([]*comicModel.Comic, error) {
	return ucc.ComicRepositoryInterface.FindAll()
}
