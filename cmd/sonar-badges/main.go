package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var username = os.Getenv("USERNAME")
var token = os.Getenv("TOKEN")
var password = os.Getenv("PASSWORD")
var serverURL = os.Getenv("SERVER_URL")
var port = os.Getenv("PORT")

var metricTypes = []string{
	"bugs",
	"code_smells",
	"coverage",
	"duplicated_lines_density",
	"ncloc",
	"sqale_rating",
	"alert_status",
	"reliability_rating",
	"security_rating",
	"sqale_index",
	"vulnerabilities",
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func errorResponse(c *gin.Context, err error) {
	defer log.Print(err)

	c.JSON(404, gin.H{
		"error": err.Error(),
	})
}

func getBadge(c *gin.Context) {
	metric := strings.TrimSuffix(c.Param("metric"), ".svg")

	if contains(metricTypes, metric) && strings.HasSuffix(c.Param("metric"), ".svg") {

		// queries
		readerString := fmt.Sprintf(`metric=%v&project=%v`, metric, c.Param("project"))

		// add branch parameter
		if branch := c.Query("branch"); branch != "" {
			readerString += fmt.Sprintf(`&branch=%v`, branch)
		}

		req, requestError := http.NewRequest("POST", serverURL+"/api/project_badges/measure",
			// body
			strings.NewReader(readerString),
		)

		if requestError != nil {
			errorResponse(c, errors.New("there was an request error"))
		}

		if token != "" {
			req.SetBasicAuth(token, "")
		} else {
			req.SetBasicAuth(username, password)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, responseError := http.DefaultClient.Do(req)

		if responseError != nil {
			errorResponse(c, errors.New("there was an response error"))
		}

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				errorResponse(c, errors.New("there was an error"))
			}

			c.Data(200, "image/svg+xml", []byte(string(bodyBytes)))
		} else {
			errorResponse(c, errors.New("there was an error"))
		}

		defer resp.Body.Close()
	} else {
		errorResponse(c, errors.New("not valid"))
	}
}

func setupRouter() *gin.Engine {
	router := gin.New()

	if os.Getenv("ENABLE_REQUEST_LOG") == "true" {
		router = gin.Default()
	}

	router.GET("/:project/:metric", getBadge)

	return router
}

func main() {
	if username == "" {
		log.Fatal("please provide a username")
	}

	if password == "" {
		log.Fatal("please provide a password")
	}

	if serverURL == "" {
		log.Fatal("please provide a server url")
	}

	if port == "" {
		port = "8080"
	}

	router := setupRouter()

	router.Run(":" + port)
}
