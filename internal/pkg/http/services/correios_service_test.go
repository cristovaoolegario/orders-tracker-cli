package services

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFindOrderByNumber_ShouldReturnDataWhenOrderExists(t *testing.T) {
	orderNumber := "QA695731454TL"
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jsonFile, _ := os.Open("../../mock_data/valid_code.json")

			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			w.Header().Add("Content-Type", "application/json")
			w.Write(byteValue)
		}),
	)

	defer ts.Close()

	service := ProvideCorreiosService(ts.URL + "/%s")
	order, err := service.FindOrderByNumber(orderNumber)

	if err != nil && order == nil {
		t.Errorf("Error: %s", err)
	}
}

func TestFindOrderByNumber_ShouldReturnErrorWhenOrderDontExists(t *testing.T) {
	orderNumber := "bla"

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jsonFile, _ := os.Open("../../mock_data/invalid_code.json")

			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			w.Header().Add("Content-Type", "application/json")
			w.Write(byteValue)
		}),
	)

	defer ts.Close()

	service := ProvideCorreiosService(ts.URL + "/%s")
	result, err := service.FindOrderByNumber(orderNumber)

	if result != nil && err == nil {
		t.Errorf("Should return a error when there's a problem with the search.")
	}
}
