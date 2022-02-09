package datasources

type Reader interface {
	Aliases() ([]byte, error)
	Get(flag string) ([]byte, error)
	Find() ([]byte, error)
	Info(name string) ([]byte, error)
	List() ([]byte, error)
	Networks() ([]byte, error)
}
