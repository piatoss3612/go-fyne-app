# Go Fyne App

### 1. Markdown Editor

#### Testing

```cmd
$ go test -v
=== RUN   Test_MakeUI
--- PASS: Test_MakeUI (0.00s)
=== RUN   Test_RunApp
--- PASS: Test_RunApp (0.00s)
PASS
ok      fynemd  0.163s
```

#### Packaging

```cmd
$ fyne package -appVersion 1.0.0 -name Markdown -appID piatoss.tech.markdown -release
```

#### Run Application

```
$ ./Markdown
```

### 2. Gold Watcher

#### Bundling Resource

```cmd
$ fyne bundle unreachable.png >> bundled.go
```

#### Testing

```cmd
$ go test -v
=== RUN   TestGold_GetPrices
--- PASS: TestGold_GetPrices (0.00s)
=== RUN   TestApp_getPriceText
--- PASS: TestApp_getPriceText (0.00s)
=== RUN   TestApp_getToolbar
--- PASS: TestApp_getToolbar (0.00s)
PASS
ok      gold-watcher    0.127s
```

### Reference

- [Building GUI Applications with Fyne and Go (Golang)](https://www.udemy.com/course/building-gui-applications-with-fyne-and-go-golang/)
