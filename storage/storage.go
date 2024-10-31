package storage

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func NewDB(storagePath string) *sql.DB {

	db, err := sql.Open("sqlite", storagePath)
	if err != nil {
		log.Fatalf("Ошибка открытия БД: %v", err)
	}

	stmt, err := db.Prepare(`
CREATE TABLE IF NOT EXISTS motivations(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content VARCHAR(128) NOT NULL DEFAULT '',
    author VARCHAR(128) NOT NULL DEFAULT '');
`)
	if err != nil {
		return nil
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil
	}
	return db

}

func InsertInDb(db *sql.DB) error {
	stmt, err := db.Prepare(`
INSERT INTO motivations (content,author) VALUES
    ('А если ты не уверен в себе ничего хорошего никогда не получится. Ведь если ты в себя не веришь, кто же поверит?', 'Кто-то умный'),
	('Без идеи не может быть ничего великого! Без великого не может быть ничего прекрасного.', 'Гюстав Флобер'),
	('Быстрее всего учишься в трех случаях — до 7 лет, на тренингах, и когда жизнь загнала тебя в угол.', 'Стивен Кови'),
	('В вашем подсознании скрыта сила, способная перевернуть мир.', 'Уильям Джеймс'),
	('В моем словаре нет слова «невозможно».', 'Наполеон Бонапарт'),
	('Важно верить, что талант нам даётся не просто так – и что любой ценой его нужно для чего-то использовать.', 'Мари Кюри'),
	('Ваше время ограничено, не тратьте его, живя чужой жизнью', 'Стив Джобс'),
	('Велики те, кто видит, что миром правят мысли.', 'Ральф Эмерсон');                            
`)
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
