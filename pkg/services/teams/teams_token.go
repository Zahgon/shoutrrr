package teams

var uuid4Pattern = "[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}"
var hex32Pattern = "[A-Za-z0-9]{32}"

func uuidPartValid(token string) bool { _ = "STUB: not implemented"; return false }

func hashPartValid(token string) bool { _ = "STUB: not implemented"; return false }

func verifyWebhookParts(p [4]string) error { _ = "STUB: not implemented"; return nil }

func matchesRegexp(pattern string, token string) bool { _ = "STUB: not implemented"; return false }
