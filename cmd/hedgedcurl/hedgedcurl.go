package main

import (
	"context"
	"flag"
	"fmt"
	"hedgedcurl/internal/hedgedcurl"
	"io"
	"os"
	"time"
)

var (
	timeout int
	help    bool
)

func init() {
	flag.IntVar(&timeout, "t", 15, "timeout")
	flag.IntVar(&timeout, "timeout", 15, "timeout (short)")
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&help, "h", false, "Show help (short)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] URL\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExample:")
	}
}

func main() {

	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	args := flag.Args()

	hedgedcurl.Start(args, ctx)

	select {
	case result := <-hedgedcurl.GetChan():
		wr := io.Writer(os.Stdout)
		wr.Write([]byte(result))
		return
	case <-ctx.Done():
		os.Exit(228)
	}
}
