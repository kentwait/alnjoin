package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"

	aln "github.com/kentwait/conspos/alignment"
	"github.com/kentwait/conspos/sequence"
)

// FastaToAlignment reads a FASTA file into an Alignment struct
func FastaToAlignment(path string) (sequences aln.Alignment) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var id, title string
	var seqBuffer bytes.Buffer
	var splitted []string

	var line string
	for {
		line, err = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if strings.HasPrefix(line, ">") {
			if seqBuffer.Len() > 0 {
				sequences = append(sequences, sequence.NewCharSequence(id, title, seqBuffer.String()))
				seqBuffer.Reset()
			}
			splitted = strings.SplitN(line[1:], " ", 2)
			id = splitted[0]
			if len(splitted) == 2 {
				title = splitted[1]
			}
		} else if strings.HasPrefix(line, "\n") {
			continue
		} else if strings.HasPrefix(line, "#") {
			continue
		} else if strings.HasPrefix(line, ";") {
			continue
		} else {
			seqBuffer.WriteString(line)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			os.Stderr.WriteString("[Error!] alignment file may be malformed")
			os.Exit(2)
		}
	}
	if seqBuffer.Len() > 0 {
		sequences = append(sequences, sequence.NewCharSequence(id, title, seqBuffer.String()))
	}
	return
}
