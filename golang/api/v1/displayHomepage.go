package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"A11Smile/eth/gen"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func DisplayHomepage()([]gen.UploadMedicalrecords_gainergainer_upMedicalInformation,error) {
	res,err := eth.Ins.SeeGainerMedicalInformationsName(nil)
	if err != nil {
		return nil, err
	}
	return res,nil
}

func ShowDetailsPage(details model.PostDetails) (interface{},error) {
	res,err := eth.Ins.SeeGainerMedicalInformations(nil,common.HexToAddress(details.Address),details.MedicalName)
	if err != nil {
		return res, err
	}

	var resume string
	DB := db.Get()
	DB.Raw("select resume from gainers where block_address = ?" , details.Address).Find(&resume)

	DetailsPage := struct {
		Min                       *big.Int
		Max                       *big.Int
		Medicalrecordrequirements string
		Requirementdescription    string
		Resume string
	}{
		Min: res.Min,
		Max: res.Max,
		Medicalrecordrequirements: res.Medicalrecordrequirements,
		Requirementdescription: res.Requirementdescription,
		Resume: resume,
	}

	return DetailsPage,nil
}