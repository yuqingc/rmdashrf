package manager

import (
	"testing"
)

type TestListDirCase struct {
	dir                string
	all                bool
	max                int
	ext                string
	wantedResultsCount int
	wantedTotal        int
	wantedHavingErr    bool
}

var (
	testCase0 = TestListDirCase{
		dir:                "../../test/mydir/",
		all:                true,
		max:                5000,
		ext:                "",
		wantedResultsCount: 5,
		wantedTotal:        5,
		wantedHavingErr:    false,
	}
	testCase1 = TestListDirCase{
		dir:                "../../test/mydir/",
		all:                false,
		max:                5000,
		ext:                "",
		wantedResultsCount: 3,
		wantedTotal:        5,
		wantedHavingErr:    false,
	}
	testCase2 = TestListDirCase{
		dir:                "../../test/mydir/src",
		all:                true,
		max:                2,
		ext:                "",
		wantedResultsCount: 2,
		wantedTotal:        4,
		wantedHavingErr:    false,
	}
	testCase3 = TestListDirCase{
		dir:                "../../test/mydir/notexist",
		all:                true,
		max:                5000,
		ext:                "",
		wantedResultsCount: 0,
		wantedTotal:        0,
		wantedHavingErr:    true,
	}
	testCase4 = TestListDirCase{
		dir:                "../../test/mydir/src/",
		all:                false,
		max:                5000,
		ext:                ".js",
		wantedResultsCount: 1,
		wantedTotal:        4,
		wantedHavingErr:    false,
	}
)

func TestListDir(t *testing.T) {
	cases := []TestListDirCase{
		testCase0,
		testCase1,
		testCase2,
		testCase3,
		testCase4,
	}
	for _, v := range cases {
		gotResults, gotTotal, gotErr := ListDir(v.dir, v.all, v.max, v.ext)
		if len(gotResults) != v.wantedResultsCount {
			t.Errorf("ListDir(%q,%v,%v,%q) results count is %d, want %d", v.dir, v.all, v.max, v.ext, len(gotResults), v.wantedResultsCount)
		}
		if gotTotal != v.wantedTotal {
			t.Errorf("ListDir(%q,%v,%v,%q) total is %d, want %d", v.dir, v.all, v.max, v.ext, gotTotal, v.wantedTotal)
		}
		if (gotErr != nil) != v.wantedHavingErr {
			t.Errorf("ListDir(%q,%v,%v,%q) havingError is %v, want %v", v.dir, v.all, v.max, v.ext, gotErr != nil, v.wantedHavingErr)
		}
	}
}
