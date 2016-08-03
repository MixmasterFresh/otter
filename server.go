package otter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"otter/document"
)

var documentPool map[string]Document // map[documentId]map[userKey]user
var masterKey string
var edited chan string

func server(port int, mainKey string) {
	masterKey = mainKey
	documentPool = make(map[string]Document)
	edited = make(chan string, 10000)
	router := gin.Default()

	router.GET("/ws", func(c *gin.Context) {
		person, authorized := identifyAndAuthorizeUser(c)
		if authorized {
			person.openConnection(c.Writer, c.Request)
		}
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Otter Version %s", VERSION)
	})

	authorized := router.Group("/")
	authorized.Use(serverAuthentication())
	{ //Server Authentication group
		authorized.GET("/document/:id", getDocumentEndpoint)
		authorized.GET("/document/:id/meta", getDocumentMetadataEndpoint)
		authorized.POST("/document/:id/create", createDocumentEndpoint)
		authorized.POST("/document/:id/edit", editDocumentEndpoint)
		authorized.DELETE("/document/:id/destroy", deleteDocumentEndpoint)

		authorized.GET("/document/:id/user/create", createUserEndpoint)
		authorized.DELETE("/document/:id/user/:token/destroy", deleteDocumentEndpoint)

		authorized.GET("/edited", getEditedDocumentsEndpoint)
	}

	router.Run(":" + strconv.Itoa(port))
}

func serverAuthentication() gin.HandlerFunc { //gin middleware
	return func(c *gin.Context) {
		key := c.Request.FormValue("key")

		if key == "" {
			respondWithError(401, "Key Required", c)
			return
		}

		if key != masterKey {
			respondWithError(401, "Invalid Key", c)
			return
		}

		c.Next()
	}
}

//Document Interactions

func getDocumentEndpoint(c *gin.Context) {
	name := c.Param("id")
	document, present := documentPool[name]
	if !present {
		notFound(c)
		return
	}
	contents := document.getString()
	//JSON
}

func getDocumentMetadataEndpoint(c *gin.Context) {
	name := c.Param("id")
	document, present := documentPool[name]
	if !present {
		notFound(c)
		return
	}
	data := document.getMetadata()
	//JSON
}

func createDocumentEndpoint(c *gin.Context) {
	name := c.Param("id")
	document, present := documentPool[name]
	if present {
		//JSON with 409:Conflict
		return
	}
	//extract document from incoming JSON
	//create document
	//JSON
}

func editDocumentEndpoint(c *gin.Context) {
	name := c.Param("id")
	document, present := documentPool[name]
	if !present {
		notFound(c)
		return
	}
	//extract document from incoming JSON
	//overwrite document
	//JSON
}

func deleteDocumentEndpoint(c *gin.Context) {
	name := c.Param("id")
	document, present := documentPool[name]
	if !present {
		notFound(c)
		return
	}

}

//User Creation and Deletion
func createUserEndpoint(c *gin.Context) {
	name := c.Param("id")
	document, present := documentPool[name]
}

func deleteUserEndpoint(c *gin.Context) {
	name := c.Param("id")
	document, present := documentPool[name]
}

//Get a list of documents edited since last
func getEditedDocumentsEndpoint(c *gin.Context) {

}

func notFound(c *gin.Context) {

}

func respondWithError(status int, message string, c *gin.Context) {

}

func identifyAndAuthorizeUser(c *gin.Context) (person *user.User, authorized bool) {

}
