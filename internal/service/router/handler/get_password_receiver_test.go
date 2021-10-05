package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/lisabestteam/password-svc/internal/database/mock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetPasswordReceiver(t *testing.T) {
	var passwords = []*database.Password{
		{Id: 1, HashOfFile: "1", SenderAddress: "2", ReceiverAddress: "1", EncryptsPassword: "1"},
		{Id: 2, HashOfFile: "2", SenderAddress: "2", ReceiverAddress: "1", EncryptsPassword: "2"},
		{Id: 3, HashOfFile: "3", SenderAddress: "2", ReceiverAddress: "1", EncryptsPassword: "3"},
		{Id: 4, HashOfFile: "4", SenderAddress: "2", ReceiverAddress: "1", EncryptsPassword: "4"},
	}

	var tests = []struct {
		name      string
		receiver  string
		passwords database.Passwords
		status    int
		expected  string
	}{
		{
			name:     "Success",
			receiver: "",
			passwords: &mock.Passwords{
				SelectByReceiverFn: func(address string) ([]*database.Password, error) {
					result := passwords
					return result, nil
				}},
			status:   http.StatusOK,
			expected: "{\"data\":[{\"type\":\"password\",\"id\":\"1\",\"attributes\":{\"encrypts_password\":\"1\",\"hash_of_file\":\"1\",\"receiver_address\":\"1\",\"sender_address\":\"2\"}},{\"type\":\"password\",\"id\":\"2\",\"attributes\":{\"encrypts_password\":\"2\",\"hash_of_file\":\"2\",\"receiver_address\":\"1\",\"sender_address\":\"2\"}},{\"type\":\"password\",\"id\":\"3\",\"attributes\":{\"encrypts_password\":\"3\",\"hash_of_file\":\"3\",\"receiver_address\":\"1\",\"sender_address\":\"2\"}},{\"type\":\"password\",\"id\":\"4\",\"attributes\":{\"encrypts_password\":\"4\",\"hash_of_file\":\"4\",\"receiver_address\":\"1\",\"sender_address\":\"2\"}}]}",
		},
		{
			name:     "NotingNotFound",
			receiver: "",
			passwords: &mock.Passwords{
				SelectByReceiverFn: func(address string) ([]*database.Password, error) {
					return nil, nil
				}},
			status:   http.StatusOK,
			expected: "{\"data\":[]}",
		},
		{
			name:     "ErrInDB",
			receiver: "",
			passwords: &mock.Passwords{
				SelectByReceiverFn: func(address string) ([]*database.Password, error) {
					return nil, sql.ErrNoRows
				}},
			status:   http.StatusInternalServerError,
			expected: "{\"errors\":[{\"code\":500,\"detail\":\"sql: no rows in result set\",\"title\":\"Database return error\"}]}",
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			handler := NewPasswordHandler(tt.passwords, logrus.New())

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("localhost:5555/%s", tt.receiver), nil)

			handler.GetPasswordReceiver(w, r)


			t.Log(w.Body.String())
			assert.Equal(t, tt.status, w.Code, "the expected code differs from the received code")
			assert.JSONEq(t, tt.expected, w.Body.String(), "The expected body differs from the received body")
		})
	}
}
