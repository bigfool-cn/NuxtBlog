package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	App       App
	AdminUser AdminUser
	Db        Db
	Cache     Cache
}


type App struct {
	Host      string
	Port      int
	JwtSecret string
	LogPath   string
}

type AdminUser struct {
	UserName  string
	Pwd       string
}

type Cache struct {
	Host   string
	Pwd    string
	Port   int
	Select int
}

type Db struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	DbType    string
}

var Configs *Config

func init()  {
	file, err := os.Open("./configs/config.yaml")
	if err != nil {
		panic(err)
	}

	bytes,err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes,&Configs)
	if err != nil {
		panic(err)
	}
}




