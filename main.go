package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"myIndia/state"
)

func india(app *fiber.App) {
	app.Get("/allstates", state.AllStates)
	app.Get("/alldistricts", state.AllDistricts)
	app.Get("/allcities", state.AllCities)
}

func main() {
	app := fiber.New()

	dsn := "root:Niran@123@tcp(localhost:3306)/universe?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("Error: ", err)
		panic("Failed to connect to the database")
	}
	log.Infof("Database: ", db)
	log.Info("Database connected succesffuly")

	result, err := db.Query("SELECT id, name FROM state")

	if err != nil {
		panic(err)
	}

	for result.Next() {

		var id int
		var name string

		// The result object provided Scan  method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		err = result.Scan(&id, &name)

		// handle error
		if err != nil {
			panic(err)
		}

		fmt.Printf("Id: %d Name: %s\n", id, name)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Infof("Error while closing the db: ", err)
		}
	}(db)

	india(app)

	err = app.Listen(":3000")
	if err != nil {
		return
	}

}
