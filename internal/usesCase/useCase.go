package useCase

import comicModel "github.com/brayanlopez8/GolangComicsStore/internal/model"

type RegistrationUseCaseInterface interface {
	ListComics() ([]comicModel.Comic, error)
}
