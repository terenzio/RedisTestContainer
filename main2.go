package main

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"log"
//
//	"github.com/go-redis/redis/v8"
//)
//
//var ctx = context.Background()
//
//func main() {
//	// Create a new Redis client
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379", // Redis server address
//		Password: "",               // No password set
//		DB:       0,                // Use default DB
//	})
//
//	// Function to set employee data
//	setEmployee(rdb, "123", "John Doe")
//	setEmployee(rdb, "124", "Jane Smith")
//
//	// Function to get employee data
//	getEmployee(rdb, "123")
//	getEmployee(rdb, "124")
//	getEmployee(rdb, "125") // This ID does not exist
//}
//
//func setEmployee(rdb *redis.Client, id string, name string) error {
//	err := rdb.Set(ctx, id, name, 0).Err()
//	if err != nil {
//		log.Fatalf("Could not set employee: %v", err)
//		return err
//	}
//	fmt.Printf("Employee with ID %s set to %s\n", id, name)
//	return nil
//}
//
//func getEmployee(rdb *redis.Client, id string) (string, error) {
//	name, err := rdb.Get(ctx, id).Result()
//	if errors.Is(err, redis.Nil) {
//		fmt.Printf("Employee with ID %s does not exist\n", id)
//		return "", err
//	} else if err != nil {
//		log.Fatalf("Could not get employee: %v", err)
//		return "", err
//	} else {
//		fmt.Printf("Employee with ID %s is %s\n", id, name)
//	}
//	return name, nil
//}
