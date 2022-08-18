package parser

import (
	"testing"
)

func TestReadFromString(t *testing.T) {
	p := NewParser()

	s := `
[Profile]
name = jarvis

# credential
password = secret
`
	p.ReadFromString(s)

	got := p.config["Profile"]["password"]
	exp := "secret"
	if got != exp {
		t.Errorf("Got: %s. Expected: %s", got, exp)
	}
}

func TestReadFromFile(t *testing.T) {
	p := NewParser()

	file_path := "./sample.ini"

	p.ReadFromFile(file_path)

	got := p.config["Profile"]["password"]
	exp := "secret"
	if got != exp {
		t.Errorf("Got: %s. Expected: %s", got, exp)
	}

}

func TestWriteToFile(t *testing.T) {
	p := NewParser()

	source := "./sample.ini"
	dist := "./sample2.ini"

	p.ReadFromFile(source)
	exp := p.config["Profile"]["password"]

	p.WriteToFile(dist)
	p.ReadFromFile(dist)
	got := p.config["Profile"]["password"]

	if got != exp {
		t.Errorf("Got: %s. Expected: %s", got, exp)
	}
}

func TestSet(t *testing.T) {
	p := NewParser()

	file_path := "./sample.ini"

	p.ReadFromFile(file_path)

	p.Set("General", "network", "Dev")

	got := p.config["General"]["network"]
	exp := "Dev"

	if got != exp {
		t.Errorf("Got: %s. Expected: %s", got, exp)
	}
}

func TestGet(t *testing.T) {
	p := NewParser()

	file_path := "./sample.ini"

	p.ReadFromFile(file_path)

	got := p.Get("Profile", "name")
	exp := p.config["Profile"]["name"]

	if got != exp {
		t.Errorf("Got: %s. Expected: %s", got, exp)
	}
}

func TestFlow(t *testing.T) {
	p := NewParser()

	p.ReadFromFile("./sample.ini")

	p.Set("General", "network", "dev")

	p.WriteToFile("./sample2.ini")
}
