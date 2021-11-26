package sentence_similarity_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, false, areSentencesSimilarTwo([]string{"an", "extraordinary", "meal", "meal"},
		[]string{"one", "good", "dinner"},
		[][]string{{"great", "good"}, {"extraordinary", "good"}, {"well", "good"}, {"wonderful", "good"}, {"excellent", "good"}, {"fine", "good"}, {"nice", "good"}, {"any", "one"}, {"some", "one"}, {"unique", "one"}, {"the", "one"}, {"an", "one"}, {"single", "one"}, {"a", "one"}, {"truck", "car"}, {"wagon", "car"}, {"automobile", "car"}, {"auto", "car"}, {"vehicle", "car"}, {"entertain", "have"}, {"drink", "have"}, {"eat", "have"}, {"take", "have"}, {"fruits", "meal"}, {"brunch", "meal"}, {"breakfast", "meal"}, {"food", "meal"}, {"dinner", "meal"}, {"super", "meal"}, {"lunch", "meal"}, {"possess", "own"}, {"keep", "own"}, {"have", "own"}, {"extremely", "very"}, {"actually", "very"}, {"really", "very"}, {"super", "very"}},
	))

	assert.Equal(t, false, areSentencesSimilarTwo([]string{"great", "acting", "skills"},
		[]string{"fine", "drama", "talent"},
		[][]string{},
	))
}
