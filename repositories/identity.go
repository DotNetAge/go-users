package repositories

import (
	"github.com/go-pg/pg/v10"
	uuid "github.com/iris-contrib/go.uuid"
	"go-uds/models"
)

type PgIdentityRepository struct {
	IdentityRepository
	db pg.DB
}

func (r *PgIdentityRepository) Find(id string) (*models.Identity, error) {
	identity := &models.Identity{ID: id}
	err := r.db.Model(identity).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *PgIdentityRepository) FindByName(name string) (*models.Identity, error) {
	identity := models.NewIdentity(name)
	err := r.db.Model(identity).Where("name=?").Select()
	if err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *PgIdentityRepository) Valid(name, password string) (bool, error) {
	identity := &models.Identity{Name: name, Password: password}

	err := r.db.Model(identity).
		Where("name=?").
		Where("password=?").
		Select()
	if err != nil {
		return false, err
	}

	return identity.ID != "", nil
}
func (r *PgIdentityRepository) Delete(id string) (bool, error) {
	res, err := r.db.Model(models.Identity{ID: id}).Where("id=?").Delete()
	return res.RowsAffected() > 0, err
}

func (r *PgIdentityRepository) Save(identity *models.Identity) (string, error) {
	if identity.ID == "" {
		identity.ID = uuid.Must(uuid.NewV1()).String()
	}

	_, err := r.db.Model(identity).
		OnConflict("(id) DO UPDATE").
		Insert()

	if err != nil {
		return "", err
	}
	return identity.ID, nil
}
