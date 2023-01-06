package main

import (
	"net/http"
	// "fmt"
	"errors"
	"github.com/gin-gonic/gin"
)

type Expendecture struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Date        string `json:"date"`
}

var Expends = []Expendecture{
	{ID: "1", Description: "Chakula", Amount: "3000", Date: "02-10-2023"},
	{ID: "2", Description: "Usafir", Amount: "2500", Date: "02-10-2023"},
	{ID: "3", Description: "Maji", Amount: "1000", Date: "02-10-2023"},
	{ID: "4", Description: "Mafuta", Amount: "1200", Date: "02-10-2023"},
	{ID: "5", Description: "Matunda", Amount: "5000", Date: "02-10-2023"},
}

func getExpends(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Expends)
}

func expendById(c *gin.Context) {
	id := c.Param("id")
	Expendecture, err := getExpendById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Such expendecture"})
		return
	}
	c.IndentedJSON(http.StatusOK, Expendecture)
}
func getExpendById(id string) (*Expendecture, error){
	for i, expends := range Expends {
		if expends.ID == id {
			return &Expends[i], nil
		}
	}
	return nil, errors.New("No Such expendecture")
}

func createExpends(c *gin.Context) {
	var newExpends Expendecture
	if err := c.BindJSON(&newExpends); err != nil {
		return
	}
	Expends = append(Expends, newExpends)
	c.IndentedJSON(http.StatusCreated, newExpends)
}
func main() {
	router := gin.Default()
	router.GET("/expends", getExpends)
	router.GET("/expends/:id", expendById)
	router.POST("/expends", createExpends)
	router.Run("localhost:8000")
}
