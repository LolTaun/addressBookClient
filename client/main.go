package main

import (
	"HW1_http/models/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	for {
		fmt.Println("Please enter the action you want to perform:\n1. Add new record\n2. Get records by any data\n3. Update record by phone number\n4. Delete record by phone number\n5. Exit")
		action := 0
		fmt.Scanln(&action)
		switch action {
		case 1:
			recordAdd()
		case 2:
			recordsGet()
		case 3:
			recordUpdate()
		case 4:
			recordDelete()
		case 5:
			return
		default:
			fmt.Println("Wrong action")
		}
	}
}

func recordAdd() {
	record := dto.Record{}
	fmt.Println("Please enter the record data:")

	for record.Name == "" {
		fmt.Println("Name:")
		fmt.Scanln(&record.Name)
	}

	for record.LastName == "" {
		fmt.Println("Last name:")
		fmt.Scanln(&record.LastName)
	}

	fmt.Println("Middle name (may leave blank):")
	fmt.Scanln(&record.MiddleName)

	for record.Address == "" {
		fmt.Println("Address:")
		fmt.Scanln(&record.Address)
	}

	for record.Phone == "" {
		fmt.Println("Phone:")
		fmt.Scanln(&record.Phone)
	}

	jsonData, err := json.Marshal(record)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// fmt.Println("jsonData =", string(jsonData)) // FOR DEBUGGING

	resp, err := http.Post("http://localhost:8080/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body := make([]byte, resp.ContentLength)
	_, err = io.ReadFull(resp.Body, body)
	if err != nil {
		fmt.Println("Error Reading Body:", err)
		return
	}

	ResponseHandler(body)

	// fmt.Println("Response:", string(body))

	// fmt.Println("Critical Success")

}

func recordsGet() {
	record := dto.Record{}
	fmt.Println("Please enter the record data (you may leave some fields blank, but at least one mustn't be blank):")

	fmt.Println("Name (leave blank if it's unknown):")
	fmt.Scanln(&record.Name)

	fmt.Println("Last name (leave blank if it's unknown):")
	fmt.Scanln(&record.LastName)

	fmt.Println("Middle name (leave blank if it's unknown):")
	fmt.Scanln(&record.MiddleName)

	fmt.Println("Address (leave blank if it's unknown):")
	fmt.Scanln(&record.Address)

	fmt.Println("Phone (leave blank if it's unknown):")
	fmt.Scanln(&record.Phone)

	if record.Name == "" && record.LastName == "" && record.MiddleName == "" && record.Address == "" && record.Phone == "" {
		fmt.Println("Error: all fields are empty")
		return
	}

	jsonData, err := json.Marshal(record)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// fmt.Println("jsonData =", string(jsonData)) // FOR DEBUGGING

	resp, err := http.Post("http://localhost:8080/get", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body := make([]byte, resp.ContentLength)
	_, err = io.ReadFull(resp.Body, body)
	if err != nil {
		fmt.Println("Error Reading Body:", err)
		return
	}

	ResponseHandler(body)
	// fmt.Println("Response:", string(body))

	// fmt.Println("Critical Success")

}

func recordUpdate() {
	record := dto.Record{}

	for record.Phone == "" {
		fmt.Println("Phone of record to update:")
		fmt.Scanln(&record.Phone)
	}

	fmt.Println("Please enter updated data:")

	fmt.Println("Name:")
	fmt.Scanln(&record.Name)

	fmt.Println("Last name:")
	fmt.Scanln(&record.LastName)

	fmt.Println("Middle name:")
	fmt.Scanln(&record.MiddleName)

	fmt.Println("Address:")
	fmt.Scanln(&record.Address)

	if record.Name == "" && record.LastName == "" && record.MiddleName == "" && record.Address == "" {
		fmt.Println("Error: all fields are empty")
		return
	}

	jsonData, err := json.Marshal(record)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// fmt.Println("jsonData =", string(jsonData)) // FOR DEBUGGING

	resp, err := http.Post("http://localhost:8080/update", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body := make([]byte, resp.ContentLength)
	_, err = io.ReadFull(resp.Body, body)
	if err != nil {
		fmt.Println("Error Reading Body:", err)
		return
	}

	ResponseHandler(body)
	// fmt.Println("Response:", string(body))

	// fmt.Println("Critical Success")
}

func recordDelete() {
	record := dto.Record{}
	fmt.Println("Please enter the record phone to delete:")
	fmt.Scanln(&record.Phone)

	jsonData, err := json.Marshal(record)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// fmt.Println("jsonData =", string(jsonData)) // FOR DEBUGGING

	resp, err := http.Post("http://localhost:8080/delete", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body := make([]byte, resp.ContentLength)
	_, err = io.ReadFull(resp.Body, body)
	if err != nil {
		fmt.Println("Error Reading Body:", err)
		return
	}

	ResponseHandler(body)
	// fmt.Println("Response:", string(body))

	// fmt.Println("Critical Success")
}

func ResponseHandler(body []byte) {
	resp := dto.Response{}
	err := json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println("Error in response:", err)
		return
	}

	if resp.Error != "" {
		fmt.Println(resp.Result)
		return
	}

	fmt.Println(resp.Result)

	if resp.Data != nil {
		data := []dto.Record{}
		err = json.Unmarshal(resp.Data, &data)
		if err != nil {
			fmt.Println("Error in response:", err)
			return
		}

		for _, record := range data {
			fmt.Println("Name:", record.Name)
			fmt.Println("Last name:", record.LastName)
			fmt.Println("Middle name:", record.MiddleName)
			fmt.Println("Address:", record.Address)
			fmt.Println("Phone:", record.Phone)
			fmt.Println()
		}
	}
	fmt.Println()
}
