// Code generated by zanzibar
// @generated

package googlenow

import (
	"bytes"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/zanzibar/test/lib/test_gateway"
)

func TestCheckCredentialsSuccessfulRequestOKResponse(t *testing.T) {
	var counter int

	gateway, err := testGateway.CreateGateway(t, nil, &testGateway.Options{
		KnownBackends: []string{"googleNow"},
		TestBinary: filepath.Join(
			getDirName(), "..", "..", "main.go",
		),
	})
	if !assert.NoError(t, err, "got bootstrap err") {
		return
	}
	defer gateway.Close()

	fakeCheckCredentials := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		// TODO(zw): generate client response.
		if _, err := w.Write([]byte("{status:\"200 OK\"}")); err != nil {
			t.Fatal("can't write fake response")
		}
		counter++
	}

	gateway.Backends()["googleNow"].HandleFunc(
		"POST", "/check-credentials", fakeCheckCredentials,
	)

	res, err := gateway.MakeRequest(
		"POST", "/googlenow/check-credentials", bytes.NewReader([]byte("{\"authcode\":\"test\"}")),
	)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	assert.Equal(t, "200 OK", res.Status)
	assert.Equal(t, 1, counter)
}
