package otter

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TheAustinSeven/otter/auth"
	"github.com/TheAustinSeven/otter/document"
	"github.com/gin-gonic/gin"
)

var documentPool map[string]*document.Document // map[documentId]map[userKey]user
var edited chan string

func server(port int, mainKey string) {
	auth.Initialize(mainKey)
	documentPool = make(map[string]*document.Document)
	edited = make(chan string, 10000)
	router := gin.Default()

	router.GET("/ws/:id/:userId", func(c *gin.Context) {
		person := identifyUser(c)
		person.OpenConnection(c.Writer, c.Request)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Otter Version %s", VERSION)
	})

	authorized := router.Group("/")
	authorized.Use(auth.ServerAuthentication())
	{ //Server Authentication group
		authorized.GET("/document/:id", getDocumentEndpoint)
		authorized.GET("/document/:id/meta", getDocumentMetadataEndpoint)
		authorized.POST("/document/:id/create", createDocumentEndpoint)
		authorized.POST("/document/:id/edit", editDocumentEndpoint)
		authorized.DELETE("/document/:id/destroy", deleteDocumentEndpoint)

		authorized.GET("/document/:id/user/create", createUserEndpoint)
		authorized.DELETE("/document/:id/user/:userId/destroy", deleteDocumentEndpoint)

		authorized.GET("/edited", getEditedDocumentsEndpoint)
		authorized.GET("/empty", getEmptyDocumentsEndpoint)
	}

	router.Run(":" + strconv.Itoa(port))
}

//Document Interactions

func getDocumentEndpoint(c *gin.Context) {
	name := c.Param("id")
	document, present := documentPool[name]
	if !present {
		auth.NotFound(c)
		return
	}
	contents := document.GetString()
	fmt.Println(contents)
	//TODO: this
}

func getDocumentMetadataEndpoint(c *gin.Context) {
	name := c.Param("id")
	document, present := documentPool[name]
	if !present {
		auth.NotFound(c)
		return
	}
	data := document.GetMetadata()
	fmt.Println(data["length"])
	//TODO: this
}

func createDocumentEndpoint(c *gin.Context) {
	name := c.Param("id")
	_, present := documentPool[name]
	if present {
		//JSON with 409:Conflict
		return
	}
	//extract document from incoming JSON
	//create document
	//TODO: this
}

func editDocumentEndpoint(c *gin.Context) {
	name := c.Param("id")
	_, present := documentPool[name]
	if !present {
		auth.NotFound(c)
		return
	}
	//extract document from incoming JSON
	//overwrite document
	//TODO: this
}

func deleteDocumentEndpoint(c *gin.Context) {
	name := c.Param("id")
	_, present := documentPool[name]
	if !present {
		auth.NotFound(c)
		return
	}
	//TODO: this
}

//User Creation and Deletion
func createUserEndpoint(c *gin.Context) {
	name := c.Param("id")
	_, present := documentPool[name]
	if !present {
		auth.NotFound(c)
		return
	}
	//TODO: this
}

func deleteUserEndpoint(c *gin.Context) {
	name := c.Param("id")
	_, present := documentPool[name]
	if !present {
		auth.NotFound(c)
		return
	}
	//TODO: this
}

//Get a list of documents edited since last
func getEditedDocumentsEndpoint(c *gin.Context) {
	//TODO: this
}

//Get a list of documents edited since last
func getEmptyDocumentsEndpoint(c *gin.Context) {
	//TODO: this
}

func identifyUser(c *gin.Context) *document.User {
	//TODO: this
}
