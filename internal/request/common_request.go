package request

const (
	LocEn      = "en"
	LocUk      = "uk"
	DefaultLoc = LocEn
)

// LocInWhiteList Перевіряє що передана користувачем локаль дозволена.
func LocInWhiteList(locale string) bool {
	if locale == LocEn {
		return true
	}
	if locale == LocUk {
		return true
	}
	return false
}
