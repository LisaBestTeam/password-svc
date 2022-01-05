package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	passwordTable         = "password"
	senderAddressColumn   = "sender_address"
	receiverAddressColumn = "receiver_address"
)

func NewPassword(db *pgdb.DB) database.Passwords {
	return &passwords{
		sql: sq.Select("*").From(passwordTable).PlaceholderFormat(sq.Dollar),
		ins: sq.Insert(passwordTable).PlaceholderFormat(sq.Dollar),
		db:  db,
	}
}

type passwords struct {
	sql sq.SelectBuilder
	ins sq.InsertBuilder
	db  *pgdb.DB
}

func (p passwords) New() database.Passwords {
	return &passwords{
		sql: sq.Select("*").From(passwordTable),
		ins: sq.Insert(passwordTable).PlaceholderFormat(sq.Dollar),
		db:  p.db,
	}
}

func (p passwords) SelectBySender(address string) ([]*database.Password, error) {
	var passwordList []*database.Password

	request := p.sql.Where(sq.Eq{senderAddressColumn: address})

	err := p.db.Select(&passwordList, request)
	return passwordList, errors.Wrap(err, "failed to select from tests")
}

func (p passwords) SelectByReceiver(address string) ([]*database.Password, error) {
	var passwordList []*database.Password

	request := p.sql.Where(sq.Eq{receiverAddressColumn: address})

	err := p.db.Select(&passwordList, request)
	return passwordList, errors.Wrap(err, "failed to select from tests")
}

func (p passwords) CreatePassword(password database.Password) error {
	clauses := structs.Map(password)

	stmt := p.ins.SetMap(clauses)
	err := p.db.Exec(stmt)

	return err
}

func (p passwords) MaxId() (*uint64, error) {
	stmt := sq.Select("max(id)").From(passwordTable)

	var result *uint64
	err := p.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return result, errors.Wrap(err, "failed to get max id from tests")
}

func (p passwords) Pagination(pagination database.Pagination) database.Passwords {
	p.sql = p.sql.Offset(pagination.Page * pagination.Limit).Limit(pagination.Limit)
	return p
}
