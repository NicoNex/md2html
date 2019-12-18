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

func addStyle(in string) string {
    var style = `<style>
    h1,h2,h3,h4,h5,h6,p,blockquote{margin:0;padding:0}body{font-family:"Helvetica Neue",Helvetica,Arial,sans-serif;font-size:13px;line-height:18px;color:#fff;background-color:#110F14;margin:10px 13px}table{margin:10px 0 15px;border-collapse:collapse}td,th{border:1px solid #ddd;padding:3px 10px}th{padding:5px 10px}a{color:#59acf3}a:hover{color:#a7d8ff;text-decoration:none}a img{border:none}p{margin-bottom:9px}h1,h2,h3,h4,h5,h6{color:#fff;line-height:36px}h1{margin-bottom:18px;font-size:30px}h2{font-size:24px}h3{font-size:18px}h4{font-size:16px}h5{font-size:14px}h6{font-size:13px}hr{margin:0 0 19px;border:0;border-bottom:1px solid #ccc}blockquote{padding:13px 13px 21px 15px;margin-bottom:18px;font-family:georgia,serif;font-style:italic}blockquote:before{content:"\201C";font-size:40px;margin-left:-10px;font-family:georgia,serif;color:#eee}blockquote p{font-size:14px;font-weight:300;line-height:18px;margin-bottom:0;font-style:italic}code,pre{font-family:Menlo,Monaco,Andale Mono,Courier New,monospace}code{padding:1px 3px;font-size:12px;-webkit-border-radius:3px;-moz-border-radius:3px;border-radius:3px;background:#334}pre{display:block;padding:14px;margin:0 0 18px;line-height:16px;font-size:11px;border:1px solid #334;white-space:pre;white-space:pre-wrap;word-wrap:break-word;background-color:#282a36;border-radius:6px}pre code{font-size:11px;padding:0;background:transparent}sup{font-size:.83em;vertical-align:super;line-height:0}*{-webkit-print-color-adjust:exact}@media screen and (min-width: 914px){body{width:854px;margin:10px auto}}@media print{body,code,pre code,h1,h2,h3,h4,h5,h6{color:#000}table,pre{page-break-inside:avoid}}
</style>`
    return fmt.Sprintf("%s\n%s", style, in)
}

func main() {
    var html string
    var args []string
    var outfile string
    var style bool

    flag.StringVar(&outfile, "o", "", "Output file")
    flag.BoolVar(&style, "s", true, "Adds the style")
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
    if style {
        html = addStyle(html)
    }

    if outfile != "" {
        ofile, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, 0666)
        check(err)
        ofile.WriteString(html)
        ofile.Close()
    } else {
        fmt.Println(html)
    }
}
