package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
)

//编辑征求者资料
func GainerEditData(gid int, resume string) (err error) {
	var sqlName = `UPDATE gainers set resume = ? where id = ?`
	DB := db.Get()
	DB.Exec(sqlName, resume, gid)

	return
}

func GainerDataSeeUpdate(gid int) (gainer model.GainerAuthentication, err error) {
	var sqlName = `SELECT id,gid,reg_num,serial,enterprise_name FROM gainer_authentication WHERE gid = ?`
	// 查询数据
	DB := db.Get()
	DB.Select(sqlName,gid).Find(&gainer)

	return
}

//查询企业信息
func GainerAuthenticationSee(id int) interface{}{
	gainer := struct {
		EnterpriseName string `json:"enterprise_name"`
		Block_address  string `json:"block_address"`
		Resume         string `json:"resume"`
	}{}
	DB := db.Get()
	DB.Table("gainers").Select("gainer_authentication.enterprise_name,gainers.block_address,gainers.resume").Joins("left join gainer_authentication on gainer_authentication.gid = gainers.id",id).Scan(&gainer)

	return gainer
}
