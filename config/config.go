package config

import "time"

type Server struct {
	AppName string
	Version string

	System     SystemConfig `mapstructure:"system" yaml:"system"`
	MySQL      MySQLConfig  `mapstructure:"mysql" yaml:"mysql"`
	Redis      RedisConfig  `mapstructure:"redis" yaml:"redis"`
	Log        LogConfig    `mapstructure:"log" yaml:"log"`
	SQLite     SQLiteConfig `mapstructure:"sqlite" yaml:"sqlite"`
	JWT        JWTConfig    `yaml:"jwt"`
	AppSetting AppConfig    `mapstructure:"appconfig" yaml:"appconfig"`
	SMSConfig  SMSConfig    `mapstructure:"sms" yaml:"sms"`
}
type SystemConfig struct {
	Env  string
	IP   string
	Port int
	DB   string
}

type MySQLConfig struct {
	Username     string `mapstructure:"username" yaml:"username"`
	Password     string `mapstructure:"password" yaml:"password"`
	Addr         string `mapstructure:"addr" yaml:"addr"`
	Database     string `mapstructure:"database" yaml:"database"`
	Config       string `mapstructure:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" yaml:"log-mode"`
}

type RedisConfig struct {
	Addr        string        `mapstructure:"addr" yaml:"addr"`
	Password    string        `mapstructure:"password" yaml:"password"`
	MaxIdle     int           `mapstructure:"maxidle" yaml:"maxidle"`
	MaxActive   int           `mapstructure:"maxactive" yaml:"maxactive"`
	IdleTimeout time.Duration `mapstructure:"idletimeout" yaml:"idletimeout"`
}

type LogConfig struct {
	Prefix  string
	Logfile bool
	Stdout  string
	File    string
}

type SQLiteConfig struct {
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Path     string `mapstructure:"path" yaml:"path"`
	Config   string `mapstructure:"config" yaml:"config"`
	LogMode  bool   `mapstructure:"log-mode" yaml:"log-mode"`
}

type JWTConfig struct {
	SigningKey string `yaml:"signing-key"`
}

type AppConfig struct {
	PrefixUrl       string   `mapstructure:"prefixurl" yaml:"prefixurl"`
	ImageSavePath   string   `mapstructure:"imagesavepath" yaml:"imagesavepath"`
	ImageMaxSize    int64    `mapstructure:"imagemaxsize" yaml:"imagemaxsize"`
	ImageAllowExts  []string `mapstructure:"imageallowexts" yaml:"imageallowexts"`
	VideoSavePath   string   `mapstructure:"videosavepath" yaml:"videosavepath"`
	VideoMaxSize    int64    `mapstructure:"videomaxsize" yaml:"videomaxsize"`
	VideoAllowExts  []string `mapstructure:"videoallowexts" yaml:"videoallowexts"`
	RuntimeRootPath string   `mapstructure:"runtimerootpath" yaml:"runtimerootpath"`
	PageSize        int      `mapstructure:"pagesize" yaml:"pagesize"`
}

type SMSConfig struct {
	AccessKey string `mapstructure:"accesskey" yaml:"accesskey"`
	Secret    string `mapstructure:"secret" yaml:"secret"`
	SignName  string `mapstructure:"signname" yaml:"signname"`
	TemplateCode  string `mapstructure:"templatecode" yaml:"templatecode"`
}
