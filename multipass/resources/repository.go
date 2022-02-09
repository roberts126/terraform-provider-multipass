package resources

type Repository interface {
	Alias(instance, command, alias string) ([]byte, error)
	Aliases() ([]byte, error)
	Delete(name string) ([]byte, error)
	Get(flag string) ([]byte, error)
	Find() ([]byte, error)
	Info(name string) ([]byte, error)
	Launch(image, name string, args ...string) ([]byte, error)
	List() ([]byte, error)
	Networks() ([]byte, error)
	Set(flag, value string) ([]byte, error)
	Unalias(alias string) ([]byte, error)
}
