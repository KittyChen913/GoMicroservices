package models

type ELKLogDetail struct {
	ServiceName string
	Request     RequestDetail
	Response    ResponseDetail
	Message     string
}

type RequestDetail struct {
	Method string
	Url    string
	Ip     string
	Body   string
}

type ResponseDetail struct {
	Status int
	Body   string
}
