package transaction

import (
	"context"
	"github.com/FreylGit/auth/internal/client/db"
	"github.com/FreylGit/auth/internal/client/db/pg"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type manager struct {
	db db.Transactor
}

func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}

func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, f db.Handler) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		f(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}
	ctx = pg.MakeContextTx(ctx, tx)
	// настраиваем фунцию отсрочки для отката или комита транзакции
	defer func() {
		// восстанавливаемся после паники
		if r := recover(); r != nil {
			err = errors.Errorf("panic revovered")
		}

		// откатываем транзакцию если ошибка произошла
		if err != nil {
			if errRollBack := tx.Rollback(ctx); errRollBack != nil {
				err = errors.Wrapf(err, "errRoleBack")
			}

			return
		}
		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "tx commit failed")
			}
		}
	}()

	// Выполните код внутри транзакции
	// Если функция терпит неудачу, возвращаем ошибку и фуркция отсрочки выполняет откат
	// или в проивном случае транзация комитится
	if err = f(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return nil
}
