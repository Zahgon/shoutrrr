package smtp

import "github.com/containrrr/shoutrrr/internal/failures"

const (
	// FailUnknown is the default FailureID
	FailUnknown failures.FailureID = iota
	// FailGetSMTPClient is returned when a SMTP client could not be created
	FailGetSMTPClient
	// FailEnableStartTLS is returned when failing to enable StartTLS
	FailEnableStartTLS
	// FailAuthType is returned when the Auth type could not be identified
	FailAuthType
	// FailAuthenticating is returned when the authentication fails
	FailAuthenticating
	// FailSendRecipient is returned when sending to a recipient fails
	FailSendRecipient
	// FailClosingSession is returned when the server doesn't accept the QUIT command
	FailClosingSession
	// FailPlainHeader is returned when the text/plain multipart header could not be set
	FailPlainHeader
	// FailHTMLHeader is returned when the text/html multipart header could not be set
	FailHTMLHeader
	// FailMultiEndHeader is returned when the multipart end header could not be set
	FailMultiEndHeader
	// FailMessageTemplate is returned when the message template could not be written to the stream
	FailMessageTemplate
	// FailMessageRaw is returned when a non-templated message could not be written to the stream
	FailMessageRaw
	// FailSetSender is returned when the server didn't accept the sender address
	FailSetSender
	// FailSetRecipient is returned when the server didn't accept the recipient address
	FailSetRecipient
	// FailOpenDataStream is returned when the server didn't accept the data stream
	FailOpenDataStream
	// FailWriteHeaders is returned when the headers could not be written to the data stream
	FailWriteHeaders
	// FailCloseDataStream is returned when the server didn't accept the data stream contents
	FailCloseDataStream
	// FailConnectToServer is returned when the TCP connection to the server failed
	FailConnectToServer
	// FailCreateSMTPClient is returned when the smtp.Client initialization failed
	FailCreateSMTPClient
	// FailApplySendParams is returned when updating the send config failed
	FailApplySendParams
	// FailHandshake is returned when the initial HELLO handshake returned an error
	FailHandshake
)

func fail(failureID failures.FailureID, err error, v ...interface{}) failure {
	_ = "STUB: not implemented"
	return *new(failure)
}

// case FailUnknown:

type failure interface {
	failures.Failure
}
