package trees

type Tree interface {
	GetKey() string
	GetType() string
	Add(value int64, shouldReport bool)
	Exists(value int64) bool
	Remove(value int64, shouldReport bool)
	GetAll() []int64
}
