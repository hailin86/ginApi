package common

import (
	"fmt"
	"ginApi/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"os"
)

var (
	Enforcer *casbin.Enforcer
)

//创建casbin的enforcer
func InitCasbinEnforcer() error {
	a, err := gormadapter.NewAdapterByDB(model.DB)
	if err != nil {
		fmt.Println("err:",err)
	}
	dir, _ := os.Getwd()
	modelPath := dir + "/conf/rbac_model.conf"
	//fmt.Println("modelPath:"+modelPath)
	var errC error
	Enforcer, errC = casbin.NewEnforcer(modelPath, a)
	if errC != nil {
		fmt.Println("errC:",errC)
		return errC
	} else {
		Enforcer.LoadPolicy()
		Enforcer.EnableLog(false)
		return nil
	}
}
