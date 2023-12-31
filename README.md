## Problem statement

You are working on a microservice for an e-commerce application
that receives orders from customers.
The microservice receives orders in JSON format,
and each order consists of
an order ID, customer ID, item ID, quantity, and status.
Status can be one of the following: ‘Received’, ‘Processing’, or ‘Delivered’.

The input is a JSON file with the following structure:

```
[
    {
        "orderId": "247149",
        "customerId": "822963",
        "itemId": "531647",
        "quantity": "2",
        "status": "Delivered"
    },
    {
        "orderId": "320673",
        "customerId": "386022",
        "itemId": "694648",
        "quantity": "2",
        "status": "Processing"
    },
    {
        "orderId": "325430",
        "customerId": "357055",
        "itemId": "326592",
        "quantity": "2",
        "status": "Received"
    },
    // More orders...
]
```

Note that all field values are in string format.

Your task is to write a Go program that
- reads these orders
- filters out the orders marked as ‘Processing’
- and calculates the total quantity of each item

The output should be a Go map with the item ID as the key
and the total quantity ordered for that item as the value.

## Prompt 1

```
You are to create a Go program that parses and analyzes a set of orders from an e-commerce microservice. The orders are in JSON
format, and each order includes an order ID, customer ID, item ID, quantity, and status ('Received', 'Processing' or 'Delivered').

Your specific task is to:
Read these orders from a JSON file (orders_data.json). Filter out the orders not marked as 'Processing'.
Calculate the total quantity ordered for each item. The JSON input file structure is as follows:
[
    {
        "orderId": "247149",
        "customerId": "822963",
        "itemId": "531647",
        "quantity": "2",
        "status": "Delivered"
    },
    // More orders...
]

Note that all field values are of type string.

The output should be a Go map. The map's key is the item ID of type string, and the value is the total quantity ordered for that
item of type int.
The code should be divided into functions: one should read orders, the other should process orders, so it is convenient to write
unit tests.
```

## Prompt 2

```
add according test cases here:

func Test_processOrders(t *testing.T) {
	type args struct {
		orders []Order
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processOrders(tt.args.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

## Prompt 3

```
I added a test case:

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

now it gives following error:

=== RUN   Test_processOrders/Test_with_invalid_Quantity
Failed to parse quantity for order 1: strconv.Atoi: parsing "invalid": invalid syntax
    main_test.go:94: processOrders() = map[], want map[1:0]

update processOrders function so that this test case would pass
```