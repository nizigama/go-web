package controller

import (
	"github.com/RobyFerro/go-web/helper"
	"net/http"
)

type HomeController struct {
	BaseController
}

// Main method
// Todo: implement gid and uid to HTML view
func (c *HomeController) Main() {
	c.Response.WriteHeader(http.StatusOK)

	//uid := os.Getuid()
	//gid := os.Getgid()

	//uidString, _ := user.LookupId(strconv.Itoa(uid))
	//gidString, _ := user.LookupGroupId(strconv.Itoa(gid))
	//outString := fmt.Sprintf("You are running as user \"%s\" (%d) with group \"%s\" (%d)", uidString.Name, uid, gidString.Name, gid)

	helper.View("index.html", c.Response, nil)
}