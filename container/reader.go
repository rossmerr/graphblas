package container

// Reader for importing data into a container
type Reader interface {
	Read() (record []string, err error)
}
