package configstruct

type UserInfo struct {
	userName string
	password string
	token    string
}
type Users struct {
	Credentials []UserInfo
}
