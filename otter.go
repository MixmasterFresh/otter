package otter

import(
  "os"
  "fmt"
  "flag"
)

const(
  VERSION = "0.1.0"
  help_usage = "prints the help message"
  version_usage = "prints the version of otter you are currently using"
  port_usage = "The port to expose the otter server on"
  key_usage = "provide a strong key here for authenticating your server"
  help_text = "text goes here" //TODO: fix this
)

func main(){
  var help bool
  flag.BoolVar(&help, "help", false, help_usage)

  var port int
  flag.IntVar(&port, "port", 1729, port_usage)
  flag.IntVar(&port, "p", 1729, port_usage)

  var version_flag bool
  flag.BoolVar(&version_flag, "version", false, version_usage)

  var key string
  flag.StringVar(&key, "key", "", key_usage)
  flag.StringVar(&key, "k", "", key_usage)

  flag.Parse()

  if help {
    fmt.Printf("%s\n", help_text)
    os.Exit(0)
  }
  if version_flag {
    fmt.Printf("%s\n", VERSION)
    os.Exit(0)
  }

  if port < 0 || port > 65535{
    panic("Port must be a valid number between 1 and 65535 and must be free\n")
  }

  server(port, key)
}
