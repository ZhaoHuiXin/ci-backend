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
)

var cfgFile string

//func Init() {
//	flag.StringVar(&cfgFile, "config", os.Getenv("CONFIG_PATH"), "mysql uri")
//}

func main() {
	cfgFile = os.Getenv("CONFIG_PATH")
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	f, _ := ioutil.ReadFile(cfgFile)

	//any approach to require this configuration into your program.
	err := viper.ReadConfig(bytes.NewBuffer(f))
	if err != nil {
		fmt.Println(err)
	}
	viper.AutomaticEnv()
	// this would be "steve"
	conf, _ := configs.GetConf()
	db, err := utils.GetDB(conf.DBType, conf.DBConnectString)
	if err != nil{
		log.Fatalf("error while GetDB, %s \n", err.Error())
	}
	fmt.Println(db)
	fmt.Println(db.DB())
	rds, err := utils.GetRDS(conf.RDBConnectString)
	if err != nil{
		log.Fatalf("error while GetRDAS, %s \n", err.Error())
	}
	fmt.Println(rds.String())

	wand := utils.NewWand(db, rds)
	daoImpl.NewUserImpl(wand)
}