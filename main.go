package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var declaredVars = make(map[string]bool)
var loopCount = 0

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: go run main.go [filename.simple] [-src | -build]")
		fmt.Println("  -src   : Converts to .go file only (does not run)")
		fmt.Println("  -build : Compiles to a standalone executable")
		return
	}

	inputFile := ""
	mode := "run"

	for _, arg := range args {
		if arg == "-src" {
			mode = "src"
		} else if arg == "-build" {
			mode = "build"
		} else {
			inputFile = arg
		}
	}

	if inputFile == "" {
		fmt.Println("Error: No input file provided.")
		return
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	goCode := compile(string(data))

	baseName := strings.TrimSuffix(inputFile, filepath.Ext(inputFile))
	goFilename := baseName + ".go"
	exeFilename := baseName
	if os.PathSeparator == '\\' {
		exeFilename += ".exe"
	}

	if err := os.WriteFile(goFilename, []byte(goCode), 0644); err != nil {
		fmt.Println("Error writing Go file:", err)
		return
	}

	exec.Command("go", "fmt", goFilename).Run()

	switch mode {
	case "src":
		fmt.Printf("Successfully converted '%s' to '%s'.\n", inputFile, goFilename)

	case "build":
		fmt.Printf("Building executable '%s'...\n", exeFilename)
		cmd := exec.Command("go", "build", "-o", exeFilename, goFilename)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err == nil {
			fmt.Println("Build successful.")
			os.Remove(goFilename)
		} else {
			fmt.Println("Build failed.")
		}

	case "run":
		cmd := exec.Command("go", "run", goFilename)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()

		os.Remove(goFilename)
	}
}

func compile(source string) string {
	var output strings.Builder

	output.WriteString("package main\n")
	output.WriteString("import (\n\"fmt\"\n\"time\"\n\"math/rand\"\n\"os\"\n\"os/exec\"\n\"strconv\"\n\"strings\"\n\"net/http\"\n\"io\"\n\"encoding/json\"\n)\n")

	output.WriteString(`
func clearScreen() { 
	c := exec.Command("clear"); 
	if os.PathSeparator == '\\' { c = exec.Command("cmd", "/c", "cls") }
	c.Stdout = os.Stdout; c.Run() 
}
func toInt(i interface{}) int {
	switch v := i.(type) {
	case int: return v
	case float64: return int(v)
	case string: s, _ := strconv.Atoi(v); return s
	}
	return 0
}
func contains(list interface{}, item interface{}) bool {
	switch l := list.(type) {
	case []interface{}:
		for _, v := range l { if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", item) { return true } }
	case string:
		return strings.Contains(l, fmt.Sprintf("%v", item))
	}
	return false
}
func fetchUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil { return "" }
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
func parseJson(data string) interface{} {
	var result interface{}
	json.Unmarshal([]byte(data), &result)
	return result
}
`)

	output.WriteString("func main() {\n")

	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			res := parseLine(line)
			if res != "" {
				output.WriteString(res + "\n")
			}
		}
	}

	output.WriteString("}\n")
	return output.String()
}

func parseLine(line string) string {
	if strings.HasPrefix(line, "#") {
		return ""
	}

	re := regexp.MustCompile(`[^\s"]+|"[^"]*"`)
	tokens := re.FindAllString(line, -1)
	if len(tokens) == 0 {
		return ""
	}

	cmd := tokens[0]

	if cmd == "say" {
		return fmt.Sprintf("fmt.Println(%s)", strings.Join(tokens[1:], " "))
	}

	if cmd == "wait" {
		return fmt.Sprintf("time.Sleep(time.Duration(toInt(%s)) * time.Second)", tokens[1])
	}
	if cmd == "clear" {
		return "clearScreen()"
	}
	if cmd == "exit" {
		return "os.Exit(0)"
	}
	if cmd == "system" {
		return fmt.Sprintf("exec.Command(\"sh\", \"-c\", %s).Run()", tokens[1])
	}

	if cmd == "set" {
		v := tokens[1]
		if len(tokens) == 6 {
			val1, op, val2 := tokens[3], tokens[4], tokens[5]
			if declaredVars[v] {
				return fmt.Sprintf("%s = toInt(%s) %s toInt(%s)", v, val1, op, val2)
			}
			declaredVars[v] = true
			return fmt.Sprintf("%s := toInt(%s) %s toInt(%s)", v, val1, op, val2)
		}
		val := tokens[3]
		if declaredVars[v] {
			return fmt.Sprintf("%s = %s", v, val)
		}
		declaredVars[v] = true
		return fmt.Sprintf("%s := %s", v, val)
	}

	if cmd == "inc" {
		return fmt.Sprintf("%s++", tokens[1])
	}
	if cmd == "dec" {
		return fmt.Sprintf("%s--", tokens[1])
	}
	if cmd == "add" {
		return fmt.Sprintf("%s += %s", tokens[3], tokens[1])
	}
	if cmd == "subtract" {
		return fmt.Sprintf("%s -= %s", tokens[3], tokens[1])
	}
	if cmd == "multiply" {
		return fmt.Sprintf("%s *= %s", tokens[3], tokens[1])
	}
	if cmd == "divide" {
		return fmt.Sprintf("%s /= %s", tokens[3], tokens[1])
	}
	if cmd == "modulo" {
		target, v1, v2 := tokens[5], tokens[1], tokens[3]
		if declaredVars[target] {
			return fmt.Sprintf("%s = toInt(%s) %% toInt(%s)", target, v1, v2)
		}
		declaredVars[target] = true
		return fmt.Sprintf("%s := toInt(%s) %% toInt(%s)", target, v1, v2)
	}

	if cmd == "if" || cmd == "elif" || cmd == "while" {
		v1, op, v2 := tokens[1], tokens[2], tokens[3]
		goOp := "=="
		if op == "is" {
			goOp = "=="
		}
		if op == "is_not" {
			goOp = "!="
		}
		if op == ">" {
			goOp = ">"
		}
		if op == "<" {
			goOp = "<"
		}

		prefix := "if"
		if cmd == "elif" {
			prefix = "} else if"
		}
		if cmd == "while" {
			prefix = "for"
		}

		if op == "contains" {
			return fmt.Sprintf("%s contains(%s, %s) {", prefix, v1, v2)
		}
		return fmt.Sprintf("%s %s %s %s {", prefix, v1, goOp, v2)
	}

	if cmd == "else" {
		return "} else {"
	}
	if cmd == "end" {
		return "}"
	}

	if cmd == "repeat" {
		loopCount++
		return fmt.Sprintf("for _i%d := 0; _i%d < toInt(%s); _i%d++ {", loopCount, loopCount, tokens[1], loopCount)
	}
	if cmd == "each" {
		itemVar, listName := tokens[1], tokens[3]
		declaredVars[itemVar] = true
		return fmt.Sprintf("for _, %s := range %s.([]interface{}) {", itemVar, listName)
	}

	if cmd == "list" {
		name := tokens[1]
		items := strings.Join(tokens[3:], ",")
		declaredVars[name] = true
		return fmt.Sprintf("%s := []interface{}{%s}", name, items)
	}
	if cmd == "push" {
		val, name := tokens[1], tokens[3]
		return fmt.Sprintf("%s = append(%s.([]interface{}), %s)", name, name, val)
	}
	if cmd == "get" {
		target, src, index := tokens[1], tokens[3], tokens[5]
		if declaredVars[target] {
			return fmt.Sprintf("%s = %s.([]interface{})[%s]", target, src, index)
		}
		declaredVars[target] = true
		return fmt.Sprintf("%s := %s.([]interface{})[%s]", target, src, index)
	}

	if cmd == "map" {
		name := tokens[1]
		declaredVars[name] = true
		return fmt.Sprintf("%s := make(map[string]interface{})", name)
	}
	if cmd == "put" {
		val, name, key := tokens[1], tokens[3], tokens[5]
		return fmt.Sprintf("%s.(map[string]interface{})[%s] = %s", name, key, val)
	}
	if cmd == "key" {
		target, name, key := tokens[1], tokens[3], tokens[5]
		if declaredVars[target] {
			return fmt.Sprintf("%s = %s.(map[string]interface{})[%s]", target, name, key)
		}
		declaredVars[target] = true
		return fmt.Sprintf("%s := %s.(map[string]interface{})[%s]", target, name, key)
	}

	if cmd == "define" {
		if len(tokens) > 2 && tokens[2] == "with" {
			args := tokens[3:]
			argsStr := ""
			for i, a := range args {
				argsStr += a + " interface{}"
				if i < len(args)-1 {
					argsStr += ", "
				}
			}
			return fmt.Sprintf("%s := func(%s) {", tokens[1], argsStr)
		}
		return fmt.Sprintf("%s := func() {", tokens[1])
	}
	if cmd == "run" {
		if len(tokens) > 2 && tokens[2] == "with" {
			return fmt.Sprintf("%s(%s)", tokens[1], strings.Join(tokens[3:], ", "))
		}
		return fmt.Sprintf("%s()", tokens[1])
	}

	if cmd == "ask" {
		Type, v, prompt := tokens[1], tokens[2], tokens[3]
		declaredVars[v] = true
		if Type == "number" {
			return fmt.Sprintf("fmt.Print(%s + \" \"); var _s string; fmt.Scanln(&_s); %s, _ := strconv.Atoi(_s)", prompt, v)
		}
		return fmt.Sprintf("fmt.Print(%s + \" \"); %s := \"\"; fmt.Scanln(&%s)", prompt, v, v)
	}
	if cmd == "write" {
		return fmt.Sprintf("os.WriteFile(%s, []byte(fmt.Sprintf(\"%%v\", %s)), 0644)", tokens[3], tokens[1])
	}
	if cmd == "read" {
		v, filename := tokens[3], tokens[1]
		declaredVars[v] = true
		return fmt.Sprintf("_b, _ := os.ReadFile(%s); %s := string(_b)", filename, v)
	}
	if cmd == "exists" {
		path, target := tokens[1], tokens[3]
		declaredVars[target] = true
		return fmt.Sprintf("_, _e := os.Stat(%s); %s := !os.IsNotExist(_e)", path, target)
	}
	if cmd == "fetch" {
		url, target := tokens[1], tokens[3]
		declaredVars[target] = true
		return fmt.Sprintf("%s := fetchUrl(%s)", target, url)
	}
	if cmd == "json" {
		src, target := tokens[1], tokens[3]
		declaredVars[target] = true
		return fmt.Sprintf("%s := parseJson(%s)", target, src)
	}

	if cmd == "random" {
		target, min, max := tokens[1], tokens[3], tokens[4]
		if declaredVars[target] {
			return fmt.Sprintf("%s = rand.Intn(%s - %s) + %s", target, max, min, min)
		}
		declaredVars[target] = true
		return fmt.Sprintf("%s := rand.Intn(%s - %s) + %s", target, max, min, min)
	}

	return ""
}
