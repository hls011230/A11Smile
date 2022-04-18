package model

import "github.com/ethereum/go-ethereum/common"

type User_AddMedicalInformation struct {
	Min_ int64  //征求者设置奖励最小金额
	Max int64   //征求者设置奖励最大金额
	Ercnum_ int64  //征求者转账金额
	Proute  string  //用户上传图片路径
	MedicalName string //征求者征求病历名称
	MedicalNeed string //征求者病历要求
	RequirementDescription string //征求者需求描述
	User_ common.Address //用户地址
	Soliciter_ common.Address //征求者地址
	People_  common.Address //所有人
	Whether bool  //征求者是否通过奖励
}

