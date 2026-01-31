package config

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}
type Log struct {
	Level         string
	Path          string
	MaxSize       int
	MaxBackups    int
	MaxAge        int
	Compress      bool
	Format        string
	ConsoleEnable bool
}
type Nacos struct {
	Host        string
	Port        int
	NamespaceId string
	DataId      string
	Group       string
}
type AppConfig struct {
	Nacos
	Mysql
	Redis
	Log
}
