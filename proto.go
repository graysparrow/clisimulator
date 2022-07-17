package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// field type
type FieldType int

const (
	T_DOUBLE FieldType = iota
	T_FLOAT
	T_INT32
	T_INT64
	T_UINT32
	T_UINT64
	T_BOOL
	T_STRING
	T_BYTES
	T_STRUCT
)

func (v FieldType) String() string {
	switch v {
	case T_DOUBLE:
		return "T_DOUBLE"
	case T_FLOAT:
		return "T_FLOAT"
	case T_INT32:
		return "T_INT32"
	case T_INT64:
		return "T_INT64"
	case T_UINT32:
		return "T_UINT32"
	case T_UINT64:
		return "T_UINT64"
	case T_BOOL:
		return "T_BOOL"
	case T_STRING:
		return "T_STRING"
	case T_BYTES:
		return "T_BYTES"
	case T_STRUCT:
		return "T_STRUCT"
	default:
		return "unknown"
	}
}

// option type
type OptionType int

const (
	OPTIONAL OptionType = iota
	REQUIRED
	REPEATED
)

func (v OptionType) String() string {
	switch v {
	case OPTIONAL:
		return "OPTIONAL"
	case REQUIRED:
		return "REQUIRED"
	case REPEATED:
		return "REPEATED"
	default:
		return "unknown"
	}
}

// Field reprensents a proto field
type Field struct {
	Option OptionType
	Type   FieldType
	Name   string
}

// String return a description string of field
func (f *Field) String() string {
	return fmt.Sprintf("%v %v %v", f.Option, f.Type, f.Name)
}

// Parse
func (f *Field) Parse(line string) {

	line = strings.TrimLeftFunc(line, func(ch rune) bool {
		return ch == indent
	})

	tokens := strings.Split(line, string(indent))

	fmt.Println(tokens)

	if (len(tokens)) >= 3 {

		switch tokens[1] {
		case "optional":
			f.Option = OPTIONAL
		case "required":
			f.Option = REQUIRED
		case "repeated":
			f.Option = REPEATED
		default:
			log.WithField("line", line).Error("parse option error")
		}

		switch tokens[2] {
		case "double":
			f.Type = T_DOUBLE
		case "float":
			f.Type = T_FLOAT
		case "int32":
			f.Type = T_INT32
		case "uint32":
			f.Type = T_UINT32
		case "int64":
			f.Type = T_INT64
		case "uint64":
			f.Type = T_UINT64
		case "bool":
			f.Type = T_BOOL
		case "string":
			f.Type = T_STRING
		case "bytes":
			f.Type = T_BYTES
		default:
			f.Type = T_STRUCT
		}

		if len(tokens[3]) > 0 {
			f.Name = tokens[3]
		} else {
			log.WithField("line", line).Error("parse name error")
		}

	} else {
		log.WithField("line", line).Error("parse line error")
	}
}

// Message reprensents a proto message
type Message struct {
	Name   string
	Fields []*Field
}

// String return a description string of field
func (m *Message) String() string {
	return fmt.Sprintf("%v", m.Name)
}

func (m *Message) Parse(line string) {

	line = strings.TrimLeftFunc(line, func(ch rune) bool {
		return ch == indent
	})

	tokens := strings.Split(line, string(indent))

	if (len(tokens)) >= 2 {
		m.Name = tokens[1]
	} else {
		log.WithField("line", line).Error("parse message error")
	}
}

func ParseFile(path string) []*Message {
	file, err := os.Open(path)
	if err != nil {
		log.WithFields(logrus.Fields{
			"path":  path,
			"error": err,
		}).Fatal("open parse file error")
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	const MAX_NESTED_SIZE = 100
	msgStack := make([]*Message, MAX_NESTED_SIZE, MAX_NESTED_SIZE) // a stack for nested message definitions
	msgIdx := -1

	msgResult := make([]*Message, 0, 0)
	for {
		line, err := reader.ReadString('\n')
		isEnd := false
		if err != nil {
			if err == io.EOF {
				isEnd = true
			} else {
				break
				log.WithField("line", line).Error("read line error")
			}
		}

		if isEnd {
			break
		}

		if strings.Index(line, "message") > -1 {

			msgIdx++
			fmt.Println("++", msgIdx)

			msg := &Message{Fields: make([]*Field, 0, 0)}
			msgStack[msgIdx] = msg
			msgStack[msgIdx].Parse(line)
			msgResult = append(msgResult, msg)

		} else if strings.Index(line, "msgend") > -1 {

			// find message definition end, pop a message
			msgIdx--
			fmt.Println("--", msgIdx)

		} else if strings.Index(line, "field") > -1 {

			fmt.Println("..", msgIdx)
			msg := msgStack[msgIdx]

			field := &Field{}
			field.Parse(line)
			msg.Fields = append(msg.Fields, field)

		}
	}
	return msgResult
}
