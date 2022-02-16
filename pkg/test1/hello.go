package test1
import "fmt"
type TestSay struct {
	Name string
}

func (ts *TestSay) TestSayHello() {
	fmt.Printf("%#v\n", *ts)
}
