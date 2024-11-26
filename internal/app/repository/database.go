package repository

type IRepository interface{}
type Repository struct{}

func NewRepoInit() IRepository {
	return &Repository{}
}
