package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestRootCMD(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Overwrite config using args.
		rootCMD.SetArgs(
			[]string{"--username", "userB"},
		)
		if err := rootCMD.Execute(); err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(1 * time.Second) // Waiting for gin engineer setup.
	resp, err := http.Get("http://127.0.0.1:8080/user")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	assert.Equal(t, result["message"], "userB")
}
