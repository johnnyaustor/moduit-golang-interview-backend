package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"moduit/app/model"
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	baseUrl = `https://screening.moduit.id`
)

var (
	transport = &http.Transport{}
	client    = &http.Client{
		Transport: transport,
	}
)

func get(uri string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl+uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "go")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	log.Printf(`GET %v - %v`, baseUrl+uri, res.Status)

	return ioutil.ReadAll(res.Body)
}

func One(c echo.Context) error {
	body, err := get("/backend/question/one")

	var oneResponse model.OneResponse
	if err = json.Unmarshal(body, &oneResponse); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, oneResponse)
}

func Two(c echo.Context) error {
	body, err := get("/backend/question/two")

	var oneResponse []model.OneResponse
	if err = json.Unmarshal(body, &oneResponse); err != nil {
		return err
	}

	// Filter
	var responses []model.OneResponse
	for _, one := range oneResponse {
		if titleAndDescContains(one, "Ergonomic") && tagsContains(one.Tags, "Sports") {
			responses = append(responses, one)
		}
	}

	// Sorting Id Desc
	sort.Slice(responses, func(i, j int) bool {
		return responses[i].Id > responses[j].Id
	})

	// Slice
	responses = responses[:3]

	return c.JSON(http.StatusOK, responses)
}

func Three(c echo.Context) error {
	body, err := get("/backend/question/three")

	var threeResponse []model.ThreeResponse
	if err = json.Unmarshal(body, &threeResponse); err != nil {
		log.Printf("%v", err)
		return err
	}

	// Mapper ThreeResponse to OneResponse
	var responses []model.OneResponse
	var one model.OneResponse
	for _, three := range threeResponse {
		for _, threeItem := range three.Items {
			one = model.OneResponse{
				Id:          three.Id,
				Category:    three.Category,
				Title:       threeItem.Title,
				Description: threeItem.Description,
				Footer:      threeItem.Footer,
				Tags:        three.Tags,
				CreatedAt:   three.CreatedAt,
			}
			responses = append(responses, one)
		}
	}

	return c.JSON(http.StatusOK, responses)
}

func titleAndDescContains(one model.OneResponse, value string) bool {
	return strings.Contains(one.Title, value) || strings.Contains(one.Description, value)
}

func tagsContains(tags []string, value string) bool {
	for _, tag := range tags {
		if tag == value {
			return true
		}
	}
	return false
}
