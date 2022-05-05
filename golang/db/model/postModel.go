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

type PostExamine struct {
	Certificate string `json:"certificate"`
	MedicalName string `json:"medical_name"`
	Whether bool `json :"whether"`
	Address string `json:"address"`
	Ercnum int64 `json:"ercnum"`
}

type PostSubmitCertificate struct {
	Certificate string `json:"certificate_"`
	Soliciter   string  `json:"soliciter_"`
	MedicalName string `json:"medical_name_"`

}