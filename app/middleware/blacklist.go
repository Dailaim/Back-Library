package middleware

import "sync"

var tokenList []string
var mu sync.Mutex

func AddToken(token string) error {
    mu.Lock()
    defer mu.Unlock()

    // Verificar si el token ya está en la lista negra
    if IterTokenList(token){
        return nil
    }

    // Agregar el token a la lista negra
    tokenList = append(tokenList, token)
    return nil
}

func IsTokenBlacklisted(token string) bool {
    mu.Lock()
    defer mu.Unlock()

    return IterTokenList(token)
}



func IterTokenList(token string) bool {
    for _, t := range tokenList {
        if t == token {
            // El token ya está en la lista negra
            return true
        }
    }
    // El token no está en la lista negra
    return false
}