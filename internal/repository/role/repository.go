package role

import (
	"context"
	"github.com/FreylGit/auth/internal/client/db"
	"github.com/FreylGit/auth/internal/model"
	"github.com/FreylGit/auth/internal/repository"
	modelRepo "github.com/FreylGit/auth/internal/repository/role/model"
	"github.com/FreylGit/auth/internal/repository/user/converter"
	sq "github.com/Masterminds/squirrel"
	"strings"
)

const (
	tableName       = "roles"
	idColumn        = "id"
	nameUpperColumn = "name_upper"
	nameLowerColumn = "name_lower"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.RoleRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Role, error) {
	builder := sq.Select(idColumn, nameUpperColumn, nameLowerColumn).
		From(tableName).Where(sq.Eq{idColumn: id}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "role_repository.Get",
		QueryRaw: query,
	}
	var role modelRepo.Role

	err = r.db.DB().ScanOneContext(ctx, &role, q, args...)
	if err != nil {
		return nil, err
	}
	roleModel := converter.ToRoleFromRepo(role)

	return &roleModel, nil
}

func (r *repo) GetByName(ctx context.Context, name string) (*model.Role, error) {
	nameLower := strings.ToLower(name)
	builder := sq.Select(idColumn, nameUpperColumn, nameLowerColumn).
		From(tableName).
		Where(sq.Eq{nameLowerColumn: nameLower}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "role_repository.GetByName",
		QueryRaw: query,
	}
	var role modelRepo.Role

	err = r.db.DB().ScanOneContext(ctx, &role, q, args...)
	if err != nil {
		return nil, err
	}
	roleModel := converter.ToRoleFromRepo(role)

	return &roleModel, nil
}

func (r *repo) Create(ctx context.Context, role *model.Role) (int64, error) {
	nameLower := strings.ToLower(role.Name)
	nameUpper := strings.ToUpper(role.Name)
	builder := sq.Insert(tableName).
		Columns(nameLowerColumn, nameUpperColumn).
		Values(nameLower, nameUpper).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	q := db.Query{
		Name:     "role_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Update(ctx context.Context, role *model.Role) error {
	nameLower := strings.ToLower(role.Name)
	nameUpper := strings.ToUpper(role.Name)

	builder := sq.Update(tableName).
		Where(sq.Eq{idColumn: role.Id}).
		Set(nameLowerColumn, nameLower).
		Set(nameUpperColumn, nameUpper).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "role_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "role_repository.Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return err
	}

	return nil
}
