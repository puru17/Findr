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

var supabaseClient *supabase.Client
var err error
var startLat = 49.266567
var startLng = -122.968769
var maxDistance float64 = 200

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
	c.IndentedJSON(http.StatusOK, gin.H{
		"TODO":  "User login page",
		"hello": "buddy",
	})
}

func apiSetRadius(c *gin.Context) {
	// Read the raw request body
	rawData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataString := string(rawData)
	num, err := strconv.ParseFloat(dataString, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("search radius ", num)
	maxDistance = num * 1000

	// http response to client browser
	c.JSON(http.StatusOK, gin.H{
		"message":  "Data received by server",
		"raw_data": dataString,
	})
}

func apiGetUsers(c *gin.Context) {

	result := supabaseClient.Rpc("get_nearby_locations", "1", map[string]interface{}{
		"input_longitude":    startLng,
		"input_latitude":     startLat,
		"input_max_distance": maxDistance,
	})

	// fmt.Println(result)

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

func init() {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseUrl == "" || supabaseKey == "" {
		log.Fatal("Error: SUPABASE_URL and SUPABASE_KEY environment variables must be set")
	}

	supabaseClient, err = supabase.NewClient(supabaseUrl, supabaseKey, &supabase.ClientOptions{})

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
}

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
	r.Static("/client", "./client")

	r.GET("/", root)

	r.GET("/api/users", apiGetUsers)
	r.POST("/api/users", apiSetRadius)
	// r.POST("/api/clicks", apiPostClicks)

	r.POST("/signup", signup)
	r.GET("/signup", signup)

	r.POST("/login", login)
	r.GET("/login", login)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
