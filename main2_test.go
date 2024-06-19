package main

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

//var ctx = context.Background()

func TestSetAndGetEmployee(t *testing.T) {
	// Start a local redis server for testing
	s, err := miniredis.Run()
	if err != nil {
		t.Fatalf("Could not start miniredis: %v", err)
	}
	defer s.Close()

	// Create a new Redis client pointing to the in-memory server
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

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
}

//func setEmployee(rdb *redis.Client, id string, name string) error {
//	err := rdb.Set(ctx, id, name, 0).Err()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func getEmployee(rdb *redis.Client, id string) (string, error) {
//	name, err := rdb.Get(ctx, id).Result()
//	if err == redis.Nil {
//		return "", err
//	} else if err != nil {
//		return "", err
//	}
//	return name, nil
//}
