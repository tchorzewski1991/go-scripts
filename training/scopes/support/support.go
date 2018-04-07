package support

// Variables / functions declared with capital letter are accessible
// from different package by direct reference like package_name.FunctionName()
var Conf = "Variable visible when support package imported (public to the scope)"
var conf = "Variable not visible when support package imported (private to the scope)"

// Setup 'getter' like function for local variable, that cannot be accessed directly
// with reference to support package. Ex. support.conf local variable call won't
// succeed.
func ConfGetter() string {
	return conf
}