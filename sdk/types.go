package sdk

type logger interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
}

// Add your shared types here
// Example:
// type User struct {
//     ID    string `json:"id"`
//     Email string `json:"email"`
// }
