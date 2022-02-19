package main

import (
	"reflect"
	"testing"
)

func TestStore_AddCustomerData(t *testing.T) {
	s := &Store{}
	s.AddCustomerData("Annika", "Kexchoklad", 10)
	s.AddCustomerData("Annika", "Kexchoklad", 10)
	s.AddCustomerData("Jonas", "Geisha", 10)

	expected := map[string]map[string]int{
		"Annika": {"Kexchoklad": 20},
		"Jonas":  {"Geisha": 10},
	}

	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Got unexpected customer data %#v. Want %#v", s.data, expected)
	}
}

func TestStore_GetSortedCustomers(t *testing.T) {
	s := &Store{
		data: map[string]map[string]int{
			"Annika": {"Kexchoklad": 5},
			"Aadya":  {"Center": 100},
			"Jonas":  {"Geisha": 20},
		},
	}

	customers := s.GetSortedCustomers()
	first := customers[0]
	second := customers[1]
	third := customers[2]

	if first.Name != "Aadya" {
		t.Errorf("Expected first to be Aadya got %s", first.Name)
	}
	if second.Name != "Jonas" {
		t.Errorf("Expected second to be Jonas got %s", second.Name)
	}
	if third.Name != "Annika" {
		t.Errorf("Expected third to be Aadya got %s", first.Name)
	}
}

func TestStore_FavouriteAndTotalEaten(t *testing.T) {
	s := &Store{
		data: map[string]map[string]int{
			"Annika": {"Kexchoklad": 5, "Center": 7, "Nötchoklad": 2},
		},
	}

	favouriteSnack, totalEaten := s.FavouriteAndTotalEaten("Annika")
	if favouriteSnack != "Center" {
		t.Errorf("Expected favourite to be Center got %s", favouriteSnack)
	}
	if totalEaten != 14 {
		t.Errorf("Expected total to be 14 got %d", totalEaten)
	}
}
