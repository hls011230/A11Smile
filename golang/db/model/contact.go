package model

type User_solidity struct {
	Proute  string  `json:"route"`//用户上传证书路径
	User_  string  `json:"user_"` //用户地址
	Soliciter_ string  `json:"solicitor_"` //征求者地址
}

type Soliciter_solidity struct {
	Soliciter_ string  `json:"solicitor_"` //征求者地址
	HospitalName string `json:"hospital_name"` //征求医院名字
	User_  string  `json:"user_"` //用户地址
	Proute  string  `json:"route"`//用户上传证书路径
	Min_ int64  `json:"min"`  //征求者设置奖励最小金额
	Max int64  `json:"max"`  //征求者设置奖励最大金额
	Ercnum_ int64  `json:"enum"`  //征求者转账金额
	Account int64 `json:"account"` //征求数量
	MedicalName string `json:"medical_name"` //征求者征求病历名称
	MedicalNeed string `json:"medical_need"`//征求者病历要求
	RequirementDescription string `json:"requirement_description"`//征求者需求描述
	Whether bool `json:"whether"`  //征求者是否通过奖励
	Department string `json:"department"`
}

type AllPeople_solidity struct {
	People_  string  `json:"people_"` //所有人
}


// 合约部署者
var Deployer = struct {
	Address string
	PrivateKey string
}{
	Address: "0xA62C4A01eBCF3728624274C364DCa0eDD75634b7",
	PrivateKey:"1888236f4354da8514397142a62fa21f0f2555d9719136206ced8c888b127cb6",
}
