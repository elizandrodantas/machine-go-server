package loggers

type LoggerType int

const (
	UPDATE_PASSWORD LoggerType = iota
	CREATE_USER
	MACHINE_VERIFY
	LOGIN
	ACCEPT_MACHINE
	BLOCK_MACHINE
	SUSPECT_MACHINE
)

var TypesNames = [...]string{
	"UPDATE_PASSWORD ",
	"CREATE_USER",
	"MACHINE_VERIFY",
	"LOGIN",
	"ACCEPT_MACHINE",
	"BLOCK_MACHINE",
	"SUSPECT_MACHINE",
}

func (y LoggerType) String() string {
	if y < UPDATE_PASSWORD || y > SUSPECT_MACHINE {
		return "Unknown"
	}
	return TypesNames[y]
}

func IsValidLoggerType(y LoggerType) bool {
	return y >= UPDATE_PASSWORD && y <= SUSPECT_MACHINE
}

func IsValidString(typ string) bool {
	res := false

	for _, k := range TypesNames {
		if k == typ {
			res = true
			break
		}
	}

	return res
}
