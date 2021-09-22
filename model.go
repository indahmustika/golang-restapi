package main

type Users struct{
	Username	string `form:"username" json:"username"`
	Password	string `form:"password" json:"password"`
}

type Response struct{
	Status		int 	`json:"status"`
	Message		string	`json:"message"`
	Data		[]Users
}