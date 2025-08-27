package services

import (
	"fmt"
	"time"

	"github.com/morkid/gocache"
	cache_redis "github.com/morkid/gocache-redis/v8"
	"github.com/morkid/paginate"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB Main database connection
var DB *gorm.DB

// PG Pagination library
var PG *paginate.Pagination

// InitDatabase initialize database connection
func InitDatabase() {
	if nil == DB {
		db := dbConnect()
		if nil != db {
			DB = db

			var cache *gocache.AdapterInterface
			cacheSeconds := viper.GetInt64("CACHE_TTL_SECONDS")

			if nil != REDIS && cacheSeconds > 0 {
				cache = cache_redis.NewRedisCache(cache_redis.RedisCacheConfig{
					Client:    REDIS,
					ExpiresIn: time.Duration(cacheSeconds) * time.Second,
				})
			}

			PG = paginate.New(&paginate.Config{
				CacheAdapter:         cache,
				FieldSelectorEnabled: true,
			})
			dbMigrate()
		}
	}
}

func dbConnect() *gorm.DB {
	logLevel := logger.Info

	switch viper.GetString("ENVIRONMENT") {
	case "staging":
		logLevel = logger.Error
	case "production":
		logLevel = logger.Silent
	}

	config := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("DB_TABLE_PREFIX"),
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	// Create DSN (Data Source Name)

	// dsn := "root:root@tcp(localhost:8889)/hrisapps_employee?charset=utf8mb4&parseTime=True&loc=Local"
	dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn2), &config)

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	if db != nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(1)
		sqlDB.SetConnMaxLifetime(time.Second * 5)
	}

	return db
}

// dbMigrate performs the database migration and seeding
func dbMigrate() {
	// db := dbConnect()
	// if db != nil && len(migrations.ModelMigrations) > 0 {
	// 	if err := db.AutoMigrate(migrations.ModelMigrations...); err != nil {
	// 		panic(err)
	// 	}

	// 	tx := db.Begin()
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			tx.Rollback()
	// 		}
	// 	}()

	// 	if err := migrations.InitialSeeds(tx); err != nil {
	// 		tx.Rollback()
	// 		panic(err) // Or handle error accordingly
	// 	}

	// 	seeds := migrations.DataSeeds()
	// 	for _, seed := range seeds {
	// 		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(seed).Error; err != nil {
	// 			tx.Rollback()
	// 			panic(err) // Or handle error accordingly
	// 		}
	// 	}

	// 	if err := tx.Commit().Error; err != nil {
	// 		tx.Rollback()
	// 		panic(err) // Or handle error accordingly
	// 	}

	// 	db.Migrator().DropTable("schema_migration")

	// 	sqlDB, _ := db.DB()
	// 	defer sqlDB.Close()
	// }
}
