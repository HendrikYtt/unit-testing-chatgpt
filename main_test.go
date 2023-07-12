package main

import (
	"reflect"
	"testing"
)

func Test_processOrders(t *testing.T) {
	type args struct {
		orders []Order
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Test with single Processing order",
			args: args{orders: []Order{
				{
					OrderID:    "1",
					CustomerID: "1",
					ItemID:     "1",
					Quantity:   "5",
					Status:     "Processing",
				},
			}},
			want: map[string]int{"1": 5},
		},
		{
			name: "Test with multiple Processing orders, same item",
			args: args{orders: []Order{
				{
					OrderID:    "1",
					CustomerID: "1",
					ItemID:     "1",
					Quantity:   "5",
					Status:     "Processing",
				},
				{
					OrderID:    "2",
					CustomerID: "2",
					ItemID:     "1",
					Quantity:   "3",
					Status:     "Processing",
				},
			}},
			want: map[string]int{"1": 8},
		},
		{
			name: "Test with multiple Processing and non-Processing orders",
			args: args{orders: []Order{
				{
					OrderID:    "1",
					CustomerID: "1",
					ItemID:     "1",
					Quantity:   "5",
					Status:     "Processing",
				},
				{
					OrderID:    "2",
					CustomerID: "2",
					ItemID:     "1",
					Quantity:   "3",
					Status:     "Processing",
				},
				{
					OrderID:    "3",
					CustomerID: "3",
					ItemID:     "1",
					Quantity:   "10",
					Status:     "Delivered",
				},
			}},
			want: map[string]int{"1": 8},
		},
		{
			name: "Test with invalid Quantity",
			args: args{orders: []Order{
				{
					OrderID:    "1",
					CustomerID: "1",
					ItemID:     "1",
					Quantity:   "invalid",
					Status:     "Processing",
				},
			}},
			want: map[string]int{"1": 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processOrders(tt.args.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
