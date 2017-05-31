package htmlhelper_test

import (
	"fmt"

	"github.com/northbright/htmlhelper"
)

func ExampleClean() {
	var str = `
</td>
                            <td 
                                class="tint"
                                
                            >
   
<a
   
  href="" >fake link   

</a>
`

	str = htmlhelper.Clean(str)
	fmt.Printf("cleaned HTML: %v\n", str)

	// Output:
	//cleaned HTML: </td><td class="tint"><a href="">fake link</a>
}

func ExampleTablesToCSVs() {
	var str = `
h1 id="playlist-1">Playlist 1</h1>

<table>
<thead>
<tr>
  <th align="left">Song Name</th>
  <th align="left">Artist</th>
  <th>Album</th>
</tr>
</thead>
<tbody><tr>
  <td align="left">squall</td>
  <td align="left">Pasteboard</td>
  <td>glitter</td>
</tr>
<tr>
  <td align="left">Recollection</td>
  <td align="left">GIN</td>
  <td>MaHaLo</td>
</tr>
</tbody></table>


<h1 id="playlist-2">Playlist 2</h1>

<table>
<tr>
  <td align="left">Song Name</td>
  <td align="left">Artist</td>
  <td>Album</td>
</tr>
<tr>
  <td align="left">I Feel Love</td>
  <td align="left">Mystic Diversions</td>
  <td>Beneath Another Sky</td>
</tr>
<tr>
  <td align="left">紫陽花</td>
  <td align="left">GIN</td>
  <td>MaHaLo</td>
</tr>
</table>
`
	csvs := htmlhelper.TablesToCSVs(str)
	for i, csv := range csvs {
		fmt.Printf("%v: %v\n", i, csv)
	}

	// Output:
	//0: [[Song Name Artist Album] [squall Pasteboard glitter] [Recollection GIN MaHaLo]]
	//1: [[Song Name Artist Album] [I Feel Love Mystic Diversions Beneath Another Sky] [紫陽花 GIN MaHaLo]]
}
