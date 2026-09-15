package main

type mensaje struct {
     Type      string `json: "type"`
     Username  string `json: "username"`
     Operation string `json: "operation"`
     Result    string `json: "result"`
     Extra     string `json: "extra"`
}
