package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	sql := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)`

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (User, error) {
	var user User

	err := repo.db.QueryRow("SELECT id, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Email)
	if err != nil {
		return user, RepoErr
	}

	return user, nil
}
