package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	ErrNotImplemented = errors.New("some new error")
	ErrTruckNotFound  = errors.New("truck not found")
	UserIdKey         = contextKey("userId")
)

type contextKey string

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

type EletricTruck struct {
	id           string
	cargo        int
	batteryLevel int
}

func processTruck(ctx context.Context, t Truck) error {
	userId := ctx.Value(UserIdKey)
	fmt.Printf("Started Processing truck: %+v with user %v\n", t, userId)
	time.Sleep(1 * time.Second)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	//simulate long running process
	delay := 1 * time.Second
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo for truck %+v: %w", t, err)
	}

	// fmt.Printf("Loaded Cargo: %+v\n", t)

	if err := t.UnloadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo for truck %+v: %w", t, err)
	}

	// fmt.Printf("Unloaded Cargo: %+v\n", t)

	fmt.Printf("Finished Processing truck: %+v\n", t)
	return nil
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1000
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo -= 1000
	return nil

}
func (t *EletricTruck) LoadCargo() error {
	t.cargo += 500
	t.batteryLevel -= 10
	return nil
}

func (t *EletricTruck) UnloadCargo() error {
	t.cargo -= 500
	t.batteryLevel -= 10
	return nil
}

func processFleet(ctx context.Context, trucks []Truck) error {
	var wg = sync.WaitGroup{}

	for _, t := range trucks {
		wg.Add(1)

		go func(tt Truck) {
			if err := processTruck(ctx, tt); err != nil {
				log.Printf("Error processing truck %+v: %v\n", tt, err)
			}
			wg.Done()
		}(t)

	}

	wg.Wait()
	return nil
}

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIdKey, 42)

	trucks := []NormalTruck{
		{id: "TRK001"},
		{id: "TRK002"},
		{id: "TRK003"},
	}

	eTrucks := []EletricTruck{
		{id: "ETRK001", cargo: 0, batteryLevel: 80},
		{id: "ETRK002", cargo: 0, batteryLevel: 50},
	}

	fleet := []Truck{
		&trucks[0],
		&trucks[1],
		&trucks[2],
		&eTrucks[0],
		&eTrucks[1],
	}

	if err := processFleet(ctx, fleet); err != nil {
		log.Fatalf("Error processing fleet: %v\n", err)
	}
}
