package source

import (
	"net/http"
)

// Checker instance structure
type Checker struct {
	Client  http.Client
	Invalid int
	Flagged int
	Locked  int
	Token   string
	Valid   int
	Resp    bool
	All     int
}

// Config structure
type Config struct {
	CheckFlagged bool `json:"check_flagged"`
	Offset       int  `json:"offset_ms"`
}

// Variables
var (
	z                      = "\033[36m"
	url                    = "https://discord.com/api/v9/users/@me/affinities/guilds"
	flagged_url            = "https://discord.com/api/v9/users/@me"
	c, r, g, bg, rb, gr, u = "\u001b[30;1m", "\033[39m", "\033[32m", "\u001b[34;1m", "\u001b[0m", "\u001b[30;1m", "\u001b[4m"
)
