package main

import (
	"os"
	"fmt"
	"bufio"

	"github.com/azillion/edgcm-converter/climate"
	"github.com/azillion/edgcm-converter/externalinputservice"
)



func main() {
	exInputService := NewExternalInputService()
	f, err := os.Open("./input-data/sote.dat")
	if err != nil {
		os.Exit(-1)
	}
	soteDatas := sotelib.ReadDat(f)

	f2, err := os.OpenFile("./input-data/sote.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(-1)
	}
	
    datawriter := bufio.NewWriter(f2)
 
	for _, data := range soteDatas {
        s := fmt.Sprintf("soteData:%+v\n", data)
		_, _ = datawriter.WriteString(s)
	}
 
    datawriter.Flush()
    f2.Close()
}