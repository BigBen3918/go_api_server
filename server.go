package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"

	"go.chromium.org/luci/server/router"
)

func Logger(c *router.Context, next router.Handler) {
	fmt.Println(c.Request.URL)
	next(c)
}

func AuthCheck(c *router.Context, next router.Handler) {
	var authenticated bool
	if !authenticated {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write([]byte("Authentication failed"))
		return
	}
	next(c)
}

func GenerateSecret(c *router.Context, next router.Handler) {
	c.Context = context.WithValue(c.Context, "secret", rand.Int())
	next(c)
}

func makeRequest(client *http.Client, url string) string {
	res, err := client.Get(url + "/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	p, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if len(p) == 0 {
		return fmt.Sprintf("%d", res.StatusCode)
	}
	return fmt.Sprintf("%d %s", res.StatusCode, p)
}

func makeRequests(url string) {
	c := &http.Client{}
	fmt.Println(makeRequest(c, url+"/hello"))
	fmt.Println(makeRequest(c, url+"/hello/darknessmyoldfriend"))
	fmt.Println(makeRequest(c, url+"/authenticated/secret"))
}

// Example_createServer demonstrates creating an HTTP server using the router
// package.
func main() {
	r := router.New()
	r.Use(router.NewMiddlewareChain(Logger))
	r.GET("/hello", nil, func(c *router.Context) {
		fmt.Fprintf(c.Writer, "Hello")
	})
	r.GET("/hello/:name", nil, func(c *router.Context) {
		fmt.Fprintf(c.Writer, "Hello %s", c.Params.ByName("name"))
	})

	auth := r.Subrouter("authenticated")
	auth.Use(router.NewMiddlewareChain(AuthCheck))
	auth.GET("/secret", router.NewMiddlewareChain(GenerateSecret), func(c *router.Context) {
		fmt.Fprintf(c.Writer, "secret: %d", c.Context.Value("secret"))
	})

	server := httptest.NewServer(r)
	defer server.Close()

	makeRequests(server.URL)

}
