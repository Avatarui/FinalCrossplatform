package database

// import (
// 	// "FinalCrossplatform/models"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )
import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn = "cp_65011212186:65011212186@csmsu@tcp(202.28.34.197:3306)/cp_65011212186?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() {
	// Open database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Optional: Add GORM configuration
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Get underlying SQL DB for connection pooling config
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get sql.DB: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	log.Println("✅ Successfully connected to database!")
}

// Optional: Function to check database connection
func CheckConnection() bool {
	if DB == nil {
		return false
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return false
	}
	if err := sqlDB.Ping(); err != nil {
		return false
	}
	return true
}
