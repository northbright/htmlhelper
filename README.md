# htmlhelper

[![Build Status](https://travis-ci.org/northbright/htmlhelper.svg?branch=master)](https://travis-ci.org/northbright/htmlhelper)
[![Go Report Card](https://goreportcard.com/badge/github.com/northbright/htmlhelper)](https://goreportcard.com/report/github.com/northbright/htmlhelper)
[![GoDoc](https://godoc.org/github.com/northbright/htmlhelper?status.svg)](https://godoc.org/github.com/northbright/htmlhelper)

htmlhelper is a [Golang](https://golang.org) package which provides helper functions for HTML string operations.

#### Features
* [Clean HTML string](https://godoc.org/github.com/northbright/htmlhelper#ex-Clean).
        
        var str = `
        </td>
                                    <td 
                                        class="tint"
                                
                                    >
   
        <a
   
          href="" >fake link   

        </a>
        `

        str = htmlhelper.Clean(str)
        fmt.Printf("cleaned HTML: %v\n", str)

        // Output:
        //cleaned HTML: </td><td class="tint"><a href="">fake link</a>

* [Convert HTML Tables into CSV Records](https://godoc.org/github.com/northbright/htmlhelper#ex-TablesToCSVs).

        var str = `<table>.......</table>
                   .....<table>....</table>`

        csvs := htmlhelper.TablesToCSVs(str)
        for i, csv := range csvs {
                fmt.Printf("%v: %v\n", i, csv)
        }

        // Output:
        //0: [[Song Name Artist Album] [squall Pasteboard glitter] [Recollection GIN MaHaLo]]
        //1: [[Song Name Artist Album] [I Feel Love Mystic Diversions Beneath Another Sky] [紫陽花 GIN MaHaLo]]

#### Documentation
* [API References](https://godoc.org/github.com/northbright/htmlhelper)

#### License
* [MIT License](./LICENSE)


