# Sudoku API
## Routes


### **GET** /sudoku?hints&boxes
Returns a 9x9 2d array representing a sudoku board. 0 represents empty spaces. 

#### PARAMETERS:
<b>hints (int):</b> Number of hints the sudoku board should have, i.e. non-empty spaces.

<b>boxes ('true' || 'false'):</b> Optional, by default 'false'. Whether each sub array in the board should represent a row or a box in the board. 

#### EXAMPLE REQUEST: 
```
/sudoku?hints=30
```

#### EXAMPLE RESPONSE:
```
[[0,0,0,0,0,0,0,9,0],
[0,2,0,1,7,9,8,0,0],
[9,0,4,0,6,8,0,2,0],
[5,0,6,0,4,0,0,7,0],
[2,3,0,0,0,1,0,0,0],
[0,0,9,7,0,0,3,0,0],
[7,0,0,0,0,3,0,1,0],
[0,0,1,0,0,5,0,0,7],
[0,0,0,4,0,7,0,0,9]]
```


