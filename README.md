# Exercise #1: Quiz Game

## Details

This exercise is broken into two parts to help simplify the process of explaining it as well as to make it easier to solve. The second part is harder than the first.

### Part 1

Quiz provided via a CSV file (more details below) will be given to a user, the quiz keeps track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to `problems.csv` (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

Assume quizzes will be relatively short (< 100 questions) and will have single word/number answers.

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

**NOTE:** *CSV files may have questions with commas in them. Eg: `"what 2+2, sir?",4` is a valid row in a CSV.*

### Part 2

Adapt part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

The quiz should stop as soon as the time limit has exceeded. That is, you shouldn't wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.