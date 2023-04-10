package main

import (
	"database/sql"
	"os"
	"github.com/DamianSformo/proyecto-entregable-final/cmd/server/handler"
	"github.com/DamianSformo/proyecto-entregable-final/internal/patient"
	"github.com/DamianSformo/proyecto-entregable-final/internal/dentist"
	"github.com/DamianSformo/proyecto-entregable-final/pkg/dentist_store"

	"github.com/DamianSformo/proyecto-entregable-final/pkg/patient_store"
	
	"github.com/DamianSformo/proyecto-entregable-final/pkg/appointment_store"
	"github.com/DamianSformo/proyecto-entregable-final/internal/appointment"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	os.Setenv("TOKEN", "appointments")

	db, err := sql.Open("mysql", "root:43824382@tcp(localhost:3306)/dbcharta")
	if err != nil{
		panic(err)
	}

	dentistStorage := dentist_store.NewSqlStore(db)
	dentistRepository := dentist.NewRepository(dentistStorage)
	dentistService := dentist.NewService(dentistRepository)
	dentistHandler := handler.NewDentistHandler(dentistService)


	patientStorage := patient_store.NewSqlStore(db)
	patientRepository := patient.NewRepository(patientStorage)
	patientService := patient.NewService(patientRepository)
	patientHandler := handler.NewPatientHandler(patientService)


	appointmentStorage := appointment_store.NewSqlStore(db)
	appointmentRepository := appointment.NewRepository(appointmentStorage)
	appointmentService := appointment.NewService(appointmentRepository, patientRepository, dentistRepository)
	appointmentProductHandler := handler.NewAppointmentHandler(appointmentService)


	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	
	dentists := r.Group("/dentists")
	{
		dentists.GET(":id", dentistHandler.GetDentistById())
		dentists.POST("", dentistHandler.PostDentist())
		dentists.PATCH(":id", dentistHandler.Patch())
		dentists.PUT(":id", dentistHandler.Put())
		dentists.DELETE(":id", dentistHandler.DeleteDentist())
	}

	patients := r.Group("/patients")
	{
		patients.GET(":id", patientHandler.GetPatientByID())
		patients.GET("/dni/:dni", patientHandler.GetPatientByDni())
		patients.POST("", patientHandler.PostPatient())
		//patients.PATCH(":id", patientProductHandler.Patch())
		patients.PUT(":id", patientHandler.PutPatient())
		patients.DELETE(":id", patientHandler.DeletePatient())
	}

	appointments := r.Group("/appointments")
	{
		//appointments.GET(":id", appointmentProductHandler.GetPatientByID())
		appointments.GET("/dni/:dni", appointmentProductHandler.GetAppointmentByDni())
		appointments.POST("", appointmentProductHandler.PostAppointment())
		appointments.POST("dniAndLicense/:dni/:license", appointmentProductHandler.PostAppointmentByDniAndLicense())
		//appointments.PATCH(":id", appointmentProductHandler.Patch())
		//appointments.PUT(":id", appointmentProductHandler.PutPatient())
		//appointments.DELETE(":id", appointmentProductHandler.DeletePatient())
	}

	r.Run(":8080")
}
