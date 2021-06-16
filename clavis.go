package clavis

import "fmt"

// Stored values
var values map[string]string = make(map[string]string)

// Set a key value
func Set(key, value string) error {
	if key == "" {
		return fmt.Errorf("Missing key")
	}

	if value == "" {
		return fmt.Errorf("Missing value")
	}

	if _, ok := values[key]; ok {
		return fmt.Errorf("%s already exists", key)
	}
	values[key] = value

	return nil
}

