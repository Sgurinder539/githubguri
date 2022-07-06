/*package main

type AddResult struct{
x        int
y        int
expected int	
}
var addResults = []AddResult{
	{1, 1, 2},
}

func TestAdd(t *testing.T){
	for_, test := range addResults{
		result := Add(test.x, test.y)
		if result != test.expected{
			t.fatal("Expexted Result not Given")
		}
	}
}