package profile

type mySelf struct {
	name      string
	photo     string
	email     string
	age       int8
	phone     string
	status    string
	education education
}

type education struct {
	name  string
	major string
}
