package main

import (
  "net/smtp"
  "log"
)

func notify(message string) {
  address := "smtp.gmail.com:587"
  auth    := smtp.PlainAuth("", "pull.ss.feed@gmail.com", "pullssfeed", "smtp.gmail.com")
  from    := "pull.ss.feed@gmail.com"
  to      := []string{"magnuss.karklins@gmail.com"}
  msg     := []byte(message)

  err := smtp.SendMail(address, auth, from, to, msg)
  if err != nil { log.Fatal(err) }
}
