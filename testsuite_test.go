package vecty

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
)

var _ = func() bool {
	isTest = true
	return true
}()

// recoverStr runs f and returns the recovered panic as a string.
func recoverStr(f func()) (s string) {
	defer func() {
		s = fmt.Sprint(recover())
	}()
	f()
	return
}

type componentFunc struct {
	Core
	id         string
	render     func() ComponentOrHTML
	skipRender func(prev Component) bool
}

func (c *componentFunc) Render() ComponentOrHTML        { return c.render() }
func (c *componentFunc) SkipRender(prev Component) bool { return c.skipRender(prev) }

func TestMain(m *testing.M) {
	// Try to remove all testdata/*.got.txt files now.
	matches, _ := filepath.Glob("testdata/*.got.txt")
	for _, match := range matches {
		os.Remove(match)
	}

	os.Exit(m.Run())
}

func testSuite(t *testing.T, testName string) *testSuiteT {
	ts := &testSuiteT{
		t:         t,
		testName:  testName,
		callbacks: make(map[string]interface{}),
		strings:   &valueMocker{},
		bools:     &valueMocker{},
		floats:    &valueMocker{},
		ints:      &valueMocker{},
	}
	global = &objectRecorder{
		ts:   ts,
		name: "global",
	}
	return ts
}

// mockedValue represents a mocked value.
type mockedValue struct {
	invocation string
	value      interface{}
}

// valueMocker keeps tracked of mocked values for method invocations on
// jsObject's.
type valueMocker struct {
	values []mockedValue
}

// mock adds an entry to mock the specified invocation to return the given
// value.
func (v *valueMocker) mock(invocation string, value interface{}) {
	v.values = append(v.values, mockedValue{invocation, value})
}

// get gets the mocked value for the specified invocation.
func (v *valueMocker) get(invocation string) interface{} {
	for i, value := range v.values {
		if value.invocation == invocation {
			// Found the right invocation.
			v.values = append(v.values[:i], v.values[i+1:]...)
			return value.value
		}
	}
	panic(fmt.Sprintf("expected mocked value for invocation: %s", invocation))
}

type testSuiteT struct {
	t                            *testing.T
	testName                     string
	callbacks                    map[string]interface{}
	strings, bools, floats, ints *valueMocker

	got    string
	isDone bool
}

func (ts *testSuiteT) done() {
	ts.multiSortedDone()
}

// sortedDone is just like done(), except it sorts the specified line range first.
func (ts *testSuiteT) sortedDone(sortStartLine, sortEndLine int) {
	ts.multiSortedDone([2]int{sortStartLine, sortEndLine})
}

// multiSortedDone is just like done(), except it sorts the specified line range first.
func (ts *testSuiteT) multiSortedDone(linesToSort ...[2]int) {
	if ts.isDone {
		panic("testSuite done methods called multiple times")
	}
	ts.isDone = true
	// Read the want file or create it if it does not exist.
	wantFileName := path.Join("testdata", ts.testName+".want.txt")
	wantBytes, err := ioutil.ReadFile(wantFileName)
	if err != nil {
		if os.IsNotExist(err) {
			// Touch the file
			f, err := os.Create(wantFileName)
			f.Close()
			if err != nil {
				ts.t.Fatal(err)
			}
		} else {
			ts.t.Fatal(err)
		}
	}
	want := strings.TrimSpace(string(wantBytes))

	// Ensure output is properly sorted.
	split := strings.Split(strings.TrimSpace(ts.got), "\n")
	for _, pair := range linesToSort {
		sortStartLine := pair[0] - 1 // to match editor line numbers
		if sortStartLine < 0 {
			sortStartLine = 0
		}
		sortEndLine := pair[1]
		if sortEndLine > len(split) {
			sortEndLine = len(split)
		}
		sorted := split[sortStartLine:sortEndLine]
		ts.t.Logf("lines selected for sorting (%d-%d):\n%s\n\n", sortStartLine, sortEndLine, strings.Join(sorted, "\n"))
		sort.Strings(sorted)
		for i := sortStartLine; i < sortEndLine; i++ {
			split[i] = sorted[i-sortStartLine]
		}
	}
	got := strings.Join(split, "\n")

	// Check if we got what we wanted.
	if got == want {
		// Successful test.

		// Ensure there are no unused mocked values.
		for _, v := range ts.strings.values {
			ts.t.Errorf("unused mocked string value %q %v", v.invocation, v.value)
		}
		for _, v := range ts.bools.values {
			ts.t.Errorf("unused mocked bool value %q %v", v.invocation, v.value)
		}
		for _, v := range ts.floats.values {
			ts.t.Errorf("unused mocked float value %q %v", v.invocation, v.value)
		}
		for _, v := range ts.ints.values {
			ts.t.Errorf("unused mocked int value %q %v", v.invocation, v.value)
		}
		return
	}

	// Write what we got to disk.
	gotFileName := path.Join("testdata", ts.testName+".got.txt")
	err = ioutil.WriteFile(gotFileName, []byte(got), 0777)
	if err != nil {
		ts.t.Fatal(err)
	}

	// Print a nice diff for easy comparison.
	cmd := exec.Command("git", "-c", "color.ui=always", "diff", "--no-index", wantFileName, gotFileName)
	out, _ := cmd.CombinedOutput()
	ts.t.Log("\n" + string(out))

	ts.t.Fatalf("to accept these changes:\n\n$ mv %s %s", gotFileName, wantFileName)
}

// record records the invocation to the test suite and returns the string
// unmodified.
func (ts *testSuiteT) record(invocation string) string {
	ts.got += "\n" + invocation
	return invocation
}

// addCallbacks adds the first function in args to ts.callbacks[invocation], if there is one.
func (ts *testSuiteT) addCallbacks(invocation string, args ...interface{}) {
	for _, a := range args {
		if reflect.TypeOf(a).Kind() == reflect.Func {
			ts.callbacks[invocation] = a
			return
		}
	}
}

// objectRecorder implements the jsObject interface by recording method
// invocations to the test suite.
type objectRecorder struct {
	ts   *testSuiteT
	name string
}

// Set implements the jsObject interface.
func (r *objectRecorder) Set(key string, value interface{}) {
	invocation := r.ts.record(fmt.Sprintf("%s.Set(%q, %+v)", r.name, key, stringify(value)))
	r.ts.addCallbacks(invocation, value)
}

// Get implements the jsObject interface.
func (r *objectRecorder) Get(key string) jsObject {
	invocation := r.ts.record(fmt.Sprintf("%s.Get(%q)", r.name, key))
	return &objectRecorder{
		ts:   r.ts,
		name: invocation,
	}
}

// Delete implements the jsObject interface.
func (r *objectRecorder) Delete(key string) {
	r.ts.record(fmt.Sprintf("%s.Delete(%q)", r.name, key))
}

// Call implements the jsObject interface.
func (r *objectRecorder) Call(name string, args ...interface{}) jsObject {
	invocation := r.ts.record(fmt.Sprintf("%s.Call(%q, %s)", r.name, name, stringify(args...)))
	r.ts.addCallbacks(invocation, args...)
	return &objectRecorder{
		ts:   r.ts,
		name: invocation,
	}
}

// String implements the jsObject interface.
func (r *objectRecorder) String() string { return r.ts.strings.get(r.name).(string) }

// Bool implements the jsObject interface.
func (r *objectRecorder) Bool() bool { return r.ts.bools.get(r.name).(bool) }

// Int implements the jsObject interface.
func (r *objectRecorder) Int() int { return r.ts.ints.get(r.name).(int) }

// Float implements the jsObject interface.
func (r *objectRecorder) Float() float64 { return r.ts.floats.get(r.name).(float64) }

func stringify(args ...interface{}) string {
	var s []string
	for _, a := range args {
		if reflect.TypeOf(a).Kind() == reflect.Func {
			s = append(s, reflect.TypeOf(a).String())
			continue
		}
		switch v := a.(type) {
		case string:
			s = append(s, fmt.Sprintf("%q", v))
		case *objectRecorder:
			s = append(s, fmt.Sprintf("jsObject(%s)", v.name))
		default:
			s = append(s, fmt.Sprintf("%v", v))
		}
	}
	return strings.Join(s, ", ")
}
