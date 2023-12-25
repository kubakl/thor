package errors

// Error is an interface representing a general compilation error.
type Error interface {
	// Default returns the error in the default format.
	Default() string
	// Verbose returns the error in the verbose format.
	Verbose() string
}
