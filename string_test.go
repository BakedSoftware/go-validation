package validation

import (
	"fmt"
	"testing"
)

type LengthTestStruct struct {
	LongTitle  string `validation:"max_length=5"`
	ShortTitle string `validation:"min_length=3"`
	Title      string `validation:"min_length=3 max_length=5"`
}

type FormatTestStruct struct {
	Email string `validation:"format=email"`
}

func TestMaxLengthValid(t *testing.T) {
	object := LengthTestStruct{
		LongTitle:  "123",
		ShortTitle: "123",
		Title:      "1234",
	}

	ok, errs := IsValid(object)

	if !ok {
		t.Fatal(errs)
	}
}

func TestMaxLengthInvalid(t *testing.T) {
	object := LengthTestStruct{
		Title: "123456",
	}

	ok, errs := IsValid(object)

	if ok {
		t.Fatal("Max length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Max length errs should have 1 item not: %d", len(errs))
	}
}

func TestMinLengthValid(t *testing.T) {
	object := LengthTestStruct{
		Title:      "12345",
		ShortTitle: "123",
	}

	ok, errs := IsValid(object)

	if !ok {
		t.Fatal(errs)
	}
}

func TestMinLengthInvalid(t *testing.T) {
	object := LengthTestStruct{
		Title: "1234",
	}

	ok, errs := IsValid(object)

	if ok {
		t.Fatal("Min length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Min length errs should have 1 item not: %d", len(errs))
	}
}

func TestLengthValid(t *testing.T) {
	object := LengthTestStruct{
		Title:      "12345",
		ShortTitle: "123",
	}

	ok, errs := IsValid(object)

	if !ok {
		t.Fatal(errs)
	}
}

func TestLengthInvalid(t *testing.T) {
	// Check min_length=3
	object := LengthTestStruct{
		Title:      "12",
		ShortTitle: "123",
	}

	ok, errs := IsValid(object)

	if ok {
		t.Fatal("Length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Length errs should have 1 item not: %d", len(errs))
	}

	// Check max_length=5
	object = LengthTestStruct{
		Title:      "123456",
		ShortTitle: "123",
	}

	ok, errs = IsValid(object)

	if ok {
		t.Fatal("Length should have failed")
	}

	if len(errs) == 0 {
		t.Fatalf("Length errs should have 1 item not: %d", len(errs))
	}
}

func TestEmail(t *testing.T) {
	object := FormatTestStruct{
		Email: "",
	}

	ok, _ := IsValid(object)

	if ok {
		t.Fatal("Empty email should be invalid")
	}

	object.Email = "123"

	ok, _ = IsValid(object)

	if ok {
		t.Fatalf("Invalid email (%s) should be invalid", object.Email)
	}

	object.Email = "test@example.com"

	ok, errs := IsValid(object)

	if !ok {
		t.Log(errs)
		t.Fatalf("Valid email (%s) should be valid", object.Email)
	}

}

func ExampleIsValid_stringlength() {
	type Person struct {
		// Name must be between 1 and 5 characters inclusive
		Name string `validation:"min_length=1 max_length=5"`
	}

	var p Person

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
}

func ExampleIsValid_format() {
	type Person struct {
		// Email must be valid email
		Email string `validation:"format=email"`
	}

	var p Person

	ok, errs := IsValid(p)
	fmt.Println(ok, errs)
}
