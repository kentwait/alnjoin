package main

import (
	"testing"

	"github.com/kentwait/conspos/sequence"
)

func TestLeftJoin_EqualSize(t *testing.T) {
	aln1 := []sequence.Sequence{
		sequence.NewCharSequence("a", "", "ATG-CGCTCT"),
		sequence.NewCharSequence("b", "", "ATG-CGCACT"),
		sequence.NewCharSequence("c", "", "ATGGCGCACT"),
	}
	aln2 := []sequence.Sequence{
		sequence.NewCharSequence("x", "", "ATGCGC-TCT"),
		sequence.NewCharSequence("y", "", "ATGCCCATCT"),
		sequence.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	exp := []sequence.Sequence{
		sequence.NewCharSequence("a", "", "ATG-CGCTCT"),
		sequence.NewCharSequence("b", "", "ATG-CGCACT"),
		sequence.NewCharSequence("c", "", "ATGGCGCACT"),
		sequence.NewCharSequence("y", "", "ATG-CCCTCT"),
		sequence.NewCharSequence("z", "", "ATG-CCCTCT"),
	}
	joinedAln := LeftJoin(aln1, aln2, 0, 0)
	for i := range joinedAln {
		if exp[i].ID() != joinedAln[i].ID() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].ID(), joinedAln[i].ID(),
			)
		}

		if exp[i].Title() != joinedAln[i].Title() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].Title(), joinedAln[i].Title(),
			)
		}

		if exp[i].Sequence() != joinedAln[i].Sequence() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].Sequence(), joinedAln[i].Sequence(),
			)
		}
	}
}

func TestLeftJoin_Shorter(t *testing.T) {
	aln1 := []sequence.Sequence{
		sequence.NewCharSequence("a", "", "ATG-CGCTCT"),
		sequence.NewCharSequence("b", "", "ATG-CGCACT"),
		sequence.NewCharSequence("c", "", "ATGGCGCACT"),
	}
	aln2 := []sequence.Sequence{
		sequence.NewCharSequence("x", "", "ATGCG-C-TCT"),
		sequence.NewCharSequence("y", "", "ATGCC-CATCT"),
		sequence.NewCharSequence("z", "", "ATGCCCC-TCT"),
	}
	exp := []sequence.Sequence{
		sequence.NewCharSequence("a", "", "ATG-CGCTCT"),
		sequence.NewCharSequence("b", "", "ATG-CGCACT"),
		sequence.NewCharSequence("c", "", "ATGGCGCACT"),
		sequence.NewCharSequence("y", "", "ATG-CCCTCT"),
		sequence.NewCharSequence("z", "", "ATG-CCCTCT"),
	}
	joinedAln := LeftJoin(aln1, aln2, 0, 0)
	for i := range joinedAln {
		if exp[i].ID() != joinedAln[i].ID() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].ID(), joinedAln[i].ID(),
			)
		}

		if exp[i].Title() != joinedAln[i].Title() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].Title(), joinedAln[i].Title(),
			)
		}

		if exp[i].Sequence() != joinedAln[i].Sequence() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].Sequence(), joinedAln[i].Sequence(),
			)
		}
	}
}

func TestLeftJoin_Longer(t *testing.T) {
	aln1 := []sequence.Sequence{
		sequence.NewCharSequence("a", "", "ATG-CGCATCT"),
		sequence.NewCharSequence("b", "", "ATG-CGCAACT"),
		sequence.NewCharSequence("c", "", "ATGGCGC-ACT"),
	}
	aln2 := []sequence.Sequence{
		sequence.NewCharSequence("x", "", "ATGCGCATCT"),
		sequence.NewCharSequence("y", "", "ATGCCCATCT"),
		sequence.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	exp := []sequence.Sequence{
		sequence.NewCharSequence("a", "", "ATG-CGCATCT"),
		sequence.NewCharSequence("b", "", "ATG-CGCAACT"),
		sequence.NewCharSequence("c", "", "ATGGCGC-ACT"),
		sequence.NewCharSequence("y", "", "ATG-CCCATCT"),
		sequence.NewCharSequence("z", "", "ATG-CCC-TCT"),
	}
	joinedAln := LeftJoin(aln1, aln2, 0, 0)
	for i := range joinedAln {
		if exp[i].ID() != joinedAln[i].ID() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].ID(), joinedAln[i].ID(),
			)
		}

		if exp[i].Title() != joinedAln[i].Title() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].Title(), joinedAln[i].Title(),
			)
		}

		if exp[i].Sequence() != joinedAln[i].Sequence() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].Sequence(), joinedAln[i].Sequence(),
			)
		}
	}
}

func TestLeftJoin_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("LeftJoin: expected panic, but did not panic")
		}
	}()

	aln1 := []sequence.Sequence{
		sequence.NewCharSequence("a", "", "ATG-CGC-TCT"),
		sequence.NewCharSequence("b", "", "ATG-CGC-ACT"),
		sequence.NewCharSequence("c", "", "ATGGCGC-ACT"),
	}
	aln2 := []sequence.Sequence{
		sequence.NewCharSequence("x", "", "ATGCGCATCT"),
		sequence.NewCharSequence("y", "", "ATGCCCATCT"),
		sequence.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	LeftJoin(aln1, aln2, 0, 0)
}
