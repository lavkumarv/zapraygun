# Raygun Hook for Zap

## Example

``` go
package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()

	raygunAppName := "RaygunAppName"
	raygunApiKey := "RaygunApiKey"

	// set the raygun credentials
	logger.WithOptions(
		zap.Hooks(zapraygun.NewRaygunHook(raygunAppName, raygunApiKey, zap.ErrorLevel).GetHook()),
	)

	logger.Error("Log this error message")
	logger.Info("Don't Log info level messages")
}
```

## Install

```
$ go get -u github.com/appointy/zapraygun
```

## Author
** Lav Vishwakarma **
* <lav@appointy.com>
