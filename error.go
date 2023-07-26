package nex

// import (
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/fatih/color"
// )

// // http404Handler is the default 404 handler
// // Its handle the request when no route is found
// // and log the request and response and return a 404 status code.
// func http404Handler(c *Context) {
// 	log := log.New(os.Stdout, "[NEX] ", 0)

// 	timestamp := time.Now().Format(time.DateTime) // use the start time as the timestamp
// 	method := c.Req.Method
// 	status := 404
// 	path := c.Req.URL.Path
// 	ip := strings.Split(c.Req.RemoteAddr, ":")[0]

// 	// Log the details
// 	log.Printf("%s %s %s %s %s \n",
// 		color.New(color.BgHiRed).Sprintf(" %s ", method),
// 		color.New(color.BgHiRed).Sprintf(" %d ", status),
// 		timestamp,
// 		path,
// 		ip,
// 	)

// 	c.Res.WriteHeader(http.StatusNotFound)
// 	c.Res.Write([]byte("404 - Not Found"))
// }
