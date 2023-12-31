package utils

import "testing"

func TestSlug(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"Challenge Of The Commander", "challenge-of-the-commander"},
		{"指挥官的挑战", "zhi-hui-guan-de-tiao-zhan"},
	}

	for index, st := range testCases {
		got := UnicodeSlug(st.in)
		if got != st.want {
			t.Errorf(
				"%d. UnicodeSlug(%#v) = %#v; want %#v",
				index, st.in, got, st.want)
		}
	}
}
