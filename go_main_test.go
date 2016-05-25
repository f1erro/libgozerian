package main

import (
  "fmt"
  "os"
  "testing"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

const (
  testHandler = "testHandler"
  testDebug = false
)

func TestGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Test")
}

var testHTTPServer *Server
var testURL string

var _ = BeforeSuite(func() {
  SetTestRequestHandler()
  CreateHandler(testHandler, "")

  testURL = os.Getenv("WEAVER_TEST_URL")

  if testURL == "" {
    var err error
    testHTTPServer, err = StartWeaverServer(0, "", true)
    Expect(err).Should(Succeed())
    testHTTPServer.SetDebug(testDebug)
    testPort := testHTTPServer.GetPort()
    Expect(testPort).ShouldNot(BeZero())
    testURL = fmt.Sprintf("http://localhost:%d", testPort)
    go testHTTPServer.Run()
    fmt.Fprintf(GinkgoWriter, "Running test against local server on port %d\n", testPort)
  } else {
    fmt.Printf("Running test against remote server at %s\n", testURL)
  }
})

var _ = AfterSuite(func() {
  DestroyHandler(testHandler)

  if testHTTPServer != nil {
    testHTTPServer.Stop()
  }
})
