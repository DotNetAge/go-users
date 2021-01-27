package repositories

import "github.com/go-pg/pg/v10"

func NewIdentityRepository(db pg.DB) IdentityRepository {
	return &PgIdentityRepository{db: db}
}

func NewUserRepository(db pg.DB) UserRepository {
	return &PgUserRepository{db: db}
}
