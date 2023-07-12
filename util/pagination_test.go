package util

import (
	"testing"
)

func TestPagination(t *testing.T) {
	noDeclaretePage := ResolvePagination("")
	invalidPageValue := ResolvePagination("-1")
	invalidPageType := ResolvePagination("1ab2")
	pageValid := ResolvePagination("2")

	if noDeclaretePage != 0 {
		t.Errorf("undeclared pages must be the first, but it receives '%b'", noDeclaretePage)
	}

	if invalidPageValue != 0 {
		t.Errorf("expected '0' but receives '%b'", invalidPageValue)
	}

	if invalidPageType != 0 {
		t.Errorf("expected '0' but receives '%b'", invalidPageType)
	}

	if pageValid != 1 {
		t.Errorf("page should return page 2, but return '%b'", pageValid)
	}
}
