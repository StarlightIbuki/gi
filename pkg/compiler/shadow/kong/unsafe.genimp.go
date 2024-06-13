package shadow_kong

import "fmt"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
}

 func InitLua() string {
  return `
__type__.kong = {};
__type__.kong.Response = {
  SetHeader = {}
};
__type__.kong.Request = {
  GetHeader = {}
};
`}