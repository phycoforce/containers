package main

import (
	"testing"

	"github.com/phycoforce/containers/testhelpers"
)

func Test(t *testing.T) {
	image := testhelpers.GetTestImage("ghcr.io/phycoforce/owncast:rolling")
	testhelpers.TestHTTPEndpoint(t, image, testhelpers.HTTPTestConfig{Port: "8080"}, nil)
}
