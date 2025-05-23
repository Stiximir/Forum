package forum

import (
	"fmt"
	"net/http"
	"time"
)

type Cook struct {
	Cookie string
}

func SetCookie(w http.ResponseWriter, name string, value string) {

	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	//fmt.Println("Cookie " + name + " a été défini avec la valeur: " + cookie.Value)

	http.SetCookie(w, cookie)
}

func GetCookie(r *http.Request, name string) Cook {

	cookieFromRequest, err := r.Cookie(name)
	if err != nil {
		fmt.Println("cookie non trouvé")
		return Cook{Cookie: ""}
	}
	//fmt.Println("Cookie 'userTest' lu avec la valeur: " + cookieFromRequest.Value)

	cook := Cook{
		Cookie: cookieFromRequest.Value,
	}

	return cook
}
