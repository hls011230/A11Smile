package v1

import (
	"A11Smile/eth"
	"A11Smile/eth/gen"
)

func DisplayHomepage()([]gen.UploadMedicalrecordsgainer_upMedicalInformation,error) {
	res,err := eth.Ins.SeeGainerMedicalInformationsName(nil)



	return res,err
}