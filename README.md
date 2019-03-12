# alnjoin

Merge 2 multiple sequence alignments that share a common sequence.

## Usage

`alnjoin -aln1 <path> -aln2 <path> -ref1 <refname_in_aln1> -ref2 <refname_in_aln2> > <output>`

```
Usage of ./alnjoin:
  -aln1 string
        Path to first alignment (required)
  -aln2 string
        Path to second alignment (required)
  -ref1 string
        Sequence name/identifier of sequence in aln1 to join on  (required)
  -ref2 string
        Sequence name/identifier of sequence in aln2 to join on (required)
```

## Background

The program builds an ungapped positional array for each of the two reference sequences to join on. Then, the arrays are read column-by-column, advancing 1 column at a time, until the positional values match.

Positional values may not readily match due to the presence of gaps in the reference sequences. When a gap is encountered in the first reference, its positional value will lag in terms of the second:

```
      0   1   2   3   4       5   6   7
ref1  A   T   G   G   C   -   G   C   C  ...
ref2  A   T   G   G   C   G   C   C   A  ...
      0   1   2   3   4   5   6   7   8
```

To solve this, the program adds gap columns until the positional values match again.

```
      0   1   2   3   4       5   6   7
ref1  A   T   G   G   C   -   G   C   C  ...
ref2  A   T   G   G   C   -   G   C   C   A  ...
      0   1   2   3   4       5   6   7   8
```

Conversely, the seoncd reference sequence may contain gaps and will lag relative to the first reference sequence as shown in the following example:

```
      0   1   2   3   4   5   6   7   8
ref1  A   T   G   G   C   G   C   C   A  ...
ref2  A   T   G   G   C   -   G   C   C  ...
      0   1   2   3   4       5   6   7
```

In this case, the program skips columns in the second alignment until the positional values match again. As a consequence, insertions that introduced the gap in the alignment will be dropped in the joined alignment (as in the case of the "other" sequence).

```
      0   1   2   3   4   5   6   7   8
ref1  A   T   G   G   C   G   C   C   A  ...
ref2  A   T   G   G   C   -   G   C   C  ...
other A   T   G   G   C   A   G   C   C
      0   1   2   3   4   *   5   6   7
```

Thus, the joining process is similar to an `SQL LEFT JOIN` where instead of tables, there are alignments and instead of joining on a particular column, alignments are joined through a common sequence.
