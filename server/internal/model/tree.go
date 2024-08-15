package model

type Tree interface {
	GetKey() string
	GetType() string
	Add(value int64)
	Exists(value int64) bool
	Remove(value int64)
	GetAll() []int64
}
