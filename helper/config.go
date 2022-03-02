package helper

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	AppName            string `mapstructure:"app_name"`
	ServiceAddress     string `mapstructure:"service_address"`
	ServicePort        string `mapstructure:"service_port"`
	ServiceMode        string `mapstructure:"service_mode"`
	DbType             string `mapstructure:"db_type"`
	MongoDbHost        string `mapstructure:"mongo_db_host"`
	MongoDbName        string `mapstructure:"mongo_db_name"`
	MongoDbUserName    string `mapstructure:"mongo_db_username"`
	MongoDbPassword    string `mapstructure:"mongo_db_password"`
	MongoDbPort        string `mapstructure:"mongo_db_port"`
	MongoDbAuthDb      string `mapstructure:"mongo_db_auth_db"`
	ServiceName        string `mapstructure:"service_name"`
	LogFile            string `mapstructure:"log_file"`
	LogDir             string `mapstructure:"log_dir"`
	ExternalConfigPath string `mapstructure:"external_config_path"`
	PageLimit          string `mapstructure:"page_limit"`
}

// var (
// 	mode               string
// 	dbHost             string
// 	user               string
// 	dbName             string
// 	password           string
// 	address            string
// 	port               string
// 	externalConfigPath string
// )

// func LoadConfig() (string, string, string, string, string, string, string, string) {
// 	flag.StringVar(&mode, "mode", Config.Mode, "application mode, either dev or prod")
// 	flag.StringVar(&dbHost, "dbhost", Config.DbHost, "database host")
// 	flag.StringVar(&user, "user", Config.User, "access user")
// 	flag.StringVar(&dbName, "dbname", Config.DbName, "database name")
// 	flag.StringVar(&password, "password", Config.Password, "access password")
// 	flag.StringVar(&address, "address", Config.Address, "local host")
// 	flag.StringVar(&port, "port", Config.Port, "application ports")
// 	flag.StringVar(&externalConfigPath, "external_config_path", Config.DbName, "external config path")

// 	flag.Parse()
// 	for i, val := range flag.Args() {
// 		os.Args[i] = val
// 	}

// 	return mode, dbHost, user, dbName, password, address, port, externalConfigPath
// }

func loadEnv(path string) (config ConfigStruct, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("buycoin")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return ConfigStruct{}, err
	}
	err = viper.Unmarshal(&config)
	return
}

func returnConfig() ConfigStruct {
	config, err := loadEnv(".")
	if err != nil {
		log.Println(err)
	}
	if config.ExternalConfigPath != "" {
		viper.Reset()
		config, err = loadEnv(config.ExternalConfigPath)
		if err != nil {
			log.Println(err)
		}
	}
	return config
}

var Config = returnConfig()
