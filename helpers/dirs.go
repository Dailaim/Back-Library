package helpers

import "os"

func DirSafe(directorio string) {
	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.Mkdir(directorio, 0755)
		if err != nil {
			panic(err)
		}
	}
}