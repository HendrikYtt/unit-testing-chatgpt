package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Order struct {
	OrderID    string `json:"orderId"`
	CustomerID string `json:"customerId"`
	ItemID     string `json:"itemId"`
	Quantity   string `json:"quantity"`
	Status     string `json:"status"`
}

// Function to read orders from JSON
func readOrders(fileName string) ([]Order, error) {
	var orders []Order

	// Read JSON file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return orders, err
	}

	// Unmarshal JSON data
	err = json.Unmarshal(data, &orders)
	if err != nil {
		return orders, err
	}

	return orders, nil
}

// Function to process orders
func processOrders(orders []Order) map[string]int {
	items := make(map[string]int)

	// Iterate over the orders
	for _, order := range orders {
		if order.Status == "Processing" {
			// Convert string quantity to int
			quantity, err := strconv.Atoi(order.Quantity)
			if err != nil {
				fmt.Printf("Failed to parse quantity for order %s: %v\n", order.OrderID, err)
				// Assign a quantity of 0 and continue
				quantity = 0
			}
			// Add the quantity to the map
			items[order.ItemID] += quantity
		}
	}

	return items
}

func main() {
	// Read orders
	orders, err := readOrders("orders_data.json")
	if err != nil {
		fmt.Printf("Failed to read orders: %v\n", err)
		os.Exit(1)
	}

	// Process orders
	items := processOrders(orders)

	// Print results
	for item, quantity := range items {
		fmt.Printf("ItemID: %s, Quantity: %d\n", item, quantity)
	}
}
