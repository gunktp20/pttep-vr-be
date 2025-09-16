package constant

const PLAYER = "PLAYER"
const ADMIN = "ADMIN"

const (
	DataNotFound     = "DataNotFound"
	DataDeleteFailed = "DataDeleteFailed"

	CannotConvertValue = "CannotConvertValue"

	NotSupported = "NotSupported"
)

// Login Type
type TYPELoginType string

var LoginType loginType

type loginType struct{}

// username and password
func (loginType) Default() TYPELoginType {
	return TYPELoginType("DEFAULT")
}

// email and password
func (loginType) Email() TYPELoginType {
	return TYPELoginType("EMAIL")
}

// email and password
func (loginType) Telephone() TYPELoginType {
	return TYPELoginType("TELEPHONE")
}

// create new account for guest
func (loginType) Guest() TYPELoginType {
	return TYPELoginType("GUEST")
}

// create new account for Contractor
func (loginType) Contractor() TYPELoginType {
	return TYPELoginType("CONTRACTOR")
}

// varify by line token
func (loginType) Line() TYPELoginType {
	return TYPELoginType("LINE")
}

// varify by facebook token
func (loginType) Facebook() TYPELoginType {
	return TYPELoginType("FACEBOOK")
}

// varify by google token
func (loginType) Google() TYPELoginType {
	return TYPELoginType("GOOGLE")
}

// varify by PTTEP
func (loginType) PTTEP() TYPELoginType {
	return TYPELoginType("PTTEP")
}

type DatabaseType string

const (
	GORM  DatabaseType = "GORM"
	MONGO DatabaseType = "MONGO"
)
