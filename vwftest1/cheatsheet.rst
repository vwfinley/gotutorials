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

Variable Declarations
======================
* x = 5     (invalid)
* x int     (invalid)
* x int = 5 (invalid)
* var x int (valid, initialized to 0)
* var x = 5 (valid, initialized to 5)
* x := 5    (valid, equivalent to var x = 5)

"Collection" declarations
==========================

Array
-----

var myarray []int = {}-> array of integers
* y [5]float32 -> array of five floats.

Slices
------

Maps
----
var direction = map[string]float32 = {
  "north": 0.0,
  "east" : 90.0,
  "south": 180.0,
  "west": 270.0
}

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

