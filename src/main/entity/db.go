package entity

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"slogv2/src/main/utils"
	"time"
)

var Db *gorm.DB
var err error

func DbInit() {

	dbHost := utils.DbHost
	//try to get host from env
	//尝试从环境变量中获取host
	if _host := os.Getenv("DB_HOST"); _host != "" {
		dbHost = _host
	}

	//dsn Data Source Name 数据库连接字符串
	//格式：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//    用户名:密码@tcp(主机:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		dbHost,
		utils.DbPort,
		utils.DbName,
	)

	log.Println("dsn: ", dsn)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		//禁用默认事务
		SkipDefaultTransaction: true,
		//日志级别：Silent
		Logger: logger.Default.LogMode(logger.Silent),
		//表名策略：单数表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("db init failed: ", err)
	}

	err := Db.AutoMigrate(&User{}, &Category{}, &Article{}, &Options{}, &Comment{}, &Relationship{})
	if err != nil {
		log.Fatal("db migrate failed: ", err)
	}

	sqlDB, _ := Db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(100)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(1000)

	// SetConnMaxLifetime 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(30 * time.Second)
}
