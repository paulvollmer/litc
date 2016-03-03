// GENERATED BY LITC v0.2.3
// DO NOT EDIT BY HAND

// # litc


	package main

	import (
		"fmt"
		"io/ioutil"
		"os"
		"strings"
	)

// # main

// the application is really lightweight and small in implementation.

	var version = "0.2.3"
	func main() {

// # input

	totalArgs := len(os.Args)
	if totalArgs == 1 {
		fmt.Println("")
		fmt.Println("  +--------+")
		fmt.Println("  | ---    | ")
		fmt.Printf("  | ---    |  litc v%s - literate preprocessor\n", version)
		fmt.Println("  |   ---  |  ---- ------ - ---------------------")
		fmt.Println("  |   ---  | ")
		fmt.Println("  | ---    | ")
		fmt.Println("  +--------+")
		fmt.Println("")
		fmt.Println("  litc <in> <out>")
		fmt.Println("  litc foo.go.lit foo.go\n")
		os.Exit(1)
	}

	var err error
	var inputData []byte

	if totalArgs >= 1 {
		// read file
		inputData, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

// # process

	lines := strings.Split(string(inputData), "\n")
	// change each non code line
	for i, v := range lines {
		if v != "" {
			if !strings.HasPrefix(v, "    ") {
				if !strings.HasPrefix(v, "\t") {
					lines[i] = "// " + v
				}
			}
		}
	}

// # output

	out := strings.Join(lines, "\n")
	if totalArgs == 3 {
		outPath := os.Args[2]
		outPre := "// GENERATED BY LITC v"+version+"\n// DO NOT EDIT BY HAND\n\n"
		err := ioutil.WriteFile(outPath, []byte(outPre+out), 0777)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// TODO: verbose mode
		// fmt.Println("successful written file to", outPath)
	} else {
		fmt.Println(out)
	}


	}
