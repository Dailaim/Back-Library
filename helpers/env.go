package helpers

import "os"


//Comprueba variables de entorno
func Env(env string, def string) string {
	variable := os.Getenv(env)
	if len(variable) == 0 {
		return def
	}
	return variable
}