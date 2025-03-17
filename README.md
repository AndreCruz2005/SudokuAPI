# Sudoku API
## Routes


### **GET** /sudoku?hints
Returns a 9x9 2d array representing a sudoku board with the specified number of hints. 0 represents empty spaces.

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


