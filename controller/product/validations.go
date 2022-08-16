package controller

import (
	"micro/pkg/gerrors"

	"micro/api/pb/product"
	"regexp"
)

func ValidateSampleRequest(req *product.SampleRequest) (bool, []gerrors.Violation) {
	violations := gerrors.NewViolationBuilder()
	if req.GetQty() > 100 {
		violations.Add("qty", "qty should be less than or equal to 100")
	}
	if req.GetQty() < 0 {
		violations.Add("qty", "qty should be zero or positive")
	}
	if !regexp.MustCompile(`^[A-Za-z]+$`).MatchString(req.GetName()) {
		violations.Add("name", "name should not contain numbers")
	}
	return violations.IsEmpty(), violations.Make()
}
