package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type jobPosting struct {
	JobID          int64     `json:"jobID"`
	Name           string    `json:"name"`
	Company        string    `json:"company"`
	Tags           []string  `json:"tags"`
	JobDescription string    `json:"description"`
	Wage           float64   `json:"wage"`
	Qualifications []string  `json:"qualifications"`
	PostingDate    time.Time `json:"postingDate"`
}

func main() {
	router := gin.Default()
	router.GET("/jobs", getJobs)
	router.GET("/jobs/:jobID", getJobByID)
	router.POST("/jobs", postJobs)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("uh oh")
	}
}

var jobs = []jobPosting{
	{JobID: 12345678, Name: "Junior Engineer", Company: "Points", Tags: []string{"cs", "intern"}, JobDescription: "Wah",
		Wage: 15.64, Qualifications: []string{"none"}, PostingDate: time.Now()},
}

func getJobs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jobs)
}

func postJobs(c *gin.Context) {
	var newPosting jobPosting
	if err := c.BindJSON(&newPosting); err != nil {
		return
	}
	jobs = append(jobs, newPosting)
	c.IndentedJSON(http.StatusCreated, newPosting)
}

func getJobByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("jobID"), 10, 64)
	fmt.Println(id)
	for _, a := range jobs {
		if a.JobID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "job not found"})
}
