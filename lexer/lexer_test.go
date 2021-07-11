package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Lexer(t *testing.T) {
	cases := []struct {
		name     string
		wantFlow func(lexer Lexer)
		sql      string
	}{
		{
			"test 1",
			func(lexer Lexer) {
				assert.True(t, lexer.MatchKeyword("SELECT"), "expect next is SELECT")
				assert.NoError(t, lexer.EatKeyword("SELECT"))

				assert.True(t, lexer.MatchId(), "expect next 'id' is Identity")
				got, err := lexer.EatId()
				assert.Equal(t, "id", got)
				assert.NoError(t, err)

				assert.True(t, lexer.MatchDelim(","), "expect next is ','")
				assert.NoError(t, lexer.EatDelim(","))

				assert.True(t, lexer.MatchId(), "expect next 'name' is Identity")
				got, err = lexer.EatId()
				assert.Equal(t, "name", got)
				assert.NoError(t, err)

				assert.True(t, lexer.MatchKeyword("FROM"), "expect next is FROM")
				assert.NoError(t, lexer.EatKeyword("FROM"))

				assert.True(t, lexer.MatchStringConstant(), "expect next is StringConstant")
				got, err = lexer.EatStringConstant()
				assert.Equal(t, "'users.csv'", got)
				assert.NoError(t, err)

				assert.True(t, lexer.MatchKeyword("WHERE"), "expect next is WHERE")
				assert.NoError(t, lexer.EatKeyword("WHERE"))

				assert.True(t, lexer.MatchId(), "expect next 'a' is Identity")
				got, err = lexer.EatId()
				assert.Equal(t, "a", got)
				assert.NoError(t, err)

				assert.True(t, lexer.MatchDelim("<"), "expect next is '<'")
				assert.NoError(t, lexer.EatDelim("<"))

				assert.True(t, lexer.MatchDelim("="), "expect next is '='")
				assert.NoError(t, lexer.EatDelim("="))

				assert.True(t, lexer.MatchStringConstant(), "expect next is StringConstant")
				got, err = lexer.EatStringConstant()
				assert.Equal(t, "'base'", got)
				assert.NoError(t, err)

				assert.True(t, lexer.MatchKeyword("AND"), "expect next is AND")
				assert.NoError(t, lexer.EatKeyword("AND"))

				assert.True(t, lexer.MatchId(), "expect next 'b' is Identity")
				got, err = lexer.EatId()
				assert.Equal(t, "b", got)
				assert.NoError(t, err)

				assert.True(t, lexer.MatchDelim("="), "expect next is '='")
				assert.NoError(t, lexer.EatDelim("="))

				assert.True(t, lexer.MatchNumericConstant(), "expect next is NumericConstant")
				gotNum, err := lexer.EatNumericConstant()
				assert.Equal(t, float64(3), gotNum)
				assert.NoError(t, err)

				t.Log("tokens are totally consumed, should not match anything")
				assert.False(t, lexer.MatchDelim(","))
				assert.False(t, lexer.MatchId())
				assert.False(t, lexer.MatchKeyword("FROM"))
				assert.False(t, lexer.MatchNumericConstant())
				assert.False(t, lexer.MatchStringConstant())
				assert.Error(t, lexer.EatDelim(","))
				assert.Error(t, lexer.EatKeyword("FROM"))
			},
			"SELECT id,name FROM 'users.csv' WHERE a <='base' and b=3",
		},
	}
	for _, tt := range cases {
		tt.wantFlow(New(tt.sql))
	}
}
