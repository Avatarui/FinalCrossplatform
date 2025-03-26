package main

import (
	"FinalCrossplatform/database"
	"FinalCrossplatform/routes"
	"log"
)

func main() {
	database.ConnectDB() // เชื่อมต่อฐานข้อมูล
	r := routes.SetupRouter()

	log.Println("🚀 Server started on port 8080")
	r.Run(":8080")
}
