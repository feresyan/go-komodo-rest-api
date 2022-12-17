package repository

type JWTRepository interface {
	CheckPassword(userPassword, loginPassword string) error
	GenerateJWT(email, role string) (string, error)
}
