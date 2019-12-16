package main

import (
    "os"
    "fmt"
    "flag"
    "context"
    "io/ioutil"

    "github.com/google/go-github/github"
)

const PROGRAM_NAME = "md2html"

func check(e error) {
    if e != nil {
        fmt.Println(e)
        os.Exit(1)
    }
}

func usage() {
    var msg = `md2html - Convert markdown to html using GitHub api.
Options:
    -o string
        Output file
Usage:
    md2html [OPTIONS] input.md`
    fmt.Println(msg)
}

func convert(in []byte) string {
    var client = github.NewClient(nil)
    var opt = &github.MarkdownOptions{Mode: "markdown"}

    output, _, err := client.Markdown(context.Background(), string(in), opt)
    check(err)
    return output
}

func main() {
    var html string
    var args []string
    var outfile string

    flag.StringVar(&outfile, "o", "", "Output file")
    flag.Usage = usage
    flag.Parse()
    args = flag.Args()
    if len(args) == 0 {
        usage()
        return
    }
    md, err := ioutil.ReadFile(args[0])
    check(err)
    html = convert(md)

    if outfile != "" {
        ofile, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, 0666)
        check(err)
        ofile.WriteString(html)
        ofile.Close()
    } else {
        fmt.Println(html)
    }
}
