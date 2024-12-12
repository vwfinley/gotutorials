=============================
Golang Cheatsheet
=============================

Resources
==========
* https://go.dev


Golang -> C# equivalents 
=========================
* module -> package
* package -> namespace
* import -> using (namespace)
* workspace -> project
* slice -> list
* map -> dictionary
* nil -> null

Golang -> C# differences
=========================
* var x int -> int x (in golang types come after variable, var keyword needed)

Commandline
===========
Useful commands of the go utility::
go help                      ->  lists help
go run .                     -> runs a program
go mod init example/hello    -> initializes a module
go mod tidy                  -> fetches modules dependencies
go get .                     -> gets dependencies in current directory
go get golang.org/x/example/hello/reverse  -> adds dependency to module
go work init ./hello         -> creates workspace file
go work use ./example/hello  -> add local file/dir to workspace

Basic Types
============
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128


Variable Declarations
======================
* x = 5     (invalid)
* x int     (invalid)
* x int = 5 (invalid)
* var x int (valid, initialized to 0)
* var x = 5 (valid, initialized to 5)
* x := 5    (valid, equivalent to var x = 5)
* x := 0.867 + 0.5i (complex128)

Multiple variables can be initialized together::
var i, j int = 1, 2
c, python, java := true, false, "no!"

Constant Declarations
=====================
Use const instead of var::
const Pi = 3.14
const World = "世界"


Equivalent Types
=================
* golang -> golang "C#"
* int32 -> int
* uint32 -> uint
* float32 -> "float"
* float64 -> "double"

Functions
==========
* Must be declared with func keyword.
* Can return a single x or tuple (x,y) value.
* Argument list is a "tuple", arguments can be collected by type and initialized: func myfunc(a, b int, x, y float64){}
* Functions with lowercase name -> are "private": not-exported (invisible) outside package.
* Functions with Uppercase name -> are "public": exported (visible) outside package.
* Can have multiple (tuple) return values.
* Named return values can be assigned in function body and implicitly returned with "return" keyword.


"Collection" declarations
==========================

Arrays
-----
Uninitialized, undetermined size::
var a []int           // declared, variable size
a = []int{1, 2, 3, 4} // implicit size assignment

Uninitialized, specific size::
var b [4]int           // declared fixed size
b = [4]int{1, 2, 3, 4} // explicit size assignment

Implicit declaration + initialized with := assignment::
c := []int{1, 2, 3, 4}    // implicit declaration, implicit size
d := [...]int{1, 2, 3, 4} // implicit declaration, implicit size
e := [4]int{1, 2, 3, 4}   // implicit declaration, explicit size

Implicit declaration  + initialized with var::
var f = []int{1, 2, 3, 4}    // implicit declaration, implicit size
var g = [...]int{1, 2, 3, 4} // implicit declaration, implicit size
var h = [4]int{1, 2, 3, 4}   // implicit declaration, explicit size

Explicit declaration  + initialized with var::
* var i []int = []int{1, 2, 3, 4} // explicit declaration, implicit size
* // var j []int = [...]int{1, 2, 3, 4} // [...] is invalid here
* var k [4]int = [4]int{1, 2, 3, 4} // explicit declaration, explicit size


Slices
------

Maps
----
Maps are declared with [key]value pairs::
var direction = map[string]float32 = {
  "north": 0.0,
  "east" : 90.0,
  "south": 180.0,
  "west": 270.0
}


Structures
===========

Structure declaration::
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

Instance of Structures in a slice::
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}



Branching
==========

If
---

Switch
------

TypeSwitch
-----------

Looping
========

For
---

