package repository

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"go-chzzk/config"
	"go-chzzk/types/schema"
	"log"
	"strings"
)

type Repository struct {
	cfg *config.Config
	db  *sql.DB
}

const (
	room       = "chatting.room"
	chat       = "chatting.chat"
	serverInfo = "chatting.serverInfo"
)

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{cfg: cfg}
	var err error

	db, err := sql.Open(cfg.DB.Database, cfg.DB.URL)
	if err != nil {
		return nil, err
	}
	r.db = db

	return r, nil
}

func (r *Repository) InsertChatting(user, message, roomName string) error {
	log.Println("Insert chatting using wss", "from", user, "message", message, "room", roomName)
	_, err := r.db.Exec("insert into chatting.chat(room, name, message) values(?, ?, ?)", roomName, user, message)
	return err
}

func (r *Repository) GetChatList(roomName string) ([]*schema.Chat, error) {
	qs := query([]string{"select * from", chat, "where room = ? order by `when` desc limit 10"})

	cursor, err := r.db.Query(qs, roomName)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var result []*schema.Chat
	for cursor.Next() {
		d := &schema.Chat{}

		if err := cursor.Scan(&d.ID, &d.Room, &d.Name, &d.Message, &d.When); err != nil {
			return nil, err
		}
		result = append(result, d)
	}

	if len(result) == 0 {
		return []*schema.Chat{}, nil
	}
	return result, nil
}

func (r *Repository) RoomList() ([]*schema.Room, error) {
	qs := query([]string{"select * from", room})

	cursor, err := r.db.Query(qs)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var result []*schema.Room
	for cursor.Next() {
		d := &schema.Room{}

		if err = cursor.Scan(&d.ID, &d.Name, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, d)
	}

	if len(result) == 0 {
		return []*schema.Room{}, nil
	}
	return result, nil
}

func (r *Repository) MakeRoom(name string) error {
	_, err := r.db.Exec("insert into chatting.room(name) values(?)", name)
	return err
}

func (r *Repository) Room(name string) (*schema.Room, error) {
	d := &schema.Room{}
	// select * from chatting.room where name = ?
	qs := query([]string{"select * from", room, "where name = ?"})

	err := r.db.QueryRow(qs, name).Scan(&d.ID, &d.Name, &d.CreatedAt, &d.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return d, err
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
