package models

import (
	"database/sql"
	"fmt"
	"ikct-ed/config"
	"log"
)

type RBAC struct {
	Role2permission Role2permission
	Role            Role
	Permission      Permission
}

// Role data structure representing a role
type Role struct {
	ID       int    
	Role     string 
	Assigned bool   
}

//Permission  data structure that represents a permission
type Permission struct {
	ID          int
	Permission  string
	DisplayName string
	Assigned    bool 
	Create      bool 
	Read        bool 
	Update      bool 
	Delete      bool 
}

// Role2permission data structure that represents the relationship between
// role and permissions, from the role2permission table
type Role2permission struct {
	ID           int
	RoleID       int
	PermissionID int  `json:"p_id"`
	Create       bool `json:"create"`
	Read         bool `json:"read"`
	Update       bool `json:"update"`
	Delete       bool `json:"delete"`
}

func GetRoleIdByUserId(userId int) (int, error) {
	var roleId int
	db, err := config.GetDB2()
	if err != nil {
		return roleId, err
	}
	defer db.Close()
	err = db.QueryRow("SELECT role_id FROM user2role WHERE user_id = $1", userId).Scan(&roleId)
	return roleId, nil
}

func GetRoleByRoleId(roleId int) (string, error) {
	db, err := config.GetDB2()
	if err != nil {
		return "", err
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT role from role where id=$1")
	rows, err := db.Query(query, roleId)
	if err != nil {
		return "", err
	}
	var role string
	for rows.Next() {
		err = rows.Scan(&role)
		if err != nil {
			return "", err
		}
	}
	return role, nil
}

func GetPermissionId(name string) (int, error) {
	db, err := config.GetDB2()
	if err != nil {
		return 0, err
	}
	var permissionId int
	defer db.Close()
	query := `SELECT id from permission where permission=$1`
	err = db.QueryRow(query, name).Scan(&permissionId)
	if err != nil {
		log.Println("GetPermissionId: failed to scan row with error:", err)
		return 0, err
	}
	return permissionId, nil
}

// AuthorizationOfRoles2Permission is used check authorization of roles corresponding to their permission/s and
// blocking their rights
// Input = roleId and permissionID
// Output = RBAC(role based access control) array  and ERROR

func AuthorizationOfRoles2Permission(roleId, permissionID int) (RBAC, error) {
	var AuthorizationOfRoles2Permissionlist RBAC
	query := `SELECT
	              r.id as role_id,rp.create,rp.read,rp.update,rp.delete, p.id as permission_id
	          FROM 
			      user2role ur 
			  JOIN	
			      role r 
			   ON 
			      ur.role_id=r.id 
			  JOIN 
			      role2permission rp
		       ON 
			      ur.role_id=rp.role_id 
		      JOIN
		          permission p 
			   ON 
			      rp.permission_id=p.id 
		      WHERE
		          ur.role_id = $1
		      AND
		          rp.permission_id = $2`

	db, err := config.GetDB2()
	if err != nil {
		log.Println("AuthorizationOfRoles2Permission is failed while connecting to DB")
		return RBAC{}, err
	}
	defer db.Close()
	row := db.QueryRow(query, roleId, permissionID)
	if err != nil {
		log.Println("AuthorizationOfRoles2Permission failed while querying with :", err)
	}
	var role_id sql.NullInt64
	var permission_id sql.NullInt64
	var create sql.NullBool
	var delete sql.NullBool
	var update sql.NullBool
	var read sql.NullBool

	err = row.Scan(&role_id, &create, &read, &update, &delete, &permission_id)
	if err != nil {
		log.Println("AuthorizationOfRoles2Permission failed while Scanning :", err)
	}
	AuthorizationOfRoles2Permissionlist = RBAC{
		Role2permission: Role2permission{
			Create: create.Bool,
			Delete: delete.Bool,
			Update: update.Bool,
			Read:   read.Bool,
		},
		Role: Role{
			ID: int(role_id.Int64),
		},
		Permission: Permission{
			ID: int(permission_id.Int64),
		},
	}
	return AuthorizationOfRoles2Permissionlist, err
}
