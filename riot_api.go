package main

import (
  "fmt"
  //"net/http"
  "strings"
)

type RiotAPI struct{
  api_key string
  server  string
  region  string
}

func matchQueue(){
  var input int
  var input2 string

  queueTypes := map[int]string{
    400 : "Draft",
    420 : "Solo/Duo",
    430 : "Blind",
    440 : "Flex",
    450 : "ARAM",
    700 : "Clash",
  }

  queueWeight := map[string]int{
    "UNRANKED"   : -1,
    "IRON"       : 0,
    "BRONZE"     : 1,
    "SILVER"     : 2,
    "GOLD"       : 3,
    "PLATINUM"   : 4,
    "DIAMOND"    : 5,
    "MASTER"     : 6,
    "GRANDMASTER": 7,
    "CHALLENGER" : 8,
  }
  fmt.Scanln(&input) // userinput is passed by reference
  fmt.Println(queueTypes[input]) // PRINTS OUT whatever the user enter
  fmt.Scanln(&input2)
  inputTwo := strings.ToUpper(input2)
  fmt.Println(queueWeight[inputTwo])
}
