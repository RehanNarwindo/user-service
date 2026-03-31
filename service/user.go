package service

import "user-service/config"

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GetPublicMessage() string {
	return "User service jalan"
}

func GetProfile(claims map[string]interface{}) map[string]interface{} {
	email, _ := claims["email"].(string)

	user, err := GetUserByEmail(email)
	if err != nil {
		return map[string]interface{}{
			"message": "User tidak ditemukan",
		}
	}

	return map[string]interface{}{
		"message": "Profile dari DB",
		"user":    user,
	}
}

func GetUserByEmail(email string) (*User, error) {
	var user User

	query := `SELECT id, email, first_name FROM users WHERE email = $1`

	err := config.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
	)

	if err != nil {
		return nil, err
	}
	
	return &user, nil
}



func GetAllUsers() ([]User, error) {
	rows, err := config.DB.Query(`SELECT id, email, first_name FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}