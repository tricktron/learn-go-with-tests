package main_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"learn-go-with-tests/go-specs-greet/adapters/httpserver"
	"learn-go-with-tests/go-specs-greet/specifications"

	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{ //nolint: exhaustruct
		FromDockerfile: testcontainers.FromDockerfile{ //nolint: exhaustruct
			Context:       "../../..",
			Dockerfile:    "./go-specs-greet/cmd/httpserver/Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	container, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{ //nolint: exhaustruct
			ContainerRequest: req,
			Started:          true,
		})

	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	client := http.Client{
		Timeout:       1 * time.Second,
		Transport:     http.DefaultTransport,
		Jar:           nil,
		CheckRedirect: nil,
	}
	driver := httpserver.Driver{BaseURL: "http://localhost:8080", Client: &client}
	specifications.GreetSpecification(t, driver)
}
