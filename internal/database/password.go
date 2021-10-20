package database

type Passwords interface {
	New() Passwords

	SelectBySender(address string) ([]*Password, error)
	SelectByReceiver(address string) ([]*Password, error)

	CreatePassword(password Password) error

	MaxId() (uint64, error)

	Pagination(pagination Pagination) Passwords
}

type Pagination struct {
	Page  uint64 `schema:"page"`
	Limit uint64 `schema:"limit"`
}

type Password struct {
	Id               uint64 `db:"id" structs:"id" jsonapi:"primary,password"`
	HashOfFile       string `db:"hash_of_file" structs:"hash_of_file" jsonapi:"attr,hash_of_file"`
	SenderAddress    string `db:"sender_address" structs:"sender_address" jsonapi:"attr,sender_address"`
	ReceiverAddress  string `db:"receiver_address" structs:"receiver_address" jsonapi:"attr,receiver_address"`
	EncryptsPassword string `db:"encrypts_password" structs:"encrypts_password" jsonapi:"attr,encrypts_password"`
	TypeOfFile       string `db:"type_of_file" structs:"type_of_file" jsonapi:"attr,type_of_file"`
}
