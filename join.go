package main

import (
	"fmt"

	fa "github.com/kentwait/gofasta"
)

// Join joins 2 alignments using a common sequence.
func Join(aln1, aln2 fa.Alignment, refindex1, refindex2 int, method string) fa.Alignment {
	switch method {
	case "left": // SQL LEFT JOIN
	case "right": // SQL RIGHT JOIN
	// case "full": // SQL OUTER JOIN
	default:
		panic("Invalid method")
	}
	ref1UngappedSlice := aln1[refindex1].UngappedPositionSlice("-")
	ref2UngappedSlice := aln2[refindex2].UngappedPositionSlice("-")
	if method == "right" {
		ref1UngappedSlice, ref2UngappedSlice = ref2UngappedSlice, ref1UngappedSlice
		aln1, aln2 = aln2, aln1
		refindex1, refindex2 = refindex2, refindex1
	}

	ref1Rune := []rune(aln1[refindex1].Sequence())
	ref2Rune := []rune(aln2[refindex2].Sequence())

	// Initialize empty slice
	emptyAln := make([]fa.Sequence, len(aln2))
	i := 0
	for _, seq := range aln2 {
		j := 0
		runedSeq := []rune(seq.Sequence())
		newRunedSeq := make([]rune, len([]rune(aln1[0].Sequence())))
		for k, pos := range ref2UngappedSlice {
			if ref1UngappedSlice[j] < pos {
				for ref1UngappedSlice[j] < pos {
					newRunedSeq[j] = rune('-')
					j++
				}
			} else if ref1UngappedSlice[j] > pos {
				continue
			}
			if ref1Rune[j] == ref2Rune[k] {
				newRunedSeq[j] = runedSeq[k]
			} else {
				panic(fmt.Sprintf("[Error!] Reference sequence 1 at %d and 2 at %d do not match: %s != %s\n", j, k, string(ref1Rune[j]), string(ref2Rune[k])))
			}
			j++
		}
		emptyAln[i] = fa.NewCharSequence(seq.ID(), seq.Description(), string(newRunedSeq))
		i++
	}
	var joinedAln fa.Alignment
	if method == "left" {
		joinedAln = append(aln1, emptyAln[1:]...)
	} else if method == "right" {
		joinedAln = append(emptyAln, aln1[1:]...)
	}
	return joinedAln
}

// LeftJoin joins 2 alignments using a common sequence similar to an SQL LEFT JOIN.
func LeftJoin(aln1, aln2 fa.Alignment, refindex1, refindex2 int) fa.Alignment {
	return Join(aln1, aln2, refindex1, refindex2, "left")
}

// RightJoin joins 2 alignments using a common sequence similar to an SQL RIGHT JOIN.
func RightJoin(aln1, aln2 fa.Alignment, refindex1, refindex2 int) fa.Alignment {
	return Join(aln1, aln2, refindex1, refindex2, "right")
}
