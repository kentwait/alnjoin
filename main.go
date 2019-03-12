package main

import (
	"flag"
	"fmt"
	"os"

	aln "github.com/kentwait/conspos/alignment"
)

// Exists returns whether the given file or directory Exists or not,
// and accompanying errors.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func main() {
	aln1Ptr := flag.String("aln1", "", "Path to first alignment (required)")
	aln2Ptr := flag.String("aln2", "", "Path to second alignment (required)")
	aln1RefNamePtr := flag.String("ref1", "", "Sequence name/identifier of sequence in aln1 to join on  (required)")
	aln2RefNamePtr := flag.String("ref2", "", "Sequence name/identifier of sequence in aln2 to join on (required)")

	flag.Parse()

	// Check arguments
	var missingArgs bool
	if len(*aln1Ptr) == 0 {
		os.Stderr.WriteString("[Error!] Path to first alignment was not specified: missing -aln1\n")
		missingArgs = true
	} else if doesExist, _ := Exists(*aln1Ptr); doesExist == false {
		os.Stderr.WriteString(fmt.Sprintf("[Error!] Alignment does not exist: %s.\n", *aln1Ptr))
		os.Exit(1)
	}

	if len(*aln2Ptr) == 0 {
		os.Stderr.WriteString("[Error!] Path to second alignment was not specified: missing -aln2\n")
		missingArgs = true
	} else if doesExist, _ := Exists(*aln2Ptr); doesExist == false {
		os.Stderr.WriteString(fmt.Sprintf("[Error!] Alignment does not exist: %s.\n", *aln2Ptr))
		os.Exit(1)
	}

	if len(*aln1RefNamePtr) == 0 {
		os.Stderr.WriteString("[Error!] Reference sequence name/identifier for the first alignment was not specified: missing -ref1\n")
		missingArgs = true
	}

	if len(*aln2RefNamePtr) == 0 {
		os.Stderr.WriteString("[Error!] Reference sequence name/identifier for the second alignment was not specified: missing -ref2\n")
		missingArgs = true
	}

	if missingArgs == true {
		os.Stderr.WriteString("\n")
		flag.Usage()
		os.Exit(1)
	}

	// Read fasta files and check if ref names exist
	aln1 := FastaToAlignment(*aln1Ptr)
	found1 := -1
	found2 := -1
	for i, seq := range aln1 {
		if *aln1RefNamePtr == seq.ID() {
			found1 = i
		}
	}
	if found1 == -1 {
		os.Stderr.WriteString(fmt.Sprintf("[Error!] %q was not found among the sequence identifiers in the first alignment.\n", *aln1RefNamePtr))
		os.Exit(1)
	}
	aln2 := FastaToAlignment(*aln2Ptr)
	for i, seq := range aln2 {
		if *aln2RefNamePtr == seq.ID() {
			found2 = i
		}
	}
	if found2 == -1 {
		os.Stderr.WriteString(fmt.Sprintf("[Error!] %q was not found among the sequence identifiers in the second alignment.\n", *aln2RefNamePtr))
		os.Exit(1)
	}

	// Join alignments
	joinedAln := leftJoin(aln1, aln2, found1, found2)
	// print to stdout
	fmt.Println(aln.ToString(joinedAln))

}
