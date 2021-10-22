package benchmark

import "testing"

func BenchmarkMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map("魔人ブウ")
	}
}

func BenchmarkSwitch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Switch("魔人ブウ")
	}
}

var m = map[string]string{
	"はにわ":  "はにわ",
	"魔人ブウ": "魔人ブウ",
	"おさげ":  "おさげ",
	"トッポギ": "トッポギ",
	"桜前線":  "桜前線",
}

func Map(s string) string {
	return m[s]
}

func Switch(s string) string {
	switch s {
	case "はにわ":
		return "はにわ"
	case "魔人ブウ":
		return "魔人ブウ"
	case "おさげ":
		return "おさげ"
	case "トッポギ":
		return "トッポギ"
	case "桜前線":
		return "桜前線"
	default:
		return ""
	}
}
