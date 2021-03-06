package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestRootCMD(t *testing.T) {
	go func() {
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
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Error(err)
	}

	assert.Equal(t, result["message"], "userB")
}
