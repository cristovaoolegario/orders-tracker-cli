package components

// Item represents a item in the list
type Item struct {
	Text, Time string
}

// Title retrieves de Text of the item
func (i Item) Title() string { return i.Text }

// Description retrieves de Time of the item
func (i Item) Description() string { return i.Time }

// FilterValue from list.Item
func (i Item) FilterValue() string { return i.Text }
