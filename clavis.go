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

// Client struct responsible for storing config and managing storage
type Client struct {
	// Client config
	Config  Config
	storage map[string]valorem
}

type valorem struct {
	value      string
	expiration int64
}

func DefaultConfig() Config {
	return Config{
		Threads:           2,
		EnableInDisk:      true,
		InDiskPath:        "clavis/",
		EnableInMemory:    true,
		DefaultExpiration: -1,
	}
}

func NewClavis(conf Config) Client {
	return Client{
		Config:  conf,
		storage: make(map[string]valorem),
	}
}

// Set a key value
func (c *Client) Set(key, value string, unixExp time.Duration) error {
	if key == "" {
		return fmt.Errorf("Missing key")
	}

	if value == "" {
		return fmt.Errorf("Missing value")
	}

	if _, ok := c.storage[key]; ok {
		return fmt.Errorf("%s already exists", key)
	}

	c.storage[key] = valorem{
		value:      value,
		expiration: int64(unixExp),
	}

	return nil
}

// Retrieve value
func (c *Client) Get(key string) (string, error) {
	if key == "" {
		return "", fmt.Errorf("Missing key")
	}

	val, ok := c.storage[key]

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
