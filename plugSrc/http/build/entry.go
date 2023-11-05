package build

import (
	"bufio"
	"fmt"
	"github.com/google/gopacket"
	"github.com/patrickmn/go-cache"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var validIPTV = regexp.MustCompile(`(\/.+\/.+\/[0-9]+)`)

const (
	Port    = 80
	Version = "0.1"
)

const (
	CmdPort = "-p"
)

type H struct {
	port    int
	version string
	c       *cache.Cache
}

var c *cache.Cache
var hp *H

func NewInstance() *H {

	if hp == nil {
		hp = &H{
			port:    Port,
			version: Version,
			c:       cache.New(5*time.Minute, 10*time.Minute),
		}
	}
	return hp
}
func main() {

}

func (m *H) ResolveStream(net, transport gopacket.Flow, buf io.Reader) {

	bio := bufio.NewReader(buf)
	for {
		req, err := http.ReadRequest(bio)

		if err == io.EOF {
			return
		} else if err != nil {
			continue
		} else {
			log.Println("allready known! " + req.URL.String())

			if validIPTV.MatchString(req.URL.String()) {
				if _, found := c.Get(req.Host + req.URL.String()); found {
					log.Println("allready known! ")
				}
				fileAppend("temp.log", req.URL.String())
				m.c.Set(req.Host+req.URL.String(), req.Form.Encode(), cache.NoExpiration)
			}

			if strings.Contains(req.URL.String(), "/live/") {
				if _, found := c.Get(req.Host + req.URL.String()); found {
					log.Println("allready known! ")
				}
				fileAppend("temp.log", req.URL.String())
				m.c.Set(req.Host+req.URL.String(), req.Form.Encode(), cache.NoExpiration)
				//log.Println("Match")
				//fmt.Sprintf("FOUND: %s\n", req.Host+req.URL.String())
				//var msg = "[" + req.Method + "]" + "[" + req.Host + "]" + req.URL.String()
				//req.ParseForm()
				//log.Println(msg)

			}

			req.Body.Close()
		}
	}
}

func (m *H) BPFFilter() string {
	return "tcp and port " + strconv.Itoa(m.port)
}

func (m *H) Version() string {
	return Version
}

func (m *H) SetFlag(flg []string) {

	c := len(flg)

	if c == 0 {
		return
	}
	if c>>1 == 0 {
		fmt.Println("ERR : Http Number of parameters")
		os.Exit(1)
	}
	for i := 0; i < c; i = i + 2 {
		key := flg[i]
		val := flg[i+1]

		switch key {
		case CmdPort:
			port, err := strconv.Atoi(val)
			m.port = port
			if err != nil {
				panic("ERR : port")
			}
			if port < 0 || port > 65535 {
				panic("ERR : port(0-65535)")
			}
			break
		default:
			panic("ERR : mysql's params")
		}
	}
}
func fileAppend(filename, data string) (err error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.WriteString(data)
	return
}
