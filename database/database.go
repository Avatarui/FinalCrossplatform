package database

// import (
// 	// "FinalCrossplatform/models"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )
import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn = "cp_65011212186:65011212186@csmsu@tcp(202.28.34.197:3306)/cp_65011212186?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() {

	// เชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ ไม่สามารถเชื่อมต่อฐานข้อมูลได้:", err)
	}

	log.Println("✅ เชื่อมต่อฐานข้อมูลสำเร็จ!")

	// Auto Migrate (สร้างตารางถ้ายังไม่มี)
	// db.AutoMigrate(&models.User{})

	// กำหนดตัวแปร global
	DB = db
}
