package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Enum struct {
	Name  string
	Value string
}

var (
	inputFile   string
	packageName string
	outputDir   string
	languages   string
)

func init() {
	flag.StringVar(&inputFile, "input", "./enum.txt", "Path to the input file")
	flag.StringVar(&packageName, "package", "EnumPackage", "Package name for the generated files")
	flag.StringVar(&outputDir, "output", "lib", "Base directory for output files")
	flag.StringVar(&languages, "langs", "kotlin,javascript,swift,csharp,golang", "Comma-separated list of languages to generate enums for")
}

func main() {
	flag.Parse()

	enums := parseInputFile(inputFile)
	if enums == nil {
		fmt.Println("No enums found or failed to parse the input file.")
		return
	}

	generateEnums(enums)
}

func parseInputFile(filename string) []Enum {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var enums []Enum
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) == 2 {
			enums = append(enums, Enum{Name: parts[0], Value: parts[1]})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return enums
}

func createDirIfNotExists(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %s", dir)
	}
	return nil
}

func generateEnums(enums []Enum) {
	// Base directory
	if err := createDirIfNotExists(outputDir); err != nil {
		fmt.Println(err)
		return
	}

	if strings.Contains(languages, "kotlin") {
		generateKotlin(enums, outputDir)
	}
	if strings.Contains(languages, "javascript") {
		generateJavaScript(enums, outputDir)
	}
	if strings.Contains(languages, "swift") {
		generateSwift(enums, outputDir)
	}
	if strings.Contains(languages, "csharp") {
		generateCSharp(enums, outputDir)
	}
	if strings.Contains(languages, "golang") {
		generateGolang(enums, outputDir)
	}
}

func generateKotlin(enums []Enum, baseDir string) {
	dir := filepath.Join(baseDir, "kotlin")
	if err := createDirIfNotExists(dir); err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(filepath.Join(dir, packageName+".kt"))
	if err != nil {
		fmt.Println("Error creating Kotlin file:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	file.WriteString("enum class Element(val value: Int) {\n")
	for _, e := range enums {
		file.WriteString(fmt.Sprintf("    %s(%s),\n", e.Name, e.Value))
	}
	file.WriteString("}\n")
	fmt.Println("Kotlin enum generated:", filepath.Join(dir, packageName+".kt"))
}

func generateJavaScript(enums []Enum, baseDir string) {
	dir := filepath.Join(baseDir, "javascript")
	if err := createDirIfNotExists(dir); err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(filepath.Join(dir, packageName+".js"))
	if err != nil {
		fmt.Println("Error creating JavaScript file:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("export const %s = Object.freeze({\n", packageName))
	for _, e := range enums {
		file.WriteString(fmt.Sprintf("    %s: %s,\n", e.Name, e.Value))
	}
	file.WriteString("});\n")
	fmt.Println("JavaScript enum generated:", filepath.Join(dir, packageName+".js"))
}

func generateSwift(enums []Enum, baseDir string) {
	dir := filepath.Join(baseDir, "swift")
	if err := createDirIfNotExists(dir); err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(filepath.Join(dir, packageName+".swift"))
	if err != nil {
		fmt.Println("Error creating Swift file:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("public enum %s: Int {\n", packageName))
	for _, e := range enums {
		file.WriteString(fmt.Sprintf("    case %s = %s\n", e.Name, e.Value))
	}
	file.WriteString("}\n")
	fmt.Println("Swift enum generated:", filepath.Join(dir, packageName+".swift"))
}

func generateCSharp(enums []Enum, baseDir string) {
	dir := filepath.Join(baseDir, "csharp")
	if err := createDirIfNotExists(dir); err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(filepath.Join(dir, packageName+".cs"))
	if err != nil {
		fmt.Println("Error creating C# file:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("namespace %s {\n", packageName))
	file.WriteString("    public enum Element {\n")
	for _, e := range enums {
		file.WriteString(fmt.Sprintf("        %s = %s,\n", e.Name, e.Value))
	}
	file.WriteString("    }\n")
	file.WriteString("}\n")
	fmt.Println("C# enum generated:", filepath.Join(dir, packageName+".cs"))
}

func generateGolang(enums []Enum, baseDir string) {
	dir := filepath.Join(baseDir, "golang")
	if err := createDirIfNotExists(dir); err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(filepath.Join(dir, packageName+".go"))
	if err != nil {
		fmt.Println("Error creating Go file:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	file.WriteString("type Element int\n\n")
	file.WriteString("const (\n")
	for _, e := range enums {
		file.WriteString(fmt.Sprintf("    %s Element = %s\n", e.Name, e.Value))
	}
	file.WriteString(")\n")
	fmt.Println("Go enum generated:", filepath.Join(dir, packageName+".go"))
}
