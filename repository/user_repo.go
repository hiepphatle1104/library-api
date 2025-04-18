package repository

import (
	"context"
	"library-api/database"
	"library-api/model"
)

type UserRepository struct {
	db *database.Database
}

func NewUserRepo(db *database.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User, ctx context.Context) error {
	_, err := r.db.DB.NewInsert().Model(user).Exec(ctx)

	return err
}

func (r *UserRepository) FindByUsername(userName string, ctx context.Context) (*model.User, error) {
	var user *model.User
	err := r.db.DB.NewSelect().Model(user).Where("user_name = ?", userName).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) ExistByUsername(userName string, ctx context.Context) (bool, error) {
	var user *model.User
	exists, err := r.db.DB.NewSelect().Model(user).Where("user_name = ?", userName).Exists(ctx)
	if err != nil {
		return false, err
	}

	return exists, err
}
