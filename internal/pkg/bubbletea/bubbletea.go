package bubbletea

type (
	errMsg error
)

func GetCliString(str string, color string) string {
	code := "30"
	switch color {
	case "black", "dark":
		code = "30"
	case "red":
		code = "31"
	case "green":
		code = "32"
	case "yellow":
		code = "33"
	case "blue":
		code = "34"
	case "purple":
		code = "35"
	case "cyan":
		code = "36"
	case "white":
		code = "37"
	}
	return "\033[" + code + "m" + str + "\033[0m"
}
