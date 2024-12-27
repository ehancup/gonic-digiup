package service

import (
	"base-gin/app/repository"
	"base-gin/config"
)

var (
	accountService *AccountService
	personService  *PersonService
	publisherService *PublisherService
	authorsService *AuthorsService
	formService *FormService
)

func SetupServices(cfg *config.Config) {
	accountService = newAccountService(cfg, repository.GetAccountRepo())
	personService = newPersonService(repository.GetPersonRepo())
	publisherService = newPublisherService(repository.GetPublisherRepo())
	authorsService = newAuthorsService(repository.GetAuthorsRepo())
	formService = newFormService(repository.GetFormRepo())
}

func GetAccountService() *AccountService {
	if accountService == nil {
		panic("account service is not initialised")
	}

	return accountService
}

func GetPersonService() *PersonService {
	return personService
}

func GetPublisherService() *PublisherService{
	return publisherService
}

func GetAuthorsService() *AuthorsService{
	return authorsService
}

func GetFormService() *FormService{
	return formService
}