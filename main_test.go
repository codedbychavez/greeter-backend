package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIEndpoints(t *testing.T) {
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	GETRequestMethodTests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string

		// ReturnHello Expected Output


	}{
		{
			description:   "index route",
			route:         "/api/v1/status",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
		{
			description:   "non existing route",
			route:         "/api/v1/status/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /api/v1/status/i-dont-exist",
		},

	}

	// Setup the app as it is done in the main function
	app := Setup()



	// Iterate through test single test cases
	for _, test := range GETRequestMethodTests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}



	// Test POST Request Methods
	POSTRequestMethodTests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string

		}{
			
		// ReturnHello Endpoint
		{
			description: "index route",
			route:        "/api/v1/hello",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "Hello, John Doe!",
		},

	}

	type Person struct {
		Name string `json:"name"`
	}

	person := Person{
		Name: "John Doe",
	}

	personJson, _ := json.Marshal(person)


	for _, test := range POSTRequestMethodTests {

	req, _ := http.NewRequest(
		"POST",
		test.route,
		bytes.NewBuffer(personJson),
	)

	req.Header.Set("Content-Type", "application/json")

	// Perform the request plain with the app.
	// The -1 disables request latency.
	res, err := app.Test(req, -1)

	// verify that no error occured, that is not expected
	assert.Equalf(t, test.expectedError, err != nil, test.description)

	// test case needs to be processed
	if test.expectedError {
		continue
	}


	// Verify if the status code is as expected
	assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)

	// Reading the response body should work everytime, such that
	// the err variable should be nil
	assert.Nilf(t, err, test.description)

	// Verify, that the reponse body equals the expected body
	assert.Equalf(t, test.expectedBody, string(body), test.description)

}

}



