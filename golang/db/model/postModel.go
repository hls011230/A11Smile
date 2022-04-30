package model

type PostEmail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type PostUserFile struct {
	FileName string `json:"file_name"`
}

type PostDetails struct {
	Address string 	`json:"address"`
	MedicalName string `json:"medical_name"`
}