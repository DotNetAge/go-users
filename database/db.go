package database

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"go-uds/models"
	"os"
)

// 连接数据库
func ConnectDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	})
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

// 创建数据库
func InitDB() error {
	db := ConnectDB()
	_models := []interface{}{
		(*models.User)(nil),
		(*models.Identity)(nil),
		(*models.RefreshToken)(nil),
	}

	for _, _model := range _models {
		err := db.Model(_model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
