# go-wasm

## やったこと

- 関数を作成する
  - js から引数を受け取る
  - go コードを呼び出す
  - js のレスポンスとして返却する
- js 側で呼び出せるように Set する

## Commands

``` sh
go install github.com/mattn/serve@latest
serve -a :9999

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

## Links

- https://github.com/golang/go/issues/41310
- https://github.com/golang/vscode-go/issues/1799#issuecomment-979060844
