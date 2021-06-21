package database

type PostDB interface {
	Open() error
	Close() error
}
