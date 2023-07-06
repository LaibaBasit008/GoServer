# GoServer
<p>A simple server setup using GO</p>
<p>Nothing FANCY SHMANCY :P </p>
<p>If you have started learning and GO and wants to understand how the http world works along with GO syntax. This is a simple way to do it</p>


<h5>Go Functions</h5>
A function is a block of statements that can be used repeatedly in a program.

A function will not execute automatically when a page loads.

A function will be executed by a call to the function.

`func FunctionName() {
  // code to be executed
}`
<h5>Go Struct</h5>
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.

`type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}`
<h5>Go Http</h5>
ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:

`http.Handle("/foo", fooHandler)`

`http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})`

`log.Fatal(http.ListenAndServe(":8080", nil))`
