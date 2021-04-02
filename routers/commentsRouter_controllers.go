package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["blog/controllers:HelloControllers"] = append(beego.GlobalControllerRouter["blog/controllers:HelloControllers"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/hello`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:HelloControllers"] = append(beego.GlobalControllerRouter["blog/controllers:HelloControllers"],
        beego.ControllerComments{
            Method: "JsonFunc",
            Router: `/hello/JsonFunc`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:HelloControllers"] = append(beego.GlobalControllerRouter["blog/controllers:HelloControllers"],
        beego.ControllerComments{
            Method: "PostStruct",
            Router: `/hello/PostStruct`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
