package main
import (
    "net/http"

    "github.com/gin-gonic/gin"
)
// album represents data about a record album.
type user struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Age int  `json:"age"`
    Course  int `json:"course"`
}

// users slice to seed record album data.
var users = []user{
    {ID: "1", Name: "Abdussalam", Age: 19, Course: 3},
    {ID: "2", Name: "Abdussalyam", Age: 19, Course: 2},
    {ID: "3", Name: "Abdusalam", Age: 19, Course: 1},
}

// getusers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

func main() {
    router := gin.Default()
    router.GET("/users", getUsers)

    router.Run("localhost:8080")
}
