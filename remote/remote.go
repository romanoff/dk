package remote

type Remote interface {
	Push(filepath, destination string) error
	Pull(filepath, destination string) error
	FilesList() ([]string, error)
}
