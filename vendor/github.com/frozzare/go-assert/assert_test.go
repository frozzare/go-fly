package assert

import "testing"

func TestEqual(t *testing.T) {
	True(t, Equal(t, "foo", "foo", "Foo will match foo"))
}

func TestNotEqual(t *testing.T) {
	True(t, NotEqual(t, "foo", "boo", "Foo will not match boo"))
}

func TestTrue(t *testing.T) {
	True(t, True(t, true, "true will match true"))
}

func TestFalse(t *testing.T) {
	True(t, False(t, false, "false will match false"))
}

func TestNotNil(t *testing.T) {
	True(t, NotNil(t, true, "true is not nil"))
}

func TestEmpty(t *testing.T) {
	values := []interface{}{
		nil,
		"",
		"",
		[]string{},
		0,
		0.0,
	}

	for _, v := range values {
		True(t, Empty(t, v))
	}

	values = []interface{}{
		"hello",
		1,
		[]string{""},
	}

	for _, v := range values {
		True(t, NotEmpty(t, v))
	}
}
