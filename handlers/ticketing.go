package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/intel/structs"
)

//Ticket Handlers
/*
// Ticketing Handlers
func ListTickets(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementportal.html", gin.H{"Username": userID})
}

func CreateTicket(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementadmin.html", gin.H{"Username": userID})
}
func UpdateTicket(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementx.html", gin.H{"Username": userID})
}
func DeleteTicket(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementx.html", gin.H{"Username": userID})
}
*/

// List all tickets
func ListTickets(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, title, description, status FROM tickets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tickets []structs.Ticket
	for rows.Next() {
		var t structs.Ticket
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tickets = append(tickets, t)
	}

	c.JSON(http.StatusOK, tickets)
}

// Create a new ticket
func CreateTicket(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	var t structs.Ticket
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO tickets (title, description, status) VALUES (?, ?, ?)", t.Subject, t.Description, t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	lastInsertID, _ := result.LastInsertId()
	t.ID = int(lastInsertID)
	c.JSON(http.StatusCreated, t)
}

// Get a ticket by ID
func GetTicket(c *gin.Context) {
	id := c.Param("id")

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	var t structs.Ticket
	err := db.QueryRow("SELECT id, subject, description, status FROM tickets WHERE id = ?", id).
		Scan(&t.ID, &t.Subject, &t.Description, &t.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	c.JSON(http.StatusOK, t)
}

// Update a ticket by ID
func UpdateTicket(c *gin.Context) {
	id := c.Param("id")

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	var t structs.Ticket
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE tickets SET title = ?, description = ?, status = ? WHERE id = ?", t.Subject, t.Description, t.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Ticket updated successfully")
}

// Delete a ticket by ID
func DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	_, err := db.Exec("DELETE FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Ticket deleted successfully")
}
