# litc

the application is really lightweight and small in implementation.

	package main

	import (
		"fmt"
		"flag"
		"io/ioutil"
		"os"
		"github.com/paulvollmer/litc/golang/lit"
	)

# commandline interface

set flags to global var

	var (
		flagRemove = flag.Bool("remove", false, "remove non code lines from source")
	)

the commandline usage output

	func usage() {
		fmt.Println("")
		fmt.Println("  +--------+")
		fmt.Println("  | ---    | ")
		fmt.Printf("  | ---    |  litc v%s - literate preprocessor\n", lit.Version)
		fmt.Println("  |   ---  |  ---- ------ - ---------------------")
		fmt.Println("  |   ---  | ")
		fmt.Println("  | ---    | ")
		fmt.Println("  +--------+")
		fmt.Println("")
		fmt.Println("  litc <in> <out> [args]")
		fmt.Println("")
		fmt.Println("  arguments:")
		flag.PrintDefaults()
	}

parse the flags

	func init() {
		flag.Parse()
	}

# main

	func main() {

check the input arguments

	totalArgs := len(os.Args)

without arg...

	if totalArgs == 1 {
		usage()
		os.Exit(1)
	}

more than one arg...

	if totalArgs >= 1 {

		switch os.Args[1] {
		case "version":
			fmt.Println("litc v"+lit.Version)
			os.Exit(0)
			break
		case "help":
			usage()
			os.Exit(1)
		}

		d, err := lit.ProcessFile(os.Args[1], *flagRemove)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if totalArgs >= 3 {
			outPath := os.Args[2]
			err := ioutil.WriteFile(outPath, []byte(d), 0777)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// fmt.Println("successful written file to", outPath)
		} else {
			fmt.Println(string(d))
		}
	}


main end

	}
