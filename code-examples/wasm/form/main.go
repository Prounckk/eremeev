package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
	"syscall/js"

	"github.com/prounckk/eremeev/code-examples/wasm/dom"
)

var CF_API_KEY string

type Form struct {
	Email        string `json:"email"`
	Message      string `json:"message"`
	Name         string `json:"name"`
	Company      string `json:"company"`
	Companyniche string `json:"companyniche"`
	Position     string `json:"position"`
	Location     string `json:"location"`
	Salary       string `json:"salary"`
	Error        string `json:"error"`
	Thanks       string `json:"thanks"`
}

func main() {
	println("Go WebAssembly - Form submission")
	js.Global().Set("SubmitForm", js.FuncOf(SubmitForm))
	<-make(chan bool)
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
	form.Thanks = "Thank you for your message!<br><a href=\"/articles/\">Click here to check my latest posts</a>"
	
	ch := make(chan Form)
	go form.sendEmail(ch)
	error := <-ch
	if error.Error != "" {
		println(error.Error)
		form.Thanks = "Oops! It looks like my wasm code failed: <br>" + error.Error + "<br><a class=\"section-button\" href=\"/contact-me/\">Back to the form</a>"
	}

	dom.Hide("formcontact")
	dom.Show("after-form")
	dom.SetValue("after-form", "innerHTML", form.Thanks)

	return nil
}

func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return emailRegex.MatchString(e)
}

func (form *Form) sendEmail(ch chan Form) {
	if !isEmailValid(form.Email) {
		form.Error = "A valid email is required, sorry."
		ch <- *form
		return
	}
	body, err := json.Marshal(form)
	if err != nil {
		form.Error = err.Error()
		ch <- *form
		return
	}

	req, err := http.NewRequest("POST", "https://eremeev.ca/contact-form", bytes.NewBuffer(body))
	if err != nil {
		form.Error = err.Error()
		ch <- *form
		return
	}
	req.Header.Set("Authorization", "Bearer "+CF_API_KEY)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	go func() {
		res, err := client.Do(req)
		if err != nil {
			form.Error = err.Error()
			ch <- *form
			return
		}
		defer res.Body.Close()

	}()

	ch <- *form
	return
}
