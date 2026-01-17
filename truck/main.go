package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("some new error")
	ErrTruckNotFound  = errors.New("truck not found")
)

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

func processTruck(t Truck) error {
	fmt.Printf("Processing truck: %+v\n", t)

	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo for truck %+v: %w", t, err)
	}

	fmt.Printf("Loaded Cargo: %+v\n", t)

	if err := t.UnloadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo for truck %+v: %w", t, err)
	}

	fmt.Printf("Unloaded Cargo: %+v\n", t)

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

func main() {
	trucks := []NormalTruck{
		{id: "TRK001"},
		{id: "TRK002"},
		{id: "TRK003"},
	}

	eTrucks := []EletricTruck{
		{id: "ETRK001", cargo: 0, batteryLevel: 80},
		{id: "ETRK002", cargo: 0, batteryLevel: 50},
	}

	if err := processTruck(&trucks[0]); err != nil {
		switch err {
		case ErrNotImplemented:
			log.Fatalf("Truck with ID %s: Feature not implemented yet.\n", trucks[0].id)
		case ErrTruckNotFound:
			log.Fatalf("Truck with ID %s: Truck not found.\n", trucks[0].id)
		default:
			log.Fatalf("Error processing truck with ID %s: %v\n", trucks[0].id, err)
		}
	}

	if err := processTruck(&eTrucks[0]); err != nil {
		switch err {
		case ErrNotImplemented:
			log.Fatalf("Truck with ID %s: Feature not implemented yet.\n", eTrucks[0].id)
		case ErrTruckNotFound:
			log.Fatalf("Truck with ID %s: Truck not found.\n", eTrucks[0].id)
		default:
			log.Fatalf("Error processing truck with ID %s: %v\n", eTrucks[0].id, err)
		}
	}
}
