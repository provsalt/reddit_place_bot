package proxies

import (
	"go.uber.org/atomic"
	"math/rand"
	"os"
	"strings"
	"time"
)

var proxies []*atomic.String

func LoadProxies() error {
	f, err := os.ReadFile("proxies.txt")
	if err != nil {
		return err
	}
	t := strings.Split(strings.ReplaceAll(string(f), "\r\n", "\n"), "\n")
	for _, s := range t {
		proxies = append(proxies, atomic.NewString(s))
	}
	return nil
}

func GetRandomProxy() string {
	rand.Seed(time.Now().UnixNano())
	return proxies[rand.Intn(len(proxies))].Load()
}
