package hard

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolveSudoku(t *testing.T) {
	tests := []struct {
		name     string
		in       [][]byte
		expected [][]byte
	}{
		{
			name: "bord from example 1",
			in: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			expected: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.in)
			assert.Equal(t, tt.expected, tt.in)
		})
	}
}

func Test_Trap(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "foundation",
			in:       []int{4, 2, 3},
			expected: 1,
		}, {
			name:     "starts with 0",
			in:       []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		}, {
			name:     "bowl",
			in:       []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		}, {
			name:     "starts with 0 + foundation + bowl",
			in:       []int{0, 7, 1, 4, 6},
			expected: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, trap(tt.in))
		})
	}
}

func Test_FirstMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "from example 1",
			in:       []int{1, 2, 0},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, firstMissingPositive(tt.in))
		})
	}
}

func Test_IsMatch(t *testing.T) {
	tests := []struct {
		str      string
		pattern  string
		expected bool
	}{
		{
			str:      "",
			pattern:  "***?",
			expected: false,
		}, {
			str:      "",
			pattern:  "***",
			expected: true,
		}, {
			str:      "",
			pattern:  "*",
			expected: true,
		}, {
			str:      "",
			pattern:  "?",
			expected: false,
		}, {
			str:      "adda",
			pattern:  "ad?",
			expected: false,
		}, {
			str:      "adda",
			pattern:  "ad?ad",
			expected: false,
		}, {
			str:      "",
			pattern:  "",
			expected: true,
		}, {
			str:      "aa",
			pattern:  "a",
			expected: false,
		}, {
			str:      "aa",
			pattern:  "*",
			expected: true,
		}, {
			str:      "aa",
			pattern:  "?a",
			expected: true,
		}, {
			str:      "aa",
			pattern:  "?",
			expected: false,
		}, {
			str:      "cb",
			pattern:  "?a",
			expected: false,
		}, {
			str:      "caecfvacd",
			pattern:  "?a*?cd",
			expected: true,
		}, {
			str:      "caecfvacdcdcd",
			pattern:  "?a*?cd",
			expected: true,
		}, {
			str:      "caacd",
			pattern:  "?a*?cd",
			expected: true,
		}, {
			str:      "bbbbbbbabaab",
			pattern:  "b*b*ab**ba",
			expected: false,
		}, {
			str:      "bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab",
			pattern:  "b*b*ab**ba*b**b***bba",
			expected: false,
		}, {
			str:      "bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabbba",
			pattern:  "b*b*ab**ba*b**b***bba",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("(%s)/(%s)(%v)", tt.str, tt.pattern, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, isMatch(tt.str, tt.pattern))
		})
	}
}

func Test_SolveNQueens(t *testing.T) {
	tests := []struct {
		in       int
		expected [][]string
	}{
		{
			in: 4,
			expected: [][]string{{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			}, {
				"..Q.",
				"Q...",
				"...Q",
				".Q.."}},
		},
		{
			in:       1,
			expected: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("N count: %d", tt.in), func(t *testing.T) {
			assert.Equal(t, tt.expected, solveNQueens(tt.in))
		})
	}
}

func Test_TotalNQueens(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{
			in:       5,
			expected: 10,
		}, {
			in:       4,
			expected: 2,
		}, {
			in:       1,
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("N count: %d", tt.in), func(t *testing.T) {
			assert.Equal(t, tt.expected, totalNQueens(tt.in))
		})
	}
}

func Test_GetPermutation(t *testing.T) {
	tests := []struct {
		inN      int
		inK      int
		expected string
	}{
		{
			inN:      3,
			inK:      2,
			expected: "132",
		}, {
			inN:      3,
			inK:      3,
			expected: "213",
		}, {
			inN:      2,
			inK:      2,
			expected: "21",
		}, {
			inN:      8,
			inK:      6593,
			expected: "24186735",
		}, {
			inN:      4,
			inK:      9,
			expected: "2314",
		}, {
			inN:      3,
			inK:      1,
			expected: "123",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("N:%d,K:%d,(%s)", tt.inN, tt.inK, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, getPermutation(tt.inN, tt.inK))
		})
	}
}

func Test_IsNumber(t *testing.T) {
	valid := []string{"46.e3", "0", "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	for _, tt := range valid {
		t.Run(tt, func(t *testing.T) {
			assert.True(t, isNumber(tt))
		})
	}
	notValid := []string{".e1", "+.", "4e+", "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53", "e", "."}
	for _, tt := range notValid {
		t.Run(tt, func(t *testing.T) {
			assert.False(t, isNumber(tt))
		})
	}
}

func Test_FullJustify(t *testing.T) {
	tests := []struct {
		maxWidth int
		words    []string
		expected []string
	}{
		{
			maxWidth: 16,
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			expected: []string{
				"This    is    an",
				"example  of text",
				"justification.  "},
		},
		{
			maxWidth: 16,
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			expected: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        "},
		},
		{
			maxWidth: 20,
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			expected: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("maxWidth: %d", tt.maxWidth), func(t *testing.T) {
			assert.Equal(t, tt.expected, fullJustify(tt.words, tt.maxWidth))
		})
	}
}

func Test_FindWords(t *testing.T) {
	tests := []struct {
		board    [][]byte
		words    []string
		expected []string
	}{{
		board:    [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}},
		words:    []string{"eaabcdgfa"},
		expected: []string{"eaabcdgfa"},
	}, {
		board:    [][]byte{{'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'}, {'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'}, {'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'}, {'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'}, {'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'}, {'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'}, {'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'}, {'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'}, {'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'}, {'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'}},
		words:    []string{"ababababaa", "ababababab", "ababababac", "ababababad", "ababababae", "ababababaf", "ababababag", "ababababah", "ababababai", "ababababaj", "ababababak", "ababababal", "ababababam", "ababababan", "ababababao", "ababababap", "ababababaq", "ababababar", "ababababas", "ababababat", "ababababau", "ababababav", "ababababaw", "ababababax", "ababababay", "ababababaz", "ababababba", "ababababbb", "ababababbc", "ababababbd", "ababababbe", "ababababbf", "ababababbg", "ababababbh", "ababababbi", "ababababbj", "ababababbk", "ababababbl", "ababababbm", "ababababbn", "ababababbo", "ababababbp", "ababababbq", "ababababbr", "ababababbs", "ababababbt", "ababababbu", "ababababbv", "ababababbw", "ababababbx", "ababababby", "ababababbz", "ababababca", "ababababcb", "ababababcc", "ababababcd", "ababababce", "ababababcf", "ababababcg", "ababababch", "ababababci", "ababababcj", "ababababck", "ababababcl", "ababababcm", "ababababcn", "ababababco", "ababababcp", "ababababcq", "ababababcr", "ababababcs", "ababababct", "ababababcu", "ababababcv", "ababababcw", "ababababcx", "ababababcy", "ababababcz", "ababababda", "ababababdb", "ababababdc", "ababababdd", "ababababde", "ababababdf", "ababababdg", "ababababdh", "ababababdi", "ababababdj", "ababababdk", "ababababdl", "ababababdm", "ababababdn", "ababababdo", "ababababdp", "ababababdq", "ababababdr", "ababababds", "ababababdt", "ababababdu", "ababababdv", "ababababdw", "ababababdx", "ababababdy", "ababababdz", "ababababea", "ababababeb", "ababababec", "ababababed", "ababababee", "ababababef", "ababababeg", "ababababeh", "ababababei", "ababababej", "ababababek", "ababababel", "ababababem", "ababababen", "ababababeo", "ababababep", "ababababeq", "ababababer", "ababababes", "ababababet", "ababababeu", "ababababev", "ababababew", "ababababex", "ababababey", "ababababez", "ababababfa", "ababababfb", "ababababfc", "ababababfd", "ababababfe", "ababababff", "ababababfg", "ababababfh", "ababababfi", "ababababfj", "ababababfk", "ababababfl", "ababababfm", "ababababfn", "ababababfo", "ababababfp", "ababababfq", "ababababfr", "ababababfs", "ababababft", "ababababfu", "ababababfv", "ababababfw", "ababababfx", "ababababfy", "ababababfz", "ababababga", "ababababgb", "ababababgc", "ababababgd", "ababababge", "ababababgf", "ababababgg", "ababababgh", "ababababgi", "ababababgj", "ababababgk", "ababababgl", "ababababgm", "ababababgn", "ababababgo", "ababababgp", "ababababgq", "ababababgr", "ababababgs", "ababababgt", "ababababgu", "ababababgv", "ababababgw", "ababababgx", "ababababgy", "ababababgz", "ababababha", "ababababhb", "ababababhc", "ababababhd", "ababababhe", "ababababhf", "ababababhg", "ababababhh", "ababababhi", "ababababhj", "ababababhk", "ababababhl", "ababababhm", "ababababhn", "ababababho", "ababababhp", "ababababhq", "ababababhr", "ababababhs", "ababababht", "ababababhu", "ababababhv", "ababababhw", "ababababhx", "ababababhy", "ababababhz", "ababababia", "ababababib", "ababababic", "ababababid", "ababababie", "ababababif", "ababababig", "ababababih", "ababababii", "ababababij", "ababababik", "ababababil", "ababababim", "ababababin", "ababababio", "ababababip", "ababababiq", "ababababir", "ababababis", "ababababit", "ababababiu", "ababababiv", "ababababiw", "ababababix", "ababababiy", "ababababiz", "ababababja", "ababababjb", "ababababjc", "ababababjd", "ababababje", "ababababjf", "ababababjg", "ababababjh", "ababababji", "ababababjj", "ababababjk", "ababababjl", "ababababjm", "ababababjn", "ababababjo", "ababababjp", "ababababjq", "ababababjr", "ababababjs", "ababababjt", "ababababju", "ababababjv", "ababababjw", "ababababjx", "ababababjy", "ababababjz", "ababababka", "ababababkb", "ababababkc", "ababababkd", "ababababke", "ababababkf", "ababababkg", "ababababkh", "ababababki", "ababababkj", "ababababkk", "ababababkl", "ababababkm", "ababababkn", "ababababko", "ababababkp", "ababababkq", "ababababkr", "ababababks", "ababababkt", "ababababku", "ababababkv", "ababababkw", "ababababkx", "ababababky", "ababababkz", "ababababla", "abababablb", "abababablc", "ababababld", "abababable", "abababablf", "abababablg", "abababablh", "ababababli", "abababablj", "abababablk", "ababababll", "abababablm", "ababababln", "abababablo", "abababablp", "abababablq", "abababablr", "ababababls", "abababablt", "abababablu", "abababablv", "abababablw", "abababablx", "abababably", "abababablz", "ababababma", "ababababmb", "ababababmc", "ababababmd", "ababababme", "ababababmf", "ababababmg", "ababababmh", "ababababmi", "ababababmj", "ababababmk", "ababababml", "ababababmm", "ababababmn", "ababababmo", "ababababmp", "ababababmq", "ababababmr", "ababababms", "ababababmt", "ababababmu", "ababababmv", "ababababmw", "ababababmx", "ababababmy", "ababababmz", "ababababna", "ababababnb", "ababababnc", "ababababnd", "ababababne", "ababababnf", "ababababng", "ababababnh", "ababababni", "ababababnj", "ababababnk", "ababababnl", "ababababnm", "ababababnn", "ababababno", "ababababnp", "ababababnq", "ababababnr", "ababababns", "ababababnt", "ababababnu", "ababababnv", "ababababnw", "ababababnx", "ababababny", "ababababnz", "ababababoa", "ababababob", "ababababoc", "ababababod", "ababababoe", "ababababof", "ababababog", "ababababoh", "ababababoi", "ababababoj", "ababababok", "ababababol", "ababababom", "ababababon", "ababababoo", "ababababop", "ababababoq", "ababababor", "ababababos", "ababababot", "ababababou", "ababababov", "ababababow", "ababababox", "ababababoy", "ababababoz", "ababababpa", "ababababpb", "ababababpc", "ababababpd", "ababababpe", "ababababpf", "ababababpg", "ababababph", "ababababpi", "ababababpj", "ababababpk", "ababababpl", "ababababpm", "ababababpn", "ababababpo", "ababababpp", "ababababpq", "ababababpr", "ababababps", "ababababpt", "ababababpu", "ababababpv", "ababababpw", "ababababpx", "ababababpy", "ababababpz", "ababababqa", "ababababqb", "ababababqc", "ababababqd", "ababababqe", "ababababqf", "ababababqg", "ababababqh", "ababababqi", "ababababqj", "ababababqk", "ababababql", "ababababqm", "ababababqn", "ababababqo", "ababababqp", "ababababqq", "ababababqr", "ababababqs", "ababababqt", "ababababqu", "ababababqv", "ababababqw", "ababababqx", "ababababqy", "ababababqz", "ababababra", "ababababrb", "ababababrc", "ababababrd", "ababababre", "ababababrf", "ababababrg", "ababababrh", "ababababri", "ababababrj", "ababababrk", "ababababrl", "ababababrm", "ababababrn", "ababababro", "ababababrp", "ababababrq", "ababababrr", "ababababrs", "ababababrt", "ababababru", "ababababrv", "ababababrw", "ababababrx", "ababababry", "ababababrz", "ababababsa", "ababababsb", "ababababsc", "ababababsd", "ababababse", "ababababsf", "ababababsg", "ababababsh", "ababababsi", "ababababsj", "ababababsk", "ababababsl", "ababababsm", "ababababsn", "ababababso", "ababababsp", "ababababsq", "ababababsr", "ababababss", "ababababst", "ababababsu", "ababababsv", "ababababsw", "ababababsx", "ababababsy", "ababababsz", "ababababta", "ababababtb", "ababababtc", "ababababtd", "ababababte", "ababababtf", "ababababtg", "ababababth", "ababababti", "ababababtj", "ababababtk", "ababababtl", "ababababtm", "ababababtn", "ababababto", "ababababtp", "ababababtq", "ababababtr", "ababababts", "ababababtt", "ababababtu", "ababababtv", "ababababtw", "ababababtx", "ababababty", "ababababtz", "ababababua", "ababababub", "ababababuc", "ababababud", "ababababue", "ababababuf", "ababababug", "ababababuh", "ababababui", "ababababuj", "ababababuk", "ababababul", "ababababum", "ababababun", "ababababuo", "ababababup", "ababababuq", "ababababur", "ababababus", "ababababut", "ababababuu", "ababababuv", "ababababuw", "ababababux", "ababababuy", "ababababuz", "ababababva", "ababababvb", "ababababvc", "ababababvd", "ababababve", "ababababvf", "ababababvg", "ababababvh", "ababababvi", "ababababvj", "ababababvk", "ababababvl", "ababababvm", "ababababvn", "ababababvo", "ababababvp", "ababababvq", "ababababvr", "ababababvs", "ababababvt", "ababababvu", "ababababvv", "ababababvw", "ababababvx", "ababababvy", "ababababvz", "ababababwa", "ababababwb", "ababababwc", "ababababwd", "ababababwe", "ababababwf", "ababababwg", "ababababwh", "ababababwi", "ababababwj", "ababababwk", "ababababwl", "ababababwm", "ababababwn", "ababababwo", "ababababwp", "ababababwq", "ababababwr", "ababababws", "ababababwt", "ababababwu", "ababababwv", "ababababww", "ababababwx", "ababababwy", "ababababwz", "ababababxa", "ababababxb", "ababababxc", "ababababxd", "ababababxe", "ababababxf", "ababababxg", "ababababxh", "ababababxi", "ababababxj", "ababababxk", "ababababxl", "ababababxm", "ababababxn", "ababababxo", "ababababxp", "ababababxq", "ababababxr", "ababababxs", "ababababxt", "ababababxu", "ababababxv", "ababababxw", "ababababxx", "ababababxy", "ababababxz", "ababababya", "ababababyb", "ababababyc", "ababababyd", "ababababye", "ababababyf", "ababababyg", "ababababyh", "ababababyi", "ababababyj", "ababababyk", "ababababyl", "ababababym", "ababababyn", "ababababyo", "ababababyp", "ababababyq", "ababababyr", "ababababys", "ababababyt", "ababababyu", "ababababyv", "ababababyw", "ababababyx", "ababababyy", "ababababyz", "ababababza", "ababababzb", "ababababzc", "ababababzd", "ababababze", "ababababzf", "ababababzg", "ababababzh", "ababababzi", "ababababzj", "ababababzk", "ababababzl", "ababababzm", "ababababzn", "ababababzo", "ababababzp", "ababababzq", "ababababzr", "ababababzs", "ababababzt", "ababababzu", "ababababzv", "ababababzw", "ababababzx", "ababababzy", "ababababzz"},
		expected: []string{"ababababab"},
	}, {
		board:    [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}},
		words:    []string{"eaafgdcba"},
		expected: []string{"eaafgdcba"},
	}, {
		board:    [][]byte{{'a', 'a'}, {'b', 'a'}},
		words:    []string{"aaba"},
		expected: []string{"aaba"},
	}, {
		board:    [][]byte{{'a', 'a'}, {'b', 'a'}},
		words:    []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"},
		expected: []string{"aba", "baa", "aaab", "aaa", "aaba"},
	}, {
		board:    [][]byte{{'a'}, {'a'}},
		words:    []string{"aaa"},
		expected: nil,
	}, {
		board:    [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
		words:    []string{"oath", "pea", "eat", "rain"},
		expected: []string{"oath", "eat"},
	}, {
		board:    [][]byte{{'a', 'b'}, {'c', 'd'}},
		words:    []string{"abcb"},
		expected: nil,
	},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("words len: %d", len(tt.words)), func(t *testing.T) {
			assert.Equal(t, tt.expected, findWords(tt.board, tt.words))
		})
	}
}

func Test_OrderlyQueue(t *testing.T) {
	tests := []struct {
		in       string
		k        int
		expected string
	}{
		{
			in:       "bcadacikam",
			k:        1,
			expected: "acikambcad",
		}, {
			in:       "bcadacikam",
			k:        2,
			expected: "aaabccdikm",
		}, {
			in:       "baaca",
			k:        3,
			expected: "aaabc",
		}, {
			in:       "cba",
			k:        1,
			expected: "acb",
		}, {
			in:       "edcba",
			k:        1,
			expected: "aedcb",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("(%d):%s", tt.k, tt.in), func(t *testing.T) {
			assert.Equal(t, tt.expected, orderlyQueue(tt.in, tt.k))
		})
	}
}
