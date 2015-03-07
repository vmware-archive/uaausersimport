package token_test

import (
	"net/http"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/gtils/http"
	"github.com/pivotalservices/gtils/http/httptest"
	. "github.com/uaaldapimport/token"
)

var requestAdaptor RequestAdaptor = func() (*http.Response, error) {
	file, _ := os.Open("fixtures/response.json")
	fakeResponse := http.Response{
		Body: file,
	}
	return &fakeResponse, nil
}

var mockGateway *httptest.MockGateway = &httptest.MockGateway{
	FakePostAdaptor: requestAdaptor,
	Capture: func(HttpRequestEntity) {
	},
}

var _ = Describe("Token", func() {
	Describe("Retrieve token from UAA", func() {
		NewGateway = func() HttpGateway {
			return mockGateway
		}
		Context("With correct response", func() {
			It("should get a valid token", func() {
				token, err := GetToken("cf", "secret", "http://uaa.domain")
				Ω(err).Should(BeNil())
				Ω(token).Should(Equal("eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiIyNTBhZWM0YS01MTEyLTRkNTMtOGFhZC0zNjQ5OTg2YjE3MjgiLCJzdWIiOiJidWxraW1wb3J0IiwiYXV0aG9yaXRpZXMiOlsiY2xvdWRfY29udHJvbGxlci5hZG1pbiIsInNjaW0uY3JlYXRlIl0sInNjb3BlIjpbImNsb3VkX2NvbnRyb2xsZXIuYWRtaW4iLCJzY2ltLmNyZWF0ZSJdLCJjbGllbnRfaWQiOiJidWxraW1wb3J0IiwiY2lkIjoiYnVsa2ltcG9ydCIsImdyYW50X3R5cGUiOiJjbGllbnRfY3JlZGVudGlhbHMiLCJpYXQiOjE0MjU3NDUwMDgsImV4cCI6MTQyNTc4ODIwOCwiaXNzIjoiaHR0cHM6Ly91YWEuc3lzLmhvbWVsYWIuaW8vb2F1dGgvdG9rZW4iLCJhdWQiOlsic2NpbSIsImNsb3VkX2NvbnRyb2xsZXIiXX0.XswRo5djNN3BrNUaGzYVYi0eMqSP0_yDUJa2iBDpiYj17PjschDKOE15KDY8vcEd94PyRbEhOLdEbSXYFxmiBG_DCMr8Ggq4s0QKkiyAvrEv8k8IIXvOTdbwCFTYWDqSMY2EdhGvxbjE5xcilOiWWkM-JPR60Q5ke86ZM2h3X40inPE4YykhTmhBhkzeqDRdoo2wgN_arAQdZT_9sV1pBCCZbl-Z6pTeoLh7wCun1201dr6Nw2RyNe6h5JKuQyKiDWX6q5HuBtfw3RLPHiPNNwTJPmw4Ozp3y_ACxyST_osQOD1UAhHn956Tt_UMw0MNnuJt12Z4H_cH1uPwNlK8qA"))
			})
		})
	})
})
