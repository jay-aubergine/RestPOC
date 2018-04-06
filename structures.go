package main

//RequestMapForm  ..
type RequestMapForm struct {
	Cmd    string  `json:"cmd"`
	Recid  int     `json:"recid",omitempty`
	Name   string  `json:"name",omitempty`
	Record Records `json:"record",omitempty`
}

//ResponseData holds the data that resides inside response.
type ResponseData struct {
	Recid int    `json:"recid"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
	Sdate string `json:"sdate"`
}

//ResponseStruct  ..
type ResponseStruct struct {
	Status       string         `json:"status"`
	Total        int            `json:"total",omitempty`
	ResponseData []ResponseData `json:"records",omitempty`
}

//ResponseStruct  ..
type SingleResponseStruct struct {
	Status       string       `json:"status"`
	ResponseData ResponseData `json:"record",omitempty`
}

//RestResponseStruct  ..
type RestResponseStruct struct {
	ResponseData []ResponseData `json:"users",omitempty`
}
