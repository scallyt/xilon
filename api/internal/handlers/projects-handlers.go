package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/scallyt/neko/internal/db"
	"github.com/scallyt/neko/internal/repository"
	"github.com/scallyt/neko/internal/utils"
)

func CreateProject(c *gin.Context) {
	repo := repository.New(db.Conn)

	var project repository.CreateProjectParams
	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProject, err := repo.CreateProject(db.Ctx, repository.CreateProjectParams{
		CustomerName: project.CustomerName,
		CustomerID:   project.CustomerID,
		ProjectName:  project.ProjectName,
		ChatID:       project.ChatID,
		Description:  project.Description,
		DeveloperID:  project.DeveloperID,
		Status:       "in progress",
		StartDate:    project.StartDate,
		EndDate:      project.EndDate,
		Budget:       project.Budget,
	})
	if err != nil {
		log.Printf("Failed to create project: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, newProject)
}

func GetProjectByUserId(c *gin.Context) {
	repo := repository.New(db.Conn)

	idStr, err := utils.GetUserIdByJWT(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token"})
		return
	}

	var userID pgtype.UUID
	err = userID.Scan(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid UUID format"})
		return
	}
	projects, err := repo.GetAllProjectByUserId(db.Ctx, userID)
	if err != nil {
		log.Printf("Failed to get projects: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}
