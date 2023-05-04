# go-wasm

sample: https://kokoichi206-sandbox.github.io/go-wasm/

## やったこと

- github api を go で呼び出し、js を使って html に表示する

### wasm の基本的な流れ

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

## 注意点

- `all goroutines are asleep - deadlock!`
  - api 呼び出しは別スレッドで行う必要がある
  - [goroutine で実行して promise で返す](https://github.com/golang/go/issues/41310)
- `could not import syscall/js`
  - [.vscode settings](https://github.com/golang/vscode-go/issues/1799#issuecomment-979060844)
- js の値として返せるものはかなり限られている！
  - https://pkg.go.dev/syscall/js#ValueOf

## Links

- https://github.com/golang/go/issues/41310
- https://github.com/golang/vscode-go/issues/1799#issuecomment-979060844
