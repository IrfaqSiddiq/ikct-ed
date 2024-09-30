package utility

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"os"
	"strings"
	"time"
)

// SQLNullStringToString convert sql null string to string
func SQLNullStringToString(s sql.NullString) string {
	if s.Valid {
		return strings.TrimSpace(s.String)
	}

	return ""

}

// SQLNullIntToInt convert sql null int to int
func SQLNullIntToInt(i sql.NullInt64) int64 {
	if i.Valid {
		return i.Int64
	}

	return 0

}

// SQLNullFloatToFloat convert sql null float to float
func SQLNullFloatToFloat(f sql.NullFloat64) float64 {
	if f.Valid {
		return f.Float64
	}

	return 0.0

}

// SQLNullTimeToTime convert sql null time to time
func SQLNullTimeToTime(t sql.NullTime) time.Time {
	if t.Valid {
		return t.Time
	}

	return time.Now()

}

func SQLNullBoolToBool(b sql.NullBool) bool {
	if b.Valid {
		return b.Bool
	}
	return false
}

// Hash password with salt using SHA-256
func HashPassword(password string) string {
	salt := GetSaltDetails()
	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Verify the password
func VerifyPassword(hashedPassword, password string) bool {
	return HashPassword(password) == hashedPassword
}

func GetSaltDetails() string {
	return os.Getenv("SALT")
}

func GetHostURL()string{
	return os.Getenv("HOST_URL")
}
