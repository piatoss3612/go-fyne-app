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

#### Packaging on Windows

```cmd
$ fyne package -appVersion 1.0.0 -name Markdown -appID piatoss.tech.markdown -release
```

#### Run Application

```cmd
$ ./Markdown
```

### 2. Gold Watcher

#### Bundling Resource

```cmd
$ fyne bundle unreachable.png >> bundled.go
```

#### Testing Main Package

```cmd
$ go test -v
=== RUN   TestGold_GetPrices
=== RUN   TestConfig_getHoldings
--- PASS: TestConfig_getHoldings (0.00s)
=== RUN   TestConfig_getHoldingSlice
--- PASS: TestConfig_getHoldingSlice (0.15s)
=== RUN   TestApp_getPriceText
--- PASS: TestApp_getPriceText (0.00s)
=== RUN   TestApp_getToolbar
--- PASS: TestApp_getToolbar (0.00s)
=== RUN   TestApp_addHoldingsDialog
--- PASS: TestApp_addHoldingsDialog (0.01s)
PASS
ok      gold-watcher    0.300s
```

#### Testing Repository Package

```cmd
$ go test -v
=== RUN   TestSQLiteRepo_Migrate
--- PASS: TestSQLiteRepo_Migrate (0.01s)
=== RUN   TestSQLiteRepo_InsertHolding
--- PASS: TestSQLiteRepo_InsertHolding (0.01s)
=== RUN   TestSQLiteRepo_AllHoldings
--- PASS: TestSQLiteRepo_AllHoldings (0.00s)
=== RUN   TestSQLiteRepo_GetHondingByID
--- PASS: TestSQLiteRepo_GetHondingByID (0.01s)
=== RUN   TestSQLiteRepo_UpdateHolding
--- PASS: TestSQLiteRepo_UpdateHolding (0.05s)
=== RUN   TestSQLiteRepo_DeleteHolding
--- PASS: TestSQLiteRepo_DeleteHolding (0.01s)
PASS
ok      gold-watcher/repository 0.147s
```

#### Packaging with Makefile

```cmd
$ make build
```

#### Run Application

```cmd
$ ./GoldWatcher
```

### Reference

- [Building GUI Applications with Fyne and Go (Golang)](https://www.udemy.com/course/building-gui-applications-with-fyne-and-go-golang/)
