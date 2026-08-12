package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"dozen":               12,
		"quarter_of_a_dozen":  3,
		"half_of_a_dozen":     6,
		"small_gross":         120,
		"gross":               144,
		"great_gross":         1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if !exists {
		return false
	}

	bill[item] += unitValue
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billItem, itemExists := bill[item]
	if !itemExists {
		return false
	}

	unitItem, unitExists := units[unit]
	if !unitExists {
		return false
	}

	newQuantity := billItem - unitItem

	if newQuantity < 0 {
		return false
	}

	if newQuantity == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = newQuantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]

	if !exists {
		return 0, false
	}

	return qty, true
}