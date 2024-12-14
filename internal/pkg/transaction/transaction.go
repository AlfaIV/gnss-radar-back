package transaction

//
//import (
//	"context"
//	"github.com/Gokert/gnss-radar/internal/store"
//
//	"github.com/jackc/pgconn"
//	"github.com/jackc/pgx/v4"
//)
//
////
//type Tx interface {
//	Tx() pgx.Tx
//
//	Begin(ctx context.Context) (Tx, error)
//	BeginFunc(ctx context.Context, f func(Tx) error) error
//	Commit(ctx context.Context) error
//	Rollback(ctx context.Context) error
//	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
//	LargeObjects() pgx.LargeObjects
//	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
//	Conn() *pgx.Conn
//
//	store.Executor
//}
//
//type txImpl struct {
//	tx pgx.Tx
//	*store.ExecutorImpl
//}
//
//// NewTx возвращает реализацию интерфейса Tx.
//func NewTx(tx pgx.Tx) *txImpl {
//	return &txImpl{
//		tx:           tx,
//		ExecutorImpl: store.NewExecutor(tx),
//	}
//}
//
//func (t *txImpl) Begin(ctx context.Context) (Tx, error) {
//	tx, err := t.tx.Begin(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return NewTx(tx), nil
//}
//
//func (t *txImpl) Tx() pgx.Tx {
//	return t.tx
//}
//
//func (t *txImpl) BeginFunc(ctx context.Context, f func(Tx) error) error {
//	return t.tx.BeginFunc(ctx, func(tx pgx.Tx) error {
//		return f(NewTx(tx))
//	})
//}
//
//func (t *txImpl) Commit(ctx context.Context) error {
//	return t.tx.Commit(ctx)
//}
//
//func (t *txImpl) Rollback(ctx context.Context) error {
//	return t.tx.Rollback(ctx)
//}
//
//func (t *txImpl) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
//	return t.tx.SendBatch(ctx, b)
//}
//
//func (t *txImpl) LargeObjects() pgx.LargeObjects {
//	return t.tx.LargeObjects()
//}
//
//func (t *txImpl) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
//	return t.tx.Prepare(ctx, name, sql)
//}
//
//func (t *txImpl) Conn() *pgx.Conn {
//	return t.tx.Conn()
//}
//
//var _ Tx = (*txImpl)(nil)
