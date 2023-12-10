package lib

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func File_part_from_cmdline(args []string) (file *os.File,part int) {
    file,err := os.Open(args[1])
    reg := regexp.MustCompile(`\d+`)
    day := reg.FindAllString(args[1],-1)
    if err != nil && (len(day) == 0 || slices.Contains(day,"test")) {
        log.Fatal("Failed to parse first cmd arg (or opening file)",err)
    } else {
        // assume this means input file doesn't exist, try to fetch from WEB
        client := &http.Client{}
        file,err = os.Open("../ADVENTCOOKIE.txt")
        ADVENTKEY := make([]byte,128)
        file.Read(ADVENTKEY); file.Close()
        req,_ := http.NewRequest("GET","https://adventofcode.com/2023/day/"+
                                day[0]+"/input",nil)
        req.AddCookie(&http.Cookie{Name: "session", Value: string(ADVENTKEY)})
        resp,err := client.Do(req)
        if err != nil {
            log.Fatal("Failed to fetch input from web")
        } else {
            fmt.Println("Fetching input from web..")
        }
        defer resp.Body.Close()
        input,_ := io.ReadAll(resp.Body)
        newfile,err := os.Create(args[1])
        n,err := newfile.Write(input)
        if err != nil {
            log.Fatal("Failed to write to input file")
        } else {
            fmt.Println("Wrote ",n," bytes to input file ",args[1])
            fmt.Println("It is now available for local reading on next call")
        }
        newfile.Close()
        file,err = os.Open(args[1]) 
        if err != nil {
            log.Fatal("problem parsing first cmd line arg (or opening file)",err)
        }
    }

    part,err = strconv.Atoi(args[2])
    if err != nil {
        log.Fatal("problem parsing second cmd line arg",err)
    }
    return file,part
}
