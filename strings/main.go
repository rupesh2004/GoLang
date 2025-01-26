package main

import (
	"fmt"
	"strings"
)

func main() {
	str1:= "rupesh,rupa,reshma"
	fmt.Println(strings.Split(str1,","))
	
	str2 := "one two three two four two"
	cnt:= strings.Count(str2,"two")
	fmt.Println(cnt)

	str3:="rupesh"
	str4:= "bhosale"
	fullString:=strings.Join([]string{str3,str4}," ")
	fmt.Println(fullString)

	fmt.Println("Contains : ",strings.Contains("rupesh","sh"))
	fmt.Println("Prefix: ",strings.HasPrefix("rupesh","r"))
	fmt.Println("Suffix : ",strings.HasSuffix("rupesh","sh"))
	fmt.Println("Index : ",strings.Index("rupesh","p"))
	fmt.Println("Repeate : ",strings.Repeat("rupesh",4))
	fmt.Println("Replace All:",strings.Replace("foo","o","a",-1))
	fmt.Println("Replace : ",strings.Replace("foooooooooooo","o","a",5))
	fmt.Println("To Upper : ",strings.ToUpper("rupesh"))
	fmt.Println("To lower : ",strings.ToLower("RUPESH"))
}
