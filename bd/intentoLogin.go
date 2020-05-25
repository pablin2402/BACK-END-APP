package bd

import (
	"github.com/snetwork/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin login que ve el usuario y compara contraseñas*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)

	if encontrado == false {
		return usu, false
	}

	//Una contraseña encriptada y otra no
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	//desencripta
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
