package components

import (
	"testing"
)

func TestItem(t *testing.T) {
	item := Item{
		Text: "test text",
		Time: "02 Jan 06 15:04",
	}

	if item.Title() != "test text" {
		t.Fatalf("Title should be test text")
	}
	if item.Description() != "02 Jan 06 15:04" {
		t.Fatalf("Description should be 02 Jan 06 15:04")
	}
	if item.FilterValue() != "test text" {
		t.Fatalf("FilterValue should be test text")
	}
}
