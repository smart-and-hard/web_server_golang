package sl

import (
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/exp/slog"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
