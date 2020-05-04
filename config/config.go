package config


type Server struct{

	AppName string
	Version string
	
	System SystemConfig
	MySQL MySQLConfig
	Redis RedisConfig
	Log LogConfig

}
type SystemConfig struct{
	Env string
	IP string
	Port int
	DB string
}

type MySQLConfig struct{
	IP string
	Port int
	User string
	PassWord string
	Database string
}

type RedisConfig struct{
	IP string 
	Port int
}


type LogConfig struct{
	Prefix string
	LogFile bool
	Stdout string
	File string
}