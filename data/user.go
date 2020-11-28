package data

type User struct {
	Password string
	Name     string
	Email    string
}

func Encrypt(input string) string {
	return input
}

func UserByEmail(email string) (result *User, err error) {
	result = &User{
		Password: "",
		Name:     "chengchao",
		Email:    email,
	}

	return
}

func (u *User) CreateSession() *Session {
	return &Session{}

}
