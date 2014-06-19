package remote

type Remote interface {
	Setup(*Config)
	Push(filepath, destination string) error
	Pull(filepath, destination string) error
	Remove(filepath string) error
	FilesList() ([]string, error)
}
