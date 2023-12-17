package main_test

//nolint: goimports
import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	go_specs_greet "learn-go-with-tests/go-specs-greet"
	"learn-go-with-tests/go-specs-greet/specifications"
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
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080", Client: &client}
	specifications.GreetSpecification(t, driver)
}
