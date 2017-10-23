package editor

import (
	"bytes"
	"io"
)

// Editor is the line-processor interface
type Editor interface {
	Edit(line string) (string, error)
}

// Stream build io.WriteCloser that process lines with editor and write to base io.Writer
func Stream(base io.Writer, editor ...Editor) io.WriteCloser {
	return &stream{editors: editor, base: base}
}

type stream struct {
	editors []Editor
	base    io.Writer
	buffer  []byte
}

func (s *stream) writeLines(lines [][]byte) error {
	for _, line := range lines {
		line := bytes.TrimSuffix(line, []byte{'\r'})
		text := string(append(line, '\n'))
		for _, e := range s.editors {
			t, err := e.Edit(text)
			if err != nil {
				return err
			}
			text = t
		}
		if _, err := s.base.Write([]byte(text)); err != nil {
			return err
		}
	}
	return nil
}

func (s *stream) Write(b []byte) (int, error) {
	lines := bytes.Split(append(s.buffer, b...), []byte("\n"))
	s.buffer = lines[len(lines)-1]
	lines = lines[:len(lines)-1]
	if err := s.writeLines(lines); err != nil {
		return 0, err
	}
	return len(b), nil
}

func (s *stream) Close() error {
	lines := bytes.Split(s.buffer, []byte(`\n`))
	s.buffer = nil
	return s.writeLines(lines)
}
