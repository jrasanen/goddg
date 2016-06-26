package ddg

import "testing"

func expectToBeEqual(t *testing.T, expected string, result string) {
	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func expectToBeEqualOrGreater(t *testing.T, expected int, result int) {
	if expected > result {
		t.Errorf("Expected at least %v, got %v", expected, result)
	}
}

func TestQueryRedirect(t *testing.T) {
	querySet := Query("!imdb%20titanic")
	expected := "http://www.imdb.com/find?s=all&q=titanic"
	expectToBeEqual(t, expected, querySet.Redirect)
}

func TestQueryRelatedTopics(t *testing.T) {
	querySet := Query("Pokémon")
	expectToBeEqual(t, "Pokémon", querySet.Heading)
	expectToBeEqualOrGreater(t, 1, len(querySet.RelatedTopics))
}

func TestQueryResults(t *testing.T) {
	querySet := Query("imdb")
	expectToBeEqualOrGreater(t, 1, len(querySet.Results))
}
