package model

import "log"

type Note struct {
	ID int
	Content string
	Time string
}

func ListNotes() ([]Note, error) {
	var ns []Note
	rows, err := db.Query("select * from notes;")
	if err != nil {
		log.Printf("query db failed: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var n Note
		err = rows.Scan(&n.ID, &n.Content, &n.Time)
		if err != nil {
			log.Printf("scan row failed: %v\n", err)
			return nil, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}