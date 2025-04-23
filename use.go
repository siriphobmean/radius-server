// package main

// import (
// 	"context"
// 	"errors"
// 	"layeh.com/radius"
// 	"layeh.com/radius/rfc2865"
// 	"time"
// )

// // AuthRADIUS ใช้ยืนยันรหัสผ่านกับ RADIUS server
// func AuthRADIUS(username, password string) (bool, error) {
// 	packet := radius.New(radius.CodeAccessRequest, []byte("radius-secret")) // เปลี่ยน secret ให้ตรงกับฝั่งมหาวิทยาลัย

// 	rfc2865.UserName_SetString(packet, username)
// 	rfc2865.UserPassword_SetString(packet, password)

// 	// ระบุ IP และพอร์ตของ RADIUS server ที่มหาวิทยาลัยใช้ เช่น "radius.uni.ac.th:1812"
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	response, err := radius.Exchange(ctx, packet, "radius.uni.ac.th:1812")
// 	if err != nil {
// 		return false, err
// 	}

// 	switch response.Code {
// 	case radius.CodeAccessAccept:
// 		return true, nil
// 	case radius.CodeAccessReject:
// 		return false, errors.New("username or password is incorrect")
// 	default:
// 		return false, errors.New("unexpected response from RADIUS server")
// 	}
// }

package main

import (
	"context"
	"errors"
	"fmt"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"time"
)

func AuthRADIUS(username, password string) (bool, error) {
	packet := radius.New(radius.CodeAccessRequest, []byte("radius-secret"))

	rfc2865.UserName_SetString(packet, username)
	rfc2865.UserPassword_SetString(packet, password)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := radius.Exchange(ctx, packet, "radius.test.ac.th:1812")
	if err != nil {
		return false, err
	}

	switch response.Code {
	case radius.CodeAccessAccept:
		return true, nil
	case radius.CodeAccessReject:
		return false, errors.New("username or password is incorrect")
	default:
		return false, errors.New("unexpected response")
	}
}

func main() {
	ok, err := AuthRADIUS("tim", "12345")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if ok {
		fmt.Println("✅ Login success")
	} else {
		fmt.Println("❌ Login failed")
	}
}