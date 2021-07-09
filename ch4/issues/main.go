package main

import (
	"fmt"
	"go-study/ch4/github"
	"log"
	"os"
)

func main()  {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

/**
# zhangquan @ MacBook-Pro in ~/goworkspace/src [14:38:25]
$ go build go-study/ch4/issues


# zhangquan @ MacBook-Pro in ~/goworkspace/src [14:40:09]
$ ./issues repo:golang/go is:open json decoder
61 issues:
#33416   bserdar encoding/json: This CL adds Decoder.InternKeys
#45628 pgundlach encoding/xml: add Decoder.InputPos
#43716 ggaaooppe encoding/json: increment byte counter when using decode
#42571     dsnet encoding/json: clarify Decoder.InputOffset semantics
#46065   vipally proposal: encoding/json: add FlexObject for encoding/de
#11046     kurin encoding/json: Decoder internally buffers full input
#32779       rsc encoding/json: memoize strings during decode
#5901        rsc encoding/json: allow per-Encoder/per-Decoder registrati
#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
#36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
#14750 cyberphon encoding/json: parser ignores the case of member names
#43401  opennota proposal: encoding/csv: add Reader.InputOffset method
#29035    jaswdr proposal: encoding/json: add error var to compare  the
#31701    lr1980 encoding/json: second decode after error impossible
#40128  rogpeppe proposal: encoding/json: garbage-free reading of tokens
#40127  rogpeppe encoding/json: add Encoder.EncodeToken method
#45512 colin-sit encoding/json: cannot unmarshal custom interface value
#40982   Segflow encoding/json: use different error type for unknown fie
#40983   Segflow encoding/json: return a different error type for unknow
#28923     mvdan encoding/json: speed up the decoding scanner
#43513 Alexander encoding/json: add line number to SyntaxError
#41144 alvaroale encoding/json: Unmarshaler breaks DisallowUnknownFields
#16212 josharian encoding/json: do all reflect work before decoding
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#29750  jacoelho cmd/vet: stdmethods check gets confused if run on a pac
#34564  mdempsky go/internal/gcimporter: single source of truth for deco
#33835     Qhesz encoding/json: unmarshalling null into non-nullable gol
#33854     Qhesz encoding/json: unmarshal option to treat omitted fields
#30301     zelch encoding/xml: option to treat unknown fields as an erro
#26946    deuill encoding/json: clarify what happens when unmarshaling i



*/