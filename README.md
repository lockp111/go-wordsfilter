# go-wordsfillter

## Usage
```golang
package main

import (
	"fmt"
	"github.com/lockp111/go-wordsfillter"
)

func main() {
    trie := wordsfillter.NewTrie()
    trie.LoadWordDict("path/to/dict")
	filter := wordsfillter.New(trie)
	// do something
}
```

## Trie AddWords
```golang
trie.AddWords("dirtyword")
trie.AddWords("shit","bullshit")
```

## Trie LoadNetWordDict
```golang
trie.LoadNetWordDict("http://xxxx.com/dict")
```

## Trie Show
```golang
trie.Show() // will show all nodes by tree
```

## Filter Filter
remove words
```golang
newText := filter.Filter("you bullshit")
// output => you
```

## Filter Replace
replace words
```golang
newText := filter.Replace("you bullshit", '*')
// output => you ********
```

## Filter FindIn
find and return first word
```golang
newText := filter.FindIn("you bullshit")
// output => true, bullshit
```

## Filter Validate
validate and return first word
```golang
newText := filter.Validate("you bullshit")
// output => false, bullshit
```

## Filter FindAll
find and return all words
```golang
newText := filter.FindAll("you bullshit")
// output => [bullshit]
```

## Filter UpdateNoisePattern
set and update noise word
```golang
// failed
filter.FindIn("you bull-shit")      // false
filter.UpdateNoisePattern(`-`)
// success
filter.FindIn("you bull-shit")      // true, bullshit
```

## Filter UpdateTrie
update trie
```golang
trie := wordsfillter.NewTrie()
trie.LoadWordDict("path/to/newDict")
filter.UpdateTrie(trie)
```

## Filter RemoveNoise
remove noise word
```golang
filter.UpdateNoisePattern(`-`)
filter.RemoveNoise("you bull-shit")
// output => you bullshit
```