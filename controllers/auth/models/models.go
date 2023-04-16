package AuthModels

type TokenAccess struct {
	TokenAccess string `json:"token_access"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
}

type Response struct {
	Data *Data `json:"data"`
	Error *Error `json:"error"` 
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Data TokenAccess


