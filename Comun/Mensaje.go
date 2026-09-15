package comun

type Mensaje struct {
     Type      string `json:"type"`
     Username  string `json:"username,omitempty"`
     Operation string `json:"operation,omitempty"`
     Result    string `json:"result,omitempty"`
     Extra     string `json:"extra,omitempty"`
}
