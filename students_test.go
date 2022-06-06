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
