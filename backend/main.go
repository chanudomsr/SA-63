package main

import (
  "context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/chanudomsr/app/controllers"
	_ "github.com/chanudomsr/app/docs"
	"github.com/chanudomsr/app/ent"
)
// Customers  defines the struct for the customers
type Customers struct {
	Customer []Customer
}

// Customer  defines the struct for the customer
type Customer struct {
	NAME string
	AGE int
	EMAIL string
	USERNAME string
	PASSWORD string
}

// Employees  defines the struct for the employees
type Employees struct {
	Employee []Employee
}

// Employee  defines the struct for the employee
type Employee struct {
	EMPLOYEENAME     string
	EMPLOYEEPASSWORD string
}

// Rooms  defines the struct for the rooms
type Rooms struct {
	Room []Room
}

// Room  defines the struct for the room
type Room struct {
	ROOMNUMBER string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewCheckoutController(v1, client)
	controllers.NewCustomerController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewRoomController(v1, client)
	
	// Set Customers Data
	customers := Customers{
		Customer: []Customer{
			Customer{"Marty Mcfly",20,"Marty85@gmail.com","marty07","1590"},
			Customer{"Cristiano Ronaldo",34,"RonaldoCT@gmail.com","ronaldo08","3570"},
			Customer{"Lionel Messi",31,"MessiLN@gmail.com","messi09","0258"},
			Customer{"David Beckham",45,"BeckhamDV@gmail.com","beckham10","1563"},
			Customer{"Harry Potter",25,"PotterHR@gmail.com","potter11","4852"},
			Customer{"Ron Weasley",25,"Weasley@gmail.com","weasley12","4752"},
			Customer{"Hermione Granger",25,"GrangerHR@gmail.com","granger13","1236"},
		},
	}

	for _, cus := range customers.Customer {
		client.Customer.
			Create().
			SetNAME(cus.NAME).
			SetAGE(cus.AGE).
			SetEMAIL(cus.EMAIL).
			SetUSERNAME(cus.USERNAME).
			SetPASSWORD(cus.PASSWORD).
			Save(context.Background())
	}

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"EM1", "em1231"},
			Employee{"EM2", "em1232"},
			Employee{"EM3", "em1233"},
			Employee{"EM4", "em1234"},
			Employee{"EM5", "em1235"},
			Employee{"EM6", "em1236"},
			Employee{"EM7", "em1237"},
			Employee{"EM8", "em1238"},
			Employee{"EM9", "em1239"},
			Employee{"EM10", "em12310"},
		},
	}

	for _, em := range employees.Employee {
		client.Employee.
			Create().
			SetEMPLOYEENAME(em.EMPLOYEENAME).
			SetEMPLOYEEPASSWORD(em.EMPLOYEEPASSWORD).
			Save(context.Background())
	}

	// Set Rooms Data
	rooms := Rooms{
		Room: []Room{
			Room{"A101"},
			Room{"A102"},
			Room{"A103"},
			Room{"A104"},
			Room{"A105"},
			Room{"A106"},
			Room{"A107"},
			Room{"A108"},
			Room{"A109"},
			Room{"A110"},
			Room{"A111"},
			Room{"A112"},
			Room{"A113"},
			Room{"A113"},
			Room{"A114"},
			Room{"A115"},
			Room{"A116"},
			Room{"A117"},
			Room{"A118"},
			Room{"A119"},
			Room{"A120"},
		},
	}

	for _, r := range rooms.Room {
		client.Room.
			Create().
			SetROOMNUMBER(r.ROOMNUMBER).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
