package source

type DataSource interface {
	List() ([]interface{}, error)
	Create(item interface{}) error
}
