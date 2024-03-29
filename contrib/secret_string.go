package contrib

import (
	"crypto/sha256"
	"fmt"
)

// SecretString return sha256 hash of target string
func SecretString(target string, salt string) string {
	value := sha256.Sum256([]byte(salt + target))
	return fmt.Sprintf("%x", value)
}
