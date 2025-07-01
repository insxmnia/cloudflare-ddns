package recovery

import (
	"cloudflare-ddns/pkg/slogger"
)

func Infinite(fn func()) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					slogger.Error("go-recovery", "application recovered from panic", "", "stacktrace", r, "restart", "initiated")
				}
			}()
			fn()
		}()
	}

}
