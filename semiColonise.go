package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	flag.Parse() // get the arguments from command line

	arg := flag.Arg(0) // get the source directory from 1st argument

	if len(arg) > 3 && arg[len(arg)-3:len(arg)] == ".js" {
		log.Printf("Reading javascript file %s", arg)
		content, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Println(err)
		}

		re := regexp.MustCompile(".*[^\n]")
		jsLines := re.FindAllString(string(content), -1)
		output := content
		for i, _ := range jsLines {

			match, err := regexp.MatchString("(=)|(\\(.*\\))", jsLines[i])
			if err != nil {
				log.Printf("Error matching string")
			}
			if match && !strings.Contains(jsLines[i], "{") && !strings.Contains(jsLines[i], "if") && !strings.Contains(jsLines[i], "else") {
				//check for semi colon
				if !strings.Contains(jsLines[i], ";") {
					r := jsLines[i][0 : len(jsLines[i])-1]
					output = bytes.Replace(output, []byte(jsLines[i]), []byte(r+";"), -1)
				}
			}
		}

		out, err := os.Create("new" + arg)
		if err != nil {
			log.Println(err)
		}
		defer out.Close()

		out.Write(output)
		log.Printf("Success. File 'new%s' created",arg)
	} else {
		log.Printf("Wrong file type. Need .js")
		return
	}

}
