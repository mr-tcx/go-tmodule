package test1
import "fmt"
import "github.com/mr-tcx/go-tmodule/internal/test2"
type TestSay struct {
	Name string
}

func (ts *TestSay) TestSayHello() {
	fmt.Printf("%#v\n", *ts)
}

func Say() {
	fmt.Println("this is tmodule-test1")
	test2.Say()
}

func Say2() {
	fmt.Println("this is say2")
}
