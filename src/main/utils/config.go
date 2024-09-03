package utils

var (
	// DbType 数据库类型 暂时只支持mysql
	DbType string
	// DbHost 数据库主机
	DbHost string
	// DbPort 数据库端口
	DbPort string
	// DbUser 数据库用户名
	DbUser string
	// DbPassword 数据库密码
	DbPassword string
	// DbName 数据库名称
	DbName string

	TestDefaultImg string = "https://img.xuanpt2.com/27.png"
)

func init() {
	DbTestInit()
}

func DbTestInit() {
	DbType = "mysql"
	DbHost = "localhost"
	DbPort = "3306"
	DbUser = "root"
	DbPassword = "wangxing"
	DbName = "slogv2"
}
