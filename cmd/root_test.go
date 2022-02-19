package cmd

import (
	"fmt"
	"testing"
)

func TestRootHandler(t *testing.T) {
	assertNoError := func(t testing.TB, err error) {
		t.Helper()

		if err == nil {
			t.Error("want an error but didn't got one")
		}
	}

	t.Run(fmt.Sprintf("wrong %s", concurrentHostsName), func(t *testing.T) {
		max := newRootCmd()
		max.Flags().IntVar(&concurrentHosts, concurrentHostsName, maxConcurrentHosts+1, "")
		assertNoError(t, max.Execute())

		min := newRootCmd()
		min.Flags().IntVar(&concurrentHosts, concurrentHostsName, minConcurrentHosts-1, "")
		assertNoError(t, min.Execute())
	})
}
