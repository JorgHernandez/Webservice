package usuario

// User : struct of the users table
type User struct {
	ID       int    `json:"id"`
	Edad     int    `json:"edad"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
}
