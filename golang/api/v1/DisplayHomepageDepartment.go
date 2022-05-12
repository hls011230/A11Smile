package v1

import (
	"A11Smile/db/model"
	"A11Smile/eth"
	"A11Smile/eth/gen"
)

func DisplayHomepageDepartment(Meical *model.POSTDepartment) ([]gen.UploadMedicalrecords_gainergainer_upMedicalInformation,error)  {
	res,err := eth.Ins.SeeGainerMedicalInformationsNameDepartment(nil,Meical.Department)
	if err!=nil{
		return nil,err
	}
	return res,nil

}
