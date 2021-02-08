package repositories

import (
	"github.com/go-pg/pg/v10"
	"go-uds/models"
)

type PgRefreshTokenRepository struct {
	RefreshTokenRepository
	db pg.DB
}

func (r *PgRefreshTokenRepository) Find(key string) (*models.RefreshToken, error) {
	refreshToken := &models.RefreshToken{Key: key}
	err := r.db.Model(refreshToken).
		WherePK().
		Select()
	return refreshToken, err
}

func (r *PgRefreshTokenRepository) Save(refreshToken *models.RefreshToken) error {
	_, err := r.db.Model(refreshToken).
		Insert()
	return err
}

func (r *PgRefreshTokenRepository) Delete(refreshTokenKey string) (bool, error) {
	res, err := r.db.Model(models.RefreshToken{Key: refreshTokenKey}).Where("key=?").Delete()
	if res != nil {
		return res.RowsAffected() > 0, err
	} else {
		return false, err
	}
}
