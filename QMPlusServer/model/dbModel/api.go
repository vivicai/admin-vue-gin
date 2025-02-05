package dbModel

import (
	"github.com/jinzhu/gorm"
	"main/controller/servers"
	"main/init/qmsql"
	"main/model/modelInterface"
)

type Api struct {
	gorm.Model  `json:"-"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

func (a *Api) CreateApi() (err error) {
	findOne := qmsql.DEFAULTDB.Where("path = ?", a.Path).Find(&Menu{}).Error
	if findOne != nil {

	} else {
		err = qmsql.DEFAULTDB.Create(a).Error
	}
	return err
}

func (a *Api) DeleteApi() (err error) {
	err = qmsql.DEFAULTDB.Where("path = ?", a.Path).Delete(a).Delete(&ApiAuthority{}).Error
	return err
}

func (a *Api) EditApi() (err error) {
	err = qmsql.DEFAULTDB.Update(a).Update(&Authority{}).Error
	return err
}

// 分页获取数据  需要分页实现这个接口即可
func (a *Api) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []Api
		err = db.Find(&apiList).Error
		return err, apiList, total
	}
}
