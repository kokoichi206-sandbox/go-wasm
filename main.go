package main

import (
	"context"
	"syscall/js"

	"github.com/google/go-github/v52/github"
	"github.com/kokoichi206-sandbox/go-wasm/mygh"
)

var repo = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// https://github.com/golang/go/issues/41310#issuecomment-725809881
	if len(args) != 1 {
		return js.ValueOf("invalid number of arguments")
	}

	name := args[0].String()

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		resolve := args[0]
		// reject object
		// reject := args[1]

		go func() {
			repos, err := mygh.Repositories(context.Background(), name)

			resp := make(map[string]interface{})

			if err != nil {
				resp["error"] = err.Error()
			} else {
				resp["result"] = convertToInterface(repos)
			}

			resolve.Invoke(js.ValueOf(resp))
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
})

// js の値として返せるものはかなり限られている！
// https://pkg.go.dev/syscall/js#ValueOf
func convertToInterface(repos []*github.Repository) []interface{} {
	resp := make([]interface{}, 0, len(repos))
	for _, repo := range repos {
		resp = append(resp, *repo.Name)
	}

	return resp
}

func main() {
	ch := make(chan struct{}, 0)

	js.Global().Set("repositories", repo)
	<-ch
}
