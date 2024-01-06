package main

import (
  "log"

  "github.com/joeyaflores/go-chat/internal/web"
  "github.com/joeyaflores/go-chat/internal/rbmq"
)

func main() {
  log.Printf("app started")
  go rbmq.Start()
  web.Start()
}
