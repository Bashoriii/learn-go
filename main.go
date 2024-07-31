package main

import (
	itemsController "github.com/bashori/learn-go/controllers"
	"github.com/bashori/learn-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// ArrayAndOther()
	// LoopAndIf()
	// SomeSandbox()

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/items", itemsController.Index)   // Show all data
	r.GET("/api/item/:id", itemsController.Show) // Retrieve single data
	r.POST("/api/item", itemsController.Create)
	r.PUT("/api/item/:id", itemsController.Update)
	r.DELETE("/api/item", itemsController.Delete)

	r.Run()

	// Fetching example on below
	// url := "https://jsonplaceholder.typicode.com/posts"
	// client := &http.Client{}

	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatalf("GET Request Error: %v", err)
	// }

	// response, err := client.Do(req)
	// if err != nil {
	// 	log.Fatalf("Error sending request: %v", err)
	// }
	// defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatalf("Error reading response body: %v", err)
	// }

	// fmt.Println("Response status:", response.Status)
	// fmt.Println("Response body:", string(body))

}
