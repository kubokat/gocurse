package shorter

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CreateSLink(origin string) string {
	hash := sha256.Sum256([]byte(origin))
	linkHash := hex.EncodeToString(hash[:])[:8]

	return fmt.Sprintf("%s?hash=%s", "test.link", linkHash)

}
