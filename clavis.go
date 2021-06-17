package clavis

import (
	"fmt"
	"time"
)

// Config for Clavis Client
type Config struct {
	// Amount of threads to run Clavis
	// Default value: 1
	Threads uint

	// If Clavis will save the value in files
	// Default value: true
	EnableInDisk bool

	// Which folder the files will be saved
	// Default value: clavis/
	InDiskPath string

	// If Clavis will save the value in memory
	// Default value: true
	EnableInMemory bool

	// Expiration time for every stored value
	// Default value(for unset expiration): true
	DefaultExpiration int64
}

type valorem struct {
	value      string
	expiration int64
}

// Stored values
var values map[string]valorem = make(map[string]valorem)

// Set a key value
func Set(key, value string, unixExp time.Duration) error {
	if key == "" {
		return fmt.Errorf("Missing key")
	}

	if value == "" {
		return fmt.Errorf("Missing value")
	}

	if _, ok := values[key]; ok {
		return fmt.Errorf("%s already exists", key)
	}

	values[key] = valorem{
		value:      value,
		expiration: int64(unixExp),
	}

	return nil
}

// Retrieve value
func Get(key string) (string, error) {
	if key == "" {
		return "", fmt.Errorf("Missing key")
	}

	val, ok := values[key]

	if !ok {
		return "", fmt.Errorf("%s not found", key)
	}

	exp := val.expiration

	if exp == -1 {
		return val.value, nil
	}

	if (exp - time.Now().Unix()) < 0 {
		return "", fmt.Errorf("%s is expirated", key)
	}

	return val.value, nil
}
