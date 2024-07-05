package env_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/voxelite-ai/env"
)

type StringSuite struct {
	suite.Suite

	key          string
	value        string
	defaultValue string
}

func (s *StringSuite) SetupTest() {
	s.key = uuid.New().String()
	s.value = uuid.New().String()
	s.defaultValue = uuid.New().String()
}

func TestStringTestSuite(t *testing.T) {
	suite.Run(t, new(StringSuite))
}

func TestInt64TestSuite(t *testing.T) {
	suite.Run(t, new(Int64Suite))
}

func (s *StringSuite) Test_WhenSet() {
	s.T().Setenv(s.key, s.value)

	got := env.String(s.key)

	s.Equal(s.value, got)
}

func (s *StringSuite) Test_WhenNotSet() {
	s.Panics(func() {
		env.String(s.key)
	})
}

func (s *StringSuite) Test_WhenSetWithDefaultValue() {
	s.T().Setenv(s.key, s.value)

	got := env.String(s.key, s.defaultValue)

	s.Equal(s.value, got)
}

func (s *StringSuite) Test_WhenNotSetWithDefaultValue() {
	got := env.String(s.key, s.defaultValue)

	s.Equal(s.defaultValue, got)
}

// INT64 Tests
type Int64Suite struct {
	suite.Suite

	key          string
	value        int64
	defaultValue int64
}

func (s *Int64Suite) SetupTest() {
	s.key = uuid.New().String()
	s.value = rand.Int64()
	s.defaultValue = rand.Int64()
}

func (s *Int64Suite) SetupEnv() {
	s.T().Setenv(s.key, fmt.Sprint(s.value))
}

func (s *Int64Suite) Test_WhenSet() {
	s.SetupEnv()

	got := env.Int64(s.key)

	s.Equal(s.value, got)
}

func (s *Int64Suite) Test_WhenNotSet() {
	s.Panics(func() {
		env.Int64(s.key)
	})
}

func (s *Int64Suite) Test_WhenSetWithDefaultValue() {
	s.SetupEnv()

	got := env.Int64(s.key, s.defaultValue)

	s.Equal(s.value, got)
}

func (s *Int64Suite) Test_WhenNotSetWithDefaultValue() {
	got := env.Int64(s.key, s.defaultValue)

	s.Equal(s.defaultValue, got)
}