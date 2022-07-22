//)Take a Paragraphs of text Max word  up to 200 words in a Slice or String and then Pass each  Sentence(up to full stop) to another go routine/routines   1)print it in reverse order(1 goroutine) 2 )count its words(2 goroutine)

package sentence 

import (
	"fmt"
)

go func sentence(a string) {
	fmt.Println(a)
}
