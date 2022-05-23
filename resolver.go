package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"io/ioutil"
	"os"
	"strconv"
)

type MarkDownFile struct {
	Heading   Heading
	Paragraph Detail
	Text      Detail
	Link      Detail
	Image     Detail
	List      Detail
	Code      Detail
	CodeBlock Detail
	Math      Detail
	MathBlock Detail
	Citation  Detail
}

type Heading struct {
	LevelOne   Detail
	LevelTwo   Detail
	LevelThree Detail
	LevelFour  Detail
	LevelFive  Detail
	LevelSix   Detail
}

type Detail struct {
	Num  int64 // number of node
	Item map[string]string // key: index of node(parent node index:child node index), value: content
}

func main() {
	md := new(MarkDownFile)
	md.Heading.LevelOne.Item = make(map[string]string)
	md.Heading.LevelTwo.Item = make(map[string]string)
	md.Heading.LevelThree.Item = make(map[string]string)
	md.Heading.LevelFour.Item = make(map[string]string)
	md.Heading.LevelFive.Item = make(map[string]string)
	md.Heading.LevelSix.Item = make(map[string]string)

	md.Paragraph.Item = make(map[string]string)
	md.Link.Item = make(map[string]string)
	md.Text.Item = make(map[string]string)
	md.Code.Item = make(map[string]string)
	md.Image.Item = make(map[string]string)
	md.CodeBlock.Item = make(map[string]string)
	md.List.Item = make(map[string]string)
	md.Citation.Item = make(map[string]string)
	md.MathBlock.Item = make(map[string]string)
	md.Math.Item = make(map[string]string)

	input := "test/data.md"
	source, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}
	p := parser.NewWithExtensions(parser.CommonExtensions)
	doc := p.Parse(source)
	var buf bytes.Buffer
	nodeList := doc.GetChildren()
	filename := input + ".json"
	for i := 0; i < len(nodeList); i++ {
		switch nodeList[i].(type) {
		case *ast.Heading:
			headingNode := nodeList[i].(*ast.Heading)
			if headingNode.Level == 1 {
				md.Heading.LevelOne.Num += 1
				//fmt.Println("heading 1 content:"+string(headingNode.Container.Children[0].AsLeaf().Literal))
				md.Heading.LevelOne.Item[strconv.Itoa(i)] = string(headingNode.Container.Children[0].AsLeaf().Literal)
				//level1Num += 1
			}
			if headingNode.Level == 2 {
				md.Heading.LevelTwo.Num += 1
				//fmt.Println("heading 2 content:"+string(headingNode.Container.Children[0].AsLeaf().Literal))
				md.Heading.LevelTwo.Item[strconv.Itoa(i)] = string(headingNode.Container.Children[0].AsLeaf().Literal)
				//level2Num += 1
			}
			if headingNode.Level == 3 {
				md.Heading.LevelThree.Num += 1
				md.Heading.LevelThree.Item[strconv.Itoa(i)] = string(headingNode.Container.Children[0].AsLeaf().Literal)
				//level3Num += 1
			}
			if headingNode.Level == 4 {
				md.Heading.LevelFour.Num += 1
				md.Heading.LevelFour.Item[strconv.Itoa(i)] = string(headingNode.Container.Children[0].AsLeaf().Literal)
				//level4Num += 1
			}
			if headingNode.Level == 5 {
				md.Heading.LevelFive.Num += 1
				md.Heading.LevelFive.Item[strconv.Itoa(i)] = string(headingNode.Container.Children[0].AsLeaf().Literal)
				//level5Num += 1
			}
			if headingNode.Level == 6 {
				md.Heading.LevelSix.Num += 1
				md.Heading.LevelSix.Item[strconv.Itoa(i)] = string(headingNode.Container.Children[0].AsLeaf().Literal)
				//level6Num += 1
			}
		case *ast.Paragraph:
			paragraphNode := nodeList[i].(*ast.Paragraph)
			md.Paragraph.Num += 1
			for j := 0; j < len(paragraphNode.Children); j++ {
				switch paragraphNode.Children[j].(type) {
				case *ast.Text:
					textNode := paragraphNode.Children[j].(*ast.Text)
					//fmt.Println("paragraph text node index:",i)
					//fmt.Println("paragraph text:", string(textNode.Literal))
					md.Text.Num += 1
					md.Text.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)] = string(textNode.Literal)
				case *ast.Link:
					linkNode := paragraphNode.Children[j].(*ast.Link)
					//fmt.Println("paragraph link:", string(linkNode.Destination))
					md.Link.Num += 1
					md.Link.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)] = string(linkNode.Destination)
				case *ast.Image:
					imageNode := paragraphNode.Children[j].(*ast.Image)
					//fmt.Println("paragraph image:", string(imageNode.Destination))
					md.Image.Num += 1
					md.Image.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)] = string(imageNode.Destination)
				case *ast.Code:
					codeNode := paragraphNode.Children[j].(*ast.Code)
					//fmt.Println("paragraph code:", string(codeNode.Leaf.Literal))
					md.Code.Num += 1
					md.Code.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)] = string(codeNode.Leaf.Literal)
				case *ast.CodeBlock:
					codeBlockNode := paragraphNode.Children[j].(*ast.CodeBlock)
					//fmt.Println("paragraph code block:", string(codeBlockNode.Leaf.Literal))
					md.CodeBlock.Num += 1
					md.CodeBlock.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)] = string(codeBlockNode.Leaf.Literal)
				case *ast.List:
					// list node wont be contained in the paragraph node
				case *ast.Math:
					mathNode := paragraphNode.Children[j].(*ast.Math)
					//fmt.Println("paragraph math:", string(mathNode.Literal))
					md.Math.Num += 1
					md.Math.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)] = string(mathNode.Literal)
				case *ast.MathBlock:
					mathBlockNode := paragraphNode.Children[j].(*ast.MathBlock)
					//fmt.Println("paragraph math block:", string(mathBlockNode.Literal))
					md.MathBlock.Num += 1
					md.MathBlock.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)] = string(mathBlockNode.Literal)
				case *ast.Citation:
					citationNode := paragraphNode.Children[j].(*ast.Citation)
					//fmt.Println("paragraph math block:", string(citationNode.Literal))
					md.Citation.Num += 1
					md.Citation.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)] = string(citationNode.Literal)
				}
			}
		case *ast.Text:
			textNode := nodeList[i].(*ast.Text)
			//fmt.Println("text node index:",i)
			//fmt.Println("text:", string(textNode.Literal))
			md.Text.Num += 1
			md.Text.Item[strconv.Itoa(i)] = string(textNode.Literal)
		case *ast.Link:
			linkNode := nodeList[i].(*ast.Link)
			//fmt.Println("link:", string(linkNode.Destination))
			md.Link.Num += 1
			md.Link.Item[strconv.Itoa(i)] = string(linkNode.Destination)
		case *ast.Image:
			imageNode := nodeList[i].(*ast.Image)
			//fmt.Println("image:", string(imageNode.Destination))
			md.Image.Num += 1
			md.Image.Item[strconv.Itoa(i)] = string(imageNode.Destination)
		case *ast.Code:
			codeNode := nodeList[i].(*ast.Code)
			//fmt.Println("code:", string(codeNode.Leaf.Literal))
			md.Code.Num += 1
			md.Code.Item[strconv.Itoa(i)] = string(codeNode.Leaf.Literal)
		case *ast.CodeBlock:
			codeBlockNode := nodeList[i].(*ast.CodeBlock)
			//fmt.Println("code block:", string(codeBlockNode.Leaf.Literal))
			md.CodeBlock.Num += 1
			md.CodeBlock.Item[strconv.Itoa(i)] = string(codeBlockNode.Leaf.Literal)
		case *ast.List:
			listNode := nodeList[i].(*ast.List)
			//fmt.Println("list length:", len(listNode.Children))
			md.List.Num += 1
			for j:=0; j< len(listNode.Children); j++ {
				listItem := listNode.Children[j].(*ast.ListItem)
				paragraphNode := listItem.Container.Children[0]
				for k := 0; k < len(paragraphNode.AsContainer().Children); k++ {
					switch paragraphNode.AsContainer().Children[k].(type) {
						case *ast.Text:
							textNode := paragraphNode.AsContainer().Children[k].(*ast.Text)
							//md.Text.Num += 1
							md.List.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)+":"+strconv.Itoa(k)] = string(textNode.Literal)
						case *ast.Link:
							linkNode := paragraphNode.AsContainer().Children[k].(*ast.Link)
							//fmt.Println("paragraph link:", string(linkNode.Destination))
							//md.Link.Num += 1
							md.List.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)+":"+strconv.Itoa(k)] = string(linkNode.Destination)
						case *ast.Image:
							imageNode := paragraphNode.AsContainer().Children[k].(*ast.Image)
							//fmt.Println("paragraph image:", string(imageNode.Destination))
							//md.Image.Num += 1
							md.List.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)+":"+strconv.Itoa(k)] = string(imageNode.Destination)
						case *ast.Code:
							codeNode := paragraphNode.AsContainer().Children[k].(*ast.Code)
							//fmt.Println("paragraph code:", string(codeNode.Leaf.Literal))
							//md.Code.Num += 1
							md.List.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)+":"+strconv.Itoa(k)] = string(codeNode.Leaf.Literal)
						case *ast.CodeBlock:
							codeBlockNode := paragraphNode.AsContainer().Children[k].(*ast.CodeBlock)
							//fmt.Println("paragraph code block:", string(codeBlockNode.Leaf.Literal))
							//md.CodeBlock.Num += 1
							md.List.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)+":"+strconv.Itoa(k)] = string(codeBlockNode.Leaf.Literal)
						case *ast.Math:
							mathNode := paragraphNode.AsContainer().Children[k].(*ast.Math)
							//fmt.Println("paragraph math:", string(mathNode.Literal))
							//md.Math.Num += 1
							md.List.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)+":"+strconv.Itoa(k)] = string(mathNode.Literal)
						case *ast.MathBlock:
							mathBlockNode := paragraphNode.AsContainer().Children[k].(*ast.MathBlock)
							//fmt.Println("paragraph math block:", string(mathBlockNode.Literal))
							//md.MathBlock.Num += 1
							md.List.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)+":"+strconv.Itoa(k)] = string(mathBlockNode.Literal)
						case *ast.Citation:
							citationNode := paragraphNode.AsContainer().Children[k].(*ast.Citation)
							//fmt.Println("paragraph math block:", string(citationNode.Literal))
							//md.Citation.Num += 1
							md.List.Item[strconv.Itoa(i)+":"+strconv.Itoa(j)+":"+strconv.Itoa(k)] = string(citationNode.Literal)
					}
				}
			}
		case *ast.Math:
			mathNode := nodeList[i].(*ast.Math)
			//fmt.Println("paragraph math:", string(mathNode.Literal))
			md.Math.Num += 1
			md.Math.Item[strconv.Itoa(i)] = string(mathNode.Literal)
		case *ast.MathBlock:
			mathBlockNode := nodeList[i].(*ast.MathBlock)
			//fmt.Println("paragraph math block:", string(mathBlockNode.Literal))
			md.MathBlock.Num += 1
			md.MathBlock.Item[strconv.Itoa(i)] = string(mathBlockNode.Literal)
		case *ast.Citation:
			citationNode := nodeList[i].(*ast.Citation)
			//fmt.Println("paragraph math block:", string(citationNode.Literal))
			md.Citation.Num += 1
			md.Citation.Item[strconv.Itoa(i)] = string(citationNode.Literal)
		}
	}
	ast.Print(&buf, doc)
	//got := buf.String()
	//mdFileJson := MarkDownFile{level1Num, level2Num, level3Num, level4Num, level5Num, level6Num}
	//fmt.Println("md struct:", md)
	mdJson, err := json.Marshal(md)
	if err != nil {
		fmt.Println("Error is ", err)
	}
	fmt.Println("markdown statistics:", string(mdJson))

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err2 := f.Write(mdJson)
	if err2 != nil {
		panic(err2)
	}
}
