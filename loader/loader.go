package loader

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//Config Start
type Config struct {
	Debug    bool
	Database struct {
		Driver     string
		Connection string
		Debug      bool
	}
	Dbcb struct {
		Driver     string
		Connection string
		Bucket     string
		UserName   string
		Password   string
	}

	Host string
	Port string
	Cors struct {
		Hosts []string
	}
}

var Conf Config

//LoadConfig Start
func LoadConfig() {

	args := os.Args[1:]
	env := ""

	if len(args) > 0 {
		env = args[0]
	}

	var confile string

	fmt.Println("---- config :" + env)

	if env == "prod" {
		confile = "./config/config.prod.yml"
	} else {
		confile = "./config/config.dev.yml"
	}

	file, err := os.Open(confile)
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	defer file.Close()
	viper.MergeConfig(file)
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	fmt.Printf("--- Load config from %s ---\n", confile)
}
