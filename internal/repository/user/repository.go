package user

import (
	"context"
	"github.com/FreylGit/platform_common/pkg/db"

	"github.com/FreylGit/auth/internal/model"
	"github.com/FreylGit/auth/internal/repository"
	"github.com/FreylGit/auth/internal/repository/user/converter"
	modelRepo "github.com/FreylGit/auth/internal/repository/user/model"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	tableNameUser      = "users"
	idColumn           = "id"
	nameColumn         = "name"
	emailColumn        = "email"
	passwordHashColumn = "password_hash"
	createdAtColumn    = "created_at"
	updatedAtColumn    = "updated_at"

	tableNameUserRole = "user_role"
	userIdColumn      = "user_id"
	roleIdColumn      = "role_id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableNameUser).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}
	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(user), nil
}

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {

	builder := sq.Insert(tableNameUser).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordHashColumn).
		Values(user.Name, user.Email, user.Password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}
	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Update(ctx context.Context, user *model.User) error {
	builder := sq.Update(tableNameUser).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, user.Name).
		Set(emailColumn, user.Email).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: user.Id})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableNameUser).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) AddRole(ctx context.Context, userId int64, roleId int64) error {
	builder := sq.Insert(tableNameUserRole).
		Columns(userIdColumn, roleIdColumn).
		Values(userId, roleId).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "user_repository.AddRole",
		QueryRaw: query,
	}
	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) RemoveRole(ctx context.Context, userId int64, roleId int64) error {
	builder := sq.Delete(tableNameUserRole).
		Where(sq.And{sq.Eq{userIdColumn: userId}, sq.Eq{roleIdColumn: roleId}}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "user_repository.RemoveRole",
		QueryRaw: query,
	}
	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
