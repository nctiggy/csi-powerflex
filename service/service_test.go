package service

import (
	"fmt"
	"github.com/cucumber/godog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	go http.ListenAndServe("localhost:6060", nil)
	fmt.Printf("starting godog...\n")

	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"features"},
		//Tags:   "wip",
	}

	st := godog.TestSuite{
		Name:                "godog",
		ScenarioInitializer: FeatureContext,
		Options:             &opts,
	}.Run()
	m.Run()

	fmt.Printf("godog test status %d\n", st)

	os.Exit(st)

}
