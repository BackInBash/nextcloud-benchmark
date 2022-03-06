package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/studio-b12/gowebdav"
)

type cliArgs struct {
	url  string
	user string
	pass string
	path string
	num  int
}

func printHelp() {
	fmt.Println("Commandline Parameter:\n" +
		"-url 	Nextcloud Root URL\n" +
		"-user 	Nextcloud Username\n" +
		"-pass	Nextcloud Password\n" +
		"-path	Nextcloud Folder Path\n" +
		"-num	Number of Files to Upload")
}

func main() {
	fmt.Println("Nextcloud Benchmark Tools - Random File Creator")
	var cli cliArgs
	// Pase CLI Arguments
	for index := 0; index < len(os.Args); index++ {
		switch os.Args[index] {
		case "-url":
			cli.url = os.Args[index+1]
		case "-user":
			cli.user = os.Args[index+1]
		case "-pass":
			cli.pass = os.Args[index+1]
		case "-path":
			cli.path = os.Args[index+1]
		case "-num":
			i, err := strconv.Atoi(os.Args[index+1])
			if err != nil {
				fmt.Print("Error: Parsing -num Input")
				os.Exit(1)
			}
			cli.num = i
		}
	}
	// Check CLI Args
	if len(os.Args) == 0 {
		printHelp()
		return
	}
	// Connect to WebDAV
	// Remove Ending / if Exists
	if strings.HasSuffix(cli.url, "/") {
		re := regexp.MustCompile(`/$`)
		cli.url = re.ReplaceAllString(cli.url, "")
	}

	root := cli.url + "/remote.php/dav/files/" + cli.user + "/"
	fmt.Println(root)

	c := gowebdav.NewClient(root, cli.user, cli.pass)

	// Create File Path
	err := c.MkdirAll(cli.path, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// Temp File
	f := make([]byte, 0x500000)
	for i := 0; i < cli.num; i++ {
		path := cli.path + "/x" + strconv.Itoa(i)
		err := c.MkdirAll(path, 0644)
		if err != nil {
			fmt.Println(err)
		}
		for x := 0; x < 5; x++ {
			c.Write(path+"/test"+strconv.Itoa(x)+".bin", f, 0644)
		}
	}
}
