package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/rabigautam/go_crm_basic/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func NewLead(c *fiber.Ctx) {
	db := database.DBconn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}
func GetLeads(c *fiber.Ctx) {
	var lead []Lead
	db := database.DBconn
	db.Find(&lead)
	c.JSON(lead)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBconn
	var lead Lead

	db.Find(&lead, id)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBconn
	var lead Lead
	db.Delete(&lead, id)
	c.JSON(lead)

}
