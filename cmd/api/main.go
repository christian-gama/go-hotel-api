package main

import (
	"os"

	"github.com/christian-gama/go-booking-api/internal"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func init() {
	// Clear terminal screen
	os.Stdout.Write([]byte("\033[H\033[2J"))
}

func main() {
	internal.Bootstrap()
}
