package models

func NewAPITestCase[T any](id T, name string, priority string, version int, editedBy string,
	method string, url string, endpoint string,
	inputBody interface{}, inputHeaders interface{}, inputCookies interface{},
	expectedStatus int, expectedBody interface{}, expectedHeaders interface{}, expectedCookies interface{},
) *APITestCase[T] {
	testCase := NewTestCase(id, name, priority, "APITestCase", version, editedBy)

	return &APITestCase[T]{
		TestCase: *testCase,

		Method:   method,
		URL:      url,
		Endpoint: endpoint,

		InputBody:    inputBody,
		InputHeaders: inputHeaders,
		InputCookies: inputCookies,

		ExpectedStatus:  expectedStatus,
		ExpectedBody:    expectedBody,
		ExpectedHeaders: expectedHeaders,
		ExpectedCookies: expectedCookies,
	}
}

type APITestCase[T any] struct {
	TestCase[T] `bson:"inline"`

	Method   string `json:"method" bson:"method"`
	URL      string `json:"url" bson:"url"`
	Endpoint string `json:"endpoint" bson:"endpoint"`

	InputBody    interface{} `json:"body" bson:"body"`
	InputHeaders interface{} `json:"headers" bson:"headers"`
	InputCookies interface{} `json:"cookies" bson:"cookies"`

	ExpectedStatus  int         `json:"expected_status" bson:"expected_status"`
	ExpectedBody    interface{} `json:"expected_body" bson:"expected_body"`
	ExpectedHeaders interface{} `json:"expected_headers" bson:"expected_headers"`
	ExpectedCookies interface{} `json:"expected_cookies" bson:"expected_cookies"`
}
