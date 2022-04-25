package v1

import (
	"A11Smile/db"
)

//编辑征求者资料
func GainerEditData(gid int, resume string) (err error) {
	var sqlName = `UPDATE gainers set resume = ? where id = ?`
	DB := db.Get()
	DB.Exec(sqlName, resume, gid)

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
	DB.Table("gainers").Select("gainer_authentication.enterprise_name,gainers.block_address,gainers.resume").Joins("left join gainer_authentication on gainer_authentication.gid = gainers.id where gainers.id = ?",id).Scan(&gainer)

	return gainer
}
