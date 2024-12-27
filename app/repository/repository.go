package repository

import "base-gin/storage"

var (
	accountRepo *AccountRepository
	personRepo  *PersonRepository
	publisherRepo *PublisherRepository
	authorsRepo *AuthorsRepository
	formRepo *FormRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = newAccountRepository(db)
	personRepo = newPersonRepository(db)
	publisherRepo = newPublisherRepo(db)
	authorsRepo = newAuthorsRepo(db)
	formRepo = newFormRepo(db)
}

func GetAccountRepo() *AccountRepository {
	return accountRepo
}

func GetPersonRepo() *PersonRepository {
	return personRepo
}

func GetPublisherRepo() *PublisherRepository {
	return publisherRepo
}

func GetAuthorsRepo() *AuthorsRepository {
	return authorsRepo
}

func GetFormRepo() *FormRepository{
	return formRepo
}