package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Course int    `json:"course"`
}

type course struct {
	Title         string  `json:"title"`
	Grade         float64 `json:"grade"`
	Description   string  `json:"description"`
	Prerequisites string  `json: "prerequisites"`
}

// users
var users = []user{
	{ID: "1", Name: "Abdussalam", Age: 19, Course: 3},
	{ID: "2", Name: "Abdussalyam", Age: 19, Course: 2},
	{ID: "3", Name: "Abdusalam", Age: 19, Course: 1},
}

var courses = []course{
	{Title: "OOP", Grade: 4.0, Description: "Object Oriented Programming using Java", Prerequisites: "PP1, PP2"},
	{Title: "Angular", Grade: 4.0, Description: "Frontend Framework Angular TS", Prerequisites: "PP1, PP2, WD"},
	{Title: "DSA", Grade: 4.0, Description: "Data Storage and Analysis", Prerequisites: "no prerequisites"},
}

// getusers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
func getUserDetail(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func postUsers(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func getCourseDetail(c *gin.Context) {
	title := c.Param("title")

	for _, course := range courses {
		if course.Title == title {
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postCourses(c *gin.Context) {
	var newCourse course
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func main() {
	router := gin.Default()
	router.GET("users", getUsers)
	router.GET("users/:id", getUserDetail)
	router.POST("users", postUsers)

	router.GET("courses", getCourses)
	router.GET("courses/:title", getCourseDetail)
	router.POST("courses", postCourses)

	router.Run("localhost:8000")
}
