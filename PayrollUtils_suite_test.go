package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPayrollUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PayrollUtils Suite")
}
