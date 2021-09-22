package database

type Passwords interface {
	New() Passwords

	SelectBySender(address string) ([]Password, error)
	SelectByReceiver(address string) ([]Password, error)

	CreatePassword(password Password) error

	MaxId() (uint64, error)
}

type Password struct {
	Id               uint64 `db:"id" structs:"id" json:"id"`
	HashOfFile       string `db:"hash_of_file" structs:"hash_of_file" json:"hash_of_file"`
	SenderAddress    string `db:"sender_address" structs:"sender_address" json:"sender_address"`
	ReceiverAddress  string `db:"receiver_address" structs:"receiver_address" json:"receiver_address"`
	EncryptsPassword string `db:"encrypts_password" structs:"encrypts_password" json:"encrypts_password"`
}
