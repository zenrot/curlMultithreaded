package main

import (
	"context"
	"flag"
	"fmt"
	"hedgedcurl/internal/hedgedcurl"
	"hedgedcurl/internal/secure/antidbg"
	"hedgedcurl/internal/secure/detectVm"
	"hedgedcurl/internal/secure/pass"
	"hedgedcurl/internal/secure/secfmt"
	"io"
	"os"
	"time"
)

var (
	timeout int
	help    bool
)

func init() {
	flag.IntVar(&timeout, secfmt.Sprintf("AQ=="),
		15,
		secfmt.Sprintf("AQYFAwAHEQ=="))
	flag.IntVar(&timeout, secfmt.Sprintf("AQYFAwAHEQ=="),
		15,
		secfmt.Sprintf("AQYFAwAHEVJHFh4dFwZL"))
	flag.BoolVar(&help, secfmt.Sprintf("HQoEFg=="),
		false,
		secfmt.Sprintf("JgcHEU8aAB4f"))
	flag.BoolVar(&help, secfmt.Sprintf("HQ=="),
		false,
		secfmt.Sprintf("JgcHEU8aAB4fRV4BDR0QEUY="))
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			secfmt.Sprintf("IBwJAQpIRVccRS09NSYrKiElMkIgPSQ6AS4L"),
			os.Args[0])
		fmt.Fprintln(os.Stderr,
			secfmt.Sprintf("Oh8cDwAcFkg="))
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr,
			secfmt.Sprintf("KQEtHg4fFR4KXw=="))
	}
}

func main() {

	junkCode()

	detectVm.CheckVM()
	antidbg.DebuggerSec()

	pass.Sec()

	var password string
	secfmt.Printf("MAEcAx1SFRMcFgEdFxZY")
	fmt.Scanf("%s", &password)

	if err := pass.FullPasswordCheck(password); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

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

func junkCode() {
	x := 0
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			_ = i * i
		} else if i%5 == 0 {
			_ = i + x
		}
	}
}
