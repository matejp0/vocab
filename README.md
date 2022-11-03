# Vocabulary exam preparer (vocab)
It picks random word from the list and prompts you to write the definition, that is on the other side of `=`. I find it a really effective method of preparing for vocabulary exams, because the ordering is always different.

## Installation
```
git clone https://github.com/matejp0/definitions
```
```
cd definitions
```
```
go build .
```
```
go install .
```
Now, you can edit the file `vocabulary.txt` or run it with a different file using the `-f` flag.

## Usage
```
Usage of vocab:
  -f string
    	Text file location (default "vocabulary.txt")
  -r	Reverse the questions and answers
```

### Text file formatting
```
question = answer
question=answer
```

