This little program reads the file "words.txt" and counts the occurrence of valid UTF8 words.
It displays the output ordered alphabetically.

To run it, you must have go installed on your pc and then build it:
go build WordCounter.go

To execute it:
go run WordCounter.go


Known issues:
-The program is expecting a valid UTF8 txt file.
-The program assumes that a file named "words.txt" exists on its root dir.