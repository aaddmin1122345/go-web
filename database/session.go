package database

//import (
//	"database/sql"
//	"fmt"
//	"github.com/gorilla/sessions"
//	"go-web/model"
//)
//
//type Session interface {
//	NewSession(session *model.SessionModel) error
//	DeleteSession()
//}
//
//type SessionImpl struct {
//	Session sessions.Store
//	db      *sql.DB
//}
//
//func (s *SessionImpl) SetDb(db *sql.DB) {
//	s.db = db
//}
//
//func (s *SessionImpl) NewSession(session *model.SessionModel) error {
//
//	query := "INSERT INTO sessions (ID, UserName, Session, Expiration) VALUES (?, ?, ?, ?)"
//	_, err := s.db.Exec(query, session.ID, session.UserName, session.SessionKey, session.Expiration)
//	if err != nil {
//		return err
//	}
//
//	fmt.Println(*session)
//	return nil
//}
//
//func (s *SessionImpl) DeleteSession() {
//	return
//}
