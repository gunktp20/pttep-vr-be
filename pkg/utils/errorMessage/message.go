package errorMessage

import (
	"strings"
)

const (
	TH = "TH"
	EN = "EN"
)

type ErrorMessage struct {
	Code    string  `json:"code"`
	Message Message `json:"message"` // index to lang
}

type Message map[string]string

func (m Message) Language(l string) string {
	var message string
	for k, v := range m {
		if k == l {
			return v
		}
		if strings.ToUpper(k) == EN {
			message = v
		}
		if message == "" {
			message = v
		}
	}
	return message
}

var Raw = `
{
  "success": {
    "code": "0",
    "message": {
      "th": "success",
      "en": "success"
    }
  },
  "errorGroup1": {
    "error_1": {
      "code": "1",
      "message": {
        "th": "error_1",
        "en": "error_1"
      }
    },
    "error_2": {
      "code": "2",
      "message": {
        "th": "error_2",
        "en": "error_2"
      }
    }
  },
  "error_3": {
    "code": "3",
    "message": {
      "th": "error_3",
      "en": "error_3"
    }
  },
  "ping_error": {
    "code": "4",
    "message": {
      "th": "ping_error",
      "en": "ping_error"
    }
  }
}
`
