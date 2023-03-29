package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abai/organizer/models"
	"github.com/abai/organizer/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateEvent(context *fiber.Ctx) error {
	timeTableItem := models.TimeTableItem{}

	err := context.BodyParser(&timeTableItem)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&timeTableItem).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to create an event"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "event has been added"})

	return nil
}

func (r *Repository) DeleteEvent(context *fiber.Ctx) error {
	timeTableItem := models.TimeTableItem{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})

		return nil
	}

	err := r.DB.Delete(timeTableItem, id)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete event",
		})

		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "event deleted successfully",
	})

	return nil
}

func (r *Repository) GetEvent(context *fiber.Ctx) error {
	id := context.Params("id")
	eventModel := &models.TimeTableItem{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})

		return nil
	}

	fmt.Println("ID is ", id)

	err := r.DB.Where("id = ?", id).First(eventModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get the event",
		})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "event fetched successfully",
		"data":    eventModel,
	})

	return nil
}

func (r *Repository) GetUser(context *fiber.Ctx) error {
	id := context.Params("id")
	eventModel := &models.User{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})

		return nil
	}

	fmt.Println("ID is ", id)

	err := r.DB.Where("id = ?", id).First(eventModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get the user",
		})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user fetched successfully",
		"data":    eventModel,
	})

	return nil
}

func (r *Repository) GetUserEvents(context *fiber.Ctx) error {
	userTimeTableItems := &[]models.TimeTableItem{}

	err := r.DB.Find(userTimeTableItems).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get events"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "events fetched successfully",
		"data":    userTimeTableItems,
	})

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/create_event", r.CreateEvent)
	api.Delete("/delete_event/:id", r.DeleteEvent)
	api.Get("/get_event/:id", r.GetEvent)
	api.Get("/get_user/:id", r.GetUser)
	api.Get("/get_user_events/:id", r.GetUserEvents)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateUser(db)
	if err != nil {
		log.Fatal("could not migrate User")
	}

	err = models.MigrateTimeTableItem(db)
	if err != nil {
		log.Fatal("could not migrate TimeTableItem")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
