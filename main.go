package Zopsmart_pro

import (
	"api/handler"
	"api/store"
	"gofr.dev/pkg/gofr"
)
func main() {
	app := gofr.New()

	s := store.New()
	h := handler.New(s)

	app.GET("/car/{id}", h.Get)
	app.POST("/car", h.Create)
	app.PUT("/car/{id}", h.Update)
	app.DELETE("/car/{id}", h.Delete)

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	
	_, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Connected to the database! wollllllllllllaaaxaaaaaa")

	
	app.Server.HTTP.Port = 9092
	app.Start()
}

