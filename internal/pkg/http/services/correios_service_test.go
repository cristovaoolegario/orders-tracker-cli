package services

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCorreiosService_FindOrderByNumber(t *testing.T) {

	t.Run("Should return data when order exists", func(t *testing.T) {
		orderNumber := "QA695731454TL"
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				jsonFile, _ := os.Open("../../mock/mock_data/valid_code.json")

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
	})

	t.Run("Should return error when order don't exists", func(t *testing.T) {
		orderNumber := "bla"

		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				jsonFile, _ := os.Open("../../mock/mock_data/invalid_code.json")

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
	})
}
