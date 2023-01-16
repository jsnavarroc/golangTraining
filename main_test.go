package main

import "testing"

/**
* Creaos una estructura que nos permita evaluar los diferentes esenarios
*
**/
var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}


func TestDivision(t *testing.T){
	for _, tt:= range tests {
		result, err :=divide(tt.dividend, tt.divisor)

		if tt.isErr{
			if err == nil {
				t.Error("Expected an error but did not get one")
			}

		} else {
			if err != nil {
				t.Error("Did not expect an error but gone one", err.Error())
			}
		}

		if result != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, result)
		}
	}
}


// func TestDivide(t *testing.T){
// 	_, err := divide(10.0,1.0)
// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
// }
// func TestBadDivide(t *testing.T){
// 	_, err := divide(10.0, 0.0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have")
// 	}
// }