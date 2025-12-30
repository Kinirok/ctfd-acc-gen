package gen

import (
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	t.Run("password_charset", func(t *testing.T) {
		password := GeneratePassword()
		allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

		for _, char := range password {
			if !strings.ContainsRune(allowedChars, char) {
				t.Errorf("Password contains invalid character: %c", char)
			}
		}
	})

	t.Run("password_randomness", func(t *testing.T) {
		pwd1 := GeneratePassword()
		pwd2 := GeneratePassword()

		if pwd1 == pwd2 {
			t.Errorf("Duplicate password generated: %s", pwd1)
		}
	})

}

func BenchmarkGeneratePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword()
	}
}

func TestGenerateLogin(t *testing.T) {
	t.Run("password_charset", func(t *testing.T) {
		login := GenerateLogin()
		allowedChars := "0123456789"
		if !strings.HasPrefix(login, "user_") {
			t.Errorf("Username should start with user_: %s", login)
		}

		for i := 5; i < len(login); i++ {
			if !strings.Contains(allowedChars, string(login[i])) {
				t.Errorf("Username %s contains invalid character: %c", login, login[i])
			}
		}
	})

	t.Run("password_randomness", func(t *testing.T) {
		log1 := GenerateLogin()
		log2 := GenerateLogin()

		if log1 == log2 {
			t.Errorf("Duplicate users generated: %s", log1)
		}
	})

}

func BenchmarkGenerateLogin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateLogin()
	}
}

func TestGenerateTeamName(t *testing.T) {
	t.Run("password_charset", func(t *testing.T) {
		login := GenerateTeamName()
		allowedChars := "0123456789"
		if !strings.HasPrefix(login, "team_") {
			t.Errorf("Username should start with team_: %s", login)
		}

		for i := 5; i < len(login); i++ {
			if !strings.Contains(allowedChars, string(login[i])) {
				t.Errorf("Username %s contains invalid character: %c", login, login[i])
			}
		}
	})

	t.Run("password_randomness", func(t *testing.T) {
		team1 := GenerateTeamName()
		team2 := GenerateTeamName()

		if team1 == team2 {
			t.Errorf("Duplicate teams generated: %s", team1)
		}
	})

}

func BenchmarkGenerateTeamName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateTeamName()
	}
}
