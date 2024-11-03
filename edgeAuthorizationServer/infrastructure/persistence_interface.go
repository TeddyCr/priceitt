package infrastructure

type IPersistence interface {
	// Initialize initializes the repository
	Initialize(config map[string]interface{}) (*IPersistence, error)
}
