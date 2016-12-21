package expectation

import (
	"encoding/json"
)

// Expectation represents a description of what output is to be expected when
// requesting calculations by providing some input.
type Expectation interface {
	json.Marshaler
	json.Unmarshaler
	// Output returns the configured output of the expectation. This output
	// represents the output which is expected to be calculated by the neural
	// network.
	Output() string
}
