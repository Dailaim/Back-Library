package helpers

import "os"

// Se asegura de que las carpetas necesarias existan
func DirSafe(directorio string) {
	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.Mkdir(directorio, 0755)
		if err != nil {
			panic(err)
		}
	}
}