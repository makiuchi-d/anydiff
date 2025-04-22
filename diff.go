package anydiff

import "slices"

const (
	Addition = '+'
	Deletion = '-'
	Keep     = '='
)

// Edit represents the sequence of edit operations
type Edit []byte

// Distance calculate edit distance
func (e Edit) Distance() int {
	d := 0
	for _, c := range e {
		if c != Keep {
			d++
		}
	}
	return d
}

// String implements fmt.Stringer interface
func (e Edit) String() string {
	return string(e)
}

// Diff computes a Myers diff of slices
func Diff[A, B any](a []A, b []B, cmp func(a *A, b *B) bool) Edit {
	m, n := len(a), len(b)
	v := make(map[int]int)
	eds := make(map[int]Edit)

	for d := range m + n + 1 {
		for k := -d; k <= d; k += 2 {
			var i int
			var ed Edit

			if d == 0 {
				i = 0
			} else if k == -d {
				i = v[k+1] + 1
				ed = append(slices.Clone(eds[k+1]), Deletion)
			} else if k == d {
				i = v[k-1]
				ed = append(slices.Clone(eds[k-1]), Addition)
			} else if v[k-1] < v[k+1]+1 {
				i = v[k+1] + 1
				ed = append(slices.Clone(eds[k+1]), Deletion)
			} else {
				i = v[k-1]
				ed = append(slices.Clone(eds[k-1]), Addition)
			}

			for i < m && i+k < n && cmp(&a[i], &b[i+k]) {
				i++
				ed = append(ed, Keep)
			}

			if k == n-m && i == m {
				return ed
			}

			v[k] = i
			eds[k] = ed
		}
	}

	panic("unreachable")
}

// Cmp is a simple equality check suitable for Diff's cmp argument
func Cmp[T comparable](a, b *T) bool { return *a == *b }
