package main

import (
	"api_automail/docs"
	"encoding/json"
	"fmt"

	"net/http"

	"api_automail/models"

	ser "api_automail/services"

	"github.com/gin-gonic/gin"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())
	initSwagger()

	service := ser.NewService("")

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {

		response := models.HttpResponse{
			Code:    "200",
			Massage: "ถูกอัปโหลดและบันทึกเรียบร้อยแล้ว",
			Data:    nil,
		}

		err := service.ReceiptFileUpload(w, r)
		if err != nil {
			response.Code = "500"
			response.Massage = err.Error()
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	// Serve Swagger UI
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Start the HTTP server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)

}
func initSwagger() {
	fmt.Println("initSwagger")
	docs.SwaggerInfo.Title = "Test"
	docs.SwaggerInfo.Version = "version : localhost"
	docs.SwaggerInfo.Host = "localhost:8080" + "/"
	docs.SwaggerInfo.BasePath = ""
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Tracking-Id, X-Timestamp")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
