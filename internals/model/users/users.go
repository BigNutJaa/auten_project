package users

type Request struct {
	User_name  string
	Password   string
	First_name string
	Last_name  string
	Email      string
}

type FitterReadUsers struct {
	User_name string
	Password  string
}

type FitterListUsers struct {
	User_name  string
	First_name string
	Email      string
	Page       int64
	Per_page   int64
}

type ReadResponseUsers struct {
	Id         int32
	User_name  string
	Password   string
	First_name string
	Last_name  string
	Email      string
	Role_code  string
}

type ListResponseUsers struct {
	Id         int32
	User_name  string
	First_name string
	Email      string
}
