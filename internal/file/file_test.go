package file

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {

	assertCorrect := func(t *testing.T, got, want []string, gotErr error, wantErr bool) {
		t.Helper()
		if (gotErr != nil) != wantErr {
			t.Errorf("gotErr %v, wantErr %v", gotErr, wantErr)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("Success", func(t *testing.T) {
		got, gotErr := Read("../../test_data/test.txt")
		want := []string{
			"sleep 1; sleep 1; echo 'sleep 2'",
			"sleep 1; sleep 1; sleep 1; echo 'sleep 3'",
			"echo 'hello'",
		}
		wantErr := false

		assertCorrect(t, got, want, gotErr, wantErr)
	})

	t.Run("Failed", func(t *testing.T) {
		got, gotErr := Read("../../test_data/failed")
		want := []string{}
		wantErr := true

		assertCorrect(t, got, want, gotErr, wantErr)
	})
}
