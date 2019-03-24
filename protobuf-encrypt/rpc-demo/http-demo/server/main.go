
package main
 
import (
    "net/http"
    "log"
    "net/rpc"
)
 
type Params struct {
    Width, Height int;
}
 
type Rect struct{}
 
// 1. 方法名Area必须大写
// 2. 必须有两个参数，且参数必须是外部能访问的类型或内置类型，且第二个参数必须是指针类型
// 3. 返回值必须是error类型
func (r *Rect) Area(p Params, ret *int) error {
    *ret = p.Width * p.Height
    return nil
}
 
 
func main() {
    rect := new(Rect)
    rpc.Register(rect)  // 注册
 
 
    rpc.HandleHTTP();
    err := http.ListenAndServe(":8080", nil);
    if err != nil {
        log.Fatal(err);
    }
}
