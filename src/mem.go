package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Memory is used to simulate the function of a physical memory
type Memory struct {
	content []byte
}

// Init is used to initialize the content array in Memory
func (m *Memory) Init(size int) {
	content := make([]byte, size)
	m.content = content
}

// Load function is used to load binary file into memory
func (m Memory) Load(infileName string) {
	file, _ := os.Open(infileName)
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Scan()
	line := s.Text()

	for len(line) > 0 && line[0] == '@' {
		// It is the starting address of the content
		offset, _ := strconv.Atoi(line[1:])

		// Read content
		s.Scan()
		line = s.Text()

		for len(line) > 0 && line[0] != '@' {
			bytes := strings.Split(line, " ")
			for idx, curByte := range bytes {
				if len(curByte) <= 0 {
					break
				}
				intValue := convertHexToInt(curByte[0])*16 +
					convertHexToInt(curByte[1])
				m.content[offset+idx] = byte(intValue)
			}
			s.Scan()
			line = s.Text()
		}
	}
}

func (m Memory) Read(offset int, size int) []byte {
	bytes := make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = m.content[offset+i]
	}
	return bytes
}

// Helper functions
func convertHexToInt(c byte) int {
	var ans int
	if c >= '0' && c <= '9' {
		ans = int(c - '0')
	} else {
		ans = int(c - 'A' + 10)
	}

	return ans
}
