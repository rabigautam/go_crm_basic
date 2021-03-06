package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rabigautam/go_crm_basic/database"
	"github.com/rabigautam/go_crm_basic/lead"
)

func setupRoutes(app *fiber.App) {

	app.Post("/api/v1/lead", lead.NewLead)
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Get("api/v1/lead/:id",lead.DeleteLead)

}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connected database")

	}
	fmt.Println("Connection opened to the database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}
func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(4040)
	defer database.DBconn.Close()

}
