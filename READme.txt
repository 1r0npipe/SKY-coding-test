Hello
This is read me document, which describes how to use this
small application
There are two versions for amd64 architecture:
- windows
- linux
Make sure you set up correct permissions to run for lunux app:
# chmod +x loto_linux
and use:
# ./loto_linux --help
> loto_win.exe --help
You will see that output:
Usage of loto_win.exe:
  -fileResult string
        provide the file with lotto result (default "lottoResults.csv")
  -pNumbers string
        provide the number of player comma separated
After -fileResult put the name of CSV file in correct format(more info below),
path to it or just name if in current directory
After -pNumbers, put 6 numbers of user (person) number via comma with no space between
The output (standard for each OS) output will be chosen to provide the result

Format of CSV file:
 - please make sure you have all numbers over then 0 with comma separated of 6 in each line
 - make sure no any other characters at file, if you see that message:
    " error occurs when read file: can't read the CSV file, make sure the format is correct"
    check the format please
    BE AWARE, if you put any character close to number it will be transferred to 0!
 - if there are a couple of lines with guessed numbers less then 3, the only one will be chosen:
   line with max guessed numbers even if sum of them will be less then others,
   cause there is the line with most guessed numbers (and no condition for that)
   my assumption.
