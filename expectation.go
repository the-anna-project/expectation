// Package expectation implements Expectation to provide transportable
// expectation primitives.
package expectation

// Config represents the configuration used to create a new expectation object.
type Config struct {
	// Settings.
	Output string
}

// DefaultConfig provides a default configuration to create a new expectation
// object by best effort.
func DefaultConfig() Config {
	return Config{
		// Settings.
		Output: "",
	}
}

// New creates a new configured expectation object.
func New(config Config) (Expectation, error) {
	newExpectation := &expectation{
		// Settings.
		output: config.Output,
	}

	return newExpectation, nil
}

type expectation struct {
	// Settings.
	output string
}

func (e *expectation) Output() string {
	return e.output
}
