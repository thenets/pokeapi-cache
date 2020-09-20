package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var c *cache.Cache

func pokeapiHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	url := fmt.Sprintf("https://pokeapi.co/%s", path)

	// Check if is reaching the API
	// ignore any other request
	if !strings.HasPrefix(r.URL.Path, "/api/") {
		// http.NotFound(w, r)
		http.Redirect(w, r, "https://github.com/thenets/pokeapi-cache", 301)
		return
	}

	// Make request using cache
	if x, found := c.Get(url); found {
		content := x.(string)

		// Log stdout
		logrus.WithFields(logrus.Fields{
			"Code":      200,
			"FromCache": true,
		}).Debug(path)

		io.WriteString(w, content)
	} else {
		// Make remote requests
		resp, err := http.Get(url)

		// Check error
		if err != nil {
			http.Error(w, "PokeAPI server returned error", 500)

			// Log stderr
			logrus.Error(fmt.Sprintf("PokeAPI server not responding"))

			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != 200 {
			http.Error(
				w,
				fmt.Sprintf("PokeAPI server returned error: %s", string(body)),
				resp.StatusCode,
			)

			// Log stderr
			logrus.WithFields(logrus.Fields{
				"Code": resp.StatusCode,
			}).Error(path)

			return
		}

		// Output
		io.WriteString(w, string(body))

		// Log stdout
		logrus.WithFields(logrus.Fields{
			"Code":      resp.StatusCode,
			"FromCache": false,
		}).Info(path)

		// Make cache
		c.Set(url, string(body), cache.DefaultExpiration)
	}
}

func main() {
	// Setup log
	if strings.ToUpper(os.Getenv("DEBUG")) == "TRUE" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Start cache
	c = cache.New(
		// Default expiration time of 5 hours
		5*time.Hour,
		// Purges expired items every 10 minutes
		10*time.Minute,
	)

	// Handlers
	r := mux.NewRouter()
	http.HandleFunc("/", pokeapiHandler)
	r.PathPrefix("/").Handler(r)

	// Start server
	serverPort := "8080"
	if os.Getenv("PORT") != "" {
		serverPort = os.Getenv("PORT")
	}
	http.ListenAndServe(":"+serverPort, nil)
}
