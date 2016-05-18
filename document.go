package otter

type user struct{
  var id string
  var conn *Conn
}

type document struct{
  users map[string]user
  operations chan operation
  in_edited_list bool
}

func get_document(c *gin.Context)(status int, document string){
  name := c.Param("id")
  document, present := document_pool[name]
  if present{
    status = http.StatusOK
  }else{
    document_
  }
}
