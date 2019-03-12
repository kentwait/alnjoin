package main

import (
	"fmt"

	aln "github.com/kentwait/conspos/alignment"
	"github.com/kentwait/conspos/sequence"
)

// LeftJoin joins 2 alignments using a common sequence similar to an SQL LEFT JOIN.
func LeftJoin(aln1, aln2 aln.Alignment, refindex1, refindex2 int) aln.Alignment {
	ref1UngappedSlice := aln1[refindex1].UngappedPositionSlice("-")
	ref2UngappedSlice := aln2[refindex2].UngappedPositionSlice("-")

	ref1Rune := []rune(aln1[refindex1].Sequence())
	ref2Rune := []rune(aln2[refindex2].Sequence())

	// Initialize empty slice
	emptyAln := make([]sequence.Sequence, len(aln2)-1)
	i := 0
	for idx2, seq := range aln2 {
		if idx2 == refindex2 {
			continue
		}
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
		emptyAln[i] = sequence.NewCharSequence(seq.ID(), seq.Title(), string(newRunedSeq))
		i++
	}
	joinedAln := append(aln1, emptyAln...)
	return joinedAln
}
