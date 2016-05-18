package otter

import(
  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
  "net/http"
  "strconv"
)

var document_pool map[string]document // map[document_id]map[user_key]user
var master_key string
var edited chan string

func server(port int, main_key string){
  master_key = main_key
  document_pool = make(map[string]document)
  edited = make(chan string)
  router := gin.Default()

  router.GET("/ws", func(c *gin.Context) {
    websocket_handler(c.Writer, c.Request)
  })

  router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Otter Version %s", VERSION)
  })

  authorized := r.Group("/")
  authorized.Use(server_authentication()){ //Server Authentication group

    authorized.GET("/document/:id", get_document_endpoint)
    authorized.POST("/document/:id/edit", post_document_endpoint)
    authorized.DELETE("/document/:id/destroy", delete_document_endpoint)

    authorized.GET("/document/:id/user/create", create_user_endpoint)
    authorized.DELETE("/document/:id/user/destroy", delete_document_endpoint)

    authorized.GET("/edited", get_edited_documents_endpoint)
  }

  router.Run(":" + strconv.Itoa(port))
}

func server_authentication(){//gin middleware
  return func(c *gin.Context){
    key := c.Request.FormValue("key")

    if key == "" {
      respondWithError(401, "Key Required", c)
      return
    }

    if key != master_key) {
      respondWithError(401, "Invalid Key", c)
      return
    }

    c.Next()
  }
}

//Document Interactions
func get_document_endpoint(c *gin.Context){

}

func post_document_endpoint(c *gin.Context){

}

func delete_document_endpoint(c *gin.Context){

}

//User Creation and Deletion
func create_user_endpoint(c *gin.Context){

}

func delete_document_endpoint(c *gin.Context){

}

//Get a list of documents edited since last
func get_edited_documents_endpoint(c *gin.Context){

}
