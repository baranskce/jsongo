package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var students []Student

func loadStudentsFromJSON(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		return err
	}
	return nil
}

func GetStudents(c *gin.Context) {
	c.JSON(200, students)
}

func GetStudentbyID(c *gin.Context) {
	id := c.Param("id")
	for _, student := range students {
		if student.ID == id {
			c.JSON(200, student)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Student not found"})
}

func CreateStudent(c *gin.Context) {
	var student Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	students = append(students, student)
	c.JSON(201, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var updatedStudent Student

	if err := c.BindJSON(&updatedStudent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	student, err := findStudentByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	student.Name = updatedStudent.Name
	student.Email = updatedStudent.Email

	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(200, student)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Student not found"})
}

func findStudentByID(id string) (*Student, error) {
	for i, student := range students {
		if student.ID == id {
			return &students[i], nil
		}
	}
	return nil, fmt.Errorf("Student not found")
}

func main() {

	err := loadStudentsFromJSON("data.json")
	if err != nil {
		fmt.Println("Error loading students from file:", err)
	}

	router := gin.Default()

	router.GET("/students", GetStudents)
	router.GET("/students/:id", GetStudentbyID)
	router.POST("/students", CreateStudent)
	router.PUT("/students/:id", UpdateStudent)
	router.DELETE("/students/:id", DeleteStudent)

	router.Run(":8080")
}
