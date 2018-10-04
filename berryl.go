package main

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

type Class struct {
	Num		int
	Name	string
}

type Assignment struct {
	Num			int
	Name		string
	Graded		bool
	Grade		int
	MaxGrade	int
	DueDate		string
}

type Discussion struct {
	Num				int
	Subject			string
	CommentCount	int
	DateStarted		string
	DateActive		string
}

func handleClass(c *gin.Context) {
	c.HTML(http.StatusOK, "class", gin.H {
		"School": "UW",
		"Class": Class { 1, "CSS300" },
	})
}

func handleAssignments(c *gin.Context) {
	c.HTML(http.StatusOK, "assignments", gin.H {
		"School": "UW",
		"Class": Class { 1, "CSS300" },
		"Assignments": []Assignment {
			Assignment { 1, "uwu", false, 0, 5, "tomorrow" },
			Assignment { 2, "uwu2", true, 3, 5, "yesterday" },
		},
	})
}

func handleDiscussions(c *gin.Context) {
	c.HTML(http.StatusOK, "discussions", gin.H {
		"School": "UW",
		"Class": Class { 1, "CSS300" },
		"Discussions": []Discussion {
			Discussion {1, "foo", 13, "two days ago", "yesterday"},
			Discussion {2, "bar", 3, "5 days ago", "2 seconds ago"},
			Discussion {3, "baz", 4, "7 weeks ago", "3 days ago"},
		},
	})
}

func handleWiki(c *gin.Context) {
	c.HTML(http.StatusOK, "wiki", gin.H {
		"School": "UW",
		"Class": Class { 1, "CSS300" },
	})	
}

func createRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("class", "templates/class.html")
	r.AddFromFiles("assignments", "templates/assignments.html")
	r.AddFromFiles("discussions", "templates/discussions.html")
	r.AddFromFiles("wiki", "templates/wiki.html")
	return r
}

func main() {
	r := gin.Default()
	r.HTMLRender = createRender()

	r.GET("/class/:class", handleClass)
	r.GET("/class/:class/assignments", handleAssignments)
	r.GET("/class/:class/discussions", handleDiscussions)
	r.GET("/class/:class/wiki", handleWiki)

	r.Run()
}
