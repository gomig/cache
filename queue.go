package cache

// Queue interface for queue drivers.
type Queue interface {
	// Push queue new item
	Push(value any) error
	// Pull read first queue item
	Pull() (*string, error)
}
