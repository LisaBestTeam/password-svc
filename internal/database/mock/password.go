package mock

import "github.com/lisabestteam/password-svc/internal/database"

type Passwords struct {
	SelectBySenderFn   func(address string) ([]*database.Password, error)
	SelectByReceiverFn func(address string) ([]*database.Password, error)
	CreatePasswordFn   func(password database.Password) error
	MaxIdFn            func() (uint64, error)
}

func (m *Passwords) SelectBySender(address string) ([]*database.Password, error) {
	if m != nil && m.SelectBySenderFn != nil {
		return m.SelectBySenderFn(address)
	}
	return nil, nil
}

func (m *Passwords) SelectByReceiver(address string) ([]*database.Password, error) {
	if m != nil && m.SelectByReceiverFn != nil {
		return m.SelectByReceiverFn(address)
	}
	return nil, nil
}

func (m *Passwords) CreatePassword(password database.Password) error {
	if m != nil && m.CreatePasswordFn != nil {
		return m.CreatePasswordFn(password)
	}
	return nil
}

func (m *Passwords) MaxId() (uint64, error) {
	if m != nil && m.MaxIdFn != nil {
		return m.MaxIdFn()
	}
	return 0, nil
}

func (m *Passwords) New() database.Passwords {
	return m
}

func (m *Passwords) Pagination(pagination database.Pagination) database.Passwords {
	return m
}
