package support

// Variables / functions declared with capital letter are accessible
// from different package by direct reference like package_name.FunctionName()

// Conf represents package level local variable that will be accessible from
// outside of support package scope.
var Conf = "Conf variable will be visible (public to the scope)"
var conf = "conf variable won't be visible (private to the scope)"

// ConfGetter refers to proxy function for package level local variable.
// It cannot be accessed directly with reference to support package.
func ConfGetter() string {
	return conf
}
