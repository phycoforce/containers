package main

import (
	"testing"

	helpers "github.com/phycoforce/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/phycoforce/owncast:rolling")
	helpers.RequireHTTPEndpoint(t, image, helpers.HTTPTestConfig{Port: "8080"}, nil)
}
