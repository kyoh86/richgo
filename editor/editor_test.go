package editor

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type editorFunc func(string) (string, error)

func (f editorFunc) Edit(line string) (string, error) {
	return f(line)
}

func addPrefixEditor(prefix string) Editor {
	return editorFunc(func(line string) (string, error) {
		return prefix + line, nil
	})
}

func TestWriter(t *testing.T) {
	t.Run("valid editor and writer", func(t *testing.T) {
		var buffer bytes.Buffer
		stream := Stream(&buffer, addPrefixEditor("A"), addPrefixEditor("B"))
		{
			_, err := stream.Write([]byte("ab"))
			require.NoError(t, err, "write in stream")
			assert.Empty(t, buffer.Bytes())
		}
		{
			_, err := stream.Write([]byte("\ncd"))
			require.NoError(t, err, "write in stream")
			assert.Equal(t, "BAab\n", buffer.String())
		}
		{
			require.NoError(t, stream.Close(), "close a stream")
			// It's undesirable spec :(
			// A following newline is unnecessary.
			assert.Equal(t, "BAab\nBAcd\n", buffer.String())
		}
	})
	t.Run("invalid editor", func(t *testing.T) {
		var buffer bytes.Buffer
		stream := Stream(&buffer, editorFunc(func(string) (string, error) { return "", errors.New("test") }))
		{
			_, err := stream.Write([]byte("ab"))
			require.NoError(t, err, "write in stream")
			assert.Empty(t, buffer.Bytes())
		}
		{
			_, err := stream.Write([]byte("\ncd"))
			require.EqualError(t, err, "test")
			assert.Empty(t, buffer.Bytes())
		}
	})
}
