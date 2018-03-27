package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = API("divider", func() {
	Title("Divider Service")
	Description("An example illustrating error handling in goa. See docs/ErrorHandling.md.")
})

var _ = Service("divider", func() {

	// The "div_by_zero" error is defined at the service level and
	// thus may be returned by both "divide" and "integer_divide".
	Error("div_by_zero")

	HTTP(func() {
		// Use HTTP status code 400 Bad Request for "div_by_zero"
		// errors.
		Response("div_by_zero", StatusBadRequest)
	})

	Method("integer_divide", func() {
		Payload(IntOperands)
		Result(Int)

		// The "has_remainder" error is defined at the method
		// level and is thus specific to "integer_divide".
		Error("has_remainder")

		HTTP(func() {
			GET("/idiv/{a}/{b}")
			Response(StatusOK)
			Response("has_remainder", StatusExpectationFailed)
		})
	})

	Method("divide", func() {
		Payload(FloatOperands)
		Result(Float64)
		HTTP(func() {
			GET("/div/{a}/{b}")
			Response(StatusOK)
		})
	})
})

var IntOperands = Type("IntOperands", func() {
	Attribute("a", Int, "Left operand")
	Attribute("b", Int, "Right operand")
	Required("a", "b")
})

var FloatOperands = Type("FloatOperands", func() {
	Attribute("a", Float64, "Left operand")
	Attribute("b", Float64, "Right operand")
	Required("a", "b")
})