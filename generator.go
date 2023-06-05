package generators

import "context"

// MaxChanSize is the maximum size of the channel used to generate ids.
const MaxChanSize = 1000

// Generator is an interface that all generators must implement.
type Generator[T comparable] interface {
	// GetID returns the next id.
	GetID() T
	// IsStateful returns true if the generator is stateful, ie if the ids are sequential ints.
	IsStateful() bool
	// SetState sets the state of the generator if it is stateful.
	SetState(s T)
	// GenerateIDs is a long-running method that generates ids that will be fetched by the GetID method.
	GenerateIDs(ctx context.Context)
}
