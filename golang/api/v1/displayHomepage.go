package v1

import (
	"A11Smile/eth"
	"A11Smile/eth/gen"
)

func DisplayHomepage()([]gen.UploadMedicalrecords_gainergainer_upMedicalInformation,error) {
	res,err := eth.Ins.SeeGainerMedicalInformationsName(nil)
	return res,err
}