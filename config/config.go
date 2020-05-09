package config

type Server struct {
	AppName string
	Version string

	System SystemConfig `mapstructure:"system" yaml:"system"`
	MySQL  MySQLConfig  `mapstructure:"mysql" yaml:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis" yaml:"redis"`
	Log    LogConfig    `mapstructure:"log" yaml:"log"`
	SQLite SQLiteConfig `mapstructure:"sqlite" yaml:"sqlite"`
	JWT    JWTConfig    `yaml:"jwt"`
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
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Password string `mapstructure:"password" yaml:"password"`
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
