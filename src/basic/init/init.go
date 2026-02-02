package init

func init() {
	InitViper()
	InitMysql()
	InitRedis()
	InitEs()
	LogInit()
}
