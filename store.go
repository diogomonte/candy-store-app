package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"sort"
	"strconv"
)

const candyStoreUrl = "https://candystore.zimpler.net/"

type Store struct {
	data map[string]map[string]int
}

type Customer struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnacks"`
}

func (s *Store) AddCustomerData(customerName string, candy string, eatenCount int) {
	if s.data == nil {
		s.data = make(map[string]map[string]int)
	}
	if candyCount, ok := s.data[customerName]; ok {
		candyCount[candy] = candyCount[candy] + eatenCount
	} else {
		s.data[customerName] = map[string]int{candy: eatenCount}
	}
}

func (s Store) GetSortedCustomers() []Customer {
	var customerArray []Customer
	for customerName := range s.data {
		favouriteSnack, totalEaten := s.FavouriteAndTotalEaten(customerName)
		c := Customer{
			Name:           customerName,
			FavouriteSnack: favouriteSnack,
			TotalSnacks:    totalEaten,
		}
		customerArray = append(customerArray, c)
	}
	sort.SliceStable(customerArray, func(i, j int) bool {
		return customerArray[i].TotalSnacks > customerArray[j].TotalSnacks
	})
	return customerArray
}

func (s Store) FavouriteAndTotalEaten(customer string) (favouriteSnack string, totalEaten int) {
	favouriteEatenCount := 0
	for candy, count := range s.data[customer] {
		totalEaten = totalEaten + count
		if count > favouriteEatenCount {
			favouriteSnack = candy
			favouriteEatenCount = count
		}
	}
	return favouriteSnack, totalEaten
}

func (s *Store) LoadData() {
	res, err := http.Get(candyStoreUrl)
	if err != nil {
		log.Panicln("Error fetching html page content")
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	userData := struct {
		name  string
		candy string
		eaten int
	}{}

	doc.Find(".details").Each(func(di int, container *goquery.Selection) {
		container.Find("table tr").Each(func(tri int, tr *goquery.Selection) {
			tr.Find("td").Each(func(tdi int, td *goquery.Selection) {
				if tdi == 0 {
					userData.name = td.Text()
				}
				if tdi == 1 {
					userData.candy = td.Text()
				}
				if tdi == 2 {
					count, _ := strconv.Atoi(td.Text())
					userData.eaten = count
				}
			})
			if userData.name != "" {
				s.AddCustomerData(userData.name, userData.candy, userData.eaten)
			}
		})
	})
}
