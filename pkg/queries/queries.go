package queries

import (
	sqlutils "github.com/dopaminegirl19/usership/pkg/sqlutils"
	structs "github.com/dopaminegirl19/usership/pkg/structs"

	"database/sql"
	"fmt"
)

var db *sql.DB

// UpdateUser updates the name and/or contact info of a user
func UpdateUser(u structs.User) (int64, error) {
	result, err := db.Exec("UPDATE users SET name = ?, contact = ?, contact2 = ? WHERE id = ?", u.Name, u.Contact, sqlutils.NewNullString(u.Contact2), u.ID)
	if err != nil {
		return 0, fmt.Errorf("UpdateUser: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateUser: %v", err)
	}
	return count, nil
}

// UpdateItem updates the name and/or description of an item
func UpdateItem(i structs.Item) (int64, error) {
	result, err := db.Exec("UPDATE items SET name = ?, description = ?, WHERE id = ?", i.Name, sqlutils.NewNullString(i.Description), i.ID)
	if err != nil {
		return 0, fmt.Errorf("UpdateItem: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateItem: %v", err)
	}
	return count, nil
}

// UpdateItemCurrentUser updates the current_user_id in an item's row
func UpdateItemCurrentUser(itemID int32, newUserID int32) (int64, error) {
	result, err := db.Exec("UPDATE items SET current_user_id = ? WHERE id = ?", newUserID, itemID)
	if err != nil {
		return 0, fmt.Errorf("UpdateItemCurrentUser: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateItemCurrentUser: %v", err)
	}
	return count, nil
}

// DeleteUser deletes a row from the users table
func DeleteUser(userID int32) (int64, error) {

	result, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return 0, fmt.Errorf("DeleteUser: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteUser: %v", err)
	}
	return count, nil
}

// DeleteItem deletes a row from the items table
func DeleteItem(itemID int32) (int64, error) {

	result, err := db.Exec("DELETE FROM items WHERE id = ?", itemID)
	if err != nil {
		return 0, fmt.Errorf("DeleteItem: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteItem: %v", err)
	}
	return count, nil
}

// AddUser adds a row to the users table
func AddUser(u structs.User) (int32, error) {
	result, err := db.Exec("INSERT INTO users (name, contact, contact2) VALUES (?, ?, ?)", u.Name, u.Contact, sqlutils.NewNullString(u.Contact2))
	if err != nil {
		return 0, fmt.Errorf("AddUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddUser: %v", err)
	}
	return int32(id), nil
}

// AddItem adds a row to the items table
func AddItem(i structs.Item) (int32, error) {
	result, err := db.Exec("INSERT INTO items (name, description, current_user_id) VALUES (?, ?, ?)", i.Name, sqlutils.NewNullString(i.Description), i.CurrentUserID)
	if err != nil {
		return 0, fmt.Errorf("AddItem: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddItem: %v", err)
	}
	return int32(id), nil
}

// UserByItemID queries for the user that's currently in possession of an item.
func UserByItemID(queryItemID int32) (structs.User, error) {
	var u structs.User
	// find the item corresponding to that item id
	item, err := ItemByID(queryItemID)
	if err != nil {
		return u, fmt.Errorf("UserByItemID %q: %v", queryItemID, err)
	}
	// find the user corresponding to item's current user id
	user, err := UserByID(item.CurrentUserID)
	if err != nil {
		return u, fmt.Errorf("UserByItemID %q: %v", queryItemID, err)
	}

	return user, nil
}

// ItemsByString queries for items which match a string either in name or description.
func ItemsByString(queryString string) ([]structs.Item, error) {
	var items []structs.Item

	rows, err := db.Query("SELECT * FROM items WHERE name LIKE '%?%' OR description LIKE '%?%'", queryString, queryString)
	if err != nil {
		return nil, fmt.Errorf("ItemsByString %q: %v", queryString, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var i structs.Item
		if err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.CurrentUserID); err != nil {
			return nil, fmt.Errorf("ItemsByUser %q: %v", queryString, err)
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ItemsByUser %q: %v", queryString, err)
	}
	return items, nil
}

// UserByID queries for the user with a specific ID
func UserByID(queryUserid int32) (structs.User, error) {
	// A user to hold data from the returned row.
	var u structs.User

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", queryUserid)
	if err := row.Scan(&u.ID, &u.Name, &u.Contact, &u.Contact2); err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("UserByID %d: no such user", queryUserid)
		}
		return u, fmt.Errorf("UserByID %d: %v", queryUserid, err)
	}
	return u, nil
}

// ItemByID queries for the user with a specific ID
func ItemByID(queryItemID int32) (structs.Item, error) {
	// An item to hold data from the returned row.
	var i structs.Item

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", queryItemID)
	if err := row.Scan(&i.ID, &i.Name, &i.Description, &i.CurrentUserID); err != nil {
		if err == sql.ErrNoRows {
			return i, fmt.Errorf("UserByID %d: no such user", queryItemID)
		}
		return i, fmt.Errorf("UserByID %d: %v", queryItemID, err)
	}
	return i, nil
}

// ItemsByUser queries for items that are in the possession of a user (by user ID)
func ItemsByUser(queryUserid int32) ([]structs.Item, error) {
	// an items slice to hold data from returned rows.
	var items []structs.Item

	rows, err := db.Query("SELECT * FROM items WHERE current_user_id = ?", queryUserid)
	if err != nil {
		return nil, fmt.Errorf("ItemsByUser %q: %v", queryUserid, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var i structs.Item
		if err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.CurrentUserID); err != nil {
			return nil, fmt.Errorf("ItemsByUser %q: %v", queryUserid, err)
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ItemsByUser %q: %v", queryUserid, err)
	}
	return items, nil
}

// UsersByName queries for users that have the specified name.
func UsersByName(queryName string) ([]structs.User, error) {
	// A users slice to hold data from returned rows.
	var users []structs.User

	rows, err := db.Query("SELECT * FROM users WHERE name = ?", queryName)
	if err != nil {
		return nil, fmt.Errorf("UsersByName %q: %v", queryName, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var u structs.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Contact); err != nil {
			return nil, fmt.Errorf("UsersByName %q: %v", queryName, err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("usersByName %q: %v", queryName, err)
	}
	return users, nil
}
