package vo

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRes struct {
	Logined bool   `json:"logined"`
	Token   string `json:"token"`
}

type UserRegisterReq struct {
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterRes struct {
	Signed bool   `json:"signed"`
	Token  string `json:"token"`
}

type UserGetInfoRes struct {
	NickName string `json:"nickname"`
	UserID   string `json:"userID"`
	Email    string `json:"email"`
	IconPath string `json:"iconPath"`
}

type UserChangeNickNameReq struct {
	NickName string `json:"nickname"`
}

type UserChangeNickNameRes struct {
	Changed bool `json:"changed"`
}
