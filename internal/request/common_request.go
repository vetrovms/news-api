package request

const (
	LocEn      = "en"
	LocUk      = "uk"
	DefaultLoc = LocEn
)

func LocInWhiteList(locale string) bool {
	if locale == LocEn {
		return true
	}
	if locale == LocUk {
		return true
	}
	return false
}
