package repository

import (
	"testing"
	"reflect"
	"go.mongodb.org/mongo-driver/bson"
)


func TestMapToFilter(t *testing.T) {
	attributes := map[string]string {
		"a": "1",
		"b": "2",
		"c": "3",
	}
	wanted := []bson.E{
		{"attributes.a", "1"},
		{"attributes.b", "2"},
		{"attributes.c", "3"},
	}
	filters := MapToFilter(attributes)
	if !reflect.DeepEqual(wanted, filters) {
		t.Fatalf("wanted %v but got %v\n", wanted, filters)
	}
}