package main

import (
	"context"
	"flag"
	"fmt"
	"hedgedcurl/internal/hedgedcurl"
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

	var password string
	fmt.Printf(secfmt.Sprintf("Mio8RkoBRTo7MSZdVFxTOR0qASoaHBxcT1cWLh05GCcWFxBILhEKDAFVSA4KFgIXCwYDAAkuEDkBNwwBEB8cXE9YSlgzFyocJBEBAB8CQicbDAcCBhwCSE8MEhcLBgsRFiodPhssBwgBFwYGBgoYSEURDgocEzMQKQE0FDMc"), "asd", "wdssa")
	secfmt.Printf("MAEcAx1SFRMcFgEdFxZY")
	fmt.Scanf("%s", &password)

	if err := pass.PasswordCheck(password); err != nil {
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
