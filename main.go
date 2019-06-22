package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

var flgVersion bool
var flgVerbose bool

func main() {
	rootCmd := flag.NewFlagSet("dailyrepo root", flag.ContinueOnError)
	rootCmd.BoolVar(&flgVersion, "version", false, "print version")
	rootCmd.BoolVar(&flgVersion, "v", false, "print version")
	rootCmd.BoolVar(&flgVersion, "verbose", false, "print log")

	err := rootCmd.Parse(os.Args[1:])
	if err != nil {
		if err != flag.ErrHelp {
			os.Exit(0)
		}
		log.Fatal(err)
	}
	if flgVersion {
		fmt.Println("dailyrepo v0.0.1")
	}

	addCmd := flag.NewFlagSet("add", flag.ContinueOnError)
	var fileName string
	addCmd.StringVar(&fileName, "name", time.Now().Format("2006-01-02")+".md", "specify generating filename")

	args := rootCmd.Args()
	if len(args) > 0 {
		switch args[0] {
		case "add":
			_ = addCmd.Parse(args[1:])
			handleAddCmd(fileName)
		default:
			os.Exit(2)
		}
	}
	os.Exit(0)
}

func handleAddCmd(filename string) error {
	btpl, _ := ioutil.ReadFile("./templates/report.md.tmpl")
	stpl := string(btpl)

	tpl := template.Must(template.New("report").Parse(stpl))
	rptFile, _ := os.Create(filename)
	rptData := struct {
		Today string
	}{
		Today: time.Now().Format("2006-01-02"),
	}
	_ = tpl.Execute(rptFile, rptData)
	return nil
}
