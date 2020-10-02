package api

//Credentials respresents the user login object
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
