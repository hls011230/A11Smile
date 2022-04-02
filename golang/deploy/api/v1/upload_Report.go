package v1

import "os"

func UploadReport(path string) (bool,error){
	//判断文件夹是否存在
	_,err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
