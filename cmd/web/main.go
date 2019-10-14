package main

import (
	"bytes"
	"ci-backend/configs"
	daoImpl "ci-backend/internal/dao/impl"
	"ci-backend/internal/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/viper"
	"flag"
	"ci-backend/internal/web"
)

var (
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "config", os.Getenv("CONFIG_PATH"), "config file path")
}

func main() {
	flag.Parse()
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	f, _ := ioutil.ReadFile(cfgFile)

	err := viper.ReadConfig(bytes.NewBuffer(f))
	if err != nil {
		fmt.Println(err)
	}
	viper.AutomaticEnv()

	conf, _ := configs.GetConf()

	// connect to database
	db, err := utils.GetDB(conf.DBType, conf.DBConnectString)
	if err != nil{
		log.Fatalf("error while GetDB, %s \n", err.Error())
	}
	// connect to redis
	rds, err := utils.GetRDS(conf.RDBConnectString)
	if err != nil{
		log.Fatalf("error while GetRDAS, %s \n", err.Error())
	}

	wand := utils.NewWand(db, rds)
	daoImpl.NewUserImpl(wand)

	router := web.InitRouter()
	router.Run(":8090")
}