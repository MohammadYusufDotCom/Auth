package handlers

import (
	"auth/internal/services"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
This struct is for
we can define struct and pass it to the handler function
*/
type UserHandler struct {
	DB *sql.DB
}

func (D *UserHandler) GetUsers(c *gin.Context) {
	var lastId int = 0
	for {
		users, err := services.GetDBData(D.DB, lastId)
		if err != nil {
			fmt.Fprintf(c.Writer, "[%s] ❌ DATABASE ERROR: %v\n", time.Now().Format(time.RFC3339), err.Error())
			if f, ok := c.Writer.(http.Flusher); ok {
				f.Flush()
			}
			time.Sleep(1 * time.Second)
			continue
		}
		lastId = users[len(users)-1].ID

		fmt.Fprintf(c.Writer, "✅ Found %d users at %s\n", len(users), time.Now().Format(time.RFC3339))
		for _, u := range users {
			fmt.Fprintf(c.Writer, "ID: %d | Name: %s | Email: %s\n", u.ID, u.UserName, u.Email)
		}
		fmt.Fprintf(c.Writer, "---------------------------------\n")
		if f, ok := c.Writer.(http.Flusher); ok {
			f.Flush()
		}
		time.Sleep(1 * time.Second)
	}
}

// func (D *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
// 	for {
// 		users, err := services.GetDBData(D.DB)
// 		if err != nil {
// 			fmt.Fprintf(w, "❌ DATABASE ERROR: %v\n", err.Error())
// 			if f, ok := w.(http.Flusher); ok {
// 				f.Flush()
// 			}
// 			time.Sleep(1 * time.Second)
// 			continue
// 		}
// 		fmt.Fprintf(w, "✅ Found %d users at %s\n", len(users), time.Now().Format(time.RFC3339))
// 		for _, u := range users {
// 			fmt.Fprintf(w, "ID: %d | Name: %s\n", u.ID, u.Name)
// 		}
// 		fmt.Fprintf(w, "---------------------------------\n")
// 		if f, ok := w.(http.Flusher); ok {
// 			f.Flush()
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }
