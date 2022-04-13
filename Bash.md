<div align="center">
    <img src="https://bashlogo.com/img/symbol/svg/full_colored_dark.svg" alt="GnuBashLogo" height="256">
</div>

# [CHAPTER 1 - Bash Shell Scripting for Beginners 2019](https://youtu.be/uXohpTNNP8A)

### ls - list

`-a`, `--all` Show hidden files or directories with . and ..

`-A`, `--almost-all` Show hidden files and directories without . and ..

`-l`,`--long` Shows extra details

`-t`, `--timesort` Sort by creation time

`-X`, `--extensionsort` Sort by file extensions

`*.txt` `hello.*` Show files with wildcard patterns

### cd - change directory

`cd directory` Changes current working directory to the other

`.` Current directory

`..` Parent directory

### mkdir - make directory

`-p directory`, `--parents directory` Make parent directories if needed

### rm - remove

`-r directory`, `--recursive directory`   Recursively remove a directory and all its subdirectories

`-rf directory`, `--recursive --force directory` Forcibly remove a directory, without confirmations or errors

`file` Remove file  

### cp - copy

`source target`  Copies something to something

### mv - move

`source target` Move or rename a file

### pwd - present working directory

`pwd` Print the current directory

# [CHAPTER 2 - File Permissions](https://youtu.be/RLp3M0lXhSQ)

### File Permissions

| Letter | Stands for | Value |
| ------ | ---------- | ----- |
| r      | read       | 4     |
| w      | write      | 2     |
| x      | execute    | 1     |

### Combinations

| Letters | Value |
| ------- | ----- |
| rwx     | 7     |
| rw      | 6     |
| rx      | 5     |
| r       | 4     |
| wx      | 3     |
| w       | 2     |
| x       | 1     |
| -       | 0     |

**Do not remember these, please. Simple math will do the trick.**

### Permissions Groups

| What                     | File or Directory         | User Permissions | Group Permissions | Other Permissons |
|:------------------------ | ------------------------- | ---------------- | ----------------- | ---------------- |
| -rwxrwxrwx<br>drwxrwxrwx | `-` file<br>`d` directory | `rwx`------      | ---`rwx`---       | ------`rwx`      |

# [CHAPTER 3 - Linux Environment Variables](https://youtu.be/td_td5tiKtY)

`env` Show the environment

`printenv` Display key-value pairs of all environment variables

Set a custom environment variable:
`VARIABLE_NAME="PATH"`

For example:
`CONF="/home/konradkon/.config"`

# [CHAPTER 4 - Reading and Writing Files](https://youtu.be/jXVIeyliKp8)

### head

**Default**: prints out the first 10 lines

`-n 1 file` , `--lines 1 file` Prints out the first line

`-c 8 file`,  `--bytes 8 file` Prints out first eight bytes

### tail

**Default**: prints out the last 10 lines

`-n 1 file`, `--lines 1 file` Prints out the last line

### cat - conCATenate

`file` Prints out the text file

`files` Con**cat**anates files (merges them into 1 output)

`-n file`, `--number file` Prints out the text file with lines numbers even with empty lines

`-b file` , `--numbernonblank file` Prints out the text file with lines numbers without empty lines
`-E file`, `--show-ends file` Display $ at end of each line

### more

`file` Read a large file in a chunks

**Navigation**
`space` Next page

### less

`file` Read a large file in a chunks

**Navigation**

| Key          | Does                   |
| ------------ | ---------------------- |
| `space`      | Next page              |
| `b`          | Previous page          |
| `arrow keys` | Navigates through text |
| `q`          | Quit                   |

### grep

Searches for something in a file, output of the command or both
`-n word file`, `--line-number word file` Prints occurences with line number

`-c word file`, `--count word file` Counts occurences

`-ic word file` , `--ignore-case word file` Counts occurences ignoring the case

`-l word *.txt` `--files-with-matches word *.txt` Prints out the names of the txt files that this word occured

### touch

`file.txt` Creates a empty txt file

### echo

`echo "Hello World" > file.txt` **Writes** "Hello World" to a file.txt

`echo "Hello World" >> file.txt` **Appends** "Hello World" to a file.txt

`echo` **is less consistent than printf because it gives** `\n` **at the end of the file**

`printf` **never does \n by itself at the end**

# [CHAPTER 5 - Grabbing User Input With the Read Command](https://youtu.be/ukQI_J04qJo)

### Read one variable with input before

```bash
#!/bin/bash

read -p "Enter var: " variable
echo $var
```

Read variable with restricted time

```bash
#!/bin/bash

read -t 5 var # 5s
echo $var
```

### Read many variables

```bash
#!/bin/bash

read -p "Enter variables: " var1 var2 var3
echo $var1 $var2 $var3
```

### Read an entire array

```bash
#!/bin/bash

read -a names
echo ${names}[@]
```

### Read variables silently (without printing while typing)

```bash
#!/bin/bash

read -s var
echo ${names}[@]
```

# [CHAPTER 6 - Wildcards](https://youtu.be/EJjkvwelbpU)

`?` Single character wildcard

`*` Any characters wildcard

# [CHAPTER 7 - Using Pipes](https://youtu.be/FICwAKYc0Pg)

### | - pipe

Syntax: `output | pipe_command_to_filter`

### less with pipe

`-R` = `--RAW-CONTROL-CHARS`

`less -R .vimrc | grep syntax` Prints out lines with `syntax` in them

### sort

`sort path/to/file` Sort a file in ascending order

### uniq

`uniq path/to/file` Output the unique lines from the file

### sort + pipe + uniq

`sort file | uniq` Displays each line once

### Piping to Vim

`ls --help | vim -` Pipes ls --help to vim

# [CHAPTER 8 - How to do Math in Bash](https://youtu.be/3-DCmbO9VL8)

### expr - expression

| expr    | expression  |
| ------- | ----------- |
| `2 + 2` | adding      |
| `2 - 2` | subtracting |
| `2 / 2` | dividing    |
| `2 % 2` | modulo      |
| `2 * 2` | multiplying |

**Watch out for spaces!**

### echo way

```bash
#!/bin/bash

num1 = 50
num2 = 10

echo $(( num1 + num2 )) # adding
echo $(( num1 - num2 )) # subtracting
echo $(( num1 / num2 )) # dividing
echo $(( num1 % num2 )) # modulo
echo $(( num1 * num2 )) # multiplying
```

### bc - basic calculcation

`echo "20 + 0.4" | bc` For float variables pipe echo expression to bc 

`echo "scale=2;20 / 0.5" | bc -l` Prints out 40 (20/0.5) with 2 decimal places (`40.00`)

# [CHAPTER 9 - Logical And & Logical Or](https://youtu.be/EYbJuvEfi7M)

| Symbol | Name |
| ------ | ---- |
| &&     | AND  |
| \|\|   | OR   |

```bash
#!/bin/bash

true_var=true
false_var=false

if $true_var; then
    echo 'This variable is true'
elif $false_var
    echo 'Something'
else
    echo 'This variable is false'
fi
```

`[true] && [true]` Use square brackes when more than 1 comparisons

```bash
#!/bin/bash

number=20
if [ $number -gt 15 ] && [ $number -lt 30 ]
then
    echo "30>$number>15"
fi
```

# [CHAPTER 10 - Flow Control Using If, Elif, Else, and While](https://youtu.be/Q1gUq4UvnLE)

### Booleans

| Name  | Other Name | Number | State |
| ----- | ---------- | ------ | ----- |
| False | `false`    | 0      | off   |
| True  | `true`     | 1      | on    |

### Comparisons

| Syntax | Equavilent to | Meaning          |
| ------ | ------------- | ---------------- |
| -ge    | >=            | greater or equal |
| -gt    | >             | greater than     |
| -le    | <=            | less or equal    |
| -lt    | <             | less than        |
| -eq    | ==            | equal            |
| -ne    | !=            | not equal        |

### Examples:

```bash
#!/bin/bash

a=10
b=15

if [ "$a" -lt "$b" ]; then
    echo "$a < $b"
fi

if (( "$a" < "$b" )); then
    echo "$a < $b"
fi
```

```bash
#!/bin/bash

count=0
until [ $count -eq 10 ]
do
    echo "$count not equals 10"
    ((count++))
done

echo "$count equals 10"
```

# [CHAPTER 11 - How to use Functions Functions](https://youtu.be/ZXcbwEZV6bc)

### function function_name()

```bash
#!/bin/bash

function function_name()
{
    echo "Called me, sir?"
}

function_name
```

### function_name()

```bash
#!/bin/bash

function_name()
{
    echo "Called me, sir?"
}

function_name
```

# [CHAPTER 12 - Local and Global Variables](https://youtu.be/4GR0wum_pOQ)

**By default bash sets everything globally.**

**Local can be only used in a function**

```bash
#!/bin/bash

greet_user()
{
    local name=$1 # Takes first argument from the function
    echo "Hello $name how are you doing?"
}

greet_user "Konrad" # Only visible in the function
name="Kenny" # Visible globally
echo "Hello $name how are you doing?"
```

# [CHAPTER 13 - Case Statement](https://youtu.be/DA-Ilf15_r8)

### Switch case

```bash
#!/bin/bash

echo -n "Enter the name of a country: "
read COUNTRY

echo -n "The official language of $COUNTRY is "

case $COUNTRY in
 England | United_States | Canada) # | = or
 echo "English"
 ;;
Poland)
 echo "Polish"
;;

*)
 echo "Unknown"
 ;;

esac
```
