package models

import (
    "log"
    db "../database"
)
type Admin struct {
 Id     int    `json:"id" form:"id"`
 Name   string `json:"name" form:"name"`
 Mail   string `json:"mail" form:"mail"`
}

func GetAllAdmin() (admins []Admin){
    rows, err := db.SqlDB.Query("SELECT id, user_name, user_mail FROM t_admin_user")
    defer rows.Close()
    admins = make([]Admin, 0)
    for rows.Next() {
        var admin Admin
        rows.Scan(&admin.Id, &admin.Name, &admin.Mail)
        admins = append(admins, admin)
    }
    if err = rows.Err(); err != nil {
        log.Fatalln(err)
    }
    return admins
}