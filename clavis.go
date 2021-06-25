package clavis

import (
	"os"
	"time"
)

// Config for Clavis Client
type Config struct {
	// Amount of threads to run Clavis
	// Default value: 1
	Threads uint

	// If Clavis will save the value in files
	// Default value: true
	EnabledInDisk bool

	// Which folder the files will be saved
	// Default value: ./clavis/
	InDiskPath string

	// If Clavis will save the value in memory
	// Default value: true
	EnabledInMemory bool

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
	key        string
	keyEncrypt string
	value      string
	expiration int64
}

// Return default config for Clavis client
func DefaultConfig() Config {
	return Config{
		Threads:           2,
		EnabledInDisk:     true,
		InDiskPath:        "./clavis/",
		EnabledInMemory:   true,
		DefaultExpiration: -1,
	}
}

// Return default client for given Config
// In case of nil field, will be assumed the default value
func NewClavis(conf Config) (Client, errat) {
	if err := os.MkdirAll(conf.InDiskPath, 0777); err != nil {
		return Client{}, ErratUnknown(err.Error())
	}

	return Client{
		Config:  conf,
		storage: make(map[string]valorem),
	}, NilErrat()
}

// Set a key value
func (c *Client) Set(key, value string, unixExp time.Duration) errat {
	if key == "" {
		return ErratMissing("key")
	}

	if value == "" {
		return ErratMissing("value")
	}

	if _, ok := c.storage[key]; ok {
		return ErratExists(key)
	}

	val := valorem{
		key:        key,
		value:      value,
		expiration: int64(unixExp),
	}

	if c.Config.EnabledInMemory {
		c.storage[key] = val
	}

	if c.Config.EnabledInDisk {
		val.createFile(c.Config.InDiskPath)
	}

	return NilErrat()
}

// Register a file with both value and expiration
func (v *valorem) createFile(path string) errat {
	v.keyEncrypt = Sha1Encrypt(v.key)

	filePath := fmt.Sprintf("%s%s", path, v.keyEncrypt)

	_, err := os.Stat(filePath)

	if !os.IsNotExist(err) {
		return ErratExists(v.key)
	}

	f, err := os.Create(filePath)

	defer f.Close()

	if err != nil {
		return ErratUnknown(err.Error())
	}

	_, err = f.Write(v.setFileContent())

	if err != nil {
		return ErratUnknown(err.Error())
	}

	return NilErrat()
}

// Return expiration:value to store
func (v *valorem) setFileContent() []byte {
	content := fmt.Sprintf("%d:%s", v.expiration, v.value)

	return []byte(content)
}

// Retrieve value
func (c *Client) Get(key string) (string, errat) {
	if key == "" {
		return "", ErratMissing("key")
	}

	val, ok := c.storage[key]

	if !ok {
		return "", ErratNotFound(key)
	}

	exp := val.expiration

	if exp == -1 {
		return val.value, NilErrat()
	}

	if (exp - time.Now().Unix()) < 0 {
		return "", ErratExpired(key)
	}

	return val.value, NilErrat()
}

// Retrieve value and remove it
func (c *Client) Pop(key string) (string, errat) {
	val, err := c.Get(key)

	if !err.Nil() {
		return "", err
	}

	delete(c.storage, key)

	return val, NilErrat()
}
