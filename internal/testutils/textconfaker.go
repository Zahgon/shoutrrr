package testutils

import (
	"bufio"
	"bytes"
	"net/textproto"
	"strings"
)

type textConFaker struct {
	inputBuffer  *bytes.Buffer
	inputWriter  *bufio.Writer
	outputReader *bufio.Reader
	responses    []string
	delim        string
}

func (tcf *textConFaker) GetInput() string { _ = "STUB: not implemented"; return "" }

// GetConversation returns the input and output streams as a conversation
func (tcf *textConFaker) GetConversation(includeGreeting bool) string {
	_ = "STUB: not implemented"
	return ""
}

// GetClientSentences returns all the input recieved from the client separated by the delimiter
func (tcf *textConFaker) GetClientSentences() []string { _ = "STUB: not implemented"; return nil }

// CreateReadWriter returns a ReadWriter from the textConFakers internal reader and writer
func (tcf *textConFaker) CreateReadWriter() *bufio.ReadWriter {
	_ = "STUB: not implemented"
	return nil
}

func (tcf *textConFaker) init() {
	tcf.inputBuffer = &bytes.Buffer{}
	stringReader := strings.NewReader(strings.Join(tcf.responses, tcf.delim))
	tcf.outputReader = bufio.NewReader(stringReader)
	tcf.inputWriter = bufio.NewWriter(tcf.inputBuffer)
}

// CreateTextConFaker returns a textproto.Conn to fake textproto based connections
func CreateTextConFaker(responses []string, delim string) (*textproto.Conn, Eavesdropper) {
	_ = "STUB: not implemented"
	return nil, *new(Eavesdropper)
}

// rx := iotest.NewReadLogger("TextConRx", tcfaker.outputReader)
// tx := iotest.NewWriteLogger("TextConTx", tcfaker.inputWriter)
// faker := CreateIOFaker(rx, tx)
