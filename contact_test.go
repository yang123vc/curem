package main

import (
	"testing"

	"github.com/joshsoftware/curem/config"
	"labix.org/v2/mgo/bson"
)

// This gives a benefit that we can use a test database when `go test` is run.
func init() {
	c := make(map[string]string)
	c["name"] = "test"
	c["url"] = "localhost"
	c["leads"] = "newlead"
	c["contacts"] = "newcontact"

	config.Configure(c)
}

func TestNewContact(t *testing.T) {
	collection := config.Db.C("newcontact")
	fakeContact, err := NewContact(
		"Encom Inc.",
		"Flynn",
		"flynn@encom.com",
		"",
		"",
		"USA",
	)
	if err != nil {
		t.Errorf("%s", err)
	}
	var fetchedContact contact
	err = collection.Find(bson.M{}).One(&fetchedContact)
	if err != nil {
		t.Errorf("%s", err)
	}

	// fakeContact is a pointer, because NewContact returns a pointer to a struct of contact type.
	// That's why we check fetchedContact with *fakeContact.

	if fetchedContact != *fakeContact {
		t.Errorf("inserted contact is not the fetched contact")
	}
	err = collection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetContact(t *testing.T) {
	collection := config.Db.C("newcontact")
	fakeContact, err := NewContact(
		"Encom Inc.",
		"Flynn",
		"flynn@encom.com",
		"",
		"",
		"USA",
	)
	if err != nil {
		t.Errorf("%s", err)
	}
	id := fakeContact.Id
	fetchedContact, err := GetContact(id)
	if err != nil {
		t.Errorf("%s", err)
	}
	if *fakeContact != *fetchedContact {
		t.Errorf("Expected %+v, but got %+v\n", *fakeContact, *fetchedContact)
	}
	err = collection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDelete(t *testing.T) {
	collection := config.Db.C("newcontact")
	fakeContact, err := NewContact(
		"Encom Inc.",
		"Flynn",
		"flynn@encom.com",
		"",
		"",
		"USA",
	)
	if err != nil {
		t.Errorf("%s", err)
	}
	err = fakeContact.Delete()
	if err != nil {
		t.Errorf("%s", err)
	}
	n, err := collection.Count()
	if err != nil {
		t.Errorf("%s", err)
	}
	if n != 0 {
		t.Errorf("expected 0 documents in the collection, but found %d", n)
	}
	err = collection.DropCollection()
	if err != nil {
		t.Errorf("%s", err)
	}
}
