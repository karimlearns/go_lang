package main

// "math/rand"
// "fmt"
// "os"
// "log"
// "time"
// "strings"
// "flag"
// "strconv"
// "unicode"
// "io"
// "strings"
// "time"
// "bufio"

//-------------------------------------------------------------------------------------------------------------------------
// STRINGS LIBRARY #PART 1 																		video 19				386
// STRINGS LIBRARY #PART 2 																		video 20				404
// STRINGS LIBRARY #PART 3 																		video 21				423
// STRINGS LIBRARY #PART 4 																		video 22				447
// STRINGS LIBRARY #PART 5 																		video 23				461
// STRINGS LIBRARY #PART 6 																		video 24				493
// STRINGS LIBRARY #PART 7 																		video 25				511
// STRINGS LIBRARY #PART 8 																		video 26				531
// Arrays				   																		video 27				574
// Slices				   																		video 28				605
// Maps					   																		video 29				644
// STructues			   																		video 30				677
// Switchers via flag library 																	video 31				703
// Assignement Operators 																		video 32				730
// DATA Types Converting																		video 33				752
// if / else if / else																			video 34				853
// Affirmation & Denial in "if Conditions"														video 35				890
// Switches Cases																				video 36				919
// Logical Operators "|| = or" "&& = and"														video 37				945
// Logical Operators More than a value as a result 												video 38				966
// While Loop																					video 39				972
// For Loop																						video 40				988
// Advanced For Loop																			video 41				1002
// Continue																						video 42				1023
// Break																						video 43				1042
// Looping into file line																		video 44				1060
// What does if err != nil means?																video 45				1083
// Functions																					video 46				1089
// Functions (Values)																			video 47				1113
// Functions (return) Pt, 1																		video 48				1130
// Functions (return) Pt, 2																		video 49				1148
// Panic																						video 50				1169
// Defer																						video 51				1189
// Recover																						video 52				1223
// Pointers																						video 53				1248
// os (operation system) library introduction 													video 54				1270
// os (operation system) library File Handing Checking file exist & Reading file 				video 55				1280
// os (operation system) library File Handing Creating Files & Writing Into Files 				video 56				1325
// os (operation system) library os.Exit()														video 57				1349
// os (operation system) library																video 58				1368
// os. library																					video 59				1389
// os library os.Rename() os.Remove()															video 60				1451
// os library os.Get // Set / Unset env()														video 61				1466
//---------------------------------------------------------------------------------------------------------------------------
// func main()  {
// 	MyLang := "golang"
// 	// fmt.Printf("Hellow World From %v !!", MyLang)
// 	MyLang2 := MyLang+"2"
// 	fmt.Printf("Hellow World From %v !!", MyLang2)
// 	time.Sleep(10 * time.Second)
// }

// https://pkg.go.dev/std@go1.20.2

// ------------------------------------------------

// DATA TYPES
// func main()  {
// // boolean
// myBool := false
// fmt.Printf("%v %n\n\n", myBool, myBool)

// // integer
// myInt := 2 * 34
// fmt.Printf("%v %T\n\n", myInt, myInt)

// // String

// myStr := "my" + " str"
// fmt.Printf("%v %T", myStr, myStr)

// // Floating
// myFloat := 34.56789
// fmt.Printf("%v %T", myFloat, myFloat)

// // []String = Array Slice
// var myString = []string{"a", "b"}
// fmt.Printf("%v %T", myString, myString)
// }
// -----------------------------
// Variables / Constants
// func main() {
// const myCOn = "ALGO"
// A := "fruit"
// fmt.Println(A)
// A = "legume"
// fmt.Println(A)

// var myVar string = "This is my variable"
// fmt.Println(myVar)

// var myVar2 string  // ndeclarih wmnba3d nasta3amlo
// myVar2 = "This is my second variable"
// fmt.Println(myVar2)

// myVar3 := "This is my third variable"  // min naktab variable bhad tarii9a mama7tajch n7adad naw3o
// fmt.Println(myVar3)

// Ziggo, Miggo := "Z", "M"
// fmt.Println(Ziggo, Miggo)

// var B, M = "Bob", "Marli"
// fmt.Println(B, M)

// var a, b = 1, 2
// fmt.Println(a, b)

// myNumber := 1
// fmt.Println(myNumber)

// var myFloat = 1.1
// fmt.Println(myFloat)

// var myBool1 bool = true
// fmt.Println(myBool1)

// myBool2 := false
// fmt.Println(myBool2)

// mySlice := []string{"A", "B"}
// fmt.Println(mySlice)

// var mySlice2 = []string{"C", "D"}
// fmt.Println(mySlice2)
// }
// ------------------------------
// Constant
// func main()  {
// Variable => var
// Constant => const

// const MyConst = "Btats" // ila bghit nbadar const l variable const nradha var
// fmt.Println(MyConst)
// // MyConst = "Batata"
// }
// ----------------------------
// Espace Sequences Chars
// func main()  {

// // \b = Back Space
// fmt.Println("Hello \bWorld")

// fmt.Println("")

// // \n = New Line
// fmt.Println("Hello\nWorld")
// fmt.Println(`Hello
// World`)

// fmt.Println("")

// // \r = Return
// fmt.Println("AAAAAAAAA\rFour")

// fmt.Println("")

// // \t = Tab
// fmt.Println("\tHello After Tap !!")

// fmt.Println("")

// // 1. naktab back slash
// fmt.Println("Hello \\ World")
// // 2. naktab back slash
// fmt.Println(`Hello \ World`)

// // 1. naktab double quotes
// fmt.Println("Hello \" World")
// // 2. naktab double quotes
// fmt.Println(`Hello " World`)
// }
// ---------------------------------------------------
// USER INPUT
// func main() {

// 	fmt.Printf("Enter your username: ")

// 	var username string
// 	fmt.Scanln(&username)

// 	fmt.Printf("\n\nThis is my name: %v" , username)

// }

// ------------------------------------------------------
// NUMBERS ARITHMETIC OPERATORS
// func main() {

// +
// -
// *
// /
// %

// // +
// A := 1.5
// A = A + 1
// final := 2 + A + 4 +4
// fmt.Printf("%T, %v\n\n", final, final)

// // -
// R := 55
// X := 100 - R - R - 15
// fmt.Printf("%T, %v\n\n", X, X)

// // / *
// Q := 50
// O := Q / 2
// fmt.Printf("%T, %v\n\n", O, O)

// // % Modules
// fmt.Println(8 % 2)
// fmt.Println(9 % 2)
// }

// -------------------------------
// RELATIONAL OPERATORS

// func main() {
// A := 6
// B := 5
// fmt.Println(A <= B)
// }

// -------------------------------
// STRINGS CONCATENATION
// func main() {
// 	A := "Hello"
// 	B := "World"
// 	C := "!!"
// 	fmt.Println(A + " " + B + " " + C)
// 	fmt.Println(A, B, C)
// 	fmt.Printf("%v %v %v", A, B, C)
// }

// -------------------------------
// FMT LIBRARY #PART 1

// func main() {
// 	A := "ABC"
// 	print("ABC2", A)
// 	print("ABC3")

// 	println("\n\n")

// 	println("DEF")
// 	println("DEF2")
// 	println("DEF3\n\n")

// 	fmt.Printf("%v\n", A) // %v == value
// 	fmt.Printf("%T\n", A) // %T == Type
// 	fmt.Printf("%q\n", A) // %q == drint the value enter double quotes
// 	fmt.Printf("%s\n", A) // %s == yatba3 string
// 	fmt.Printf("%x\n", A) // %x == yatba3 ramz nta3 variable lower-case
// 	fmt.Printf("%X\n", A) // %X == yatba3 ramz nta3 variable upper-case
// }

//////////////////////
// Exercice: fmt 1 //
////////////////////

// func main()  {
// 	A := "kareeem"
// 	B := "hassan"
// 	// print(A)
// 	println(A, B)
// 	fmt.Printf("%v\n", B)
// 	fmt.Printf("%T\n", B)
// 	fmt.Printf("%q\n", B)
// 	fmt.Printf("%s\n", B)
// 	fmt.Printf("%x\n", B)
// 	fmt.Printf("%X\n", B)
// }

// -------------------------------
// FMT LIBRARY #PART 2

// func main() {

// 	K := 20

// Integer:
// fmt.Printf("%b\n", K) // systeme binaire 0/1
// fmt.Printf("%c\n", K)
// fmt.Printf("%d\n", K) // systeme decimale 10
// fmt.Printf("%o\n", K) // systeme octal 8
// fmt.Printf("%O\n", K) // base 8 with 0o prefix
// fmt.Printf("%q\n", K) // "\b"
// fmt.Printf("%x\n", K) // 8 -- 0/1/2/3/4/5/6/7/8/9/a/b/c/d/e/f --
// fmt.Printf("%X\n", K) // 8 -- 0/1/2/3/4/5/6/7/8/9/A/B/C/D/E/F --
// fmt.Printf("%U\n", K) // U+0008

// Pointer:
// fmt.Printf("%p\n", &K) // ramz takhzin f ram // yakhdam 7ata m3a %b, %d, %o, %x and %X

// The default format for %v is:
////////////////////////////////
// bool:                    %t
// int, int8 etc.:          %d
// uint, uint8 etc.:        %d, %#x if printed with %#v
// float32, complex64, etc: %g
// string:                  %s
// chan:                    %p
// pointer:                 %p

// }

//////////////////////
// Exercice: fmt 2 //
////////////////////

// func main()  {
// 	K := 23
// 	fmt.Printf("%b\n", K)
// 	fmt.Printf("%c\n", K)
// 	fmt.Printf("%d\n", K)
// 	fmt.Printf("%o\n", K)
// 	fmt.Printf("%O\n", K)
// 	fmt.Printf("%x\n", K)
// 	fmt.Printf("%X\n", K)
// 	fmt.Printf("%U\n", K)
// }

// -------------------------------
// FMT LIBRARY #PART 3

// func main() {

// D := "QQQQQQ"
// fmt.Printf("%v %s\n\n\n", D, D)

// W := "R1"
// WW := "R2"
// fmt.Println(WW, W)

// 	fmt.Printf("Enter Your Name: ")
// 	var name string
// 	fmt.Scanf("%v", &name) // == fmt.Scanln(&name)
// 	fmt.Println(name)

// 	var namee, agee = "tarek", 22
// 	s := fmt.Sprintf("My Name is %s and my Age is %d\n\n", namee, agee)
// 	io.WriteString(os.Stdout, s)
// }

//////////////////////
// Exercice: fmt 3 //
////////////////////

// func main()  {

// 	H := "Hello"
// 	fmt.Printf("%v %s\n", H, H)

// 	W := "World"
// 	WW := "World2"
// 	fmt.Println(W, WW)

// 	fmt.Printf("Enter Your Name: ")
// 	var name string
// 	fmt.Scanf("%v", &name)
// 	fmt.Println(name)
// }

// -------------------------------
// VIDEO 19
// STRINGS LIBRARY #PART 1

// func main() {
// A := "string11gg2"
// B := "srting2"
// fmt.Println(strings.Compare(A, B))
// fmt.Println(strings.Contains(A, "g"))
// fmt.Println(strings.ContainsAny(A, "1gg34567890fhhd"))
// fmt.Println(strings.Count(A, "g"))
// fmt.Println(strings.Cut(A, "11"))
// bf, af, fo := strings.Cut(A, "11")
// fmt.Println(bf)
// fmt.Println(af)
// fmt.Println(fo)
// }

// -------------------------------

// VIDEO 20
// STRINGS LIBRARY #PART 2

// func main() {
// fmt.Println(strings.EqualFold("GO", "go"))
// A := strings.Fields("myelm1   myele2   myele3")
// fmt.Println(A[1])

// 	f := func(c rune) bool {
// 		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
// 	}

// 	A := strings.FieldsFunc("ele1, ele2 ... ele3-ele4", f)
// 	fmt.Println(A)
// }

// -------------------------------

// VIDEO 21
// STRINGS LIBRARY #PART 3

// func main() {
// 	root13 := func (r rune) rune {
// 		switch {
// 			case r >= 'A' && r <= 'Z':
// 				return 'A' + (r - 'A' + 14)%26
// 			case r >= 'a' && r <= 'z':
// 				return 'a' + (r - 'a'+14)%26
// 		}
// 		return r
// 	}
// 	fmt.Println(strings.Map(root13, "'Twas brillig and the slithy gopher...")) // 'Hkog pfwzzwu obr hvs gzwhvm ucdvsf...

// mySlice := []string{"one", "two", "three"}
// fmt.Println(strings.Join(mySlice, "-")) // one-two-three
// fmt.Printf("%T", mySlice) // []string

// 	fmt.Println("ba" + strings.Repeat("na", 2)) // banana
// }

// -------------------------------

// VIDEO 22
// STRINGS LIBRARY #PART 4
// func main() {

// 	fmt.Println(strings.Index("chiken", "k")) // 3
// 	fmt.Println(strings.LastIndex("chikenk", "k")) // 6

// 	fmt.Println(strings.IndexAny("chiikenk", "k")) // 4
// 	fmt.Println(strings.LastIndexAny("chiikenk", "k")) // 7

// }

// -------------------------------

// VIDEO 23
// STRINGS LIBRARY #PART 5
// func main() {
// 	// IndexFunc
// 	f := func(c rune) bool {
// 		return unicode.Is(unicode.Han, c)
// 	}
// 	fmt.Println(strings.IndexFunc("Hello, 世界", f))
// 	fmt.Println(strings.IndexFunc("Hello, world", f))

// 	// LastIndexFunc
// 	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
// 	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
// 	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))

// 	// IndexByte
// 	fmt.Println(strings.IndexByte("golang", 'g'))
// 	fmt.Println(strings.IndexByte("gophers", 'h'))
// 	fmt.Println(strings.IndexByte("golang", 'x'))

// 	// LastIndexByte
// 	fmt.Println(strings.LastIndexByte("Hello, world", 'l'))
// 	fmt.Println(strings.LastIndexByte("Hello, world", 'o'))
// 	fmt.Println(strings.LastIndexByte("Hello, world", 'x'))

// 	// IndexRune
// 	fmt.Println(strings.IndexRune("chicken", 'k'))
// 	fmt.Println(strings.IndexRune("chicken", 'd'))
// }

// -------------------------------

// VIDEO 24
// STRINGS LIBRARY #PART 6

// func main() {
// // Replace And ReplaceAll
// fmt.Println(strings.Replace("Thismy is my text", "my", "your", 1)) // Thisyour is my text
// fmt.Println(strings.ReplaceAll("Thismy is my text", "my", "your")) // Thisyour is your text

// // Split And SplitN
// fmt.Println(strings.Split("This,is,a,string", ",")) // [This is a string])
// fmt.Println(strings.SplitN("This,is,a,string", ",", 2)) // [This is,a,string]

// fmt.Println(strings.SplitAfter("A,aB,aC", ",")) // [A, aB, aC]
// fmt.Println(strings.SplitAfterN("A,B,C,D,E", ",", 4)) // [A, B, C, D,E]
// }

// -------------------------------

// VIDEO 25
// STRINGS LIBRARY #PART 7

// func main() {
// 	// Title, ToTitle And ToTitleSpecial
// 	fmt.Println(strings.Title("hello, world")) // Hello, World
// 	fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir")) // DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
// 	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir")) // DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR

// 	// ToLower And ToLowerSpecial
// 	fmt.Println(strings.ToLower("HELLO, WORLD")) // hello, world
// 	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş")) // önnek iş

// 	// ToUpper And ToUpperSpecial
// 	fmt.Println(strings.ToUpper("hello, world")) // HELLO, WORLD
// 	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "Önnek İş")) // ÖNNEK İŞ
// }

// -------------------------------

// VIDEO 26
// STRINGS LIBRARY #PART 8

// func main() {
// // Trim
// fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers
// // TrimFunc
// fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
// 	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
// })) // Hello, Gophers

// // TrimLeft And TrimRight
// fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers!!!
// fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡")) // ¡¡¡Hello, Gophers

// // TrimLeftFunc
// fmt.Println(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
// 	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
// })) // Hello, Gophers!!!

// // TrimLeftFunc
// fmt.Println(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
// 	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
// })) // ¡¡¡Hello, Gophers

// TrimPrefix
// var s = "¡¡¡Hello, Gophers!!!"
// s = strings.TrimPrefix(s, "¡¡¡Hello, ")
// s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
// fmt.Print(s) // Gophers!!!

// TrimSuffix
// var s = "¡¡¡Hello, Gophers!!!"
// s = strings.TrimSuffix(s, ", Gophers!!!")
// s = strings.TrimSuffix(s, ", Marmots!!!")
// fmt.Print(s) // ¡¡¡Hello

// TrimSpace
// fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) // Hello, Gophers
// }

// -------------------------------

// VIDEO 27
// Arrays

// func main() {
// var arr1 [3]string = [3]string{"A", "B", "C"}
// var arr1  = [3]string{"A", "B", "C"}
// arr1 := [3]string{"A", "B", "C"}
// fmt.Println(arr1)

// DELETE ELEMENTS
// elm_to_del := "C"
// my_new_arr := []string{}

// for _, x := range arr1 {
//     if x == elm_to_del {
//         fmt.Printf("")
//     } else {
// 		my_new_arr = append(my_new_arr, x)
// 	}
// }

// fmt.Println(my_new_arr)
// fmt.Println(arr1)

// arr1[1] = "D"
// fmt.Println(arr1)
// }

// -------------------------------

// VIDEO 28

// Slices

// func main()  {
// 	slice1 := []string{"foo","bar"}
// var slice1 = []string{"foo", "bar"}
// var slice1 []string = []string{"foo", "bar"}
// fmt.Println(slice1)

// ADD ELEMENTS
// slice1 = append(slice1, "too")
// fmt.Println(slice1)

// DELETE ELEMENTS
// slice1[1] = ""
// fmt.Println(slice1)

// elm_to_del := "bar"
// my_new_arr := []string{}

// for _, x := range slice1 {
//     if x == elm_to_del {
//         fmt.Printf("")
//     } else {
// 		my_new_arr = append(my_new_arr, x)
// 	}
// }
// fmt.Println(my_new_arr)
// fmt.Println(slice1)

// CHANGE ELEMENTS
// slice1[0] = "first"
// fmt.Println(slice1)

// }

// -------------------------------

// VIDEO 29

// MAPS

// func main()  {
// 	mymap := map[string]string {
// 		"Cat":    "1",
// 		"Dog":    "2",
// 		"Botato":  "3",
// 	}
// fmt.Println(mymap) // map[Botato:3 Cat:1 Dog:2]

// mymap["Dog"] = "40"
// fmt.Println(mymap) // [Botato:3 Cat:1 Dog:40]

// var mappy = make(map[string]string)
// mappy["Cat"] = "40"
// mappy["Dog"] = "20"
// mappy["Botato"] = "5"

// delete(mappy, "Cat")
// fmt.Println(mappy) // map[Botato:5 Dog:20]

// _, avl := mymap["Cat"]
// fmt.Println("is this element available? :" , avl) // true

// for elm, val := range mymap {
// 	fmt.Println("element: ", elm, "Has value:", val) // element:  Cat Has value: 1
// }
// }

// -------------------------------

// VIDEO 30

// Structures

// type Cars struct {
// 	name string
// 	color string
// 	model int
// 	price int
// }

// func main()  {
// var CarA Cars = Cars{"Peogeot", "Black", 2021, 60000}
// CarA := Cars{name: "Tesla", color: "Black", model: 2021, price: 60000}
// fmt.Println(CarA) // {Peogeot Black 2021 60000}
// fmt.Println(CarA.name) // Peogeot

// // & = code takhzin f la ram
// CarA := &Cars{name: "Tesla", color: "Black", model: 2021, price: 60000}

// // * : ta3ti value nta3 code takhzin f la ram
// fmt.Println("Name of the car", (*CarA).name) // Tesla
// }

// -------------------------------

// VIDEO 31

// switches via flag Library

// flag.DataType
// flag.String("name", "fares Walid", "Full Name Entry")
// flag.Int("age", 30, "Your Age Entry")
// flag.Bool("verbose", false, "Enable Verbose")
// flag.Float

// func main()  {

// 	name := flag.String("name", "fares Walid", "Full Name Entry")
// 	age := flag.Int("age", 30, "Your Age Entry")
// 	verbose := flag.Bool("verbose", false, "Enable Verbose")
// 	flag.Parse()

// 	fmt.Println(name) // 0xc0000200b0
// 	fmt.Println(age) // 0xc00000a0c8
// 	fmt.Println(verbose) // 0xc00000a0e0

// 	fmt.Println(*name) // fares Walid
// 	fmt.Println(*age) // 30
// 	fmt.Println(*verbose) // false

// }

// -------------------------------

// VIDEO 32

// Assignment Operators

// func main()  {

// 	// // = += -= *= /=

// 	a := 30

// 	a += 10 // a = 40
// 	a -= 5 // a = 35
// 	a *= 2 // a = 70
// 	a /= 2 // a = 35

// 	fmt.Println(a)
// }

// -------------------------------

// VIDEO 33

// Data Types Converting

// func main()  {

// int8 int16 int32 int64 int
// var a int8 = 20
// fmt.Printf("%v %T\n", a, a) // 20 int8

// var b int16 = int16(a)
// fmt.Printf("%v %T\n", b, b) // 20 int16

// var a int32 = 3289647
// fmt.Printf("%v %T\n", a, a) // 3289647 int32

// var b int8 = int8(a)
// fmt.Printf("%v %T\n", b, b) // 47 int8

// c := int(a)
// fmt.Printf("%v %T\n", c, c) // 3289647 int

// int => float64
// a := 15
// fmt.Printf("%v %T\n", a, a) // 15 int

// var b float32 = float32(a)
// fmt.Printf("%v %T\n", b, b) // 15 float32

// float => int
// c := 5.5
// fmt.Printf("%v %T\n", c, c) // 5.5 float64

// var d int = int(c)
// fmt.Printf("%v %T\n", d, d) // 5 int8

// a := 5 / 2
// fmt.Printf("%v %T\n", a, a) // 2 int

// b := 5.0 / 2
// fmt.Printf("%v %T", b, b) // 2.5 float64

// int => string
// a := 15
// fmt.Printf("%v %T\n", a, a) // 15 int

// b := strconv.Itoa(a)
// fmt.Printf("%v %T\n", b, b) // 15 string

// string => int
// c := "15"
// fmt.Printf("%v %T\n", c, c) // 15 string

// d, err := strconv.Atoi(c)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("%v %T\n", d, d) // 15 int

// string => bytes
// a := "Hello, World!"
// fmt.Printf("%v %T\n", a, a) // Hello, World! string

// b := []byte(a)
// fmt.Printf("%v %T\n", b, b) // [72 101 108 108 111 44 32 87 111 114 108 100 33] []uint8

// bytes => string
// c := string(b)
// fmt.Printf("%v %T\n", c, c) // Hello, World! string

// INTEGER
//			---------------------------------------------------------------------------------
//			|	 Type		|			Minimum				|			Maximum				|
//			---------------------------------------------------------------------------------
//			|	int8		|			-128				|			127					|
//			---------------------------------------------------------------------------------
//			|	int16		|			-32768				|			32 767				|
//			---------------------------------------------------------------------------------
//			|	int32		|			-2 147 483 648		|			2 147 483 647		|
//			---------------------------------------------------------------------------------
//			|	int64		|	-9 223 372 036 854 775 808	|	9 223 372 036 854 775 807	|
//			---------------------------------------------------------------------------------
//			|	int			|	-9 223 372 036 854 775 808	|	9 223 372 036 854 775 807	|
//			---------------------------------------------------------------------------------
//			|				|								|								|
//			---------------------------------------------------------------------------------
//			|	int8		|				0				|			255					|
//			---------------------------------------------------------------------------------
//			|	int16		|				0				|			65 535				|
//			---------------------------------------------------------------------------------
//			|	int32		|				0				|			4 294 967 295		|
//			---------------------------------------------------------------------------------
//			|	int64		|				0				|	18 446 744 073 709 551 615	|
//			---------------------------------------------------------------------------------
//			|	int			|				0				|	18 446 744 073 709 551 615	|
//			---------------------------------------------------------------------------------

// }

// -------------------------------

// VIDEO 34

// if / else if / else

// func main()  {

// myInt := 10

// if 10+5 > myInt {
// 	fmt.Println("myInt is small")
// } else if 10+5 < myInt {
// 	fmt.Println("myInt is big")
// } else if 10+5 == myInt{
// 	fmt.Println("myInt is equal to 15")
// } else {
// 	fmt.Println("myInt is not in the range")
// }

// myString := "this is my clear text, Please help me!"

// if strings.Contains(myString, "clearr") {
//     fmt.Println("The string contains the word 'clear'")
// } else {
//     fmt.Println("The string does not contain the word 'clear'")
// }

// myString1 := "AAAAA"
// myString2 := "BBBBB"

// if myString1 == myString2 {
// 	fmt.Println("Yes")
// } else {
// 	fmt.Println("No")
// }
// }
// -------------------------------

// VIDEO 35

// Affirmation & Denial in "If Conditions"
// func main()  {
// myString := "ABC"

// if myString != "ABC" {
//     fmt.Println("not equal")
// } else if myString != "ABCD" {
//     fmt.Println("not ABCD") // not ABCD
// }

// myInt := 15

// if 14 < myInt && 14 != myInt {
// 	fmt.Println("Yes it's right") // Yes it's right
// }

// myString := "this is my text"

// if !strings.Contains(myString, "text") {
// 	fmt.Println("Yes, It does not contain it")
// } else {
// 	fmt.Println("No, It does contain it") // No, It does contain it
// }
// }

// -------------------------------

// VIDEO 36

// Switch Cases

// func main()  {
// 	rand.Seed(time.Now().UnixNano())
// 	randNum := rand.Intn(5) + 1

// 	switch randNum {
// 		case 1:
//             fmt.Println("The random number is 1")
//         case 2:
//             fmt.Println("The random number is 2")
//         case 3:
//             fmt.Println("The random number is 3")
//         case 4:
//             fmt.Println("The random number is 4")
//         case 5:
//             fmt.Println("The random number is 5")
//         default:
//             fmt.Println("The random number is not between 1 and 5")
// 	}
// }

// -------------------------------

// VIDEO 37

// Logical Operators
// || (OR) // && (AND)

// func main()  {

// 	myString := "stringo"
// 	myInt := 10

// 	if myString == "stringo" && myInt == 15 {
// 		fmt.Println("Both conditions are true")
// 	} else if myString == "stringo" || myInt == 10 {
// 		fmt.Println("one of them is right")
// 	} else {
// 		fmt.Println("no one of them is right !!")
// 	}
// }

// -------------------------------

// VIDEO 38

// Logical Operators More than a value as a result

// -------------------------------

// VIDEO 39

// While Loop

// func main()  {

// 	// myInt := 1
// 	// for { // Without Condition + While
// 	// 	fmt.Println(myInt)
//     //     myInt += 1
//     //     time.Sleep(1 * time.Second)
// 	// }
// }

// -------------------------------

// VIDEO 40

// For Loop

// func main()  {
// myInt := 0
// for i := 0; i <= 10; i++ { // With Condition + For
// 	fmt.Println(myInt)
// 	myInt += 1
// }
// }

// -------------------------------

// VIDEO 41

// Advanced For Loop

// func main()  {
// 	myInt := 0
// 	i := 5
// 	for { // With Condition + For
// 		if i <= 10 {
// 			i += 1
// 			fmt.Println(myInt)
// 			myInt += 1
// 			continue
// 		} else {
// 			break
// 		}
// 	}
// }

// -------------------------------

// VIDEO 42

// Continue

// func main()  {
// 	myInt := 1
// 	for {
// 		fmt.Println("i started %v...\n",myInt)
// 		time.Sleep(1 * time.Second)
// 		if myInt == 10 || myInt == 20 || myInt == 30 || myInt == 40 || myInt == 50 {
// 			continue
// 		}
// 		fmt.Println("it's the end, Will not print on 10-20-30-40-50")
// 		myInt += 1
// 	}
// }

// -------------------------------

// VIDEO 43

// Break

// func main()  {
// 	myInt := 1
// 	for {
// 		fmt.Printf("this is my var: %v\n",myInt)
// 		time.Sleep(1 * time.Second)
// 		if myInt == 10 {
//             break
//         }
// 		myInt += 1
// 	}
// }

// -------------------------------

// VIDEO 44

// Looping into file line

// func main()  {
// 	file, err := os.Open("D:\\learn_go\\syntax\file.txt")
// 	if err != nil {
//         log.Fatal(err)
//     }
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	// I/O READ I/O Writing
// 	fmt.Println(scanner)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fmt.Println(line)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// -------------------------------

// VIDEO 45

// What does if err != nil means?

// -------------------------------

// VIDEO 46

// Functions

// func FuncName() {
// 	fmt.Println("Hello, I am a function")
// }

// func _Fu_ncNa_me2() {
// 	FuncName()
// }

// func main()  {
// 	fmt.Println("This is the main function")

// 	for i := 0; i <= 5; i++ {
// 		FuncName()
// 	}

// 	_Fu_ncNa_me2()
// }

// -------------------------------

// VIDEO 47

// Functions (Values)

// func FuncName(firstName, lastName string, age int)  {
// 	fmt.Printf("My name is %v %v, and my age is %v", firstName, lastName, age)
// }

// func main()  {
// 	name := "fares"
// 	lname := "walid"
// 	myage := 18
// 	FuncName(name, lname, myage) // My name is fares walid, and my age is 18
// }

// -------------------------------

// VIDEO 48

// Functions (return) Pt, 1

// func FuncName(first_name, last_name string) string {

// 	full_name := first_name + " " + last_name
// 	return full_name
// }

// func main() {
// 	first_name := "fares"
// 	last_name := "walid"
// 	fmt.Printf("This is my full name: %v", FuncName(first_name, last_name)) // This is my full name: fares walid
// }

// -------------------------------

// VIDEO 49

// Functions (return) Pt, 2

// func FuncName(first_name, last_name, third_name string) (string, string) {

// 	two_part_name := first_name + " " + last_name
// 	full_name := first_name + " " + last_name + " " + third_name
// 	return full_name, two_part_name
// }

// func main() {
// 	first_name := "fares"
// 	last_name := "walid"
// 	third_name := "ali"
// 	full_name, two_part_name := FuncName(first_name, last_name, third_name)
// 	fmt.Printf("This is my full name: %v and this is my 2 names: %v", full_name, two_part_name) // This is my full name: fares walid ali and this is my 2 names: fares walid
// }

// -------------------------------

// VIDEO 50

// Panic

// func main() {

// fmt.Println("abc") // abc
// panic("i am panic") // panic: i am panic
// defer panic("i am panic")
// fmt.Println("def") // with defer def

// 	file, err := os.ReadFile("teto.txt")
// 	if err != nil {
// 		log.Fatal(err) // log.Fatal(err) = panic(err)
// 	}
// 	fmt.Println(file)
// }

// -------------------------------

// VIDEO 51

// Defer

// func main() {

// defer fmt.Println("World")
// fmt.Println("Hello")
// output:
// Hello
// World

// file, err := os.Open("exemple.txt")
// if err!= nil {
//     log.Fatal(err)
// }
// defer file.Close() // defer t run ftali dal function
// Write
// if Function
// ...
// blabla
// Grace defer file.Close() run here

// defer fmt.Println(1)
// defer fmt.Println(2)
// defer fmt.Println(3)
// output:
// 3
// 2
// 1
// }

// -------------------------------

// VIDEO 52

// Recover

// func main() {
// 	fmt.Println("This is a start !!")
// 	PanicMe()
// 	fmt.Println("After the panic")
// }

// func PanicMe() {
// 	defer IRecover() // recover t7ayad maf3ol panic
// 	fmt.Println("I am gonna panic now !!")
// 	panic("Panicccccccccccced") // panic twa9af lcode
// 	fmt.Println("I already Paniced !!")
// }

// func IRecover() {
//     if err := recover(); err != nil {
//         fmt.Println("Recovered from the panic: ", err)
//     }
// }

// -------------------------------

// VIDEO 53

// Pointers

// func main()  {
// 	// value = *x
// 	// memory = &x

// 	a := 0
// 	fmt.Println(a)
// 	fmt.Println(&a)

// 	y := &a
// 	*y = 1
// 	fmt.Println(*y)
// 	fmt.Println(&y)

// 	x := &y
// 	fmt.Println(x)
// }

// --------------------------------
// VIDEO 54

// os (operation system) library introduction

// Chdir // Chmod // Chown // Chtimes // Exit // Getpagesize // Hostname
// IsExist // IsNotExist // IsPathSeparator // IsPermission // IsTimeout
// Mkdir // MkdirAll // MkdirTemp // Pipe // ReadFile // Readlink // Remove
// RemoveAll // Rename // Setenv // Unsetenv //

// --------------------------------
// VIDEO 55

// os (operation system) library File Handing Checking file exist & Reading file

// func main() {
// 	myFile := "exemple5.txt"
// 	fmt.Println(CheckFileExists(myFile))
// }

// func CheckFileExists(file string) bool {
// 	_, err := os.Stat(file)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return err == nil
// }

// ReadFile
// func main() {
// 	myFile := "exemple.txt"
// 	if CheckFileExists(myFile) == true {
// 		fmt.Println(ReadFile(myFile))
// 	} else {
// 		fmt.Println("No file to open !!")
// 	}
// }

// func CheckFileExists(file string) bool {
// 	_, err := os.Stat(file)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return err == nil
// }

// func ReadFile(file string) string {
// 	content, err := os.ReadFile(file)
// 	if err!= nil {
//         panic(err)
//     }
// 	ctnt := fmt.Sprintf("%s", content)
// 	return ctnt
// }

// --------------------------------
// VIDEO 56

// os (operation system) library File Handing Creating Files & Writing Into Files

// func main()  {
// 	myFile := "fileToCreate.txt"
// 	content := "this is my content for my recent file"

// 	file, err := os.Create(myFile)
// 	if err != nil {
//         panic(err)
//     }
// 	defer file.Close()

// 	_, err = file.WriteString(content)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = file.Sync() // make sure that the file saved in hard disk
// 	if err != nil {
//         panic(err)
//     }
// }

// --------------------------------
// VIDEO 57

// os (operation system) library os.Exit()

// func main()  {
// 	status := "Okay"
// 	if status == "Error" {
// 		// fmt.Println("error")
// 		os.Exit(404)
// 		time.Sleep(5 * time.Second)
// 	} else {
// 		fmt.Print("success ")
// 		os.Exit(200)
// 	}
// }

// --------------------------------
// VIDEO 58

// os (operation system) library

// os.Stat() : ta3ti wad3iya path
// os.IsNotExist() : tchof wach file mawjood
// os.IsExist() : tchof wach file makaynch

// func main()  {
// 	_, err := os.Stat("exemple.txt")

// 	if os.IsExist(err) == true {
// 		fmt.Println("File does not exist")
// 	} else if err != nil {
// 		fmt.Println("An Error Has Occured", err)
// 	} else {
// 		fmt.Println("File Exists")
// 	}
// }

// --------------------------------
// VIDEO 59

// os library
//os.Mkdir() : make directory
// MkdirAll() : make several directories
// MkdirTemp() : make temporary directory

// func main()  {
// err := os.Mkdir("dir_name", 0755)

// err := os.MkdirAll("dir_name", 0755)

// err := os.MkdirAll("dir_name/dir2/dir3", 0755)
// tempDir, err := os.MkdirTemp("", "prefix-")
// if err!= nil {
// handle dir
// }
// fmt.Println(tempDir)

// err := os.MkdirAll("dir_name/dir2/dir3", 0755)
// tempDir, err := os.MkdirTemp("", "prefix-")
// if err!= nil {
// 	// handle dir
// }
// fmt.Println(tempDir)

// read = 4
// write = 2
// execute = 1

// 	---------------------------------------------------------------------
// 	|	constant	|	administrator 7	|	Group 5		|	Others 5	|
//	---------------------------------------------------------------------
//	|				|	read 	= 4		|	read = 4	|	read = 4	|
//	|		0		|	write 	= 2		|	execute = 1	|	execute = 1	|
//	|				|	execute = 1		|				|				|
//	---------------------------------------------------------------------

// 777 => rwxrwxrwx
// 755 => rwxr-xr-x
// 700 => rwx------
// 666 => rw-rw-rw-
// 644 => rw-r--r--
// 600 => rw-------
// 555 => r-xr-xr-x
// 500 => r-x------
// 444 => r--r--r--
// 400 => r--------
// 333 => -wx-wx-wx
// 300 => -wx------
// 222 => -w-w-w-
// 200 => -w-------
// 111 => --x--x--
// 100 => --x-----
// 000 => -----x--
// 001 => ------x-
// 010 => -----x--
// 011 => ------xx
// }

// --------------------------------
// VIDEO 60

// os library os.Rename() // os.Remove()

// func main()  {
// 	// Rename
// 	// err := os.Rename("ex.txt", "exemple.txt")
// 	// Remove
// 	err := os.Remove("abc.txt")
// 	if err != nil {
// 		// Handle error
// 	}
// }

// --------------------------------
// VIDEO 61

// os library os.Get // Set // Unset env()

// func main()  {
// 	// os.Setenv() => set env var
// 	// os.Getenv() => get env var
// 	// os.Setenv() =>

// 	// Set
// 	os.Setenv("TEST", "1234444")
// 	// Get
// 	fmt.Println("TESTXX:", os.Getenv("TEST")) // TESTXX: 1234444
// 	// Unset
// 	os.Unsetenv("TEST")
// 	fmt.Println("TESTXX:", os.Getenv("TEST")) // TESTXX:
// }

