package token

type Request struct {
	User_name string
	Password  string
}

type ReadResponseToken struct {
	User_name string
	Token     string
}

type FitterUpdateToken struct {
	LogoutRequest string
}

type FitterReadToken struct {
	TokenLogout string
}
