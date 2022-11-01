package ll

import "testing"

func TestGet(t *testing.T){
	l := linkedList[int]{}
	l.add(2)
	l.add(4)
	l.add(27)

	if l.get(4).value != 4 {
		t.Errorf("Unable to get value from linked list.\nExpected: %d, Actual: %v", 4, l.get(4))
	}

	if l.get(5) != nil {
		t.Errorf("Incorrect result from search.\nExpected: %v, Actual: %v", nil, l.get(5))
	}
}

func TestAdd(t *testing.T){
	l := linkedList[string]{}
	l.add("Raf")
	l.add("Connor")
	l.add("Shlok")

	if l.String() != "Raf Connor Shlok" {
		t.Errorf("Unable to add values to linked list\nExpected: %s, Actual: %v", "Raf Connor Shlok", l.String())
	}
}

func TestRemove(t *testing.T){
	l := linkedList[int]{}
	l.add(2)	
	l.add(4)	
	l.add(6)
	l.remove(4)

	if l.String() != "2 6"{
		t.Errorf("Unable to remove elements.\n Expected: %s, Actual: %s", "2 6", l.String()) 
	}	
}
