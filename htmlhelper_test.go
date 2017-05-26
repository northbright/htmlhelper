package htmlhelper_test

import (
	"log"

	"github.com/northbright/htmlhelper"
)

func ExampleClean() {
	str = htmlhelper.Clean(str)
	log.Printf("cleaned HTML: %v\n", str)

	// Output:
}

var str = `
</td>
                            <td 
                                class="tint"
                                
                            >
   
<a
   
  href="" >fake link   

</a>
`
