package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"syscall/js"

	"github.com/prounckk/eremeev/code-examples/wasm/dom"
)

var CF_API_KEY string

var wg sync.WaitGroup

type Form struct {
	Email        string `json:"email"`
	Message      string `json:"message"`
	Name         string `json:"name"`
	Company      string `json:"company"`
	Companyniche string `json:"companyniche"`
	Position     string `json:"position"`
	Location     string `json:"location"`
	Salary       string `json:"salary"`
}

func main() {
	fmt.Println("Go Web Assembly")
	wg.Add(1)
	js.Global().Set("SubmitForm", js.FuncOf(SubmitForm))
	wg.Wait()
}

func SubmitForm(this js.Value, args []js.Value) any {
	form := Form{}
	form.Name = dom.GetStringFromElement("name")
	form.Email = dom.GetStringFromElement("email")
	form.Company = dom.GetStringFromElement("company")
	form.Companyniche = dom.GetStringFromElement("companyniche")
	form.Message = dom.GetStringFromElement("message")
	form.Position = dom.GetStringFromElement("position")
	form.Location = dom.GetStringFromElement("location")
	form.Salary = dom.GetStringFromElement("salary")

	go form.sendEmail()
	return nil
}

func (form Form) sendEmail() {
	body, err := json.Marshal(form)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", "https://eremeev.ca/contact-form", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+CF_API_KEY)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	defer wg.Done()
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	return
}
