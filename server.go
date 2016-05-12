package main

import(
  "net"
  "strconv"
  "os"
  "time"
)

keychain := make(map[string]map[int]string)
keychain := make(map[string]chan)

func start(port string){
  port_num, err := strconv.Atoi(port)
  if err != nil{
    panic("Port must be a valid number between 1 and 65535 and must be free\n")
  }

  listener, err := net.ListenTCP("tcp", ":" + port)
  if err != nil {
    panic(err);
  }
  var buf bytes.Buffer
  logger := log.New(&buf, "LOG: ", log.Ldate | log.Ltime)

  signals := make(chan os.Signal, 1)
  signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
  for {
    listener.SetDeadline(time.Now().Add(2 * time.Second))
    conn, err := listener.Accept()
    if err != nil {

    }
    go Handle(conn, logger)
  }
}

func handle(conn net.Conn, logger Logger){


  logger.Printf("IP: %s    Authenticated: %s    Request: %s", )
}

func authenticate() bool{

}

func generate_key() string{
  //TODO: write key generation
}

func revoke_key(){

}
