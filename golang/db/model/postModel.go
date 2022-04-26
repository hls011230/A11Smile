package model

type PostEmail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type PostUserFile struct {
	FileName string `json:"file_name"`
}