package parser

import (
	"fmt"
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
				fmt.Println(line)
				field := strings.Split(line, " = ")
				fmt.Println(field)
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

func (p *Parser) Stringify(config Config) string {
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
	s := p.Stringify(p.config)
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
	p.config[section][key] = value
}
