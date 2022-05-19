package jwt

import (
	"database/sql"
	"errors"
)

type Permission string

/*PRIVILEGI: ucta
u: User = Accesso base
c: Creator = Accesso alle api di verifica dati aggiunti da utenti
t: Tester = Accesso alle api di test
a: Admin = tutti gli accessi
*/

const (
	UserPerm    Permission = "u"
	CreatorPerm Permission = "c"
	AdminPerm   Permission = "a"
)

func runeToPermission(r rune) (Permission, error) {
	switch string(r) {
	case UserPerm.ToString():
		return UserPerm, nil
	case CreatorPerm.ToString():
		return CreatorPerm, nil
	case AdminPerm.ToString():
		return AdminPerm, nil
	default:
		return Permission(""), errors.New("Error, permission not setted!")
	}
}

func (p Permission) ToString() string {
	return string(p)
}

// GetPermissionFromDB is a function that gets the user permissions
func GetPermissionFromDB(db *sql.DB, userId int64) (Permission, error) {

	var perm string

	//Get permissions
	q, err := db.Prepare(`SELECT permissions FROM Users WHERE id = ?`)
	if err != nil {
		return "", err
	}

	row, err := q.Query(userId)
	if err != nil {
		return "", err
	}

	for row.Next() {
		row.Scan(&perm)
	}

	return Permission(perm), nil
}

func IsAuthorized(perms []Permission, permsRequire ...Permission) bool {

	myFunc := func(permss []Permission, permReq Permission) bool {
		for _, p := range permss {
			if p == permReq {
				return true
			}
		}
		return false
	}

	for _, preq := range permsRequire {
		if !myFunc(perms, preq) {
			return false
		}
	}

	return true
}
