package pathname

import (
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPathname(t *testing.T) {
	assert.NoError(t, nil)
}

func TestPathname_ToPath(t *testing.T) {
	var p *Pathname
	assert.Equal(t, "", p.ToPath())

	testcases := []struct {
		p        Pathname
		expected string
	}{
		{NewPathname(""), ""},
		{NewPathname("/path"), "/path"},
		{NewPathname("path"), "path"},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.ToPath())
		})
	}
}

func TestPathname_IsAbsolute(t *testing.T) {
	var p *Pathname
	assert.False(t, p.IsAbsolute())

	testcases := []struct {
		p        Pathname
		expected bool
	}{
		{NewPathname(""), false},
		{NewPathname("/"), true},
		{NewPathname("/path"), true},
		{NewPathname("."), false},
		{NewPathname("path"), false},
		{NewPathname("./path"), false},
		{NewPathname("../path"), false},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.IsAbsolute())
		})
	}
}

func TestPathname_IsRelative(t *testing.T) {
	var p *Pathname
	assert.False(t, p.IsRelative())

	testcases := []struct {
		p        Pathname
		expected bool
	}{
		{NewPathname(""), false},
		{NewPathname("/"), false},
		{NewPathname("/path"), false},
		{NewPathname("."), true},
		{NewPathname("path"), true},
		{NewPathname("./path"), true},
		{NewPathname("../path"), true},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.IsRelative())
		})
	}
}

func TestPathname_IsRoot(t *testing.T) {
	var p *Pathname
	assert.False(t, p.IsRoot())

	testcases := []struct {
		p        Pathname
		expected bool
	}{
		{NewPathname(""), false},
		{NewPathname("/"), true},
		{NewPathname("/path"), false},
		{NewPathname("/."), false},
		{NewPathname("/path/.."), false},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.IsRoot())
		})
	}
}

func TestPathname_IsExist(t *testing.T) {
	var p *Pathname
	assert.False(t, p.IsExist())

	testcases := []struct {
		p        Pathname
		expected bool
	}{
		{NewPathname(""), false},
		{NewPathname("/"), true},
		{NewPathname("."), true},
		{NewPathname("testdata/file1.txt"), true},
		{NewPathname("./testdata/file1.txt"), true},
		{NewPathname("testdata/dir1"), true},
		{NewPathname("./testdata/dir1"), true},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.IsExist())
		})
	}
}

func TestPathname_IsFile(t *testing.T) {
	var p *Pathname
	assert.False(t, p.IsFile())

	testcases := []struct {
		p        Pathname
		expected bool
	}{
		{NewPathname(""), false},
		{NewPathname("/"), false},
		{NewPathname("."), false},
		{NewPathname("testdata/file1.txt"), true},
		{NewPathname("./testdata/file1.txt"), true},
		{NewPathname("testdata/dir1"), false},
		{NewPathname("./testdata/dir1"), false},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.IsFile())
		})
	}
}

func TestPathname_IsDirectory(t *testing.T) {
	var p *Pathname
	assert.False(t, p.IsDirectory())

	testcases := []struct {
		p        Pathname
		expected bool
	}{
		{NewPathname(""), false},
		{NewPathname("/"), true},
		{NewPathname("."), true},
		{NewPathname("testdata/file1.txt"), false},
		{NewPathname("./testdata/file1.txt"), false},
		{NewPathname("testdata/dir1"), true},
		{NewPathname("./testdata/dir1"), true},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.IsDirectory())
		})
	}
}

func TestPathname_Basename(t *testing.T) {
	var p *Pathname
	assert.Equal(t, p, p.Dirname())

	testcases := []struct {
		p        Pathname
		expected *Pathname
	}{
		{NewPathname(""), NewPathnamePtr(".")},
		{NewPathname("/"), NewPathnamePtr("/")},
		{NewPathname("."), NewPathnamePtr(".")},
		{NewPathname("testdata/file1.txt"), NewPathnamePtr("file1.txt")},
		{NewPathname("./testdata/file1.txt"), NewPathnamePtr("file1.txt")},
		{NewPathname("testdata/dir1"), NewPathnamePtr("dir1")},
		{NewPathname("testdata/dir1/"), NewPathnamePtr("dir1")},
		{NewPathname("./testdata/dir1"), NewPathnamePtr("dir1")},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.Basename())
		})
	}
}

func TestPathname_Dirname(t *testing.T) {
	var p *Pathname
	assert.Equal(t, p, p.Dirname())

	testcases := []struct {
		p        Pathname
		expected *Pathname
	}{
		{NewPathname(""), NewPathnamePtr(".")},
		{NewPathname("/"), NewPathnamePtr("/")},
		{NewPathname("."), NewPathnamePtr(".")},
		{NewPathname("testdata/file1.txt"), NewPathnamePtr("testdata")},
		{NewPathname("./testdata/file1.txt"), NewPathnamePtr("testdata")},
		{NewPathname("testdata/dir1"), NewPathnamePtr("testdata")},
		{NewPathname("testdata/dir1/"), NewPathnamePtr("testdata/dir1")},
		{NewPathname("./testdata/dir1"), NewPathnamePtr("testdata")},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.Dirname())
		})
	}
}

func TestPathname_Join(t *testing.T) {
	var p *Pathname
	assert.Equal(t, p, p.Join())

	testcases := []struct {
		p        Pathname
		parts    []string
		expected *Pathname
	}{
		{NewPathname(""), []string{}, NewPathnamePtr("")},
		{NewPathname(""), []string{"path", "file.txt"}, NewPathnamePtr("path/file.txt")},
		{NewPathname("/"), []string{"."}, NewPathnamePtr("/")},
		{NewPathname("/"), []string{".."}, NewPathnamePtr("/")},
		{NewPathname("/"), []string{"path", "file.txt"}, NewPathnamePtr("/path/file.txt")},
		{NewPathname("/"), []string{"path", ".", "..", "file.txt"}, NewPathnamePtr("/file.txt")},
		{NewPathname("/"), []string{"path", "..", "..", "file.txt"}, NewPathnamePtr("/file.txt")},
		{NewPathname("."), []string{}, NewPathnamePtr(".")},
		{NewPathname("testdata/file1.txt"), []string{}, NewPathnamePtr("testdata/file1.txt")},
		{NewPathname("./testdata/file1.txt"), []string{}, NewPathnamePtr("testdata/file1.txt")},
		{NewPathname("testdata/dir1"), []string{}, NewPathnamePtr("testdata/dir1")},
		{NewPathname("testdata/dir1/"), []string{}, NewPathnamePtr("testdata/dir1")},
		{NewPathname("./testdata/dir1"), []string{}, NewPathnamePtr("testdata/dir1")},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.Join(tc.parts...))
		})
	}
}

func TestPathname_OpenFile(t *testing.T) {
	var p *Pathname
	file, err := p.OpenFile(os.O_RDONLY, 0)
	assert.Nil(t, file)
	var expectedError *os.PathError
	assert.ErrorAs(t, err, &expectedError)

	testcases := []struct {
		p           Pathname
		expectedErr error
	}{
		{NewPathname("./testdata/file1.txt"), nil},
		{NewPathname("./testdata/file2.txt"), &fs.PathError{}},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			file, err := tc.p.OpenFile(os.O_RDONLY, 0)
			if tc.expectedErr == nil {
				assert.NotNil(t, file)
				assert.NoError(t, err)
			} else {
				assert.Nil(t, file)
				assert.ErrorAs(t, err, &tc.expectedErr)
			}
		})
	}
}

func TestPathname_Size(t *testing.T) {
	size, err := NewPathnamePtr("./testdata/file1.txt").Size()
	assert.NoError(t, err)
	assert.Equal(t, 507, size)
}
