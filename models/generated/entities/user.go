package entities

type User struct {
	BaseEntity
	Email string `json:"email"`
	Image50 string `json:"image50"`
	AuthenticationMechanism interface{} `json:"authenticationMechanism"` // map to oneOf
}
