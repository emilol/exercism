package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	itemCount, itemExists := bill[item]
	if !itemExists {
		bill[item] = unitValue
	} else {
		bill[item] = itemCount + unitValue
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemCount, itemExists := bill[item]
	if !itemExists {
		return false
	}

	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	newItemCount := itemCount - unitValue
	if newItemCount < 0 {
		return false
	}

	if newItemCount == 0 {
		delete(bill, item)
	} else {
		bill[item] = newItemCount
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemCount, itemExists := bill[item]
	return itemCount, itemExists
}
