package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"syscall/js"

	"github.com/prounckk/eremeev/code-examples/wasm/dom"
)

var SENDGRID_API_KEY = os.Getenv("SENDGRID_API_KEY")

var wg sync.WaitGroup // 1

type Form struct {
	email        string `json:"email"`
	message      string `json:"message"`
	name         string `json:"name"`
	company      string `json:"company"`
	companyniche string `json:"companyniche"`
	position     string `json:"position"`
	location     string `json:"location"`
	salary       string `json:"salary"`
}

func main() {
	fmt.Println("Go Web Assembly")
	wg.Add(1)
	js.Global().Set("SubmitForm", js.FuncOf(SubmitForm))
	wg.Wait()

}

func SubmitForm(this js.Value, args []js.Value) any {
	form := Form{}
	form.name = dom.GetStringFromElement("name")
	form.email = dom.GetStringFromElement("email")
	form.company = dom.GetStringFromElement("company")
	form.companyniche = dom.GetStringFromElement("companyniche")
	form.message = dom.GetStringFromElement("message")
	form.position = dom.GetStringFromElement("position")
	form.location = dom.GetStringFromElement("location")
	form.salary = dom.GetStringFromElement("salary")

	go form.sendEmail()
	return nil
}

func (form Form) sendEmail() {
	emailBody := map[string]interface{}{
		"from": map[string]string{
			"email": "me@eremeev.ca",
		},
		"personalizations": []map[string]interface{}{
			{
				"to": []map[string]string{
					{
						"email": "me@eremeev.ca",
					},
				},
				"subject": form.location + " " + form.position,
				"dynamic_template_data": map[string]string{
					"sender":   form.email,
					"name":     form.name,
					"message":  form.message,
					"subject":  form.location + " | " + form.position,
					"company":  form.company,
					"position": form.position,
					"salary":   form.salary,
					"location": form.location,
				},
			},
		},
		"template_id": "d-6f8ef6f0f3ec4fc0bdae401d050babee",
	}

	body, err := json.Marshal(emailBody)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+SENDGRID_API_KEY)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	defer wg.Done()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Println(res.Status)
	return
}
