package generators

// Manager will manage a single type of generator.
// Will allow id generation for multiple types of records.
type Manager[T comparable] struct {
	generators map[string]Generator[T]
}

// NewManager creates a new Manager.
func NewManager[T comparable]() *Manager[T] {
	return &Manager[T]{
		generators: make(map[string]Generator[T]),
	}
}

// GetIdForRecordType returns the next id for the given record type.
func (m *Manager[T]) GetIDForRecordType(recordType string) T {
	generator, ok := m.generators[recordType]
	if !ok {

	}
	return generator.GetID()
}
