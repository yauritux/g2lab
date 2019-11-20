// main integration tests
package main

import (
	"bytes"
	"encoding/json"
	"finalproject_ecomerce/adapters/web"
	"finalproject_ecomerce/domain"
	"finalproject_ecomerce/engine"
	"finalproject_ecomerce/providers"
	mysql2 "finalproject_ecomerce/providers/mysql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"upper.io/db.v2/lib/sqlbuilder"
	"upper.io/db.v2/mysql"

	_ "github.com/mattes/migrate/driver/mysql"
	"github.com/mattes/migrate/migrate"
)

type (
	afterFunc  func(*httptest.ResponseRecorder) error
	beforeFunc func(*http.Request) error

	testCase struct {
		name               string
		url                string
		method             string
		body               interface{}
		expectedStatusCode int
		beforeFuncs        []beforeFunc
		afterFuncs         []afterFunc
	}

	testSuite struct {
		engine engine.Factory
		server http.Handler
		db     sqlbuilder.Database
		dbURL  string
	}
)

const (
	// api endpoints
	loginURL          = "/v1/auth/login"
	registerURL       = "/v1/auth/register"
	activateURL       = "/v1/auth/activate"
	forgotPasswordURL = "/v1/password/forgot"
	resetPasswordURL  = "/v1/password/reset"

	// http request methods
	mGet  = "GET"
	mPost = "POST"
	mPut  = "PUT"
	mDel  = "DELETE"

	// global user email and password
	email    = "user@example.com"
	password = "password"
)

var (
	ts                 *testSuite
	activationToken    string
	resetPasswordToken string
	authToken          string
	testUser           *domain.User
)

func TestMain(m *testing.M) {
	ts = newTestSuite()
	ts.setup()
	c := m.Run()
	ts.teardown()
	os.Exit(c)
}

func TestUser(t *testing.T) {
	type (
		rr engine.RegisterRequest
		lr engine.LoginRequest

		uar engine.UserActivateRequest
		fpr engine.ForgotPasswordRequest
		rpr engine.ResetPasswordRequest
	)

	var (
		wrongEmail    = "wrongemail@example.com"
		badEmail      = "email"
		wrongPassword = "wrong password"
		badPassword   = "p"
	)

	testCases := []testCase{
		{"register with no params", registerURL, mPost, nil, http.StatusBadRequest, nil, nil},
		{"register with bad email", registerURL, mPost, rr{Email: badEmail, Password: wrongPassword}, http.StatusBadRequest, nil, nil},
		{"register with good email but bad password", registerURL, mPost, rr{Email: wrongEmail, Password: badPassword}, http.StatusBadRequest, nil, nil},
		{"register with good params", registerURL, mPost, rr{Email: testUser.Email, Password: testUser.Password, FirstName: testUser.FirstName, LastName: testUser.LastName}, http.StatusCreated, nil, nil},
		{"login with no cred.", loginURL, mPost, nil, http.StatusUnauthorized, nil, nil},
		{"login with bad cred.", loginURL, mPost, lr{wrongEmail, wrongPassword}, http.StatusUnauthorized, nil, nil},
		{"login with inactive user", loginURL, mPost, lr{testUser.Email, testUser.Password}, http.StatusUnauthorized, nil, nil},
		{"activate with no params", activateURL, mPost, nil, http.StatusBadRequest, nil, nil},
		{"activate with different token", activateURL, mPost, uar{authToken}, http.StatusBadRequest, nil, nil},
		{"activate with goods params", activateURL, mPost, uar{activationToken}, http.StatusNoContent, nil, nil},
		{"login with good cred.", loginURL, mPost, lr{testUser.Email, testUser.Password}, http.StatusOK, nil, nil},
		{"forgot password with no params", forgotPasswordURL, mPost, nil, http.StatusNotFound, nil, nil},
		{"forgot password with not exists email", forgotPasswordURL, mPost, fpr{Email: wrongEmail}, http.StatusNotFound, nil, nil},
		{"forgot password with good params", forgotPasswordURL, mPost, fpr{Email: testUser.Email}, http.StatusNoContent, nil, nil},
		{"reset password with no params", resetPasswordURL, mPost, nil, http.StatusBadRequest, nil, nil},
		{"reset password with different token", resetPasswordURL, mPost, rpr{authToken, password}, http.StatusBadRequest, nil, nil},
		{"reset password with good token and bad password", resetPasswordURL, mPost, rpr{resetPasswordToken, badPassword}, http.StatusBadRequest, nil, nil},
		{"reset password with good params", resetPasswordURL, mPost, rpr{resetPasswordToken, "new password"}, http.StatusNoContent, nil, nil},
		{"login after reset password with old password", loginURL, mPost, lr{testUser.Email, testUser.Password}, http.StatusUnauthorized, nil, nil},
		{"login after reset password with new password", loginURL, mPost, lr{testUser.Email, "new password"}, http.StatusOK, nil, nil},
	}

	ts.runTestCases(testCases, t)
}

func newTestSuite() *testSuite {
	var ts testSuite

	dbURL, err := parseDBURL(os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatalf("parse db url failed: %v", err)
	}
	ts.dbURL = dbURL

	connURL, err := mysql.ParseURL(dbURL)
	if err != nil {
		log.Fatalf("db url bind to mysql.ConnectionURL failed: %v", err)
	}
	sess, err := mysql.Open(connURL)
	if err != nil {
		log.Fatalf("mysql connect failed: %v", err)
	}
	ts.db = sess

	// Setup storage factory
	var sf engine.StorageFactory
	sf = mysql2.NewStorage(sess)

	// Setup service dependencies
	var (
		validator  engine.Validator
		mailSender engine.MailSender
		jwt        engine.JWTSignParser
	)

	validator = providers.NewValidator()
	mailSender = providers.NewFakeMail()
	jwt = providers.NewJWT()
	emitter := providers.NewEmitter()

	ts.engine = engine.New(sf, mailSender, validator, jwt, emitter)

	ts.server = web.NewWebAdapter(ts.engine)

	return &ts
}

func (ts *testSuite) setTestUser() {
	var u domain.User
	u.Email = email
	u.Password = password
	u.FirstName = "Jhon"
	u.LastName = "Doe"
	testUser = &u
}

func (ts *testSuite) setTokens() {
	us := ts.engine.NewUser()

	token, err := us.GenToken(testUser, engine.AuthToken)
	if err != nil {
		log.Fatalf("auth token generate failed: %v", err)
	}
	authToken = token

	token, err = us.GenToken(testUser, engine.ActivationToken)
	if err != nil {
		log.Fatalf("activation token generate failed: %v", err)
	}
	activationToken = token

	token, err = us.GenToken(testUser, engine.PasswordResetToken)
	if err != nil {
		log.Fatalf("reset password token generate failed: %v", err)
	}
	resetPasswordToken = token
}

func (ts *testSuite) upDB() {
	errs, ok := migrate.UpSync("mysql://"+ts.dbURL, "../sql")
	if !ok {
		log.Fatalf("migrate up failed: %v", errs)
	}
}

func (ts *testSuite) downDB() {
	errs, ok := migrate.DownSync("mysql://"+ts.dbURL, "../sql")
	if !ok {
		log.Fatalf("migrate down failed: %v", errs)
	}
}

func (ts *testSuite) closeDB() {
	if err := ts.db.Close(); err != nil {
		log.Fatalf("close db connection failed: %v", err)
	}
}

func (ts *testSuite) setup() {
	ts.setTestUser()
	ts.setTokens()
	ts.upDB()
}

func (ts *testSuite) teardown() {
	ts.closeDB()
	ts.downDB()
}

func (ts *testSuite) runTestCases(tcs []testCase, t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ts.runTestCase(&tc, t)
		})
	}
}

func (ts *testSuite) runTestCase(tc *testCase, t *testing.T) {
	b, err := json.Marshal(tc.body)
	if err != nil {
		t.Fatalf("test case's body Marshal failed: %v", err)
	}
	reqBody := bytes.NewReader(b)
	r, err := http.NewRequest(tc.method, tc.url, reqBody)
	if err != nil {
		t.Fatalf("new request failed: %v", err)
	}

	w := httptest.NewRecorder()

	ts.server.ServeHTTP(w, r)

	if w.Code != tc.expectedStatusCode {
		t.Errorf("%s %s (%s) status code want %v got %v", tc.method, tc.url, tc.name, tc.expectedStatusCode, w.Code)
		t.Logf("request body: %v", string(b))
		t.Logf("response body: %v", w.Body)
	}

	// run after funcs
	for _, cb := range tc.afterFuncs {
		if err := cb(w); err != nil {
			t.Error(err)
		}
	}
}
