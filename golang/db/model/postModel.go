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
	HospitalName string `json:"hospital_name"`
}


type PostCertificate struct {
	ArrayMedicalHistory []string `json:"array_medical_history"`
	ArrayMedicalExaminationReport []string `json:"array_medical_examination_report"`
}

type PostCertificateHash struct {
	Serial string `json:"serial"`
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


type  PostETHforAS struct {
	AddETH string ` json:"addeth"`
	RedETH string `json:"redeth"`
	Quantity int `json:"quantity"`
}

type PostWarehouse struct {
	Medical string `json:"medical"`
	User string `json:"user"`
	Soliciter string `json:"soliciter"`
}