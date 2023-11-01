package handlers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shuttlersIT/intel/structs"
)

// Get a user ID from database
func GetUser(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(gorm.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	session := sessions.Default(c)
	email := session.Get("user-id")
	db := d
	id := c.Param("id")
	var s structs.Staff
	err := db.QueryRow("SELECT id, first_name, last_name, staff_email, username, position_id, department_id FROM staff WHERE email = ?", email).
		Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.StaffEmail, &s.Username, &s.PositionID, &s.DepartmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update a update by ID
func UpdateUser(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(gorm.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	session := sessions.Default(c)
	email := session.Get("user-id")
	db := d
	id := c.Param("id")
	var s structs.Staff
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE staff SET first_name = ?, last_name = ?, staff_email = ?, username = ?, position_id = ?, department_id = ?, WHERE id = ?", s.FirstName, s.LastName, s.StaffEmail, s.Username, s.PositionID, s.DepartmentID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User updated successfully")
}

// Delete a user by ID
func DeleteUser(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(gorm.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	session := sessions.Default(c)
	email := session.Get("user-id")
	db := d
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM staff WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User deleted successfully")
}
