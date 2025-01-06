# go-android-log

Native Android logging bindings for Go apps.
Recommended to use with [gomobile](https://pkg.go.dev/golang.org/x/mobile).

## Install
```bash
go get github.com/Securepoint/go-android-log
```

## Usage
```go
logger := androidlog.NewLogger("MyPackage")
logger.Info("Application started")
logger.Debug("Debug message")
logger.Error("Something went wrong")
```

### Formatted messages
```go
logger.Infof("Application started at %s", time.Now())
logger.Debugf("Debug message: %d", 42)
logger.Errorf("Something went wrong: %v", err)
```

## License
[MIT](LICENSE)