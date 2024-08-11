package db

type DBAction[T any] interface {
	Insert(item T) error
	Select(id int) (T, error)
	Update(id int, item T) error
	Delete(id int) error
}
