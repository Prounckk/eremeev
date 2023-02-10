+++
title = "Golang WebAssembly: Compile Go to Wasm from scratch with examples"
date = "2023-02-11"
author = "prounckk"
tags = ["wasm", "go"]
keywords = ["wasm", "golang","go", "WebAssambly"]
description = "Golang WebAssembly: Compile Go to Wasm from scratch with examples"
showFullContent = false
+++


WebAssambly was on my wishlist of things to try for a long time. My original idea was to write Rust code compiled to WASM, and the blog was created for this purpose. But times fly fast; my knowledge of Rust stayed on the same level while Golang progressed for a long mile.   

Before I jump into implementation, I want to say that all my pre-wasm code is written in Go, but you can use what you want; I could find compilers for Rust, PHP, C++, etc. 
In this article will be a lot of Go-WASM-related comments, so please check [my cat's Instagram](https://www.instagram.com/cat_watson_mtl/) if golang is not your area of interest. 

# How to start with WASM
## Create a golang file
Here is an elementary example of code. Let's look at it line by line.
{{< highlight go "linenos=table,hl_lines=8 10-12,linenostart=199" >}}

package main

import (
    "syscall/js"
)
func main() {
	document := js.Global().Get("document")
	document.Call("getElementById", "calculate")
	<-make(chan bool)
}

{{< / highlight >}}
- `"syscall/js"` - is the package that allows us to communicate with browsers using JavaScript APIs. There are lots of packages on the internet, but the core of all of them is "syscall/js".
- `document := js.Global().Get("document"`)` - here, we are loading DOM to work with it later.
- `document.Call("getElementById", "calculate")` - and this is our call to an element with id `calculate`. JavaScript alternative to this is `document.getElementById("calculate")`
- `<-make(chan bool)`- here is a catch! To keep our app running and not exiting right after the getElementById, let's create a channel waiting for an event but never getting one.

## Compile to WASM
So far, nothing complicated. Let's compile the code and check in the browser
```bash
GOOS=js GOARCH=wasm go build -o  ../../static/js/page.wasm
```
- `GOOS=js GOARCH=wasm` - flag to generate binary file for js/wasm architecture.
- `-o <path>/name.wasm` - flag to specify the destination and name of the binary file.
Needless to say, the build should happen in the same directory as the golang file. 

## Loading in the browser
Now we have generated WASM binary file, but how to load it from the page? Sorry, we can't skip JavaScript here ¯\\_(ツ)_/¯

```html
    <script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("/js/page.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });
</script>
```  
... and we got an error:  `Uncaught ReferenceError: Go is not defined`  
It looks like a browser needs a guide on how to use Go on Javascript Wasm execution architecture. Happily, it shipped with Go. Please don't move it, but copy it!
```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" $PWD      
```
and now the code should look like this:
```html
   <script src="/js/wasm_exec.js"></script>
   <script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("/js/page.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });
</script>
```
I prefer to store the files in the /js/ directory, but nobody would blame you for using a different folder structure.

## DOM manipulation from WASM
The default package is not user-friendly. For example, to get a value from the element `calculate` the code should look something like this:
```golang
js.Global().Get("document").Call("getElementById", "calculate").Get("value").String()
```
Now you see why there are so many wrapper packages on the Internet ;) I did my wrapper as well and saved it as a file `dom.go` and now it looks a little bit shorter:
```golang
input := dom.GetStringFromElement("calculate")
```
All we did is amazing, but what are the benefits of this? Let's do a test!
```javascript
function veryBadForLoops(n) {
	res = 0;
	for (let i = 0; i < n; i++) {
		for (let j = 0; j < n; j++) {
			res = res + j + i
		}
	}
	return res;
}
```
VS
```go
func veryBadForLoopsWasm(this js.Value, args []js.Value) any {
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res = res + j + i
		}
	}
	return res	
}
```


Be aware that big numbers might freeze your browsers. Start with something small.


{{< js_calculator >}}

I don't know your feelings, but I'm impressed with this performance! 
Let'see why!
 - WASM is an assembly language for the browser but not for hardware. WebAssembly can't be executed on a machine without a browser, but still, it's closer to the hardware than JavaScript.
 - WASM is already compiled to a binary, while JavaScript needs to be compiled on runtime.
 - WASM was designed to be fast.


At the same time, WebAssembly is not a JavaScript replacement. It is designed to operate hand in hand with JavaScript to take care of performance-critical components of a web application. As you can see, browsers still need JS to load wasm properly.

# Form submission
My second experiment was with a form. I know there is nothing simpler than `<form action="/foo/bar" method="post">`, but what if the recipient needs credentials? Ah? Here it begins a little bit complicated. 


{{< highlight go "linenos=table,hl_lines=8 10-12,linenostart=199" >}}
...
var wg sync.WaitGroup

func main() {
	fmt.Println("Go Web Assembly")
	wg.Add(1)
	js.Global().Set("SubmitForm", js.FuncOf(SubmitForm))
	wg.Wait()
}
func SubmitForm(this js.Value, args []js.Value) interface{} {
	form := Form{}
	form.Name = dom.GetStringFromElement("name")
	form.Email = dom.GetStringFromElement("email")
	form.Company = dom.GetStringFromElement("company")
	go form.sendEmail()
	return nil
}

func (form Form) sendEmail() {
	body, err := json.Marshal(form)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", "https://eremeev.ca/contact-form", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
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
...
{{< / highlight >}}

This code is slightly different from the previous version. 
I've implemented WaitGroup here with the same idea of not exiting while working.
- `js.Global().Set("SubmitForm", js.FuncOf(SubmitForm))` - this is also something new. We expose the function SubmitForm to JavaScript with the same name, and it can be called as any other JavaScript function, like:
```html
 <button type="submit" class='form-submit' onsubmit="SubmitForm()">Submit</button>
```
You definitely noticed I do not return errors but print them. This was done to make debugging easier, but it's not a good practice. Please, don't be Sergei. 

Fun fact! `CF_API_KEY` is undefined in the console, which is great for security reasons. But unfortunately, it's not a working solution for public repositories like this one. 
Please, don't be shy to check the git history to see that initially, I was trying to send email from WASM but had to move email sending logic to the backend, which is Cloudflare Worker in my case.
Why? Because my SendGrid Account was blocked two times! First time I committed code with the API Key. Shame on me. But the second time, the code was added to the binary file during the build. What could go wrong? Well, the WASM file still contains the build parameters in the header. So in a minute from my git push command, I've got a second account suspension. Kudos to SendGrid for tracking leaked API keys on Github!

```text
ítA·¡Ê÷Êpath	github.com/prounckk/eremeev/code-examples/wasm
mod	github.com/prounckk/eremeev	(devel)	
build	-compiler=gc
build	-ldflags="-X main.SENDGRID_API_KEY=blablablarblablablablafoooffooo"
build	CGO_ENABLED=0
build	GOARCH=wasm
build	GOOS=js
build	vcs=git
build	vcs.revision=d6f2f0437f25b76cdb6d844301a9a342fce2f0a8
build	vcs.modified=true
```
Don't worry; the API Key is already deactivated. 

# Pros and Cons

Important note: this is my first experience with WASM. Probably some of the points here are just my misunderstanding of the concept.
CONS:
 - WASM still relies on JavaScript API to communicate with [DOM](https://developer.mozilla.org/en-US/docs/Web/API/Document)
 - `syscall/js` - is still in development, documentation is missing examples, and the package's functionality is still limited.
 - My hope to deploy secrets secured way to the front end was broken. The keys are unavailable from the console but stored in the file's header. I think it should be a way to solve it!
- WASM files are heavy! See the screenshot. But again, I think it should be a way to solve it!
![](/2023/wasm-file-size.png )
- Writing golang code with js.Value is contreintuitive. There are so many [conversions between types](https://pkg.go.dev/syscall/js#ValueOf), and error handling is less intuitive than with pure golang.
PROS:
- Write FrontEnd code with the language you are more comfortable with.
- In theory, everything can be compiled and used as FrontEnd applications, even complicated games.
- It is a Server-side, ahead-of-time compiling application, so it should use fewer resources to start and run.

# Conclusion
Will I use WASM in the future - definitely, yes! Do I know how? Not yet, but I feel it has potential and, with good support from the community, can reach a bright future!   

# Links to documentation and other resources
- [WebAssembly official website](https://webassembly.org)
- [Things Ian Says](https://ian-says.com/articles/golang-in-the-browser-with-web-assembly/) I've learned a lot from this article and copied some implementation.
- [syscall/js](https://pkg.go.dev/syscall/js) The most important package for this blog post.
- [source code](https://github.com/Prounckk/eremeev/tree/main/code-examples/wasm) Don't judge me. This is for fun.