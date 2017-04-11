package cloudword

import(
	"testing"
)

var testCasesCloud = []struct{
	Content 	[]byte
	CloudSize 	int
	Results  	CloudWord
}{
	{
		[]byte("approximation is the mother of all approximation ahaha ahaha ahaha"), 
		10,
		[]WordOccurence{
			{"ahaha", 3},
			{"approximation", 2},
			{"mother", 1},
		},
	},{
		[]byte("this is a test"), 
		10,
		[]WordOccurence{},
	},{
		[]byte("length length length other"),
		1,
		[]WordOccurence{
			{"length", 3},
		},
	},{
		[]byte("length length length other"),
		-1,
		[]WordOccurence{
			{"length", 3},
		},
	},
}

func TestWordCloud(t *testing.T) {
	for _, testCase := range testCasesCloud {
		wordCount := NewWordCounter(testCase.Content)
		cloud := wordCount.BuildCloud(testCase.CloudSize)
		for i, word := range cloud {
			if word != testCase.Results[i] {
				t.Error("expecting",testCase.Results[i] , "having", word)
			}
		}
	}
}

var testCasesMerge = []struct{
	Inputs 	[]map[string]int
	Output 	map[string]int
}{
	{
		Inputs : []map[string]int{
			map[string]int{"test" : 5, "hello" : 2},
			map[string]int{"test" : 5, "bonjour" : 1},
		},Output : map[string]int{"test" : 10, "hello" : 2, "bonjour" : 1},
	},
}

func TestMergeWordCounter(t *testing.T) {
	for _, testCase := range testCasesMerge {
		res := make(map[string]int)
		for _, input := range testCase.Inputs {
			res = Merge(res, input)
		}
		for key, expected := range testCase.Output {
			if value, exist := res[key] ; !exist || expected != value {
				t.Error("expecting", expected, "having", res)
			}
		}
	}
}