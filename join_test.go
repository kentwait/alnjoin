package main

import (
	"testing"

	fa "github.com/kentwait/gofasta"
)

func TestJoin_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Join: expected panic, but did not panic")
		}
	}()

	aln1 := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGC-TCT"),
		fa.NewCharSequence("b", "", "ATG-CGC-ACT"),
		fa.NewCharSequence("c", "", "ATGGCGC-ACT"),
	}
	aln2 := []fa.Sequence{
		fa.NewCharSequence("x", "", "ATGCGCATCT"),
		fa.NewCharSequence("y", "", "ATGCCCATCT"),
		fa.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	Join(aln1, aln2, 0, 0, "l")
}

func TestLeftJoin_EqualSize(t *testing.T) {
	aln1 := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGCTCT"),
		fa.NewCharSequence("b", "", "ATG-CGCACT"),
		fa.NewCharSequence("c", "", "ATGGCGCACT"),
	}
	aln2 := []fa.Sequence{
		fa.NewCharSequence("x", "", "ATGCGC-TCT"),
		fa.NewCharSequence("y", "", "ATGCCCATCT"),
		fa.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	exp := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGCTCT"),
		fa.NewCharSequence("b", "", "ATG-CGCACT"),
		fa.NewCharSequence("c", "", "ATGGCGCACT"),
		fa.NewCharSequence("y", "", "ATG-CCCTCT"),
		fa.NewCharSequence("z", "", "ATG-CCCTCT"),
	}
	joinedAln := LeftJoin(aln1, aln2, 0, 0)
	for i := range joinedAln {
		if exp[i].ID() != joinedAln[i].ID() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].ID(), joinedAln[i].ID(),
			)
		}

		if exp[i].Description() != joinedAln[i].Description() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].Description(), joinedAln[i].Description(),
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
	aln1 := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGCTCT"),
		fa.NewCharSequence("b", "", "ATG-CGCACT"),
		fa.NewCharSequence("c", "", "ATGGCGCACT"),
	}
	aln2 := []fa.Sequence{
		fa.NewCharSequence("x", "", "ATGCG-C-TCT"),
		fa.NewCharSequence("y", "", "ATGCC-CATCT"),
		fa.NewCharSequence("z", "", "ATGCCCC-TCT"),
	}
	exp := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGCTCT"),
		fa.NewCharSequence("b", "", "ATG-CGCACT"),
		fa.NewCharSequence("c", "", "ATGGCGCACT"),
		fa.NewCharSequence("y", "", "ATG-CCCTCT"),
		fa.NewCharSequence("z", "", "ATG-CCCTCT"),
	}
	joinedAln := LeftJoin(aln1, aln2, 0, 0)
	for i := range joinedAln {
		if exp[i].ID() != joinedAln[i].ID() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].ID(), joinedAln[i].ID(),
			)
		}

		if exp[i].Description() != joinedAln[i].Description() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].Description(), joinedAln[i].Description(),
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
	aln1 := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGCATCT"),
		fa.NewCharSequence("b", "", "ATG-CGCAACT"),
		fa.NewCharSequence("c", "", "ATGGCGC-ACT"),
	}
	aln2 := []fa.Sequence{
		fa.NewCharSequence("x", "", "ATGCGCATCT"),
		fa.NewCharSequence("y", "", "ATGCCCATCT"),
		fa.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	exp := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGCATCT"),
		fa.NewCharSequence("b", "", "ATG-CGCAACT"),
		fa.NewCharSequence("c", "", "ATGGCGC-ACT"),
		fa.NewCharSequence("y", "", "ATG-CCCATCT"),
		fa.NewCharSequence("z", "", "ATG-CCC-TCT"),
	}
	joinedAln := LeftJoin(aln1, aln2, 0, 0)
	for i := range joinedAln {
		if exp[i].ID() != joinedAln[i].ID() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].ID(), joinedAln[i].ID(),
			)
		}

		if exp[i].Description() != joinedAln[i].Description() {
			t.Errorf("LeftJoin: expected %v, actual %v",
				exp[i].Description(), joinedAln[i].Description(),
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

	aln1 := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGC-TCT"),
		fa.NewCharSequence("b", "", "ATG-CGC-ACT"),
		fa.NewCharSequence("c", "", "ATGGCGC-ACT"),
	}
	aln2 := []fa.Sequence{
		fa.NewCharSequence("x", "", "ATGCGCATCT"),
		fa.NewCharSequence("y", "", "ATGCCCATCT"),
		fa.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	LeftJoin(aln1, aln2, 0, 0)
}

func TestRightJoin_EqualSize(t *testing.T) {
	aln1 := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGCTCT"),
		fa.NewCharSequence("b", "", "ATG-CGCACT"),
		fa.NewCharSequence("c", "", "ATGGCGCACT"),
	}
	aln2 := []fa.Sequence{
		fa.NewCharSequence("x", "", "ATGCGC-TCT"),
		fa.NewCharSequence("y", "", "ATGCCCATCT"),
		fa.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	exp := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATGCGC-TCT"),
		fa.NewCharSequence("b", "", "ATGCGC-ACT"),
		fa.NewCharSequence("c", "", "ATGCGC-ACT"),
		fa.NewCharSequence("y", "", "ATGCCCATCT"),
		fa.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	joinedAln := RightJoin(aln1, aln2, 0, 0)
	for i := range joinedAln {
		if exp[i].ID() != joinedAln[i].ID() {
			t.Errorf("RightJoin: expected %v, actual %v",
				exp[i].ID(), joinedAln[i].ID(),
			)
		}

		if exp[i].Description() != joinedAln[i].Description() {
			t.Errorf("RightJoin: expected %v, actual %v",
				exp[i].Description(), joinedAln[i].Description(),
			)
		}

		if exp[i].Sequence() != joinedAln[i].Sequence() {
			t.Errorf("RightJoin: expected %v, actual %v",
				exp[i].Sequence(), joinedAln[i].Sequence(),
			)
		}
	}
}

func TestRightJoin_Shorter(t *testing.T) {
	aln1 := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGCTCT"),
		fa.NewCharSequence("b", "", "ATG-CGCACT"),
		fa.NewCharSequence("c", "", "ATGGCGCACT"),
	}
	aln2 := []fa.Sequence{
		fa.NewCharSequence("x", "", "ATGCG-C-TCT"),
		fa.NewCharSequence("y", "", "ATGCC-CATCT"),
		fa.NewCharSequence("z", "", "ATGCCCC-TCT"),
	}
	exp := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATGCG-C-TCT"),
		fa.NewCharSequence("b", "", "ATGCG-C-ACT"),
		fa.NewCharSequence("c", "", "ATGCG-C-ACT"),
		fa.NewCharSequence("y", "", "ATGCC-CATCT"),
		fa.NewCharSequence("z", "", "ATGCCCC-TCT"),
	}
	joinedAln := RightJoin(aln1, aln2, 0, 0)
	for i := range joinedAln {
		if exp[i].ID() != joinedAln[i].ID() {
			t.Errorf("RightJoin: expected %v, actual %v",
				exp[i].ID(), joinedAln[i].ID(),
			)
		}

		if exp[i].Description() != joinedAln[i].Description() {
			t.Errorf("RightJoin: expected %v, actual %v",
				exp[i].Description(), joinedAln[i].Description(),
			)
		}

		if exp[i].Sequence() != joinedAln[i].Sequence() {
			t.Errorf("RightJoin: expected %v, actual %v",
				exp[i].Sequence(), joinedAln[i].Sequence(),
			)
		}
	}
}

func TestRightJoin_Longer(t *testing.T) {
	aln1 := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGCATCT"),
		fa.NewCharSequence("b", "", "ATG-CGCAACT"),
		fa.NewCharSequence("c", "", "ATGGCGC-ACT"),
	}
	aln2 := []fa.Sequence{
		fa.NewCharSequence("x", "", "ATGCGCATCT"),
		fa.NewCharSequence("y", "", "ATGCCCATCT"),
		fa.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	exp := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATGCGCATCT"),
		fa.NewCharSequence("b", "", "ATGCGCAACT"),
		fa.NewCharSequence("c", "", "ATGCGC-ACT"),
		fa.NewCharSequence("y", "", "ATGCCCATCT"),
		fa.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	joinedAln := RightJoin(aln1, aln2, 0, 0)
	for i := range joinedAln {
		if exp[i].ID() != joinedAln[i].ID() {
			t.Errorf("RightJoin: expected %v, actual %v",
				exp[i].ID(), joinedAln[i].ID(),
			)
		}

		if exp[i].Description() != joinedAln[i].Description() {
			t.Errorf("RightJoin: expected %v, actual %v",
				exp[i].Description(), joinedAln[i].Description(),
			)
		}

		if exp[i].Sequence() != joinedAln[i].Sequence() {
			t.Errorf("RightJoin: expected %v, actual %v",
				exp[i].Sequence(), joinedAln[i].Sequence(),
			)
		}
	}
}

func TestRightJoin_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("RightJoin: expected panic, but did not panic")
		}
	}()

	aln1 := []fa.Sequence{
		fa.NewCharSequence("a", "", "ATG-CGC-TCT"),
		fa.NewCharSequence("b", "", "ATG-CGC-ACT"),
		fa.NewCharSequence("c", "", "ATGGCGC-ACT"),
	}
	aln2 := []fa.Sequence{
		fa.NewCharSequence("x", "", "ATGCGCATCT"),
		fa.NewCharSequence("y", "", "ATGCCCATCT"),
		fa.NewCharSequence("z", "", "ATGCCC-TCT"),
	}
	RightJoin(aln1, aln2, 0, 0)
}
