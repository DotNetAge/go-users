package repositories

import "go-uds/models"

type IdentityRepository interface {
	Find(id string) (*models.Identity, error)
	FindByName(name string) (*models.Identity, error)
	Valid(name, password string) (bool, error)
	Delete(id string) (bool, error)
	Save(identity *models.Identity) (string, error)
}

type UserRepository interface {
	Save(user *models.User) (string, error)
	Find(uid string) (*models.User, error)
	Delete(id string) (bool, error)
}
