package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	redisClient "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var ctx = context.Background()

// TestIntegrationWithRedis tests setting and getting an employee
func TestIntegrationWithRedis(t *testing.T) {

	//// Start a Redis container
	request := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		//Env:          map[string]string{"ALLOW_EMPTY_PASSWORD": "yes"},
		Env: map[string]string{
			//"REDIS_USERNAME": "myadmin",
			"REDIS_PASSWORD": "secret",
		},
		WaitingFor: wait.ForLog("Ready to accept connections"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
	//s.NoError(err)
	if err != nil {
		t.Fatalf("Could not start redis container: %v", err)
	}

	endpoint, err := container.Endpoint(ctx, "")
	//s.NoError(err)
	if err != nil {
		t.Fatalf("Could not get container IP: %v", err)
	}

	// Create a new Redis client pointing to the Redis container
	rdb := redisClient.NewClient(&redisClient.Options{
		//Addr: redisAddress,
		Addr: endpoint,
		//Username: "myadmin",
		//	Password: "secret",
	})

	// Wait for the Redis container to be ready
	for {
		err := rdb.Ping(ctx).Err()
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Run the integration tests
	t.Run("Set and Get Employee", func(t *testing.T) {
		// Test setting an employee
		err = setEmployee(rdb, "123", "John Doe")
		assert.NoError(t, err)

		// Test getting an employee
		name, err := getEmployee(rdb, "123")
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", name)

		// Test getting a non-existent employee
		_, err = getEmployee(rdb, "999")
		assert.Error(t, err)
	})
}

func setEmployee(rdb *redisClient.Client, id string, name string) error {
	err := rdb.Set(ctx, id, name, 0).Err()
	if err != nil {
		fmt.Printf("Could not set employee: %v\n", err)
		return err
	}
	fmt.Printf("Employee with ID %s set to %s\n", id, name)
	return nil
}

func getEmployee(rdb *redisClient.Client, id string) (string, error) {
	name, err := rdb.Get(ctx, id).Result()
	if errors.Is(err, redisClient.Nil) {
		fmt.Printf("Employee with ID %s does not exist\n", id)
		return "", err
	} else if err != nil {
		fmt.Printf("Could not get employee: %v\n", err)
		return "", err
	} else {
		fmt.Printf("Employee with ID %s is %s\n", id, name)
	}
	return name, nil
}
