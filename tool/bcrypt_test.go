package tool

import "testing"

func TestBcrypt(t *testing.T) {
	text := "text testing"
	hash, err := Bcrypt([]byte(text)).Hash()

	if err != nil {
		t.Errorf("error creating bcrypt hash: %s", err)
	}

	err = Bcrypt(nil).Compare([]byte(hash))

	if err != nil {
		t.Errorf("generated hash was not validated: %s", err)
	}
}
