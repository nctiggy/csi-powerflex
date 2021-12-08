package service

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"testing"
	"github.com/cucumber/godog"
)

func TestMain(m *testing.M) {
	go http.ListenAndServe("localhost:6060", nil)
        fmt.Printf("starting godog...\n")

        opts := godog.Options{
                Format: "pretty",
                Paths:  []string{"features"},
                Tags:   "wip",
        }

        status := m.Run()

        st := godog.TestSuite{
                Name:                "godog",
                ScenarioInitializer: FeatureContext,
                Options:             &opts,
        }.Run()


        fmt.Printf("godog test status %d\n", status)
	if st > 0 || status > 0 {
		os.Exit(1)
	}
        os.Exit(0)

}
