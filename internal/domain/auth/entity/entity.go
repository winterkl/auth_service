package auth_entity

type User struct {
	ID       int `bun:",pk,autoincrement"`
	Login    string
	Password string
}
