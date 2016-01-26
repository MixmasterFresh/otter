package otter

import(
  "os"
  "flag"
)

func main(){
  action := os.Args[1]
  switch action {
  case "help":
    
  case "server":
    var port = flag.Uint("port", 1729, "port")
  case "-v"
  default:
  }
}
