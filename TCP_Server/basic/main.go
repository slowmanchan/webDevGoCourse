package main

import (
  "fmt"
  "io"
  "log"
  "net"
)

func main() {
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Panic(err)

  }
  defer li.Close()

  for {
    conn, err := li.Accept()
    if err != nil {
      log.Println(err)
    }
    io.WriteString(conn, "\n Hi from TCP server\n")
    fmt.Fprintln(conn, "How you doing")
    fmt.Fprintf(conn, "%v", "Well, I hope!")

    conn.Close()
  }
}
