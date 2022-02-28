package command

import "testing"

func TestRun(t *testing.T) {

	assertCorrect := func(t *testing.T, got, want string, gotErr error, wantErr bool) {
		t.Helper()
		if (gotErr != nil) != wantErr {
			t.Errorf("gotErr %v, wantErr %v", gotErr, wantErr)
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("Success", func(t *testing.T) {
		got, gotErr := Run("sleep 1;echo 'sleep1'")
		want := "sleep1"
		wantErr := false

		assertCorrect(t, got, want, gotErr, wantErr)
	})

	t.Run("Success (not return result)", func(t *testing.T) {
		got, gotErr := Run("sleep 1")
		want := ""
		wantErr := false

		assertCorrect(t, got, want, gotErr, wantErr)
	})

	t.Run("Failed", func(t *testing.T) {
		got, gotErr := Run("error command")
		want := ""
		wantErr := true

		assertCorrect(t, got, want, gotErr, wantErr)
	})
}
