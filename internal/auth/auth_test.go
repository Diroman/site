package auth

import "testing"

func TestCreateNewToken(t *testing.T) {
	got, err := CreateNewToken("111")
	if err != nil {
		t.Fail()
	}

	println(got)
}
