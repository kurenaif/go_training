package main

import (
	"net/http/httptest"
	"testing"
)

func isSame(lhs map[string]dollars, rhs map[string]dollars) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for k := range lhs {
		if lhs[k] != rhs[k] {
			return false
		}
	}
	return true
}

func TestCreate(t *testing.T) {
	tests := []struct {
		db       database
		request  string
		wantBody string
		wantDB   map[string]dollars
	}{
		{database{}, "/?item=hoge&price=200", "hoge: $200.00 setted\n", map[string]dollars{"hoge": 200}},
		{database{}, "/?item=hoge", "field 'price' error: strconv.ParseFloat: parsing \"\": invalid syntax", map[string]dollars{}},
		{database{}, "/?price=200", "field 'item' must be setted.", map[string]dollars{}},
		{database{"hoge": 100}, "/?item=hoge&price=200", "hoge is already exist.", map[string]dollars{"hoge": 100}},
		{database{"fuga": 0}, "/?item=hoge&price=200", "hoge: $200.00 setted\n", map[string]dollars{"hoge": 200, "fuga": 0}},
	}

	for _, test := range tests {
		req := httptest.NewRequest("PUT", test.request, nil)
		rec := httptest.NewRecorder()
		db := test.db

		db.create(rec, req)

		if !isSame(db, test.wantDB) {
			t.Errorf("---get body---\n %v \n ---want body---\n %v \n", db, test.wantDB)
		}

		if rec.Body.String() != test.wantBody {
			t.Errorf("get request Body is %q, want %q", rec.Body, test.wantBody)
		}
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		db       database
		request  string
		wantBody string
		wantDB   map[string]dollars
	}{
		{database{"fuga": 0}, "/?item=fuga&price=200", "fuga: $200.00 setted\n", map[string]dollars{"fuga": 200}},
		{database{}, "/?item=hoge&price=200", "hoge is not exist.", map[string]dollars{}},
		{database{}, "/?item=hoge", "field 'price' error: strconv.ParseFloat: parsing \"\": invalid syntax", map[string]dollars{}},
		{database{}, "/?price=200", "field 'item' must be setted.", map[string]dollars{}},
		{database{"hoge": 100, "fuga": 200}, "/?item=hoge&price=200", "hoge: $200.00 setted\n", map[string]dollars{"hoge": 200, "fuga": 200}},
	}

	for _, test := range tests {
		req := httptest.NewRequest("PUT", test.request, nil)
		rec := httptest.NewRecorder()
		db := test.db

		db.update(rec, req)

		if !isSame(db, test.wantDB) {
			t.Errorf("---get body---\n %v \n ---want body---\n %v \n", db, test.wantDB)
		}

		if rec.Body.String() != test.wantBody {
			t.Errorf("get request Body is %q, want %q", rec.Body, test.wantBody)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		db       database
		request  string
		wantBody string
		wantDB   map[string]dollars
	}{
		{database{"fuga": 0}, "/?item=fuga", "fuga is deleted", map[string]dollars{}},
		{database{"fuga": 0}, "/?item=hoge", "hoge is not exist.", map[string]dollars{"fuga": 0}},
		{database{"hoge": 100, "fuga": 200}, "/?item=hoge", "hoge is deleted", map[string]dollars{"fuga": 200}},
	}

	for _, test := range tests {
		req := httptest.NewRequest("PUT", test.request, nil)
		rec := httptest.NewRecorder()
		db := test.db

		db.deleteItem(rec, req)

		if !isSame(db, test.wantDB) {
			t.Errorf("---get body---\n %v \n ---want body---\n %v \n", db, test.wantDB)
		}

		if rec.Body.String() != test.wantBody {
			t.Errorf("get request Body is %q, want %q", rec.Body, test.wantBody)
		}
	}
}
