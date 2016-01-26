package otter

import(
  "os"
  "flag"
)

func main(){
  action := os.Args[1]
  var port = flag.Uint("port", 1729, )
  switch action {
  case "help":
    
  case "run":
    
  default:
  }
}
