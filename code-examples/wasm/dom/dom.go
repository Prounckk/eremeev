package dom

import (
	"fmt"
	"syscall/js"
)
// ----------------------------------------------------------------
// This code is copied from the article:
// https://ian-says.com/articles/golang-in-the-browser-with-web-assembly/
// with some modifications by me (@prounckk)
// ----------------------------------------------------------------
var document js.Value

func init() {
    document = js.Global().Get("document")
}

func wrapGoFunction(fn func()) func(js.Value, []js.Value) interface {} {
    return func(_ js.Value, _ []js.Value) interface {} {
        fn()
        return nil
    }
}

func getElementById(elem string) js.Value {
    return document.Call("getElementById", elem)
}

func getElementValue(elem string, value string) js.Value {
    element := getElementById(elem)
    if element.IsNull() {
        return js.Null()
    }
    return element.Get(value)
}

func Hide(elem string) {
    getElementValue(elem, "style").Call("setProperty", "display", "none")
}

func Show(elem string) {
    getElementValue(elem, "style").Call("setProperty", "display", "block")
}

func AddEventListener(elem string, event string, fn func()) {
    getElementById(elem).Call("addEventListener", event, js.FuncOf(wrapGoFunction(fn)))
}

func GetStringFromElement(elem string) string {
    return GetStringFromElementByValue(elem, "value")
}

func GetStringFromElementByValue(elem string, value string) string {
    element := getElementValue(elem, value)
    if !element.IsNull() {
        return element.String()
    }
    fmt.Println("got empty "+ value + " for element: " + elem)
    return ""
}


func SetValue(elem string, key string, value string) {
    getElementById(elem).Set(key, value)
}
