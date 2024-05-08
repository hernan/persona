package session

import (
	"database/sql"
	"fmt"
	"persona/models"
)

type Session struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id" db:"user_id"`
	Session   *string `json:"session" db:"session"`
	CreatedAt *string `json:"created_at" db:"created_at"`
	TimeoutAt *string `json:"timeout_at" db:"timeout_at"`
}

func FindBySession(session string) (Session, error) {
	ses := Session{}
	row := models.MyDb.QueryRow("SELECT * FROM sessions WHERE session = ?", session)
	err := row.Scan(&ses.ID, &ses.UserID, &ses.Session, &ses.CreatedAt, &ses.TimeoutAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return Session{}, fmt.Errorf("session not found")
		}
		return Session{}, err
	}

	return ses, nil
}

func Create(session Session) (Session, error) {
	stmt, err := models.MyDb.Prepare("INSERT INTO sessions (user_id, session) VALUES (?, ?)")
	if err != nil {
		return Session{}, err
	}

	res, err := stmt.Exec(session.UserID, session.Session)
	if err != nil {
		return Session{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Session{}, err
	}

	session.ID = int(id)
	return session, nil
}
