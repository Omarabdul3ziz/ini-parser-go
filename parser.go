package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config map[string]map[string]string

type Parser struct {
	config Config
}

func NewParser() Parser {
	parser := Parser{}
	parser.config = make(Config)
	return parser
}

func (p *Parser) ReadFromString(s string) {
	lines := strings.Split(s, "\n")

	var current_section string

	for _, line := range lines {
		if line != "" {
			switch string(line[0]) {
			case "#":
			case ";":
				continue
			case "[":
				current_section = line[1 : len(line)-1]
				p.config[current_section] = map[string]string{}
			default:
				field := strings.Split(line, " = ")
				key, value := string(field[0]), string(field[1])
				p.config[current_section][key] = value
			}
		}
	}
}

func (p *Parser) ReadFromFile(file_path string) error {
	stream, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}

	s := string(stream)

	p.ReadFromString(s)
	return nil
}

func stringify(config Config) string {
	content := ""

	for title, body := range config {
		content += fmt.Sprintf("[%s]\n", title)
		for key, value := range body {
			content += fmt.Sprintf("%s = %s\n", key, value)
		}
		content += "\n"
	}

	return content
}

func (p *Parser) WriteToFile(file_path string) error {
	s := stringify(p.config)
	bytes := []byte(s)

	file, c_err := os.Create(file_path)

	if c_err != nil {
		return c_err
	}

	_, w_err := file.Write(bytes)
	if w_err != nil {
		return w_err
	}

	return nil
}

func (p *Parser) Get(section, key string) string {
	return p.config[section][key]
}

func (p *Parser) Set(section, key, value string) {
	inner := map[string]string{key: value}
	p.config[section] = inner
}

func ReadFromFile(file_path string) (*os.File, error) {
	file, err := os.Open(file_path)

	if err != nil {
		return nil, err
	}

	// The returned value of type *os.File implements the io.Reader interface.
	return file, nil
}

func ReadFromStirng(input string) io.Reader {
	return strings.NewReader(input)
}

func Parse(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
