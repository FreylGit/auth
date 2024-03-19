package refreshToken

import (
	"context"
	"github.com/FreylGit/auth/internal/model"
	"github.com/FreylGit/auth/internal/repository"
	"github.com/FreylGit/auth/internal/repository/refreshToken/converter"
	modelRepo "github.com/FreylGit/auth/internal/repository/refreshToken/model"
	"github.com/FreylGit/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName   = "refresh_tokens"
	idColumn    = "id"
	tokenColumn = "token"
	expColumn   = "exp"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.RefreshTokenRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, token string) (*model.RefreshToken, error) {
	tokenByte := []byte(token)
	_ = tokenByte
	builder := sq.Select(idColumn, tokenColumn, expColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{tokenColumn: token}).
		Limit(1)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "token_repository.Get",
		QueryRaw: query,
	}
	var tokenModel modelRepo.RefreshToken
	err = r.db.DB().ScanOneContext(ctx, &tokenModel, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToRefreshTokenFromRepo(tokenModel), nil
}

func (r *repo) Create(ctx context.Context, token *model.RefreshToken) (int64, error) {
	tokenByte := []byte(token.Token)
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(tokenColumn, expColumn).
		Values(tokenByte, token.Exp).
		Suffix("RETURNING id")
	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	q := db.Query{
		Name:     "token_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *repo) Update(ctx context.Context, token *model.RefreshToken) error {
	tokenByte := []byte(token.Token)
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{tokenColumn: token.Token}).
		Set(tokenColumn, tokenByte).
		Set(expColumn, token.Exp)
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "token_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, token string) error {
	tokenByte := []byte(token)
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{tokenColumn: tokenByte})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "token_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
