package main

import (
	"database/sql"
	"fmt"
	"os"
	"proyecto_final_go/cmd/handler"
	"proyecto_final_go/internal/repository"
	"proyecto_final_go/internal/service"
	"proyecto_final_go/pkg/middleware"
	storeAppointment "proyecto_final_go/pkg/store/appointment"
	storeDentist "proyecto_final_go/pkg/store/dentist"
	storePatient "proyecto_final_go/pkg/store/patient"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:33060)/my_db")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error")
	}

	storageDentists := storeDentist.NewSqlStore(db)
	storagePatients := storePatient.NewSqlStore(db)
	storageAppointments := storeAppointment.NewSqlAppointmentStore(db)

	repoDentists := repository.NewDentistRepository(storageDentists)
	serviceDentists := service.NewDentistService(repoDentists)
	handlerDentists := handler.NewDentistHandler(serviceDentists)

	repoPatients := repository.NewPatientRepository(storagePatients)
	servicePatients := service.NewPatientService(repoPatients)
	handlerPatients := handler.NewPatientHandler(servicePatients)

	repoAppointments := repository.NewAppointmentRepository(storageAppointments)
	serviceAppointments := service.NewAppointmentService(repoAppointments)
	handlerAppointments := handler.NewAppointmentHandler(serviceAppointments)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dentists := r.Group("/dentists")
	{
		dentists.POST("", middleware.Authentication(), handlerDentists.Post())
		dentists.GET(":id", handlerDentists.GetByID())
		dentists.PUT(":id", middleware.Authentication(), handlerDentists.Put())
		dentists.PATCH(":id", middleware.Authentication(), handlerDentists.Patch())
		dentists.DELETE(":id", middleware.Authentication(), handlerDentists.Delete())
		dentists.GET("", handlerDentists.GetAll())
	}

	patients := r.Group("/patients")
	{
		patients.POST("", middleware.Authentication(), handlerPatients.Post())
		patients.GET(":id", handlerPatients.GetByID())
		patients.PUT(":id", middleware.Authentication(), handlerPatients.Put())
		patients.PATCH(":id", middleware.Authentication(), handlerPatients.Patch())
		patients.DELETE(":id", middleware.Authentication(), handlerPatients.Delete())
		patients.GET("", handlerPatients.GetAll())
	}

	appointments := r.Group("/appointments")
	{
		appointments.POST("", middleware.Authentication(), handlerAppointments.Post())
		appointments.GET(":id", handlerAppointments.GetByID())
		appointments.PUT(":id", middleware.Authentication(), handlerAppointments.Put())
		appointments.PATCH(":id/description", middleware.Authentication(), handlerAppointments.PatchDescription())
		appointments.DELETE(":id", middleware.Authentication(), handlerAppointments.Delete())
		appointments.POST("/dni-license", middleware.Authentication(), handlerAppointments.PostByDNIAndLicese())
		appointments.GET("/patient/:dni", handlerAppointments.GetByPatientDNI())
		appointments.GET("", handlerAppointments.GetAll())
	}

	r.Run()

}

// * http://localhost:8080/docs/index.html --> Para leer documentación de la API.
