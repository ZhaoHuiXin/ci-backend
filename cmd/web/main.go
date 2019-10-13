package main

import (
	"ci-backend/cmd/web/cmd"
	"ci-backend/configs"
	"ci-backend/internal/utils"

	"github.com/prometheus/common/log"
	"ci-backend/internal/dao/impl"
)

var cfgFile string

func main() {
	cmd.Execute()

	conf, _ := configs.GetConf()

	db, err := utils.GetDB(conf.DBType, conf.DBConnectString)
	if err != nil {
		log.Fatalf("error while GetDB, %s", err.Error())
	}

	rds, err := utils.GetRDS(conf.RDBConnectString)
	if err != nil {
		log.Fatalf("error while GetRDS, %s", err.Error())
	}

	wand := utils.NewWand(db, rds)

	daoImpl.NewUserImpl(wand)
}
