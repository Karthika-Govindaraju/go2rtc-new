package api

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
)

var stackSkip = [][]byte{
	// debug.go
	[]byte("github.com/AlexxIT/go2rtc/cmd/debug.handler"),

	// cmd.go
	[]byte("github.com/AlexxIT/go2rtc/cmd.Run"),
	[]byte("created by os/signal.Notify"),

	// api.go
	[]byte("created by github.com/AlexxIT/go2rtc/cmd/api.Init"),
	[]byte("created by net/http.(*connReader).startBackgroundRead"),
	[]byte("created by net/http.(*Server).Serve"),

	[]byte("created by github.com/AlexxIT/go2rtc/cmd/rtsp.Init"),
}

func stackHandler(w http.ResponseWriter, r *http.Request) {
	sep := []byte("\n\n")
	buf := make([]byte, 65535)
	i := 0
	n := runtime.Stack(buf, true)
	skipped := 0
	for _, item := range bytes.Split(buf[:n], sep) {
		for _, skip := range stackSkip {
			if bytes.Contains(item, skip) {
				item = nil
				skipped++
				break
			}
		}
		if item != nil {
			i += copy(buf[i:], item)
			i += copy(buf[i:], sep)
		}
	}
	i += copy(buf[i:], fmt.Sprintf(
		"Total: %d, Skipped: %d", runtime.NumGoroutine(), skipped),
	)

	if _, err := w.Write(buf[:i]); err != nil {
		panic(err)
	}
}