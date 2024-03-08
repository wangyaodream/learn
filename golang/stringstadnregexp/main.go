package main

import (
    "fmt"
    "strings"
)

func main() {
    
    product := "Kayak"

    fmt.Println("Contains:", strings.Contains(product, "yak"))
    fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
    fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
    fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
    fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
    fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))

    description := "A boat for sailing"

    fmt.Println("Original:", description)
    fmt.Println("Title:", strings.Title(description))
    
    specialChar := "\u01c9"

    fmt.Println("Original:", specialChar, []byte(specialChar))
    
    upperChar := strings.ToUpper(specialChar)
    fmt.Println("Upper:", upperChar, []byte(upperChar))

    isLetterB := func (r rune) bool {
        return r == 'B' || r == 'b'
    } 
    fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))

    splits := strings.Split(description, " ")
    for _, x := range splits {
        fmt.Println("Split >>" + x + "<<")
    }

    splitesAfter := strings.SplitAfter(description, " ")

    for _, x := range splitesAfter {
        fmt.Println("SplitAfter >>" + x + "<<")
    }

    description = "This  is  double  spaced"

    splits = strings.Fields(description)

    for _, x := range splits {
        fmt.Println("Split >>" + x + "<<")
    }

    splitter := func(r rune) bool {
        return r == ' '
    }

    splits = strings.FieldsFunc(description, splitter)
    for _, x := range splits {
        fmt.Println("Field >>" + x + "<<")
    }

    text := "It was a boat. A small boat."

    mapper := func(r rune) rune {
        if r == 'b' {
            return 'c'
        }
        return r
    }

    mapped := strings.Map(mapper, text)
    fmt.Println("Original:", text)
    fmt.Println("Mapped:", mapped)

    replacer := strings.NewReplacer("boat", "kayak", "small", "huge")

    replacesd := replacer.Replace(text)
    fmt.Println("Replaced:", replacesd)


}
