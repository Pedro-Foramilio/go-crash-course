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

type Truck struct {
	id string
}

func processTruck(t Truck) error {
	fmt.Printf("Processing truck with ID: %s\n", t.id)

	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo for truck %s: %w", t.id, err)
	}

	return nil
}

func (t *Truck) LoadCargo() error {
	return ErrTruckNotFound
}

func main() {
	trucks := []Truck{
		{id: "TRK001"},
		{id: "TRK002"},
		{id: "TRK003"},
	}

	for _, truck := range trucks {
		if err := processTruck(truck); err != nil {

			switch err {
			case ErrNotImplemented:
				log.Fatalf("Truck with ID %s: Feature not implemented yet.\n", truck.id)
				continue
			case ErrTruckNotFound:
				log.Fatalf("Truck with ID %s: Truck not found.\n", truck.id)
				continue
			default:
				log.Fatalf("Error processing truck with ID %s: %v\n", truck.id, err)
			}

		}
	}
}
