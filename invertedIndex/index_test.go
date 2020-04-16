package invertedIndex

import (
	"os"
	"strings"
	"testing"
)

//trueStrings for inverted index +
//getInvertedIndex +
//makeMapForIndex +
//parseString +
//makeStringForIndex +

func TestTrueStringsForInvertedIndexWork(t *testing.T) {
	actual := []string{"Actual!!!", "Lol)", "GO!!!"}
	expect := []string{"actual", "lol", "go"}
	for i, actualString := range actual {
		if trueStringForInvertedIndex(actualString) != expect[i] {
			t.Errorf("%v is not equal %v", actual, expect)
		}
	}

}

func TestGetInvertedIndexWork(t *testing.T) {
	tmpfile, err := os.Create("tempFile.txt")
	if err != nil {
		t.Errorf("can't create tmp file in current dir!!!")
	}
	fileInfo, err := os.Lstat(tmpfile.Name())
	if err != nil {
		t.Errorf("can't create tmp fileInfo in current dir!!!")
	}
	content := "Stringgo test!!!\n Helooo!!!"
	_, err = tmpfile.WriteString(content)
	if err != nil {
		t.Errorf("can't write current data in tmp file, erroe is %s!!!", err)
	}
	fileInfoArr := []os.FileInfo{fileInfo}

	tokens := GetInvertedIndex("", fileInfoArr)
	newContent := strings.Fields(content)
	for _, contentString := range newContent {
		if tokens[trueStringForInvertedIndex(contentString)] == nil {
			t.Errorf("don't get token!!!")
		} else if tokens[trueStringForInvertedIndex(contentString)]["tempFile.txt"] == nil {
			t.Errorf("don't get file!!!")
		}
	}
	tmpfile.Close()
	err = os.Remove(tmpfile.Name())
}

func TestMakeMapForIndex(t *testing.T) {
	stringTest := []string{"test", "go", "good", "not", "good"}
	actual := makeMapForIndex(stringTest, "Test")
	for _, testString := range stringTest {
		if actual[testString] == nil {
			t.Errorf("don't have token!!!")
		}
	}
}

func TestParseString(t *testing.T) {
	line := "a: f1.txt{1} | f2.txt{2,3}"
	expect := map[string]Index{}
	expect["a"] = Index{}
	expect["a"]["f1.txt"] = []int{1}
	expect["a"]["f2.txt"] = []int{2, 3}
	actual := parseString(line)

	if actual["a"]["f1.txt"][0] != 1 {
		t.Errorf("can't have token!!!")
	}
	if actual["a"]["f2.txt"][1] != 2 {
		t.Errorf("can't have token!!!")
	}
}

func TestMakeStringForIndex(t *testing.T) {
	expect := []string{"hello", "hello"}
	actual := makeStringsForInvertedIndex([]string{"Hello1!", "Hello2!"})

	for i, stringIn := range expect {
		if stringIn != actual[i] {
			t.Errorf("%v not equal %v", actual[i], stringIn)
		}
	}
}
