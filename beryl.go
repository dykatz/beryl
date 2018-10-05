package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

type Class struct {
	Num  int
	Name string
}

type User struct {
	Num  int
	Name string
}

type Category struct {
	Name   string
	Weight int
}

type Assignment struct {
	Num      int
	Name     string
	Graded   bool
	Grade    int
	MaxGrade int
	DueDate  string
	Category string
}

type Discussion struct {
	Num          int
	Subject      string
	CommentCount int
	DateStarted  string
	DateActive   string
}

type Comment struct {
	Text       string
	Author     string
	Edited     bool
	DatePosted string
	DateEdited string
}

var (
	School = "University of Washington"
)

func send404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404", gin.H{
		"School": School,
	})
}

func handleClass(c *gin.Context) {
	class, err := strconv.Atoi(c.Param("class"))
	if err != nil {
		send404(c)
		return
	}

	c.HTML(http.StatusOK, "class", gin.H{
		"School": School,
		"Class":  Class{class, "CSS300"},
		"Users": []User{
			{1, "Anony Mouse"},
			{2, "Some Guy"},
		},
	})
}

func handleAssignments(c *gin.Context) {
	class, err := strconv.Atoi(c.Param("class"))
	if err != nil {
		send404(c)
		return
	}

	c.HTML(http.StatusOK, "assignments", gin.H{
		"School":       School,
		"Class":        Class{class, "CSS300"},
		"CurrentGrade": 89.5,
		"MinGrade":     50.2,
		"Categories": []Category{
			{"hw", 50},
			{"tests", 25},
			{"participation", 10},
			{"presentation", 15},
		},
		"Assignments": []Assignment{
			{1, "uwu", false, 0, 5, "tomorrow", "hw"},
			{2, "uwu2", true, 3, 5, "yesterday", "hw"},
		},
	})
}

func handleAssignment(c *gin.Context) {
	class, err := strconv.Atoi(c.Param("class"))
	if err != nil {
		send404(c)
		return
	}

	assignment, err := strconv.Atoi(c.Param("assignment"))
	if err != nil {
		send404(c)
		return
	}

	c.HTML(http.StatusOK, "assignment", gin.H{
		"School":     School,
		"Class":      Class{class, "CSS300"},
		"Assignment": Assignment{assignment, "uwu", false, 0, 5, "tomorrow", "hw"},
	})
}

func handleDiscussions(c *gin.Context) {
	class, err := strconv.Atoi(c.Param("class"))
	if err != nil {
		send404(c)
		return
	}

	c.HTML(http.StatusOK, "discussions", gin.H{
		"School": School,
		"Class":  Class{class, "CSS300"},
		"Discussions": []Discussion{
			{1, "foo", 13, "two days ago", "yesterday"},
			{2, "bar", 3, "5 days ago", "2 seconds ago"},
			{3, "baz", 4, "7 weeks ago", "3 days ago"},
		},
	})
}

func handleDiscussion(c *gin.Context) {
	class, err := strconv.Atoi(c.Param("class"))
	if err != nil {
		send404(c)
		return
	}

	discussion, err := strconv.Atoi(c.Param("discussion"))
	if err != nil {
		send404(c)
		return
	}

	c.HTML(http.StatusOK, "discussion", gin.H{
		"School":     School,
		"Class":      Class{class, "CSS300"},
		"Discussion": Discussion{discussion, "foo", 13, "two days ago", "yesterday"},
		"Comments": []Comment{
			{"why post here", "unpleasant person", false, "yesterday", ""},
		},
	})
}

func handleWiki(c *gin.Context) {
	class, err := strconv.Atoi(c.Param("class"))
	if err != nil {
		send404(c)
		return
	}

	page := c.Param("page")

	c.HTML(http.StatusOK, "wiki", gin.H{
		"School": School,
		"Class":  Class{class, "CSS300"},
		"Page":   page,
	})
}

func createRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("404", "templates/404.html")
	r.AddFromFiles("class", "templates/class.html")
	r.AddFromFiles("assignments", "templates/assignments.html")
	r.AddFromFiles("assignment", "templates/assignment.html")
	r.AddFromFiles("discussions", "templates/discussions.html")
	r.AddFromFiles("discussion", "templates/discussion.html")
	r.AddFromFiles("wiki", "templates/wiki.html")

	return r
}

func main() {
	r := gin.Default()
	r.HTMLRender = createRender()

	r.StaticFile("/style.css", "assets/style.css")

	r.GET("/class/:class", handleClass)
	r.GET("/class/:class/assignments", handleAssignments)
	r.GET("/class/:class/assignment/:assignment", handleAssignment)
	r.GET("/class/:class/discussions", handleDiscussions)
	r.GET("/class/:class/discussion/:discussion", handleDiscussion)
	r.GET("/class/:class/wiki/*page", handleWiki)

	r.Run()
}
