package postgres

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lisabestteam/password-svc/internal/database"
)

const (
	passwordTable         = "password"
	senderAddressColumn   = "sender"
	receiverAddressColumn = "receiver"
)

func NewPassword(db *sqlx.DB) database.Passwords {
	return &passwords{
		sql: sq.Select("*").From(passwordTable).PlaceholderFormat(sq.Dollar),
		ins: sq.Insert(passwordTable).PlaceholderFormat(sq.Dollar),
		db:  db,
	}
}

type passwords struct {
	sql sq.SelectBuilder
	ins sq.InsertBuilder
	db  *sqlx.DB
}

func (p passwords) New() database.Passwords {
	return &passwords{
		sql: sq.Select("*").From(passwordTable).PlaceholderFormat(sq.Dollar),
		ins: sq.Insert(passwordTable).PlaceholderFormat(sq.Dollar),
		db:  p.db,
	}
}

func (p passwords) SelectBySender(address string) ([]*database.Password, error) {
	passwordList := make([]*database.Password, 0)

	query, args := p.sql.Where(sq.Eq{senderAddressColumn: address}).MustSql()
	err := p.db.Select(&passwordList, query, args...)

	return passwordList, err
}

func (p passwords) SelectByReceiver(address string) ([]*database.Password, error) {
	var passwordList []*database.Password

	query, args := p.sql.Where(sq.Eq{receiverAddressColumn: address}).MustSql()
	err := p.db.Select(&passwordList, query, args...)

	return passwordList, err
}

func (p passwords) CreatePassword(password database.Password) error {
	clauses := structs.Map(password)
	query, args, err := p.ins.SetMap(clauses).ToSql()
	if err != nil {
		return err
	}

	_, err = p.db.Exec(query, args...)

	return err
}

func (p passwords) MaxId() (uint64, error) {
	query, args := sq.Select("max(id)").From(passwordTable).PlaceholderFormat(sq.Dollar).MustSql()

	var result *uint64
	err := p.db.Get(&result, query, args...)

	return *result, err
}

func (p passwords) Pagination(pagination database.Pagination) database.Passwords {
	p.sql = p.sql.Offset(pagination.Page * pagination.Limit).Limit(pagination.Limit)
	return p
}
