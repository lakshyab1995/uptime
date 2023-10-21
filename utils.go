package url_short

import "github.com/oklog/ulid/v2"

// generateShortKey generate unique short key of 6 characters
func generateShortKey() string {
	return ulid.Make().String()[:6]
}
