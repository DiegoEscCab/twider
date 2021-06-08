package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass string) (string, error) {
	costo := 8 //el costo elevará el numero de encriptaciones en este caso encriptará la contraseña hasta 8 veces, haciendola segura ante hackeos
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err

}
