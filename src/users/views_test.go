package users

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func (s *users.UserTestSuite) TestLoginView() {
	req, err := http.NewRequest(http.MethodGet, "/api/login/", nil)
	if err != nil {
		panic(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginView)
	handler.ServeHTTP(response, req)

	s.Router.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusMethodNotAllowed, response.Code)

	//ur := NewUserRepository(s.DB)
	//user, err := ur.Create("hilal", "123456", true)
	//assert.Nil(s.T(), err)
	//reqBody, _ := json.Marshal(map[string]string{
	//	"username": user.Username,
	//	"password": "123456",
	//})
	//req, err := http.NewRequest(http.MethodPost, "/api/login/", bytes.NewReader(reqBody))
}
