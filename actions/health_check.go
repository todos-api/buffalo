package actions

import (
    "net/http"

	"github.com/gobuffalo/buffalo"
)

// HealthCheckHandler default implementation.
func HealthCheckHandler(c buffalo.Context) error {
	// requests are wrapped in a transaction, and will fail
	// before we get here if the database connection isn't
	// healthy
	return c.Render(http.StatusOK, r.String("success"))
}

