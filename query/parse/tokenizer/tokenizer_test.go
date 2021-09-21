package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenizer(t *testing.T) {
	cases := []struct {
		name string
		str  string
		want []string
	}{
		{
			"test 1",
			"SELECT id,name FROM 'users.csv' WHERE a <='base' and b=3",
			[]string{"SELECT", "id", ",", "name", "FROM", "'users.csv'", "WHERE", "a", "<", "=", "'base'", "and", "b", "=", "3"},
		},
		{
			"test 2",
			`
	if a > 10 {
		someParsable = text
	}`,
			[]string{"if", "a", ">", "10", "{", "someParsable", "=", "text", "}"},
		},
		{
			"empty",
			"",
			[]string{},
		},
		{
			"spaces",
			"  ",
			[]string{},
		},
		{
			"symbol 1",
			",  ",
			[]string{","},
		},
		{
			"symbol 2",
			"  ,  ",
			[]string{","},
		},
		{
			"symbol 3",
			"   ,",
			[]string{","},
		},
		{
			"quote 1",
			"'   ",
			[]string{"'"},
		},
		{
			"quote 2",
			"   '",
			[]string{"'"},
		},
		{
			"quote 3",
			"'   '",
			[]string{"'   '"},
		},
		{
			"quote 4",
			"  '   '  ",
			[]string{"'   '"},
		},
	}

	for _, tt := range cases {
		got := make([]string, 0)
		tokenizer := New(tt.str)
		for tokenizer.Next() {
			got = append(got, tokenizer.Token().SVal)
		}
		assert.Equal(t, tt.want, got)
	}
}

func Test_Tokenizer_EOL(t *testing.T) {
	cases := []struct {
		name    string
		str     string
		wantTyp tokenType
	}{
		{
			"EOL",
			"SELECT id,name FROM 'users.csv' WHERE a <='base' and b=3",
			EOL,
		},
	}
	for _, tt := range cases {
		tokenizer := New(tt.str)
		for tokenizer.Next() {
			assert.NotEmpty(t, tokenizer.Token().SVal)
		}
		assert.Empty(t, tokenizer.Token().SVal)
		assert.Equal(t, tt.wantTyp, tokenizer.Token().Typ)
	}

}
