package service

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

var users = []User{
	{ID: "1", Email: "tes@mail.com", Name: "Rehan"},
	{ID: "2", Email: "user@mail.com", Name: "User"},
}

func GetPublicMessage() string {
	return "User service jalan"
}

func GetProfile(claims map[string]interface{}) map[string]interface{} {
	email, _ := claims["email"].(string)
	print("email")
	for _, user := range users {
		if user.Email == email {
			return map[string]interface{}{
				"message": "Profile ditemukan",
				"user":    user,
			}
		}
	}

	return map[string]interface{}{
		"message": "User tidak ditemukan",
	}
}