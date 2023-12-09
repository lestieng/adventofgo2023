package lib

import (
    "os"
    "log"
    "strconv"
)

func File_part_from_cmdline(args []string) (file *os.File,part int) {
    file,err := os.Open(args[1])
    if err != nil {
        log.Fatal("problem parsing first cmd line arg (or opening file)",err)
    }

    part,err = strconv.Atoi(args[2])
    if err != nil {
        log.Fatal("problem parsing second cmd line arg",err)
    }
    return file,part
}
