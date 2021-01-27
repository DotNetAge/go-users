package repositories

import (
	"github.com/go-pg/pg/v10"
	uuid "github.com/iris-contrib/go.uuid"
	"go-uds/models"
)

type PgUserRepository struct {
	UserRepository
	db pg.DB
}

func (r *PgUserRepository) Save(user *models.User) (string, error) {
	if user.ID == "" {
		user.ID = uuid.Must(uuid.NewV1()).String()
	}

	_, err := r.db.Model(user).
		OnConflict("(id) DO UPDATE").
		Insert()

	if err != nil {
		return "", err
	}
	return user.ID, nil

}

func (r *PgUserRepository) Find(uid string) (*models.User, error) {
	user := &models.User{ID: uid}
	err := r.db.Model(user).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PgUserRepository) Delete(id string) (bool, error) {
	res, err := r.db.Model(models.User{ID: id}).Where("id=?").Delete()
	return res.RowsAffected() > 0, err
}
