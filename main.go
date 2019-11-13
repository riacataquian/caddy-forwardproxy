// This package extends caddy with forwardproxy
// https://github.com/caddyserver/caddy/wiki/Plugging-in-Plugins-Yourself
package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	_ "github.com/caddyserver/forwardproxy" // plug in for http(s) forwardproxy
)

func main() {
	caddymain.Run()
}
