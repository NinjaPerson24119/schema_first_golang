package typeconversions

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func UUIDToString(uuid uuid.UUID) string {
	if uuid.IsNil() {
		return ""
	}
	return uuid.String()
}
func StringToUUID(uuidStr string) (uuid.UUID, error) {
	u, err := uuid.FromString(uuidStr)
	if err != nil {
		return uuid.UUID{}, err
	}
	return u, nil
}

func PGTypeTextToString(text pgtype.Text) string {
	if !text.Valid {
		return ""
	}
	return text.String
}
func StringToPGTypeText(str string) pgtype.Text {
	return pgtype.Text{
		String: str,
		Valid:  true,
	}
}

func PGTypeTimestampTZToTime(t pgtype.Timestamptz) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}
func TimeToPGTypeTimestampTZ(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{
			Valid: false,
		}
	}
	return pgtype.Timestamptz{
		Time:  *t,
		Valid: true,
	}
}
