# 🚀 Simple Language

<div align="center">

**A minimalist, human-readable programming language that compiles to Go**

*Write code that reads like plain English!*

[Quick Start](#-quick-start) • [Features](#-complete-command-reference) • [Examples](#-example-programs) • [Docs](#-language-syntax)

---

</div>

## 📋 Table of Contents

- [What is Simple?](#-what-is-simple)
- [Quick Start](#-quick-start)
- [Installation](#-installation--usage)
- [Complete Command Reference](#-complete-command-reference)
- [Example Programs](#-example-programs)
- [Advanced Patterns](#-advanced-patterns)
- [Language Syntax](#-language-syntax)
- [Troubleshooting](#-troubleshooting)
- [Best Practices](#-best-practices)

---

## 💡 What is Simple?

Simple is a beginner-friendly programming language designed to make coding accessible to everyone. Instead of cryptic symbols and complex syntax, Simple uses natural English-like commands.

### Why Simple?

| Traditional Code | Simple Code |
|-----------------|-------------|
| `for i := 0; i < 10; i++` | `repeat 10` |
| `arr = append(arr, val)` | `push val to arr` |
| `if x == y { ... }` | `if x is y` |
| `resp, _ := http.Get(url)` | `fetch url into response` |

**Key Features:**
- ✅ No semicolons, brackets, or confusing syntax
- ✅ Reads like natural language
- ✅ Built-in HTTP requests and JSON parsing
- ✅ Automatic type handling
- ✅ Compiles to efficient Go code
- ✅ Perfect for learning programming concepts

---

## ⚡ Quick Start

**1. Create `hello.simple`:**
```
say "Hello, World!"
ask text name "What's your name?"
say "Nice to meet you," name "!"
```

**2. Run it:**
```bash
go run main.go hello.simple
```

**3. Start coding!** 🎉

---

## 📦 Installation & Usage

### Running Simple Programs

Simple provides three modes of operation:

#### **Default Mode: Run**
Compiles and executes immediately, then cleans up:
```bash
go run main.go program.simple
```

#### **Source Mode: Generate Go Code**
Converts to .go file without running:
```bash
go run main.go program.simple -src
```
Creates `program.go` that you can inspect or modify.

#### **Build Mode: Create Executable**
Compiles to standalone executable:
```bash
go run main.go program.simple -build
```
Creates `program.exe` (Windows) or `program` (Unix) that runs without Go.

---

### ⚙️ System Requirements

- **To compile:** Go 1.16+ 
- **To run executables:** No dependencies!
- **Platforms:** Windows, Linux, macOS

---

## 📖 Complete Command Reference

### 💬 Output & Display

| Command | Description | Example |
|---------|-------------|---------|
| `say` | Print to console | `say "Hello World"` |
| `say` (multi) | Print multiple values | `say "Score:" score` |
| `clear` | Clear the screen | `clear` |

**Examples:**
```
say "Welcome to my program!"
say "Your score is" score "points"
say x y z
```

---

### 🔢 Variables & Assignment

| Command | Syntax | Description |
|---------|--------|-------------|
| `set` | `set VAR to VALUE` | Create or update variable |
| `set` (math) | `set VAR to X OP Y` | Set with calculation |

**Examples:**
```
set x to 10
set name to "Alice"
set pi to 3.14
set sum to a + b
set product to 5 * 3
set difference to 10 - 3
set quotient to 20 / 4
```

**Supported operations:** `+` `-` `*` `/`

**Important:** Math expressions must follow the exact pattern: `set VAR to VALUE OPERATOR VALUE`

---

### 🧮 Math Operations

| Command | Syntax | Description |
|---------|--------|-------------|
| `inc` | `inc VAR` | Increment by 1 |
| `dec` | `dec VAR` | Decrement by 1 |
| `add` | `add N to VAR` | Add N to variable |
| `subtract` | `subtract N from VAR` | Subtract N from variable |
| `multiply` | `multiply VAR by N` | Multiply variable |
| `divide` | `divide VAR by N` | Divide variable |
| `modulo` | `modulo X by Y store VAR` | Get remainder |
| `random` | `random VAR between MIN MAX` | Random number |

**Examples:**
```
set score to 100
inc score              # score = 101
add 50 to score        # score = 151
multiply score by 2    # score = 302
subtract 100 from score # score = 202
divide score by 2      # score = 101

modulo 17 by 5 store remainder    # remainder = 2
random dice between 1 6           # dice = 1-6
```

---

### 🎯 Conditionals

| Operator | Meaning | Example |
|----------|---------|---------|
| `is` | Equal | `if x is 5` |
| `is_not` | Not equal | `if x is_not 0` |
| `>` | Greater than | `if age > 18` |
| `<` | Less than | `if temp < 0` |
| `contains` | Contains substring/item | `if list contains "apple"` |

**If-Elif-Else:**
```
if score > 90
    say "Grade: A"
elif score > 80
    say "Grade: B"
elif score > 70
    say "Grade: C"
else
    say "Grade: F"
end
```

**Contains check:**
```
set text to "Hello World"
if text contains "World"
    say "Found it!"
end

list fruits is "apple" "banana" "orange"
if fruits contains "banana"
    say "We have bananas!"
end
```

---

### 🔁 Loops

**Repeat N times:**
```
repeat 10
    say "Loop iteration"
end
```

**While loop:**
```
set i to 0
while i < 10
    say i
    inc i
end
```

**For-each loop:**
```
list colors is "red" "green" "blue"

each color in colors
    say color
end
```

---

### 📋 Lists (Arrays)

| Command | Syntax | Description |
|---------|--------|-------------|
| `list` | `list NAME is ITEMS...` | Create list |
| `push` | `push VALUE to LIST` | Add to end |
| `get` | `get VAR from LIST at INDEX` | Get by index |

**Examples:**
```
# Create list
list numbers is 1 2 3 4 5
list names is "Alice" "Bob" "Charlie"
list empty is

# Add items
push 6 to numbers
push "Dave" to names

# Access items (0-indexed)
get first from numbers at 0
say first    # Prints: 1

get last from numbers at 5
say last     # Prints: 6

# Check contents
if numbers contains 3
    say "Found 3!"
end

# Loop through
each num in numbers
    say num
end
```

---

### 🗂️ Maps (Dictionaries)

| Command | Syntax | Description |
|---------|--------|-------------|
| `map` | `map NAME` | Create empty map |
| `put` | `put VALUE in MAP at KEY` | Set key-value |
| `key` | `key VAR from MAP at KEY` | Get value by key |

**Examples:**
```
# Create map
map person
map scores

# Add data
put "Alice" in person at "name"
put 30 in person at "age"
put "Engineer" in person at "job"

put 95 in scores at "math"
put 87 in scores at "english"

# Read data
key name from person at "name"
key age from person at "age"
say name "is" age "years old"

key mathScore from scores at "math"
say "Math score:" mathScore
```

---

### 🎭 Functions

**Simple function:**
```
define greet
    say "Hello, World!"
end

run greet
```

**Function with parameters:**
```
define introduce with name age
    say "My name is" name
    say "I am" age "years old"
end

run introduce with "Alice" 25
```

**Multiple functions:**
```
define add with a b
    set sum to a + b
    say "Sum:" sum
end

define multiply with a b
    set product to a * b
    say "Product:" product
end

run add with 5 3
run multiply with 4 7
```

---

### 📥 User Input

| Command | Syntax | Description |
|---------|--------|-------------|
| `ask text` | `ask text VAR "PROMPT"` | Get string input |
| `ask number` | `ask number VAR "PROMPT"` | Get integer input |

**Examples:**
```
# Text input
ask text name "Enter your name:"
say "Hello," name

# Number input
ask number age "Enter your age:"
if age > 18
    say "You are an adult"
else
    say "You are a minor"
end

# Multiple inputs
ask text city "Where do you live?"
ask number year "What year were you born?"
set currentAge to 2024 - year
say "You live in" city "and are about" currentAge
```

---

### 💾 File Operations

| Command | Syntax | Description |
|---------|--------|-------------|
| `write` | `write DATA to "FILE"` | Write to file |
| `read` | `read "FILE" into VAR` | Read from file |
| `exists` | `exists "FILE" store VAR` | Check if exists |

**Examples:**
```
# Write to file
set message to "Hello from Simple!"
write message to "output.txt"
say "File saved!"

# Read from file
read "data.txt" into content
say "File contents:"
say content

# Check if file exists
exists "config.txt" store fileExists
if fileExists is true
    read "config.txt" into config
    say "Config loaded"
else
    say "Config not found"
    set config to "default"
end
```

---

### 🌐 Network & HTTP

| Command | Syntax | Description |
|---------|--------|-------------|
| `fetch` | `fetch "URL" into VAR` | HTTP GET request |
| `json` | `json DATA into VAR` | Parse JSON |

**Examples:**
```
# Simple API call
fetch "https://api.github.com/users/torvalds" into response
say response

# Fetch and parse JSON
fetch "https://api.github.com/users/octocat" into jsonData
json jsonData into user
say "User data received"

# Check if API is available
fetch "https://example.com/api/health" into status
if status contains "ok"
    say "API is online!"
else
    say "API is down"
end
```

---

### 💻 System Commands

| Command | Description | Example |
|---------|-------------|---------|
| `wait` | Sleep for N seconds | `wait 5` |
| `clear` | Clear screen | `clear` |
| `exit` | Exit program | `exit` |
| `system` | Run shell command | `system "ls -la"` |

**Examples:**
```
# Countdown timer
set i to 5
while i > 0
    say i
    wait 1
    dec i
end
say "Blast off!"

# Clear and display
clear
say "Screen cleared!"

# Run system commands (Unix/Linux/Mac)
system "echo 'Hello from shell'"
system "ls -la"
```

**Note:** `system` command uses `sh -c` on all platforms, so use Unix-style commands.

---

### 💬 Comments

```
# This is a comment
say "This runs"

set x to 10    # Comments are ignored after code

# Comments can explain your code
# Multiple lines of comments
say "Comments are ignored"
```

---

## 🎨 Example Programs

### 🌟 Hello World
```
say "Hello, World!"
```

---

### 🎮 Number Guessing Game
```
say "=== NUMBER GUESSING GAME ==="
random secret between 1 100
set attempts to 0

say "I'm thinking of a number between 1 and 100"

set guessed to false
while guessed is false
    ask number guess "Your guess:"
    inc attempts
    
    if guess is secret
        say "🎉 Correct! You won in" attempts "attempts!"
        set guessed to true
    elif guess < secret
        say "📈 Too low! Try again."
    else
        say "📉 Too high! Try again."
    end
end
```

---

### 🔢 FizzBuzz
```
say "=== FIZZBUZZ ==="

set i to 1
while i < 101
    modulo i by 15 store mod15
    modulo i by 3 store mod3
    modulo i by 5 store mod5
    
    if mod15 is 0
        say "FizzBuzz"
    elif mod3 is 0
        say "Fizz"
    elif mod5 is 0
        say "Buzz"
    else
        say i
    end
    
    inc i
end
```

---

### 🧮 Calculator
```
say "=== SIMPLE CALCULATOR ==="

ask number a "Enter first number:"
ask number b "Enter second number:"

set sum to a + b
set difference to a - b
set product to a * b
set quotient to a / b

say ""
say "Results:"
say a "+" b "=" sum
say a "-" b "=" difference
say a "*" b "=" product
say a "/" b "=" quotient

modulo a by b store remainder
say a "%" b "=" remainder
```

---

### ✅ Todo List Manager
```
say "=== TODO LIST MANAGER ==="
list todos is

define showMenu
    say ""
    say "1. Add task"
    say "2. View all tasks"
    say "3. Exit"
    say ""
end

define showTasks
    say ""
    say "=== YOUR TASKS ==="
    each task in todos
        say "• " task
    end
    say "================="
end

set running to true
while running is true
    run showMenu
    ask number choice "Choose option:"
    
    if choice is 1
        ask text task "Enter new task:"
        push task to todos
        say "✓ Task added!"
    elif choice is 2
        run showTasks
    elif choice is 3
        say "Goodbye!"
        set running to false
    else
        say "Invalid option!"
    end
end
```

---

### 🎲 Dice Roller
```
say "=== DICE ROLLER ==="

define rollDice with sides
    random result between 1 sides
    say "🎲 Rolled a" result
end

set playing to true
while playing is true
    say ""
    say "Choose dice:"
    say "1. D6 (6-sided)"
    say "2. D20 (20-sided)"
    say "3. D100 (100-sided)"
    say "4. Exit"
    
    ask number choice "Your choice:"
    
    if choice is 1
        run rollDice with 6
    elif choice is 2
        run rollDice with 20
    elif choice is 3
        run rollDice with 100
    elif choice is 4
        set playing to false
        say "Thanks for playing!"
    end
end
```

---

### 🌡️ Temperature Converter
```
say "=== TEMPERATURE CONVERTER ==="

define celsiusToFahrenheit with c
    set f to c * 9
    divide f by 5
    add 32 to f
    say c "°C =" f "°F"
end

define fahrenheitToCelsius with f
    set c to f - 32
    multiply c by 5
    divide c by 9
    say f "°F =" c "°C"
end

say "1. Celsius to Fahrenheit"
say "2. Fahrenheit to Celsius"
ask number choice "Choose:"

if choice is 1
    ask number temp "Enter °C:"
    run celsiusToFahrenheit with temp
elif choice is 2
    ask number temp "Enter °F:"
    run fahrenheitToCelsius with temp
end
```

---

### 📊 Grade Calculator
```
say "=== GRADE CALCULATOR ==="

list scores is
set total to 0
set count to 0

ask number numGrades "How many grades to enter?"

repeat numGrades
    ask number grade "Enter grade:"
    push grade to scores
    add grade to total
    inc count
end

set average to total / count

say ""
say "=== RESULTS ==="
say "Total grades:" count
say "Sum:" total
say "Average:" average

if average > 90
    say "Grade: A (Excellent!)"
elif average > 80
    say "Grade: B (Good job!)"
elif average > 70
    say "Grade: C (Fair)"
elif average > 60
    say "Grade: D (Needs improvement)"
else
    say "Grade: F (Failed)"
end
```

---

### 📝 Journal with File Storage
```
say "=== DAILY JOURNAL ==="

exists "journal.txt" store hasJournal
if hasJournal is true
    read "journal.txt" into entries
else
    set entries to ""
end

say ""
say "1. New entry"
say "2. View entries"
say "3. Exit"

ask number choice "Choose:"

if choice is 1
    ask text entry "Write your entry:"
    set newEntry to entries
    write newEntry to "journal.txt"
    say "✓ Entry saved!"
elif choice is 2
    if hasJournal is true
        say ""
        say "=== YOUR JOURNAL ==="
        say entries
    else
        say "No entries yet!"
    end
elif choice is 3
    say "Goodbye!"
end
```

---

## 🚀 Advanced Patterns

### Nested Loops
```
say "Multiplication Table:"
set i to 1
while i < 11
    set j to 1
    while j < 11
        set product to i * j
        say i "×" j "=" product
        inc j
    end
    inc i
end
```

---

### Complex Conditionals
```
if age > 18
    if hasLicense is true
        if hasCar is true
            say "You can drive!"
        else
            say "You need a car"
        end
    else
        say "Get a license first"
    end
else
    say "Too young to drive"
end
```

---

### Working with Maps
```
# Create configuration
map config
put "localhost" in config at "host"
put 8080 in config at "port"
put true in config at "debug"

# Read configuration
key host from config at "host"
key port from config at "port"
key debug from config at "debug"

say "Server:" host ":" port
if debug is true
    say "Debug mode enabled"
end
```

---

## 📚 Language Syntax

### Keywords

**Control Flow:** `if`, `elif`, `else`, `end`, `while`, `repeat`, `each`

**Operators:** `is`, `is_not`, `>`, `<`, `contains`, `to`, `from`, `in`, `at`, `with`, `by`, `store`, `into`, `between`

**Commands:** `say`, `set`, `ask`, `inc`, `dec`, `add`, `subtract`, `multiply`, `divide`, `modulo`, `random`, `list`, `push`, `get`, `map`, `put`, `key`, `define`, `run`, `write`, `read`, `exists`, `fetch`, `json`, `wait`, `clear`, `exit`, `system`

**Data Types:**
- Numbers: `42`, `3.14`, `-10`
- Strings: `"hello"`, `"multi word"`
- Booleans: `true`, `false`
- Lists: Arrays of mixed types
- Maps: String-keyed dictionaries

---

### Syntax Patterns

**Variable assignment:**
```
set VAR to VALUE
set VAR to VALUE1 OPERATOR VALUE2
```

**Math operations:**
```
inc VAR
dec VAR
add VALUE to VAR
subtract VALUE from VAR
multiply VAR by VALUE
divide VAR by VALUE
modulo VALUE1 by VALUE2 store VAR
```

**Conditionals:**
```
if VALUE1 OPERATOR VALUE2
    # code
elif VALUE1 OPERATOR VALUE2
    # code
else
    # code
end
```

**Loops:**
```
repeat COUNT
    # code
end

while VALUE1 OPERATOR VALUE2
    # code
end

each ITEM in LIST
    # code
end
```

**Functions:**
```
define NAME
    # code
end

define NAME with PARAM1 PARAM2
    # code
end

run NAME
run NAME with ARG1 ARG2
```

---

### Naming Rules

✅ **Valid names:**
- `myVariable`
- `user_name`
- `score2`
- `totalAmount`

❌ **Invalid names:**
- Keywords: `if`, `end`, `say`
- Starting with numbers: `2ndPlace`
- Spaces: `my variable`
- Special chars: `total$`, `name@`

---

### String Literals

Strings use **double quotes** only:
```
set message to "Hello"
set path to "C:\Users\Documents"
set multiWord to "This is a sentence"
```

---

## 🔧 Troubleshooting

### Common Errors

**Error:** `Usage: go run main.go [filename.simple]`
- **Fix:** Provide a filename: `go run main.go test.simple`

---

**Error:** `Error reading file`
- **Fix:** Check file exists and path is correct
- **Fix:** Ensure `.simple` extension

---

**Error:** Compiler hangs or crashes
- **Fix:** Check for missing `end` keywords
- **Fix:** Verify all `if`/`while`/`repeat`/`define` have matching `end`

---

**Error:** Variable not defined
- **Fix:** Declare with `set` before using
- **Fix:** Check spelling matches exactly

---

**Error:** Math operation doesn't work
- **Fix:** Use exact syntax: `set result to value1 + value2`
- **Fix:** Ensure spaces around operator

---

**Error:** List index out of bounds
- **Fix:** Lists are 0-indexed
- **Fix:** Check list length before accessing

---

### Debug Tips

```
# Print variable values
say "Debug: x =" x

# Check conditions
if condition is true
    say "Condition passed"
else
    say "Condition failed"
end

# Trace execution
say "Reached point A"
# ... code ...
say "Reached point B"
```

---

## ✨ Best Practices

### 1. Use Descriptive Names
```
# ❌ Bad
set x to 100
set y to 5

# ✅ Good
set totalScore to 100
set lives to 5
```

---

### 2. Comment Your Code
```
# Initialize player stats
set health to 100
set mana to 50

# Game loop
while playing is true
    # Handle player input
    ask text action "What do you do?"
end
```

---

### 3. Break Code Into Functions
```
# ❌ Repeated code
say "Score:" score
say "Lives:" lives
say "Level:" level

# ✅ Use function
define showStats
    say "Score:" score
    say "Lives:" lives
    say "Level:" level
end

run showStats
```

---

### 4. Validate Input
```
ask number age "Enter age:"

if age < 0
    say "Invalid age!"
    exit
elif age > 150
    say "Invalid age!"
    exit
end
```

---

### 5. Handle Edge Cases
```
# Check before dividing
if denominator is 0
    say "Cannot divide by zero!"
else
    divide numerator by denominator
end

# Check file exists
exists "data.txt" store fileExists
if fileExists is true
    read "data.txt" into data
else
    say "File not found - using defaults"
    set data to "default"
end
```

---

## 🎯 Quick Reference Card

```
┌─────────────────────────────────────────────────┐
│                 SIMPLE LANGUAGE                 │
│              Quick Reference Guide              │
├─────────────────────────────────────────────────┤
│ OUTPUT                                          │
│  say "text"              Print to console       │
│  clear                   Clear screen           │
├─────────────────────────────────────────────────┤
│ VARIABLES                                       │
│  set x to 10             Create/assign          │
│  set x to a + b          Math assignment        │
│  inc x                   Increment              │
│  dec x                   Decrement              │
│  add 5 to x              x += 5                 │
│  subtract 5 from x       x -= 5                 │
│  multiply x by 2         x *= 2                 │
│  divide x by 2           x /= 2                 │
├─────────────────────────────────────────────────┤
│ CONDITIONS                                      │
│  if x is 5               Equal                  │
│  if x is_not 5           Not equal              │
│  if x > 10               Greater than           │
│  if x < 10               Less than              │
│  if x contains "a"       Contains               │
│  elif / else / end       Branches               │
├─────────────────────────────────────────────────┤
│ LOOPS                                           │
│  repeat 10 ... end       Loop N times           │
│  while x < 10 ... end    Conditional loop       │
│  each item in list       For-each               │
├─────────────────────────────────────────────────┤
│ LISTS                                           │
│  list x is 1 2 3         Create list            │
│  push 4 to x             Append                 │
│  get item from x at 0    Access by index        │
├─────────────────────────────────────────────────┤
│ MAPS                                            │
│  map x                   Create map             │
│  put val in x at "key"   Set value              │
│  key var from x at "key" Get value              │
├─────────────────────────────────────────────────┤
│ FUNCTIONS                                       │
│  define func ... end     Declare function       │
│  define f with x y       With parameters        │
│  run func                Call function          │
│  run f with 5 10         Call with args         │
├─────────────────────────────────────────────────┤
│ INPUT/OUTPUT                                    │
│  ask text x "prompt"     String input           │
│  ask number x "prompt"   Number input           │
│  write x to "file"       Write file             │
│  read "file" into x      Read file              │
│  exists "file" store x   Check exists           │
├─────────────────────────────────────────────────┤
│ NETWORK                                         │
│  fetch "url" into x      HTTP GET               │
│  json data into x        Parse JSON             │
├─────────────────────────────────────────────────┤
│ SYSTEM                                          │
│  wait 5                  Sleep seconds          │
│  exit                    Quit program           │
│  system "command"        Run shell command      │
│  random x between 1 10   Random number          │
│  modulo a by b store x   Remainder operation    │
└─────────────────────────────────────────────────┘
```

---

## 🔍 Implementation Details

### Compiled Go Code

Simple programs compile to Go with these features:

**Built-in Helper Functions:**
- `clearScreen()` - Cross-platform screen clearing
- `toInt()` - Automatic type conversion to integers
- `contains()` - Check list/string containment
- `fetchUrl()` - HTTP GET requests
- `parseJson()` - JSON parsing

**Automatic Imports:**
- `fmt` - Formatting and printing
- `time` - Sleep/delay functionality
- `math/rand` - Random number generation
- `os` - File operations and exit
- `os/exec` - System command execution
- `strconv` - String conversions
- `strings` - String manipulation
- `net/http` - HTTP requests
- `io` - I/O operations
- `encoding/json` - JSON parsing

---

## 🤝 Contributing

Found a bug? Want a feature? Here's how you can help:

1. **Report Issues:** Describe the problem clearly
2. **Suggest Features:** Explain the use case
3. **Share Examples:** Show cool programs you've built
4. **Improve Docs:** Help make this guide better

---

## 🌟 Philosophy

> "Code should be written for humans first, computers second."

Simple Language embodies this principle. Every command reads like an instruction you'd give to another person. No cryptic symbols, no hidden complexity—just clear, expressive code.

### Design Goals

- **Readability:** Code reads like plain English
- **Learnability:** Zero programming experience needed
- **Practicality:** Real features for real programs
- **Simplicity:** One obvious way to do things

---

<div align="center">

**Made with ❤️ for simplicity**

*Start coding in minutes, not months!*

---

**Happy Coding! 🚀**

</div>
