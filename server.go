package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	supabase "github.com/supabase-community/supabase-go"
)

// Config holds all configuration values
type Config struct {
	StartLat    float64
	StartLng    float64
	MaxDistance float64
	SupabaseURL string
	SupabaseKey string
	ServerPort  string
}

var (
	config         Config
	supabaseClient *supabase.Client
)

func loadConfig() error {
	// Load environment variables
	config.SupabaseURL = os.Getenv("SUPABASE_URL")
	config.SupabaseKey = os.Getenv("SUPABASE_KEY")
	config.ServerPort = os.Getenv("SERVER_PORT")
	if config.ServerPort == "" {
		config.ServerPort = "8080"
	}

	// Load default values
	config.StartLat = 49.266567
	config.StartLng = -122.968769
	config.MaxDistance = 200

	// Validate required environment variables
	if config.SupabaseURL == "" || config.SupabaseKey == "" {
		return fmt.Errorf("SUPABASE_URL and SUPABASE_KEY environment variables must be set")
	}

	return nil
}

func init() {
	if err := loadConfig(); err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	var err error
	supabaseClient, err = supabase.NewClient(config.SupabaseURL, config.SupabaseKey, &supabase.ClientOptions{})
	if err != nil {
		log.Fatal("Failed to initialize Supabase client:", err)
	}
}

func root(c *gin.Context) {
	c.File("./client/index.html")
}

func signup(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"TODO":  "User signup page",
		"hello": "buddy",
	})
}

func login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.File("./client/login.html")
		return
	}

	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// TODO: Implement actual authentication logic
	// For now, just return success
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func apiSetRadius(c *gin.Context) {
	rawData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	dataString := string(rawData)
	num, err := strconv.ParseFloat(dataString, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid radius value"})
		return
	}

	// Validate radius
	if num <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Radius must be greater than 0"})
		return
	}

	config.MaxDistance = num * 1000
	c.JSON(http.StatusOK, gin.H{
		"message": "Radius updated successfully",
		"radius":  num,
	})
}

func apiGetUsers(c *gin.Context) {
	result := supabaseClient.Rpc("get_nearby_locations", "1", map[string]interface{}{
		"input_longitude":    config.StartLng,
		"input_latitude":     config.StartLat,
		"input_max_distance": config.MaxDistance,
	})

	if result == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch nearby locations"})
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(result))
}

// func apiPostClicks(c *gin.Context) {
// 	// Read the raw request body
// 	rawData, err := c.GetRawData()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Convert the raw data to a string
// 	dataString := string(rawData)

// 	// Process the raw data -- write to txt file
// 	// err = appendStringToFile("coordinates.txt", dataString+"\n")
// 	// if err != nil {
// 	// 	fmt.Println("Error appending to file:", err)
// 	// 	return
// 	// }

// 	// http response to client browser
// 	c.JSON(http.StatusOK, gin.H{
// 		"message":  "Data received",
// 		"raw_data": dataString,
// 	})
// }

// func appendStringToFile(filename string, data string) error {
// 	// Open the file in append mode. Create it if it doesn't exist.
// 	// 0644 gives read and write permissions to the owner, and read
// 	// permissions to the group and others.
// 	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Write the string to the file.
// 	_, err = file.WriteString(data)
// 	if err != nil {
// 		return err
// 	}

// 	// (Optional) Force the data to be written to the underlying storage device.
// 	err = file.Sync()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func main() {
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.Static("/client", "./client")

	r.GET("/", root)
	r.GET("/api/users", apiGetUsers)
	r.POST("/api/users", apiSetRadius)
	// r.POST("/api/clicks", apiPostClicks)

	r.POST("/signup", signup)
	r.GET("/signup", signup)

	r.POST("/login", login)
	r.GET("/login", login)

	serverAddr := fmt.Sprintf("localhost:%s", config.ServerPort)
	log.Printf("Server starting on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
