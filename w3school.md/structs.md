Go Structures
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.

Declare a Struct
To declare a structure in Go, use the type and struct keywords:

Syntax
type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}
Example

Here we declare a struct type Person with the following members: name, age, job and salary:

type Person struct {
  name string
  age int
  job string
  salary int
}
Tip: Notice that the struct members above have different data types. name and job is of type string, while age and salary is of type int.

