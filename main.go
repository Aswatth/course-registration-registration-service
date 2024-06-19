package main

import (
	"context"
	"course-registration-system/registration-service/controllers"
	"course-registration-system/registration-service/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongo_database := new(services.MongoDatabase)

	mongo_database.Connect(context.Background(), os.Getenv("CONNECTION_STRING"))

	defer mongo_database.Disconnect(context.Background())

	err = mongo_database.Ping(context.Background())

	if err != nil {
		fmt.Println("Unable to connect to mongo db")
	}

	mongo_database.SetDatabase(os.Getenv("DATABASE"))

	offered_course_service := new(services.OfferedCourseService)
	offered_course_service.Init(*mongo_database)

	offered_course_controller := new(controllers.OfferedCourseController)
	offered_course_controller.Init(*offered_course_service)

	registered_course_service := new(services.RegisteredCourseService)
	registered_course_service.Init(*mongo_database)

	registered_course_controller := new(controllers.RegisteredCourseController)
	registered_course_controller.Init(*registered_course_service)

	server := gin.Default()

	base_path := server.Group("")
	offered_course_controller.RegisterRoutes(base_path)
	registered_course_controller.RegisterRoutes(base_path)

	server.Run(":" + os.Getenv("PORT"))
}
