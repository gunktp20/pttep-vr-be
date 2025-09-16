package errorMessage_test

import (
	"pttep-vr-api/pkg/utils/errorMessage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	raw := `
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
  "error_4": {
    "code": "4",
    "message": {
      "my": "error_3",
      "mk": "error_3"
    }
  },
  "ping_error": {
    "code": "5",
    "message": {
      "th": "ping_error",
      "en": "ping_error"
    }
  }
}
`

	err := errorMessage.Read(raw)
	assert.NoError(t, err)

	t.Run("Get", func(t *testing.T) {
		result := errorMessage.Get("success")
		assert.Equal(t, result.Code, "0")
	})

	t.Run("Get", func(t *testing.T) {
		result := errorMessage.Get("zxcvbnm")
		assert.Equal(t, result.Code, "9999")

		result.Message.Language("th")
		result.Message.Language("en")
		result.Message.Language("ch")

	})
	t.Run("Get", func(t *testing.T) {
		result := errorMessage.Get("error_4")

		result.Message.Language("th")
		result.Message.Language("en")
		result.Message.Language("ch")

	})

}

func Test_Read(t *testing.T) {

	t.Run("Read", func(t *testing.T) {
		err := errorMessage.Read(`errorMessage`)
		assert.Error(t, err)
	})

	t.Run("Read", func(t *testing.T) {
		raw := `{
			"error_1": {
				"code": "1",
				"message": {
					"th": "error_1",
					"en": "error_1"
				}
			},
			"errorGroup1": {
				"errorGroup1_1": {
					"error_2": {
						"code": "2",
						"message": {
							"th": {
								"message": "error_2"
							},
							"en": {
								"message": "error_2"
							}
						}
					}
				}
			}
		}`
		err := errorMessage.Read(raw)
		assert.Error(t, err)
	})
}
