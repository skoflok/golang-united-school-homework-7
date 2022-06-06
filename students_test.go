package coverage

import (
	"fmt"
	"os"
	"testing"
	"time"
)

const timeLayout = "2006-01-02"

func newBirthday(year int, month int, day int) time.Time {
	t, e := time.Parse(timeLayout, fmt.Sprintf("%d-%02d-%02d", year, month, day))
	if e != nil {
		panic(e)
	} else {
		return t
	}
}

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

func testPeople() People {
	return People{
		Person{"Alex", "Monroe", newBirthday(2001, 01, 11)},
		Person{"Kate", "Flow", newBirthday(2002, 02, 12)},
		Person{"Gregor", "Mindkeeper", newBirthday(2001, 01, 11)},
		Person{"Oustin", "Powers", newBirthday(2004, 04, 9)},
		Person{"Alex", "Powerball", newBirthday(2001, 01, 11)},
	}
}

// WRITE YOUR CODE BELOW

func Test_PeopleLength(t *testing.T) {
	expected := 3
	p := testPeople()[:expected]
	l := p.Len()

	if l != expected {
		t.Errorf("Unexpected len %d, want %d", l, expected)
	}
}

func Test_Less(t *testing.T) {
	p := testPeople()
	dataProvider := map[string]struct {
		i int
		j int
		e bool
	}{
		"Not Less by Birthday": {0, 1, false},
		"Less by FirstName":    {0, 2, true},
		"Less by Lastname":     {0, 4, true},
	}

	for k, d := range dataProvider {
		res := p.Less(d.i, d.j)
		if res != d.e {
			t.Errorf("Set %s i:(%d %s %s), j:(%d %s %s).Expected %t, actual %t",
				k,
				d.i, p[d.i].firstName, p[d.i].lastName,
				d.j, p[d.j].firstName, p[d.j].lastName,
				d.e, res)
		}
	}
}

func Test_Swap(t *testing.T) {
	p := testPeople()
	p0 := p[0]
	p3 := p[3]

	p.Swap(0, 3)

	if p[0] == p0 && p[3] == p3 {
		t.Errorf("Persons not been swaped")
	}
}

func Test_SimpleNew(t *testing.T) {
	simpleInput := "11 12\n21 22"
	simpleMatrix := Matrix{2, 2, []int{11, 12, 21, 22}}

	dataProvider := map[string]struct {
		input  string
		output Matrix
		err    error
	}{
		"simple": {
			input:  simpleInput,
			output: simpleMatrix,
			err:    nil,
		},
	}

	for k, data := range dataProvider {
		_, err := New(data.input)
		if err != nil {
			t.Errorf("%s. Matrix's error is not be nil", k)
		}
	}

}

func Test_AoitErrorNew(t *testing.T) {
	atoiErrorInput := "11A 12\n21 22"
	_, err := New(atoiErrorInput)
	if err == nil {
		t.Errorf("Matrix's error is nil. Error should be Atoi error")
	}
}

func Test_DiffrentRowsLengthNew(t *testing.T) {
	simpleInput := "11 12\n21 22\n 31"
	_, err := New(simpleInput)
	if err == nil {
		t.Errorf("Matrix's error is nil. Error should be row's length error")
	}
}
