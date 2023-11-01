package handlers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shuttlersIT/intel/structs"
)

// Get an agent id from database
func GetAgent(c *gin.Context) {
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
	var a structs.Agent
	err := db.QueryRow("SELECT id, first_name, last_name, agent_email, username, role_id, unit, supervisor_id FROM agents WHERE email = ?", email).
		Scan(&a.StaffID, &a.FirstName, &a.LastName, &a.StaffEmail, &a.Username, &a.RoleID, &a.Unit, &a.SupervisorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// Update an agent by ID
func UpdateAgent(c *gin.Context) {
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
	var t structs.Staff
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE agents SET SELECT id = ?, first_name = ?, last_name = ?, agent_email = ?, username = ?, role_id = ?, unit = ?, supervisor_id = ?", s.FirstName, s.LastName, s.AgentEmail, s.Username, s.RoleID, s.Unit, s.SupervisorID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Agent updated successfully")
}

// Delete an agent by ID
func DeleteAgent(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(gorm.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return
	}

	session := sessions.Default(c)
	email := session.Get("user-id")
	db := d
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM agents WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Agent deleted successfully")
}