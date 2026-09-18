package comun

type Mensaje struct {
     Type      string `json:"type"`
     Username  string `json:"username,omitempty"`
     Operation string `json:"operation,omitempty"`
     Result    string `json:"result,omitempty"`
     Extra     string `json:"extra,omitempty"`
     Status    string `json:"status,omitempty"`
     Text      string `json:"text,omitempty"`
     Roomname  string `json:"roomname,omitempty"`
     Users     map[string]string `json:"users,omitempty"`
     //Roomname  string `json:"roomname,omitempty"`
}
