package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			nt := NormalTruck{id: "NT-001", cargo: 0}
			et := EletricTruck{id: "ET-001", cargo: 0, batteryLevel: 100}

			if err := processTruck(&nt); err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if err := processTruck(&et); err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			//asserting
			if nt.cargo != 0 {
				t.Fatalf("Expected NormalTruck cargo to be 0 after load and unload, got %d", nt.cargo)
			}
			if (et.cargo != 0) || (et.batteryLevel != 80) {
				t.Fatalf("Expected EletricTruck cargo to be 0 and batteryLevel to be 80 after load and unload, got cargo %d and batteryLevel %d", et.cargo, et.batteryLevel)
			}
		})
	})
}
