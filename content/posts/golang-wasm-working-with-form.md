+++
title = "Golang WebAssembly: Compile Go to Wasm from scratch with examples"
date = "2023-02-10"
author = "prounckk"
tags = ["wasm", "go"]
keywords = ["wasm", "golang","go"]
description = "Golang WebAssembly: Compile Go to Wasm from scratch with examples"
showFullContent = false
draft = true
+++


WebAssambly was on my wishlist of things to try for a long time. My original idea was to write Rust code compiled to WASM, and the blog was created for this purpose. But times fly fast; my knowledge of Rust stayed on the same level while Golang progressed for a long mile. 

{{< highlight go "linenos=table,hl_lines=8 10-12,linenostart=199" >}}
func SubmitForm(this js.Value, args []js.Value) interface{} {
	form := Form{}
	form.name = js.Global().Get("document").Call("getElementById", "name").Get("value").String()
	form.email = js.Global().Get("document").Call("getElementById", "email").Get("value").String()
	form.company = js.Global().Get("document").Call("getElementById", "company").Get("value").String()
	form.companyniche = js.Global().Get("document").Call("getElementById", "companyniche").Get("value").String()
	form.message = js.Global().Get("document").Call("getElementById", "message").Get("value").String()

	return nil
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("SubmitForm", js.FuncOf(SubmitForm))
	<-make(chan bool)
}
{{< / highlight >}}