package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guiziin227/goCleanArc/internal/models"
)

func main() {

	req := models.CreateUserRequest{
		Name:  "Malu",
		Email: "Malu@example.com",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8000/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusCreated {
		panic(fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
	}

	var responseApi models.CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
		panic(err)
	}

	fmt.Println(responseApi)
}
