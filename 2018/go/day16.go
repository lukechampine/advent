package main

import (
	"fmt"

	"github.com/lukechampine/advent/utils"
)

const input = `Before: [3, 1, 2, 0]
5 1 2 0
After:  [0, 1, 2, 0]

Before: [3, 3, 0, 2]
10 2 0 1
After:  [3, 0, 0, 2]

Before: [3, 3, 1, 1]
7 3 3 3
After:  [3, 3, 1, 0]

Before: [1, 3, 2, 0]
8 0 2 2
After:  [1, 3, 0, 0]

Before: [3, 0, 0, 2]
12 2 3 2
After:  [3, 0, 1, 2]

Before: [0, 3, 1, 1]
1 2 3 1
After:  [0, 2, 1, 1]

Before: [1, 1, 0, 2]
2 0 2 2
After:  [1, 1, 0, 2]

Before: [3, 1, 3, 3]
0 1 2 0
After:  [0, 1, 3, 3]

Before: [3, 0, 2, 1]
10 1 0 3
After:  [3, 0, 2, 0]

Before: [2, 1, 3, 0]
14 0 3 0
After:  [1, 1, 3, 0]

Before: [1, 0, 0, 3]
2 0 2 1
After:  [1, 0, 0, 3]

Before: [2, 1, 3, 0]
0 1 2 3
After:  [2, 1, 3, 0]

Before: [2, 1, 0, 0]
14 0 3 0
After:  [1, 1, 0, 0]

Before: [2, 1, 2, 1]
4 3 1 2
After:  [2, 1, 0, 1]

Before: [1, 0, 1, 1]
7 3 3 1
After:  [1, 0, 1, 1]

Before: [0, 1, 0, 1]
7 3 3 2
After:  [0, 1, 0, 1]

Before: [0, 1, 3, 2]
0 1 2 1
After:  [0, 0, 3, 2]

Before: [1, 1, 2, 2]
8 0 2 3
After:  [1, 1, 2, 0]

Before: [2, 0, 2, 2]
10 1 0 0
After:  [0, 0, 2, 2]

Before: [2, 1, 3, 1]
0 1 2 3
After:  [2, 1, 3, 0]

Before: [1, 1, 2, 2]
8 0 2 1
After:  [1, 0, 2, 2]

Before: [0, 0, 2, 1]
11 3 2 1
After:  [0, 1, 2, 1]

Before: [0, 0, 2, 0]
3 2 2 0
After:  [4, 0, 2, 0]

Before: [1, 0, 0, 0]
2 0 2 3
After:  [1, 0, 0, 0]

Before: [2, 0, 0, 1]
4 0 1 1
After:  [2, 1, 0, 1]

Before: [2, 0, 3, 3]
10 1 0 3
After:  [2, 0, 3, 0]

Before: [0, 1, 3, 2]
0 1 2 0
After:  [0, 1, 3, 2]

Before: [2, 1, 3, 3]
0 1 2 3
After:  [2, 1, 3, 0]

Before: [0, 1, 2, 1]
11 3 2 3
After:  [0, 1, 2, 1]

Before: [0, 2, 2, 3]
13 0 2 1
After:  [0, 0, 2, 3]

Before: [0, 2, 3, 2]
13 0 1 3
After:  [0, 2, 3, 0]

Before: [1, 1, 2, 2]
4 3 2 1
After:  [1, 0, 2, 2]

Before: [1, 1, 2, 1]
11 3 2 2
After:  [1, 1, 1, 1]

Before: [0, 2, 1, 0]
13 0 1 1
After:  [0, 0, 1, 0]

Before: [2, 1, 2, 2]
9 2 0 0
After:  [1, 1, 2, 2]

Before: [1, 1, 1, 1]
1 2 0 3
After:  [1, 1, 1, 2]

Before: [3, 3, 0, 2]
12 2 3 2
After:  [3, 3, 1, 2]

Before: [0, 1, 2, 2]
3 3 2 1
After:  [0, 4, 2, 2]

Before: [0, 2, 2, 3]
13 0 3 0
After:  [0, 2, 2, 3]

Before: [1, 3, 1, 0]
1 2 0 0
After:  [2, 3, 1, 0]

Before: [0, 1, 1, 3]
15 1 3 3
After:  [0, 1, 1, 0]

Before: [3, 2, 3, 0]
12 3 2 3
After:  [3, 2, 3, 1]

Before: [1, 0, 1, 3]
1 2 0 3
After:  [1, 0, 1, 2]

Before: [3, 1, 2, 1]
11 3 2 2
After:  [3, 1, 1, 1]

Before: [0, 0, 0, 2]
6 0 0 3
After:  [0, 0, 0, 0]

Before: [1, 0, 0, 1]
2 0 2 3
After:  [1, 0, 0, 0]

Before: [0, 3, 0, 2]
6 0 0 0
After:  [0, 3, 0, 2]

Before: [1, 1, 3, 0]
0 1 2 0
After:  [0, 1, 3, 0]

Before: [3, 1, 0, 3]
15 1 3 0
After:  [0, 1, 0, 3]

Before: [3, 2, 0, 2]
10 2 0 1
After:  [3, 0, 0, 2]

Before: [0, 2, 0, 2]
6 0 0 1
After:  [0, 0, 0, 2]

Before: [0, 3, 3, 1]
6 0 0 1
After:  [0, 0, 3, 1]

Before: [0, 1, 2, 0]
3 2 2 3
After:  [0, 1, 2, 4]

Before: [3, 0, 0, 0]
4 0 2 1
After:  [3, 1, 0, 0]

Before: [0, 2, 2, 1]
11 3 2 3
After:  [0, 2, 2, 1]

Before: [2, 0, 3, 3]
9 3 2 0
After:  [1, 0, 3, 3]

Before: [2, 1, 2, 2]
5 1 2 2
After:  [2, 1, 0, 2]

Before: [1, 1, 2, 2]
5 1 2 2
After:  [1, 1, 0, 2]

Before: [0, 1, 0, 1]
6 0 0 3
After:  [0, 1, 0, 0]

Before: [1, 3, 2, 2]
3 2 2 1
After:  [1, 4, 2, 2]

Before: [3, 0, 2, 2]
3 2 2 2
After:  [3, 0, 4, 2]

Before: [2, 1, 2, 1]
3 2 2 0
After:  [4, 1, 2, 1]

Before: [1, 3, 2, 3]
8 0 2 2
After:  [1, 3, 0, 3]

Before: [2, 1, 1, 3]
4 2 1 0
After:  [0, 1, 1, 3]

Before: [2, 0, 1, 1]
1 2 3 0
After:  [2, 0, 1, 1]

Before: [3, 0, 0, 3]
10 1 0 0
After:  [0, 0, 0, 3]

Before: [1, 1, 3, 1]
0 1 2 1
After:  [1, 0, 3, 1]

Before: [0, 0, 1, 2]
6 0 0 2
After:  [0, 0, 0, 2]

Before: [3, 3, 0, 2]
10 2 0 2
After:  [3, 3, 0, 2]

Before: [2, 3, 2, 1]
11 3 2 0
After:  [1, 3, 2, 1]

Before: [1, 0, 0, 1]
2 0 2 0
After:  [0, 0, 0, 1]

Before: [0, 2, 2, 0]
6 0 0 3
After:  [0, 2, 2, 0]

Before: [1, 1, 2, 0]
8 0 2 0
After:  [0, 1, 2, 0]

Before: [2, 3, 2, 0]
9 2 0 2
After:  [2, 3, 1, 0]

Before: [1, 3, 0, 2]
2 0 2 0
After:  [0, 3, 0, 2]

Before: [3, 0, 0, 0]
10 1 0 3
After:  [3, 0, 0, 0]

Before: [0, 1, 0, 1]
6 0 0 0
After:  [0, 1, 0, 1]

Before: [0, 1, 3, 0]
12 3 2 2
After:  [0, 1, 1, 0]

Before: [2, 3, 1, 0]
14 0 3 2
After:  [2, 3, 1, 0]

Before: [0, 1, 2, 1]
11 3 2 1
After:  [0, 1, 2, 1]

Before: [3, 1, 1, 3]
4 2 1 0
After:  [0, 1, 1, 3]

Before: [0, 3, 0, 2]
6 0 0 2
After:  [0, 3, 0, 2]

Before: [0, 3, 2, 1]
11 3 2 3
After:  [0, 3, 2, 1]

Before: [2, 2, 2, 1]
3 1 2 3
After:  [2, 2, 2, 4]

Before: [2, 1, 0, 0]
14 0 3 1
After:  [2, 1, 0, 0]

Before: [0, 3, 3, 3]
9 3 2 2
After:  [0, 3, 1, 3]

Before: [0, 3, 3, 1]
13 0 1 2
After:  [0, 3, 0, 1]

Before: [0, 1, 1, 1]
13 0 3 1
After:  [0, 0, 1, 1]

Before: [2, 0, 2, 1]
3 2 2 2
After:  [2, 0, 4, 1]

Before: [0, 2, 2, 1]
9 2 1 1
After:  [0, 1, 2, 1]

Before: [2, 1, 0, 1]
7 3 3 3
After:  [2, 1, 0, 0]

Before: [0, 0, 1, 3]
13 0 3 0
After:  [0, 0, 1, 3]

Before: [3, 3, 1, 3]
15 2 3 3
After:  [3, 3, 1, 0]

Before: [0, 1, 3, 3]
0 1 2 0
After:  [0, 1, 3, 3]

Before: [3, 0, 3, 0]
12 3 2 0
After:  [1, 0, 3, 0]

Before: [1, 0, 2, 3]
8 0 2 1
After:  [1, 0, 2, 3]

Before: [2, 1, 3, 3]
0 1 2 1
After:  [2, 0, 3, 3]

Before: [2, 0, 3, 3]
4 2 0 1
After:  [2, 1, 3, 3]

Before: [0, 0, 2, 3]
13 0 3 1
After:  [0, 0, 2, 3]

Before: [1, 1, 3, 3]
9 3 2 0
After:  [1, 1, 3, 3]

Before: [0, 0, 0, 1]
6 0 0 2
After:  [0, 0, 0, 1]

Before: [0, 0, 3, 2]
13 0 2 3
After:  [0, 0, 3, 0]

Before: [1, 3, 2, 2]
8 0 2 1
After:  [1, 0, 2, 2]

Before: [0, 1, 0, 1]
13 0 1 2
After:  [0, 1, 0, 1]

Before: [2, 3, 0, 0]
14 0 3 2
After:  [2, 3, 1, 0]

Before: [2, 0, 2, 1]
7 3 3 1
After:  [2, 0, 2, 1]

Before: [3, 3, 2, 2]
7 3 3 2
After:  [3, 3, 0, 2]

Before: [0, 3, 1, 2]
6 0 0 3
After:  [0, 3, 1, 0]

Before: [1, 0, 1, 1]
1 2 0 2
After:  [1, 0, 2, 1]

Before: [2, 1, 1, 0]
14 0 3 1
After:  [2, 1, 1, 0]

Before: [3, 3, 3, 3]
9 3 0 3
After:  [3, 3, 3, 1]

Before: [1, 1, 1, 1]
1 2 3 0
After:  [2, 1, 1, 1]

Before: [3, 1, 3, 1]
0 1 2 2
After:  [3, 1, 0, 1]

Before: [1, 1, 2, 0]
5 1 2 2
After:  [1, 1, 0, 0]

Before: [1, 1, 2, 3]
5 1 2 2
After:  [1, 1, 0, 3]

Before: [2, 0, 0, 3]
10 1 0 1
After:  [2, 0, 0, 3]

Before: [1, 2, 0, 1]
7 3 3 3
After:  [1, 2, 0, 0]

Before: [0, 3, 1, 3]
6 0 0 1
After:  [0, 0, 1, 3]

Before: [2, 2, 3, 1]
4 2 0 3
After:  [2, 2, 3, 1]

Before: [3, 0, 2, 1]
11 3 2 2
After:  [3, 0, 1, 1]

Before: [0, 1, 2, 0]
5 1 2 2
After:  [0, 1, 0, 0]

Before: [2, 3, 1, 1]
1 2 3 2
After:  [2, 3, 2, 1]

Before: [0, 1, 1, 2]
13 0 3 3
After:  [0, 1, 1, 0]

Before: [3, 1, 2, 3]
5 1 2 0
After:  [0, 1, 2, 3]

Before: [1, 3, 3, 0]
12 3 2 0
After:  [1, 3, 3, 0]

Before: [2, 3, 0, 2]
12 2 3 2
After:  [2, 3, 1, 2]

Before: [3, 1, 0, 1]
10 2 0 1
After:  [3, 0, 0, 1]

Before: [0, 2, 0, 0]
13 0 1 0
After:  [0, 2, 0, 0]

Before: [0, 2, 2, 2]
13 0 1 2
After:  [0, 2, 0, 2]

Before: [3, 2, 2, 3]
15 2 3 2
After:  [3, 2, 0, 3]

Before: [2, 0, 2, 0]
14 0 3 1
After:  [2, 1, 2, 0]

Before: [2, 2, 0, 0]
14 0 3 2
After:  [2, 2, 1, 0]

Before: [1, 3, 2, 0]
8 0 2 0
After:  [0, 3, 2, 0]

Before: [3, 2, 3, 1]
9 2 3 1
After:  [3, 0, 3, 1]

Before: [0, 1, 1, 0]
4 2 1 1
After:  [0, 0, 1, 0]

Before: [0, 3, 2, 3]
13 0 3 0
After:  [0, 3, 2, 3]

Before: [3, 1, 0, 1]
7 3 3 3
After:  [3, 1, 0, 0]

Before: [1, 3, 3, 3]
9 3 2 3
After:  [1, 3, 3, 1]

Before: [2, 3, 3, 1]
4 2 0 2
After:  [2, 3, 1, 1]

Before: [0, 3, 3, 1]
7 3 3 0
After:  [0, 3, 3, 1]

Before: [1, 2, 1, 2]
1 2 0 1
After:  [1, 2, 1, 2]

Before: [0, 2, 2, 2]
13 0 2 2
After:  [0, 2, 0, 2]

Before: [1, 1, 3, 3]
0 1 2 1
After:  [1, 0, 3, 3]

Before: [2, 0, 3, 2]
4 0 1 0
After:  [1, 0, 3, 2]

Before: [3, 1, 3, 0]
12 3 2 2
After:  [3, 1, 1, 0]

Before: [0, 2, 1, 1]
1 2 3 3
After:  [0, 2, 1, 2]

Before: [2, 3, 1, 0]
14 0 3 3
After:  [2, 3, 1, 1]

Before: [2, 0, 1, 0]
14 0 3 0
After:  [1, 0, 1, 0]

Before: [0, 3, 1, 1]
7 3 3 3
After:  [0, 3, 1, 0]

Before: [2, 3, 3, 3]
9 3 2 3
After:  [2, 3, 3, 1]

Before: [2, 1, 2, 3]
15 2 3 1
After:  [2, 0, 2, 3]

Before: [2, 1, 3, 3]
0 1 2 2
After:  [2, 1, 0, 3]

Before: [3, 1, 2, 1]
7 3 3 0
After:  [0, 1, 2, 1]

Before: [1, 1, 0, 1]
7 3 3 1
After:  [1, 0, 0, 1]

Before: [2, 1, 2, 1]
5 1 2 2
After:  [2, 1, 0, 1]

Before: [3, 1, 3, 2]
0 1 2 1
After:  [3, 0, 3, 2]

Before: [2, 1, 1, 3]
15 1 3 3
After:  [2, 1, 1, 0]

Before: [3, 2, 0, 1]
10 2 0 3
After:  [3, 2, 0, 0]

Before: [1, 3, 0, 1]
2 0 2 2
After:  [1, 3, 0, 1]

Before: [1, 1, 3, 2]
0 1 2 2
After:  [1, 1, 0, 2]

Before: [1, 2, 2, 2]
7 3 3 0
After:  [0, 2, 2, 2]

Before: [1, 2, 2, 2]
8 0 2 0
After:  [0, 2, 2, 2]

Before: [3, 1, 0, 3]
9 3 0 3
After:  [3, 1, 0, 1]

Before: [1, 1, 0, 1]
2 0 2 2
After:  [1, 1, 0, 1]

Before: [2, 3, 3, 0]
14 0 3 0
After:  [1, 3, 3, 0]

Before: [3, 0, 2, 1]
11 3 2 1
After:  [3, 1, 2, 1]

Before: [2, 3, 3, 0]
14 0 3 3
After:  [2, 3, 3, 1]

Before: [0, 1, 2, 2]
5 1 2 2
After:  [0, 1, 0, 2]

Before: [3, 0, 2, 3]
10 1 0 3
After:  [3, 0, 2, 0]

Before: [1, 1, 2, 3]
8 0 2 2
After:  [1, 1, 0, 3]

Before: [0, 0, 0, 3]
6 0 0 1
After:  [0, 0, 0, 3]

Before: [1, 3, 2, 1]
11 3 2 2
After:  [1, 3, 1, 1]

Before: [0, 3, 1, 3]
13 0 3 1
After:  [0, 0, 1, 3]

Before: [2, 0, 3, 0]
14 0 3 1
After:  [2, 1, 3, 0]

Before: [0, 2, 0, 3]
13 0 3 3
After:  [0, 2, 0, 0]

Before: [1, 0, 2, 0]
8 0 2 2
After:  [1, 0, 0, 0]

Before: [1, 3, 2, 0]
8 0 2 1
After:  [1, 0, 2, 0]

Before: [1, 0, 2, 1]
11 3 2 1
After:  [1, 1, 2, 1]

Before: [1, 0, 0, 0]
2 0 2 1
After:  [1, 0, 0, 0]

Before: [2, 2, 2, 1]
11 3 2 1
After:  [2, 1, 2, 1]

Before: [3, 3, 0, 2]
7 3 3 1
After:  [3, 0, 0, 2]

Before: [2, 1, 3, 3]
15 1 3 0
After:  [0, 1, 3, 3]

Before: [0, 3, 3, 2]
13 0 1 1
After:  [0, 0, 3, 2]

Before: [0, 3, 3, 3]
6 0 0 1
After:  [0, 0, 3, 3]

Before: [1, 3, 2, 2]
8 0 2 2
After:  [1, 3, 0, 2]

Before: [1, 0, 0, 3]
2 0 2 3
After:  [1, 0, 0, 0]

Before: [2, 3, 2, 1]
9 2 0 2
After:  [2, 3, 1, 1]

Before: [1, 3, 1, 2]
1 2 0 0
After:  [2, 3, 1, 2]

Before: [2, 1, 3, 0]
0 1 2 0
After:  [0, 1, 3, 0]

Before: [3, 2, 2, 1]
9 2 1 2
After:  [3, 2, 1, 1]

Before: [2, 1, 3, 1]
0 1 2 2
After:  [2, 1, 0, 1]

Before: [0, 2, 1, 0]
6 0 0 1
After:  [0, 0, 1, 0]

Before: [1, 1, 0, 3]
15 1 3 1
After:  [1, 0, 0, 3]

Before: [0, 3, 3, 3]
9 3 2 3
After:  [0, 3, 3, 1]

Before: [3, 1, 3, 1]
7 3 3 3
After:  [3, 1, 3, 0]

Before: [2, 2, 2, 1]
11 3 2 3
After:  [2, 2, 2, 1]

Before: [3, 2, 2, 3]
9 3 0 1
After:  [3, 1, 2, 3]

Before: [3, 2, 2, 0]
3 1 2 1
After:  [3, 4, 2, 0]

Before: [2, 3, 2, 0]
14 0 3 1
After:  [2, 1, 2, 0]

Before: [0, 1, 3, 0]
0 1 2 2
After:  [0, 1, 0, 0]

Before: [2, 2, 3, 0]
4 2 0 3
After:  [2, 2, 3, 1]

Before: [2, 0, 0, 0]
14 0 3 1
After:  [2, 1, 0, 0]

Before: [1, 1, 0, 0]
2 0 2 0
After:  [0, 1, 0, 0]

Before: [0, 2, 3, 2]
6 0 0 1
After:  [0, 0, 3, 2]

Before: [2, 3, 2, 3]
3 2 2 0
After:  [4, 3, 2, 3]

Before: [1, 2, 2, 1]
3 2 2 0
After:  [4, 2, 2, 1]

Before: [2, 2, 2, 0]
14 0 3 0
After:  [1, 2, 2, 0]

Before: [1, 1, 2, 3]
5 1 2 0
After:  [0, 1, 2, 3]

Before: [3, 1, 3, 3]
9 3 2 0
After:  [1, 1, 3, 3]

Before: [1, 1, 0, 0]
2 0 2 2
After:  [1, 1, 0, 0]

Before: [1, 1, 2, 3]
8 0 2 3
After:  [1, 1, 2, 0]

Before: [3, 3, 2, 3]
3 2 2 1
After:  [3, 4, 2, 3]

Before: [3, 3, 1, 3]
15 2 3 2
After:  [3, 3, 0, 3]

Before: [0, 1, 0, 2]
6 0 0 1
After:  [0, 0, 0, 2]

Before: [1, 1, 3, 0]
12 3 2 0
After:  [1, 1, 3, 0]

Before: [2, 1, 2, 2]
3 3 2 2
After:  [2, 1, 4, 2]

Before: [1, 1, 3, 2]
0 1 2 1
After:  [1, 0, 3, 2]

Before: [3, 1, 2, 3]
5 1 2 1
After:  [3, 0, 2, 3]

Before: [0, 1, 3, 1]
0 1 2 1
After:  [0, 0, 3, 1]

Before: [3, 2, 2, 1]
7 3 3 1
After:  [3, 0, 2, 1]

Before: [0, 1, 2, 3]
5 1 2 2
After:  [0, 1, 0, 3]

Before: [1, 1, 0, 3]
2 0 2 1
After:  [1, 0, 0, 3]

Before: [2, 0, 2, 0]
14 0 3 0
After:  [1, 0, 2, 0]

Before: [0, 2, 1, 0]
6 0 0 2
After:  [0, 2, 0, 0]

Before: [0, 2, 0, 1]
13 0 1 0
After:  [0, 2, 0, 1]

Before: [2, 3, 0, 0]
14 0 3 0
After:  [1, 3, 0, 0]

Before: [3, 1, 1, 3]
15 1 3 3
After:  [3, 1, 1, 0]

Before: [1, 2, 1, 3]
15 2 3 3
After:  [1, 2, 1, 0]

Before: [3, 1, 2, 1]
5 1 2 1
After:  [3, 0, 2, 1]

Before: [1, 1, 0, 1]
2 0 2 1
After:  [1, 0, 0, 1]

Before: [1, 2, 2, 3]
8 0 2 0
After:  [0, 2, 2, 3]

Before: [0, 2, 2, 3]
15 2 3 2
After:  [0, 2, 0, 3]

Before: [1, 0, 2, 0]
8 0 2 0
After:  [0, 0, 2, 0]

Before: [0, 2, 2, 3]
3 2 2 0
After:  [4, 2, 2, 3]

Before: [3, 3, 2, 1]
11 3 2 2
After:  [3, 3, 1, 1]

Before: [0, 2, 2, 2]
6 0 0 2
After:  [0, 2, 0, 2]

Before: [1, 1, 2, 1]
5 1 2 3
After:  [1, 1, 2, 0]

Before: [2, 0, 2, 1]
11 3 2 1
After:  [2, 1, 2, 1]

Before: [0, 1, 2, 1]
5 1 2 3
After:  [0, 1, 2, 0]

Before: [1, 2, 2, 2]
3 1 2 0
After:  [4, 2, 2, 2]

Before: [1, 3, 2, 1]
11 3 2 3
After:  [1, 3, 2, 1]

Before: [0, 1, 1, 2]
13 0 1 2
After:  [0, 1, 0, 2]

Before: [1, 2, 2, 0]
3 1 2 1
After:  [1, 4, 2, 0]

Before: [1, 0, 0, 0]
2 0 2 0
After:  [0, 0, 0, 0]

Before: [1, 2, 1, 1]
1 2 0 3
After:  [1, 2, 1, 2]

Before: [3, 3, 2, 1]
11 3 2 1
After:  [3, 1, 2, 1]

Before: [0, 0, 2, 1]
11 3 2 0
After:  [1, 0, 2, 1]

Before: [1, 1, 1, 3]
4 2 1 3
After:  [1, 1, 1, 0]

Before: [2, 0, 2, 2]
7 3 3 1
After:  [2, 0, 2, 2]

Before: [0, 1, 2, 1]
5 1 2 0
After:  [0, 1, 2, 1]

Before: [3, 2, 0, 3]
10 2 0 0
After:  [0, 2, 0, 3]

Before: [1, 1, 2, 0]
5 1 2 1
After:  [1, 0, 2, 0]

Before: [2, 0, 3, 1]
9 2 3 2
After:  [2, 0, 0, 1]

Before: [1, 0, 2, 1]
8 0 2 1
After:  [1, 0, 2, 1]

Before: [0, 2, 0, 3]
15 1 3 3
After:  [0, 2, 0, 0]

Before: [0, 2, 2, 0]
6 0 0 0
After:  [0, 2, 2, 0]

Before: [0, 0, 1, 1]
1 2 3 2
After:  [0, 0, 2, 1]

Before: [1, 0, 1, 3]
15 2 3 3
After:  [1, 0, 1, 0]

Before: [2, 1, 2, 3]
15 1 3 2
After:  [2, 1, 0, 3]

Before: [1, 2, 2, 3]
8 0 2 2
After:  [1, 2, 0, 3]

Before: [2, 2, 3, 0]
12 3 2 2
After:  [2, 2, 1, 0]

Before: [3, 1, 3, 1]
9 2 3 3
After:  [3, 1, 3, 0]

Before: [3, 0, 3, 1]
7 3 3 1
After:  [3, 0, 3, 1]

Before: [0, 3, 0, 2]
6 0 0 3
After:  [0, 3, 0, 0]

Before: [1, 1, 2, 1]
5 1 2 0
After:  [0, 1, 2, 1]

Before: [1, 3, 3, 0]
12 3 2 3
After:  [1, 3, 3, 1]

Before: [3, 0, 1, 1]
1 2 3 0
After:  [2, 0, 1, 1]

Before: [2, 1, 2, 0]
14 0 3 3
After:  [2, 1, 2, 1]

Before: [2, 3, 2, 1]
11 3 2 2
After:  [2, 3, 1, 1]

Before: [0, 3, 3, 0]
6 0 0 3
After:  [0, 3, 3, 0]

Before: [0, 1, 3, 0]
12 3 2 3
After:  [0, 1, 3, 1]

Before: [1, 2, 3, 0]
12 3 2 3
After:  [1, 2, 3, 1]

Before: [2, 1, 0, 2]
12 2 3 3
After:  [2, 1, 0, 1]

Before: [0, 3, 0, 2]
7 3 3 3
After:  [0, 3, 0, 0]

Before: [3, 0, 3, 0]
12 3 2 2
After:  [3, 0, 1, 0]

Before: [2, 2, 0, 0]
14 0 3 3
After:  [2, 2, 0, 1]

Before: [1, 2, 0, 1]
2 0 2 1
After:  [1, 0, 0, 1]

Before: [0, 2, 3, 1]
9 2 3 2
After:  [0, 2, 0, 1]

Before: [2, 1, 2, 1]
11 3 2 1
After:  [2, 1, 2, 1]

Before: [3, 1, 3, 3]
0 1 2 3
After:  [3, 1, 3, 0]

Before: [0, 1, 2, 3]
5 1 2 3
After:  [0, 1, 2, 0]

Before: [1, 3, 1, 0]
1 2 0 2
After:  [1, 3, 2, 0]

Before: [0, 3, 3, 2]
13 0 1 3
After:  [0, 3, 3, 0]

Before: [1, 2, 0, 2]
2 0 2 0
After:  [0, 2, 0, 2]

Before: [2, 2, 0, 0]
14 0 3 0
After:  [1, 2, 0, 0]

Before: [3, 2, 0, 0]
10 2 0 0
After:  [0, 2, 0, 0]

Before: [1, 2, 2, 0]
8 0 2 0
After:  [0, 2, 2, 0]

Before: [2, 0, 2, 2]
3 3 2 1
After:  [2, 4, 2, 2]

Before: [1, 3, 0, 2]
12 2 3 1
After:  [1, 1, 0, 2]

Before: [0, 1, 1, 2]
7 3 3 3
After:  [0, 1, 1, 0]

Before: [0, 0, 1, 2]
7 3 3 1
After:  [0, 0, 1, 2]

Before: [0, 1, 1, 0]
6 0 0 1
After:  [0, 0, 1, 0]

Before: [2, 0, 2, 1]
11 3 2 0
After:  [1, 0, 2, 1]

Before: [2, 2, 3, 3]
15 1 3 1
After:  [2, 0, 3, 3]

Before: [0, 3, 2, 3]
13 0 2 3
After:  [0, 3, 2, 0]

Before: [0, 1, 2, 2]
13 0 3 3
After:  [0, 1, 2, 0]

Before: [2, 1, 2, 0]
5 1 2 0
After:  [0, 1, 2, 0]

Before: [3, 3, 0, 2]
7 3 3 3
After:  [3, 3, 0, 0]

Before: [0, 0, 1, 2]
13 0 3 2
After:  [0, 0, 0, 2]

Before: [1, 3, 1, 2]
1 2 0 3
After:  [1, 3, 1, 2]

Before: [0, 1, 3, 1]
13 0 3 1
After:  [0, 0, 3, 1]

Before: [2, 1, 3, 1]
4 2 0 0
After:  [1, 1, 3, 1]

Before: [1, 1, 2, 2]
8 0 2 2
After:  [1, 1, 0, 2]

Before: [3, 0, 0, 2]
10 1 0 3
After:  [3, 0, 0, 0]

Before: [0, 3, 1, 2]
7 3 3 2
After:  [0, 3, 0, 2]

Before: [3, 2, 2, 2]
9 2 1 2
After:  [3, 2, 1, 2]

Before: [0, 3, 2, 2]
4 3 2 1
After:  [0, 0, 2, 2]

Before: [0, 1, 2, 3]
5 1 2 0
After:  [0, 1, 2, 3]

Before: [1, 2, 0, 3]
2 0 2 3
After:  [1, 2, 0, 0]

Before: [0, 3, 2, 3]
13 0 2 0
After:  [0, 3, 2, 3]

Before: [1, 1, 2, 0]
8 0 2 2
After:  [1, 1, 0, 0]

Before: [1, 2, 0, 0]
2 0 2 2
After:  [1, 2, 0, 0]

Before: [0, 0, 2, 1]
11 3 2 2
After:  [0, 0, 1, 1]

Before: [3, 1, 3, 2]
0 1 2 2
After:  [3, 1, 0, 2]

Before: [1, 1, 1, 1]
4 3 1 1
After:  [1, 0, 1, 1]

Before: [0, 3, 3, 1]
13 0 3 1
After:  [0, 0, 3, 1]

Before: [1, 1, 3, 0]
12 3 2 1
After:  [1, 1, 3, 0]

Before: [2, 0, 0, 0]
14 0 3 0
After:  [1, 0, 0, 0]

Before: [0, 1, 3, 2]
0 1 2 3
After:  [0, 1, 3, 0]

Before: [1, 1, 2, 0]
8 0 2 3
After:  [1, 1, 2, 0]

Before: [3, 0, 2, 3]
15 2 3 2
After:  [3, 0, 0, 3]

Before: [1, 1, 1, 2]
1 2 0 0
After:  [2, 1, 1, 2]

Before: [3, 3, 2, 2]
3 2 2 2
After:  [3, 3, 4, 2]

Before: [0, 1, 0, 2]
7 3 3 2
After:  [0, 1, 0, 2]

Before: [0, 2, 2, 2]
13 0 2 0
After:  [0, 2, 2, 2]

Before: [3, 1, 0, 0]
10 2 0 3
After:  [3, 1, 0, 0]

Before: [2, 1, 3, 0]
0 1 2 1
After:  [2, 0, 3, 0]

Before: [1, 0, 2, 2]
8 0 2 1
After:  [1, 0, 2, 2]

Before: [2, 0, 1, 2]
10 1 0 1
After:  [2, 0, 1, 2]

Before: [2, 3, 2, 1]
11 3 2 1
After:  [2, 1, 2, 1]

Before: [2, 1, 1, 1]
4 2 1 3
After:  [2, 1, 1, 0]

Before: [2, 0, 2, 1]
11 3 2 2
After:  [2, 0, 1, 1]

Before: [2, 1, 3, 2]
0 1 2 1
After:  [2, 0, 3, 2]

Before: [2, 1, 2, 0]
5 1 2 3
After:  [2, 1, 2, 0]

Before: [0, 3, 2, 3]
13 0 1 3
After:  [0, 3, 2, 0]

Before: [1, 1, 2, 2]
5 1 2 0
After:  [0, 1, 2, 2]

Before: [1, 3, 1, 1]
1 2 3 3
After:  [1, 3, 1, 2]

Before: [2, 0, 2, 0]
14 0 3 2
After:  [2, 0, 1, 0]

Before: [3, 2, 2, 1]
11 3 2 1
After:  [3, 1, 2, 1]

Before: [1, 1, 3, 2]
0 1 2 0
After:  [0, 1, 3, 2]

Before: [0, 2, 2, 1]
11 3 2 1
After:  [0, 1, 2, 1]

Before: [2, 0, 0, 3]
10 1 0 3
After:  [2, 0, 0, 0]

Before: [2, 1, 0, 3]
15 1 3 1
After:  [2, 0, 0, 3]

Before: [3, 0, 2, 2]
4 3 2 1
After:  [3, 0, 2, 2]

Before: [3, 2, 2, 1]
11 3 2 2
After:  [3, 2, 1, 1]

Before: [2, 0, 0, 2]
12 2 3 2
After:  [2, 0, 1, 2]

Before: [0, 1, 2, 2]
6 0 0 0
After:  [0, 1, 2, 2]

Before: [2, 2, 2, 3]
3 2 2 1
After:  [2, 4, 2, 3]

Before: [2, 1, 2, 3]
5 1 2 3
After:  [2, 1, 2, 0]

Before: [0, 3, 3, 2]
6 0 0 0
After:  [0, 3, 3, 2]

Before: [3, 0, 3, 1]
7 3 3 0
After:  [0, 0, 3, 1]

Before: [0, 1, 1, 3]
15 1 3 1
After:  [0, 0, 1, 3]

Before: [0, 1, 3, 3]
15 1 3 3
After:  [0, 1, 3, 0]

Before: [2, 3, 2, 2]
4 3 2 0
After:  [0, 3, 2, 2]

Before: [0, 0, 2, 2]
7 3 3 2
After:  [0, 0, 0, 2]

Before: [2, 1, 1, 2]
4 2 1 1
After:  [2, 0, 1, 2]

Before: [1, 2, 0, 3]
2 0 2 0
After:  [0, 2, 0, 3]

Before: [3, 1, 2, 3]
9 3 0 3
After:  [3, 1, 2, 1]

Before: [0, 0, 0, 3]
13 0 3 0
After:  [0, 0, 0, 3]

Before: [1, 2, 2, 1]
8 0 2 2
After:  [1, 2, 0, 1]

Before: [3, 0, 1, 3]
9 3 0 3
After:  [3, 0, 1, 1]

Before: [1, 1, 2, 2]
5 1 2 1
After:  [1, 0, 2, 2]

Before: [1, 0, 2, 1]
11 3 2 0
After:  [1, 0, 2, 1]

Before: [0, 2, 3, 0]
12 3 2 3
After:  [0, 2, 3, 1]

Before: [3, 1, 3, 3]
0 1 2 1
After:  [3, 0, 3, 3]

Before: [3, 1, 0, 2]
7 3 3 0
After:  [0, 1, 0, 2]

Before: [1, 2, 0, 1]
2 0 2 3
After:  [1, 2, 0, 0]

Before: [1, 0, 2, 1]
8 0 2 3
After:  [1, 0, 2, 0]

Before: [2, 0, 2, 0]
10 1 0 1
After:  [2, 0, 2, 0]

Before: [3, 0, 3, 1]
10 1 0 2
After:  [3, 0, 0, 1]

Before: [3, 0, 3, 3]
9 3 0 1
After:  [3, 1, 3, 3]

Before: [0, 1, 3, 1]
13 0 2 1
After:  [0, 0, 3, 1]

Before: [0, 1, 1, 0]
6 0 0 3
After:  [0, 1, 1, 0]

Before: [0, 3, 3, 1]
13 0 3 3
After:  [0, 3, 3, 0]

Before: [1, 3, 0, 3]
2 0 2 0
After:  [0, 3, 0, 3]

Before: [2, 3, 3, 0]
12 3 2 2
After:  [2, 3, 1, 0]

Before: [3, 2, 1, 3]
15 1 3 0
After:  [0, 2, 1, 3]

Before: [1, 3, 2, 1]
8 0 2 2
After:  [1, 3, 0, 1]

Before: [3, 0, 3, 2]
10 1 0 1
After:  [3, 0, 3, 2]

Before: [3, 0, 2, 0]
10 1 0 2
After:  [3, 0, 0, 0]

Before: [2, 2, 3, 0]
14 0 3 1
After:  [2, 1, 3, 0]

Before: [3, 1, 2, 2]
5 1 2 1
After:  [3, 0, 2, 2]

Before: [1, 1, 2, 1]
5 1 2 1
After:  [1, 0, 2, 1]

Before: [1, 2, 2, 1]
11 3 2 2
After:  [1, 2, 1, 1]

Before: [3, 3, 2, 1]
11 3 2 0
After:  [1, 3, 2, 1]

Before: [0, 2, 0, 1]
6 0 0 3
After:  [0, 2, 0, 0]

Before: [2, 1, 2, 1]
11 3 2 2
After:  [2, 1, 1, 1]

Before: [1, 1, 3, 0]
0 1 2 1
After:  [1, 0, 3, 0]

Before: [3, 1, 2, 0]
5 1 2 2
After:  [3, 1, 0, 0]

Before: [0, 1, 2, 1]
5 1 2 2
After:  [0, 1, 0, 1]

Before: [0, 2, 3, 1]
13 0 1 1
After:  [0, 0, 3, 1]

Before: [1, 2, 3, 0]
12 3 2 1
After:  [1, 1, 3, 0]

Before: [1, 0, 3, 0]
12 3 2 0
After:  [1, 0, 3, 0]

Before: [2, 3, 0, 0]
14 0 3 1
After:  [2, 1, 0, 0]

Before: [0, 3, 1, 1]
7 3 3 1
After:  [0, 0, 1, 1]

Before: [3, 1, 3, 2]
0 1 2 3
After:  [3, 1, 3, 0]

Before: [0, 1, 3, 1]
9 2 3 1
After:  [0, 0, 3, 1]

Before: [3, 3, 0, 2]
12 2 3 3
After:  [3, 3, 0, 1]

Before: [1, 3, 2, 1]
11 3 2 0
After:  [1, 3, 2, 1]

Before: [0, 3, 1, 1]
7 3 3 0
After:  [0, 3, 1, 1]

Before: [1, 3, 3, 1]
7 3 3 3
After:  [1, 3, 3, 0]

Before: [2, 1, 2, 2]
5 1 2 3
After:  [2, 1, 2, 0]

Before: [2, 1, 3, 1]
9 2 3 1
After:  [2, 0, 3, 1]

Before: [0, 2, 2, 1]
13 0 1 0
After:  [0, 2, 2, 1]

Before: [0, 1, 2, 1]
11 3 2 0
After:  [1, 1, 2, 1]

Before: [0, 1, 3, 3]
0 1 2 3
After:  [0, 1, 3, 0]

Before: [1, 1, 3, 3]
0 1 2 0
After:  [0, 1, 3, 3]

Before: [1, 0, 1, 2]
1 2 0 1
After:  [1, 2, 1, 2]

Before: [1, 0, 1, 3]
1 2 0 2
After:  [1, 0, 2, 3]

Before: [2, 1, 2, 0]
3 0 2 3
After:  [2, 1, 2, 4]

Before: [3, 0, 2, 2]
7 3 3 0
After:  [0, 0, 2, 2]

Before: [3, 0, 0, 2]
10 1 0 0
After:  [0, 0, 0, 2]

Before: [2, 1, 3, 2]
7 3 3 3
After:  [2, 1, 3, 0]

Before: [2, 0, 2, 1]
4 0 1 1
After:  [2, 1, 2, 1]

Before: [3, 3, 3, 0]
12 3 2 1
After:  [3, 1, 3, 0]

Before: [1, 0, 2, 3]
8 0 2 0
After:  [0, 0, 2, 3]

Before: [0, 1, 1, 3]
15 1 3 0
After:  [0, 1, 1, 3]

Before: [2, 0, 3, 3]
9 3 2 2
After:  [2, 0, 1, 3]

Before: [2, 3, 0, 0]
14 0 3 3
After:  [2, 3, 0, 1]

Before: [2, 3, 2, 1]
11 3 2 3
After:  [2, 3, 2, 1]

Before: [2, 0, 0, 3]
4 0 1 2
After:  [2, 0, 1, 3]

Before: [0, 1, 1, 2]
6 0 0 0
After:  [0, 1, 1, 2]

Before: [1, 0, 3, 0]
12 3 2 1
After:  [1, 1, 3, 0]

Before: [2, 1, 1, 1]
1 2 3 2
After:  [2, 1, 2, 1]

Before: [1, 3, 2, 2]
3 2 2 3
After:  [1, 3, 2, 4]

Before: [0, 1, 1, 1]
1 2 3 1
After:  [0, 2, 1, 1]

Before: [2, 1, 3, 2]
0 1 2 0
After:  [0, 1, 3, 2]

Before: [1, 0, 2, 0]
8 0 2 3
After:  [1, 0, 2, 0]

Before: [0, 3, 2, 3]
3 2 2 2
After:  [0, 3, 4, 3]

Before: [1, 2, 1, 1]
1 2 3 3
After:  [1, 2, 1, 2]

Before: [2, 3, 2, 1]
3 2 2 3
After:  [2, 3, 2, 4]

Before: [3, 1, 1, 1]
1 2 3 2
After:  [3, 1, 2, 1]

Before: [2, 3, 2, 0]
9 2 0 1
After:  [2, 1, 2, 0]

Before: [2, 0, 2, 1]
11 3 2 3
After:  [2, 0, 2, 1]

Before: [1, 1, 2, 0]
3 2 2 2
After:  [1, 1, 4, 0]

Before: [2, 2, 2, 1]
11 3 2 0
After:  [1, 2, 2, 1]

Before: [1, 0, 0, 2]
2 0 2 0
After:  [0, 0, 0, 2]

Before: [1, 0, 3, 0]
12 3 2 3
After:  [1, 0, 3, 1]

Before: [3, 1, 0, 3]
4 0 2 3
After:  [3, 1, 0, 1]

Before: [3, 2, 2, 1]
3 2 2 0
After:  [4, 2, 2, 1]

Before: [3, 1, 2, 3]
5 1 2 3
After:  [3, 1, 2, 0]

Before: [3, 1, 2, 2]
3 2 2 2
After:  [3, 1, 4, 2]

Before: [3, 0, 1, 1]
7 3 3 0
After:  [0, 0, 1, 1]

Before: [1, 0, 3, 0]
12 3 2 2
After:  [1, 0, 1, 0]

Before: [2, 3, 3, 0]
4 2 0 0
After:  [1, 3, 3, 0]

Before: [2, 2, 1, 0]
14 0 3 3
After:  [2, 2, 1, 1]

Before: [1, 2, 0, 3]
2 0 2 1
After:  [1, 0, 0, 3]

Before: [0, 2, 2, 1]
9 2 1 0
After:  [1, 2, 2, 1]

Before: [3, 1, 3, 1]
0 1 2 0
After:  [0, 1, 3, 1]

Before: [0, 1, 2, 3]
6 0 0 0
After:  [0, 1, 2, 3]

Before: [3, 0, 1, 2]
10 1 0 1
After:  [3, 0, 1, 2]

Before: [2, 2, 2, 1]
7 3 3 3
After:  [2, 2, 2, 0]

Before: [3, 0, 1, 3]
10 1 0 0
After:  [0, 0, 1, 3]

Before: [0, 0, 3, 0]
12 3 2 3
After:  [0, 0, 3, 1]

Before: [2, 1, 2, 0]
5 1 2 1
After:  [2, 0, 2, 0]

Before: [2, 2, 2, 3]
15 2 3 2
After:  [2, 2, 0, 3]

Before: [2, 0, 3, 0]
14 0 3 3
After:  [2, 0, 3, 1]

Before: [1, 1, 3, 3]
0 1 2 2
After:  [1, 1, 0, 3]

Before: [0, 1, 1, 2]
6 0 0 3
After:  [0, 1, 1, 0]

Before: [2, 2, 2, 3]
15 1 3 3
After:  [2, 2, 2, 0]

Before: [2, 3, 2, 0]
14 0 3 3
After:  [2, 3, 2, 1]

Before: [0, 1, 3, 3]
0 1 2 1
After:  [0, 0, 3, 3]

Before: [1, 3, 2, 2]
4 3 2 0
After:  [0, 3, 2, 2]

Before: [2, 1, 1, 3]
4 2 1 3
After:  [2, 1, 1, 0]

Before: [2, 1, 0, 0]
14 0 3 3
After:  [2, 1, 0, 1]

Before: [0, 2, 1, 3]
15 1 3 1
After:  [0, 0, 1, 3]

Before: [2, 1, 3, 0]
0 1 2 2
After:  [2, 1, 0, 0]

Before: [1, 2, 0, 2]
12 2 3 0
After:  [1, 2, 0, 2]

Before: [0, 3, 0, 3]
13 0 3 3
After:  [0, 3, 0, 0]

Before: [3, 3, 0, 1]
10 2 0 1
After:  [3, 0, 0, 1]

Before: [2, 1, 1, 0]
14 0 3 3
After:  [2, 1, 1, 1]

Before: [1, 2, 2, 0]
8 0 2 2
After:  [1, 2, 0, 0]

Before: [2, 0, 0, 3]
4 0 1 0
After:  [1, 0, 0, 3]

Before: [1, 1, 2, 1]
5 1 2 2
After:  [1, 1, 0, 1]

Before: [3, 0, 3, 3]
9 3 2 0
After:  [1, 0, 3, 3]

Before: [1, 2, 2, 3]
9 2 1 0
After:  [1, 2, 2, 3]

Before: [2, 3, 2, 0]
14 0 3 0
After:  [1, 3, 2, 0]

Before: [3, 2, 3, 1]
7 3 3 3
After:  [3, 2, 3, 0]

Before: [0, 1, 3, 0]
0 1 2 1
After:  [0, 0, 3, 0]

Before: [3, 3, 0, 3]
9 3 0 2
After:  [3, 3, 1, 3]

Before: [3, 1, 3, 2]
0 1 2 0
After:  [0, 1, 3, 2]

Before: [0, 0, 2, 1]
11 3 2 3
After:  [0, 0, 2, 1]

Before: [0, 2, 1, 3]
13 0 1 3
After:  [0, 2, 1, 0]

Before: [1, 2, 2, 2]
3 1 2 3
After:  [1, 2, 2, 4]

Before: [0, 1, 2, 3]
5 1 2 1
After:  [0, 0, 2, 3]

Before: [2, 0, 3, 2]
10 1 0 0
After:  [0, 0, 3, 2]

Before: [1, 1, 0, 2]
12 2 3 0
After:  [1, 1, 0, 2]

Before: [3, 3, 1, 1]
1 2 3 0
After:  [2, 3, 1, 1]

Before: [1, 1, 0, 3]
2 0 2 3
After:  [1, 1, 0, 0]

Before: [1, 3, 1, 1]
1 2 0 2
After:  [1, 3, 2, 1]

Before: [2, 1, 2, 0]
14 0 3 1
After:  [2, 1, 2, 0]

Before: [2, 0, 2, 1]
4 0 1 3
After:  [2, 0, 2, 1]

Before: [2, 0, 2, 0]
14 0 3 3
After:  [2, 0, 2, 1]

Before: [0, 1, 3, 1]
0 1 2 2
After:  [0, 1, 0, 1]

Before: [1, 1, 2, 2]
3 2 2 1
After:  [1, 4, 2, 2]

Before: [0, 2, 3, 1]
6 0 0 3
After:  [0, 2, 3, 0]

Before: [1, 3, 0, 2]
2 0 2 3
After:  [1, 3, 0, 0]

Before: [1, 0, 1, 1]
1 2 3 3
After:  [1, 0, 1, 2]

Before: [1, 3, 1, 1]
1 2 3 0
After:  [2, 3, 1, 1]

Before: [1, 2, 2, 2]
9 2 1 0
After:  [1, 2, 2, 2]

Before: [2, 1, 3, 0]
14 0 3 1
After:  [2, 1, 3, 0]

Before: [1, 2, 1, 3]
1 2 0 0
After:  [2, 2, 1, 3]

Before: [1, 0, 0, 0]
2 0 2 2
After:  [1, 0, 0, 0]

Before: [2, 0, 2, 2]
3 2 2 0
After:  [4, 0, 2, 2]

Before: [2, 0, 3, 2]
4 0 1 3
After:  [2, 0, 3, 1]

Before: [2, 1, 2, 1]
5 1 2 3
After:  [2, 1, 2, 0]

Before: [1, 2, 2, 3]
15 2 3 1
After:  [1, 0, 2, 3]

Before: [0, 3, 2, 1]
6 0 0 2
After:  [0, 3, 0, 1]

Before: [1, 0, 0, 3]
2 0 2 2
After:  [1, 0, 0, 3]

Before: [1, 3, 2, 2]
8 0 2 0
After:  [0, 3, 2, 2]

Before: [0, 2, 1, 1]
6 0 0 1
After:  [0, 0, 1, 1]

Before: [2, 0, 2, 3]
9 2 0 1
After:  [2, 1, 2, 3]

Before: [2, 1, 2, 0]
14 0 3 2
After:  [2, 1, 1, 0]

Before: [1, 0, 1, 2]
1 2 0 3
After:  [1, 0, 1, 2]

Before: [2, 2, 3, 0]
12 3 2 3
After:  [2, 2, 3, 1]

Before: [3, 3, 3, 3]
9 3 2 1
After:  [3, 1, 3, 3]

Before: [0, 2, 2, 2]
3 1 2 3
After:  [0, 2, 2, 4]

Before: [2, 2, 2, 1]
9 2 1 1
After:  [2, 1, 2, 1]

Before: [1, 3, 0, 0]
2 0 2 1
After:  [1, 0, 0, 0]

Before: [2, 1, 2, 0]
14 0 3 0
After:  [1, 1, 2, 0]

Before: [1, 3, 0, 1]
2 0 2 3
After:  [1, 3, 0, 0]

Before: [1, 3, 1, 1]
1 2 3 1
After:  [1, 2, 1, 1]

Before: [1, 3, 0, 3]
2 0 2 1
After:  [1, 0, 0, 3]

Before: [3, 1, 2, 3]
15 2 3 2
After:  [3, 1, 0, 3]

Before: [2, 3, 3, 3]
9 3 2 0
After:  [1, 3, 3, 3]

Before: [0, 1, 3, 1]
0 1 2 3
After:  [0, 1, 3, 0]

Before: [1, 0, 1, 1]
1 2 0 1
After:  [1, 2, 1, 1]

Before: [1, 0, 2, 1]
8 0 2 0
After:  [0, 0, 2, 1]

Before: [1, 2, 2, 2]
8 0 2 1
After:  [1, 0, 2, 2]

Before: [3, 2, 1, 3]
9 3 0 0
After:  [1, 2, 1, 3]

Before: [2, 1, 3, 2]
0 1 2 3
After:  [2, 1, 3, 0]

Before: [2, 2, 2, 0]
3 2 2 0
After:  [4, 2, 2, 0]

Before: [3, 3, 2, 3]
15 2 3 1
After:  [3, 0, 2, 3]

Before: [3, 0, 1, 1]
10 1 0 3
After:  [3, 0, 1, 0]

Before: [0, 1, 3, 3]
6 0 0 1
After:  [0, 0, 3, 3]

Before: [0, 0, 3, 3]
9 3 2 0
After:  [1, 0, 3, 3]

Before: [0, 2, 1, 1]
1 2 3 0
After:  [2, 2, 1, 1]

Before: [2, 3, 1, 0]
14 0 3 1
After:  [2, 1, 1, 0]

Before: [3, 0, 0, 3]
9 3 0 1
After:  [3, 1, 0, 3]

Before: [0, 0, 0, 2]
6 0 0 0
After:  [0, 0, 0, 2]

Before: [1, 3, 2, 0]
8 0 2 3
After:  [1, 3, 2, 0]

Before: [0, 3, 1, 1]
13 0 3 1
After:  [0, 0, 1, 1]

Before: [1, 2, 1, 1]
1 2 3 0
After:  [2, 2, 1, 1]

Before: [0, 1, 1, 2]
13 0 1 3
After:  [0, 1, 1, 0]

Before: [1, 2, 0, 0]
2 0 2 0
After:  [0, 2, 0, 0]

Before: [0, 0, 0, 0]
6 0 0 1
After:  [0, 0, 0, 0]

Before: [0, 0, 3, 0]
12 3 2 0
After:  [1, 0, 3, 0]

Before: [0, 3, 2, 2]
4 3 2 3
After:  [0, 3, 2, 0]

Before: [3, 3, 3, 2]
7 3 3 2
After:  [3, 3, 0, 2]

Before: [0, 2, 0, 1]
13 0 3 1
After:  [0, 0, 0, 1]

Before: [3, 1, 3, 1]
0 1 2 3
After:  [3, 1, 3, 0]

Before: [3, 1, 2, 0]
5 1 2 1
After:  [3, 0, 2, 0]

Before: [0, 1, 2, 0]
5 1 2 0
After:  [0, 1, 2, 0]

Before: [0, 1, 3, 3]
0 1 2 2
After:  [0, 1, 0, 3]

Before: [1, 0, 0, 2]
2 0 2 1
After:  [1, 0, 0, 2]

Before: [2, 3, 1, 3]
15 2 3 2
After:  [2, 3, 0, 3]

Before: [2, 2, 1, 2]
7 3 3 1
After:  [2, 0, 1, 2]

Before: [0, 1, 0, 2]
7 3 3 3
After:  [0, 1, 0, 0]

Before: [1, 1, 2, 2]
3 3 2 2
After:  [1, 1, 4, 2]

Before: [2, 0, 1, 1]
1 2 3 1
After:  [2, 2, 1, 1]

Before: [1, 2, 0, 1]
2 0 2 2
After:  [1, 2, 0, 1]

Before: [0, 1, 1, 0]
6 0 0 2
After:  [0, 1, 0, 0]

Before: [2, 0, 2, 3]
3 0 2 1
After:  [2, 4, 2, 3]

Before: [0, 1, 2, 0]
5 1 2 1
After:  [0, 0, 2, 0]

Before: [2, 2, 2, 3]
3 0 2 0
After:  [4, 2, 2, 3]

Before: [1, 1, 2, 3]
15 2 3 0
After:  [0, 1, 2, 3]

Before: [0, 1, 2, 1]
3 2 2 2
After:  [0, 1, 4, 1]

Before: [2, 1, 2, 3]
5 1 2 2
After:  [2, 1, 0, 3]

Before: [3, 3, 0, 2]
12 2 3 0
After:  [1, 3, 0, 2]

Before: [3, 1, 3, 1]
0 1 2 1
After:  [3, 0, 3, 1]

Before: [3, 0, 3, 0]
12 3 2 1
After:  [3, 1, 3, 0]

Before: [2, 2, 2, 0]
14 0 3 1
After:  [2, 1, 2, 0]

Before: [1, 1, 0, 1]
2 0 2 0
After:  [0, 1, 0, 1]

Before: [3, 3, 1, 1]
1 2 3 3
After:  [3, 3, 1, 2]

Before: [0, 3, 0, 2]
12 2 3 1
After:  [0, 1, 0, 2]

Before: [1, 1, 2, 2]
3 2 2 0
After:  [4, 1, 2, 2]

Before: [3, 2, 2, 1]
11 3 2 3
After:  [3, 2, 2, 1]

Before: [0, 1, 2, 1]
13 0 1 3
After:  [0, 1, 2, 0]

Before: [1, 1, 1, 2]
1 2 0 1
After:  [1, 2, 1, 2]

Before: [0, 3, 2, 1]
11 3 2 2
After:  [0, 3, 1, 1]

Before: [1, 1, 3, 3]
0 1 2 3
After:  [1, 1, 3, 0]

Before: [3, 1, 0, 2]
4 0 2 0
After:  [1, 1, 0, 2]

Before: [1, 1, 3, 0]
0 1 2 2
After:  [1, 1, 0, 0]

Before: [1, 2, 0, 2]
2 0 2 1
After:  [1, 0, 0, 2]

Before: [1, 0, 1, 2]
1 2 0 0
After:  [2, 0, 1, 2]

Before: [3, 1, 2, 1]
5 1 2 0
After:  [0, 1, 2, 1]

Before: [2, 2, 2, 1]
11 3 2 2
After:  [2, 2, 1, 1]

Before: [0, 1, 2, 0]
5 1 2 3
After:  [0, 1, 2, 0]

Before: [3, 3, 1, 2]
7 3 3 1
After:  [3, 0, 1, 2]

Before: [3, 1, 2, 2]
5 1 2 2
After:  [3, 1, 0, 2]

Before: [1, 3, 0, 3]
2 0 2 3
After:  [1, 3, 0, 0]

Before: [1, 1, 2, 0]
5 1 2 0
After:  [0, 1, 2, 0]

Before: [1, 1, 2, 2]
3 3 2 1
After:  [1, 4, 2, 2]

Before: [3, 0, 2, 1]
10 1 0 0
After:  [0, 0, 2, 1]

Before: [0, 0, 1, 1]
1 2 3 1
After:  [0, 2, 1, 1]

Before: [0, 3, 2, 1]
11 3 2 0
After:  [1, 3, 2, 1]

Before: [3, 3, 0, 3]
10 2 0 1
After:  [3, 0, 0, 3]

Before: [2, 2, 3, 1]
9 2 3 2
After:  [2, 2, 0, 1]

Before: [1, 1, 2, 1]
8 0 2 2
After:  [1, 1, 0, 1]

Before: [2, 1, 2, 1]
5 1 2 0
After:  [0, 1, 2, 1]

Before: [3, 1, 3, 0]
12 3 2 0
After:  [1, 1, 3, 0]

Before: [1, 0, 0, 2]
12 2 3 2
After:  [1, 0, 1, 2]

Before: [1, 3, 2, 1]
8 0 2 0
After:  [0, 3, 2, 1]

Before: [3, 2, 2, 3]
3 1 2 2
After:  [3, 2, 4, 3]

Before: [0, 2, 2, 2]
4 3 2 3
After:  [0, 2, 2, 0]

Before: [3, 1, 2, 1]
5 1 2 3
After:  [3, 1, 2, 0]

Before: [0, 2, 0, 2]
6 0 0 2
After:  [0, 2, 0, 2]

Before: [3, 1, 0, 3]
15 1 3 1
After:  [3, 0, 0, 3]

Before: [2, 2, 2, 0]
3 0 2 0
After:  [4, 2, 2, 0]

Before: [0, 0, 0, 2]
12 2 3 1
After:  [0, 1, 0, 2]

Before: [3, 1, 3, 3]
9 3 0 1
After:  [3, 1, 3, 3]

Before: [1, 0, 2, 1]
11 3 2 3
After:  [1, 0, 2, 1]

Before: [2, 1, 0, 0]
14 0 3 2
After:  [2, 1, 1, 0]

Before: [3, 0, 2, 2]
10 1 0 1
After:  [3, 0, 2, 2]

Before: [2, 2, 1, 1]
1 2 3 1
After:  [2, 2, 1, 1]

Before: [0, 2, 2, 3]
15 1 3 2
After:  [0, 2, 0, 3]

Before: [1, 3, 0, 3]
2 0 2 2
After:  [1, 3, 0, 3]

Before: [3, 1, 2, 1]
11 3 2 0
After:  [1, 1, 2, 1]

Before: [2, 1, 2, 1]
11 3 2 0
After:  [1, 1, 2, 1]

Before: [1, 1, 2, 3]
8 0 2 1
After:  [1, 0, 2, 3]

Before: [2, 1, 2, 3]
5 1 2 0
After:  [0, 1, 2, 3]

Before: [1, 1, 0, 2]
7 3 3 1
After:  [1, 0, 0, 2]

Before: [0, 0, 2, 3]
15 2 3 2
After:  [0, 0, 0, 3]

Before: [1, 3, 0, 0]
2 0 2 0
After:  [0, 3, 0, 0]

Before: [3, 2, 2, 2]
9 2 1 3
After:  [3, 2, 2, 1]

Before: [2, 3, 3, 0]
12 3 2 3
After:  [2, 3, 3, 1]

Before: [3, 0, 1, 1]
7 2 3 0
After:  [0, 0, 1, 1]

Before: [3, 1, 3, 0]
0 1 2 0
After:  [0, 1, 3, 0]

Before: [0, 2, 2, 1]
11 3 2 2
After:  [0, 2, 1, 1]

Before: [1, 1, 1, 3]
15 1 3 2
After:  [1, 1, 0, 3]

Before: [2, 3, 3, 1]
9 2 3 2
After:  [2, 3, 0, 1]

Before: [0, 2, 0, 3]
6 0 0 0
After:  [0, 2, 0, 3]

Before: [2, 3, 1, 0]
14 0 3 0
After:  [1, 3, 1, 0]

Before: [2, 3, 2, 0]
14 0 3 2
After:  [2, 3, 1, 0]

Before: [0, 3, 1, 1]
6 0 0 2
After:  [0, 3, 0, 1]

Before: [2, 0, 1, 1]
1 2 3 3
After:  [2, 0, 1, 2]

Before: [1, 2, 1, 3]
1 2 0 1
After:  [1, 2, 1, 3]

Before: [3, 1, 2, 1]
5 1 2 2
After:  [3, 1, 0, 1]

Before: [1, 3, 2, 3]
8 0 2 0
After:  [0, 3, 2, 3]

Before: [3, 1, 0, 1]
4 3 1 0
After:  [0, 1, 0, 1]

Before: [0, 3, 2, 1]
11 3 2 1
After:  [0, 1, 2, 1]

Before: [2, 1, 3, 2]
0 1 2 2
After:  [2, 1, 0, 2]

Before: [2, 2, 1, 0]
14 0 3 0
After:  [1, 2, 1, 0]

Before: [1, 2, 2, 1]
11 3 2 1
After:  [1, 1, 2, 1]

Before: [2, 2, 2, 0]
14 0 3 3
After:  [2, 2, 2, 1]

Before: [3, 1, 1, 1]
1 2 3 3
After:  [3, 1, 1, 2]

Before: [1, 2, 2, 1]
11 3 2 0
After:  [1, 2, 2, 1]

Before: [3, 3, 0, 2]
4 0 2 1
After:  [3, 1, 0, 2]

Before: [0, 1, 0, 1]
6 0 0 1
After:  [0, 0, 0, 1]

Before: [3, 3, 2, 3]
9 3 0 0
After:  [1, 3, 2, 3]

Before: [0, 2, 1, 3]
13 0 2 3
After:  [0, 2, 1, 0]

Before: [0, 2, 1, 1]
6 0 0 3
After:  [0, 2, 1, 0]

Before: [1, 2, 0, 0]
2 0 2 3
After:  [1, 2, 0, 0]

Before: [2, 1, 0, 2]
7 3 3 0
After:  [0, 1, 0, 2]

Before: [3, 1, 2, 1]
11 3 2 3
After:  [3, 1, 2, 1]

Before: [0, 2, 0, 1]
6 0 0 2
After:  [0, 2, 0, 1]

Before: [1, 2, 2, 1]
8 0 2 3
After:  [1, 2, 2, 0]

Before: [0, 0, 2, 2]
6 0 0 0
After:  [0, 0, 2, 2]

Before: [0, 2, 3, 1]
6 0 0 2
After:  [0, 2, 0, 1]

Before: [3, 2, 0, 2]
12 2 3 3
After:  [3, 2, 0, 1]

Before: [2, 1, 3, 0]
12 3 2 1
After:  [2, 1, 3, 0]

Before: [3, 1, 1, 0]
4 2 1 1
After:  [3, 0, 1, 0]

Before: [2, 1, 2, 1]
11 3 2 3
After:  [2, 1, 2, 1]

Before: [1, 0, 2, 0]
8 0 2 1
After:  [1, 0, 2, 0]

Before: [3, 0, 2, 0]
3 2 2 0
After:  [4, 0, 2, 0]

Before: [2, 2, 1, 3]
15 2 3 2
After:  [2, 2, 0, 3]

Before: [1, 0, 3, 1]
7 3 3 1
After:  [1, 0, 3, 1]

Before: [0, 1, 2, 1]
11 3 2 2
After:  [0, 1, 1, 1]

Before: [1, 0, 1, 0]
1 2 0 1
After:  [1, 2, 1, 0]

Before: [0, 1, 3, 3]
13 0 1 1
After:  [0, 0, 3, 3]

Before: [1, 3, 0, 0]
2 0 2 2
After:  [1, 3, 0, 0]

Before: [0, 1, 1, 1]
13 0 1 0
After:  [0, 1, 1, 1]

Before: [1, 0, 1, 2]
1 2 0 2
After:  [1, 0, 2, 2]

Before: [0, 3, 2, 1]
13 0 1 3
After:  [0, 3, 2, 0]

Before: [2, 0, 0, 0]
14 0 3 2
After:  [2, 0, 1, 0]

Before: [1, 1, 1, 1]
1 2 0 0
After:  [2, 1, 1, 1]

Before: [3, 0, 1, 3]
10 1 0 1
After:  [3, 0, 1, 3]

Before: [1, 2, 1, 2]
7 3 3 1
After:  [1, 0, 1, 2]

Before: [0, 1, 2, 2]
5 1 2 1
After:  [0, 0, 2, 2]

Before: [0, 0, 2, 1]
6 0 0 0
After:  [0, 0, 2, 1]

Before: [1, 1, 0, 2]
2 0 2 3
After:  [1, 1, 0, 0]

Before: [2, 2, 3, 0]
14 0 3 2
After:  [2, 2, 1, 0]

Before: [1, 1, 3, 1]
0 1 2 2
After:  [1, 1, 0, 1]

Before: [1, 3, 1, 1]
7 3 3 0
After:  [0, 3, 1, 1]

Before: [0, 0, 3, 0]
6 0 0 1
After:  [0, 0, 3, 0]

Before: [3, 1, 3, 0]
0 1 2 1
After:  [3, 0, 3, 0]

Before: [1, 2, 2, 1]
8 0 2 1
After:  [1, 0, 2, 1]

Before: [1, 2, 0, 3]
2 0 2 2
After:  [1, 2, 0, 3]

Before: [1, 1, 0, 2]
12 2 3 1
After:  [1, 1, 0, 2]

Before: [0, 0, 3, 0]
12 3 2 2
After:  [0, 0, 1, 0]

Before: [1, 2, 0, 2]
2 0 2 3
After:  [1, 2, 0, 0]

Before: [0, 2, 3, 1]
13 0 3 3
After:  [0, 2, 3, 0]

Before: [0, 2, 2, 1]
11 3 2 0
After:  [1, 2, 2, 1]

Before: [1, 3, 2, 2]
8 0 2 3
After:  [1, 3, 2, 0]

Before: [1, 1, 1, 3]
1 2 0 1
After:  [1, 2, 1, 3]

Before: [3, 0, 0, 1]
10 1 0 0
After:  [0, 0, 0, 1]

Before: [2, 1, 1, 1]
7 2 3 0
After:  [0, 1, 1, 1]

Before: [0, 1, 2, 2]
5 1 2 0
After:  [0, 1, 2, 2]

Before: [0, 2, 1, 2]
6 0 0 1
After:  [0, 0, 1, 2]

Before: [0, 3, 3, 1]
7 3 3 2
After:  [0, 3, 0, 1]

Before: [1, 1, 2, 1]
8 0 2 1
After:  [1, 0, 2, 1]

Before: [1, 1, 2, 1]
11 3 2 3
After:  [1, 1, 2, 1]

Before: [0, 1, 1, 3]
13 0 2 2
After:  [0, 1, 0, 3]

Before: [2, 2, 2, 2]
3 1 2 0
After:  [4, 2, 2, 2]

Before: [0, 3, 1, 1]
1 2 3 0
After:  [2, 3, 1, 1]

Before: [1, 1, 0, 3]
2 0 2 2
After:  [1, 1, 0, 3]

Before: [2, 1, 3, 1]
0 1 2 0
After:  [0, 1, 3, 1]

Before: [1, 1, 0, 0]
2 0 2 1
After:  [1, 0, 0, 0]

Before: [3, 2, 1, 1]
1 2 3 2
After:  [3, 2, 2, 1]

Before: [2, 3, 2, 3]
9 2 0 3
After:  [2, 3, 2, 1]

Before: [0, 0, 0, 1]
6 0 0 3
After:  [0, 0, 0, 0]

Before: [2, 2, 3, 3]
15 1 3 3
After:  [2, 2, 3, 0]

Before: [1, 0, 2, 3]
8 0 2 3
After:  [1, 0, 2, 0]

Before: [1, 1, 0, 2]
2 0 2 0
After:  [0, 1, 0, 2]

Before: [1, 1, 0, 2]
2 0 2 1
After:  [1, 0, 0, 2]

Before: [0, 0, 0, 1]
6 0 0 1
After:  [0, 0, 0, 1]

Before: [1, 3, 2, 1]
8 0 2 1
After:  [1, 0, 2, 1]

Before: [2, 0, 1, 0]
10 1 0 2
After:  [2, 0, 0, 0]

Before: [1, 2, 0, 1]
2 0 2 0
After:  [0, 2, 0, 1]

Before: [0, 2, 2, 3]
15 1 3 3
After:  [0, 2, 2, 0]

Before: [2, 1, 2, 3]
5 1 2 1
After:  [2, 0, 2, 3]

Before: [1, 3, 2, 1]
8 0 2 3
After:  [1, 3, 2, 0]

Before: [1, 1, 2, 2]
5 1 2 3
After:  [1, 1, 2, 0]

Before: [0, 1, 2, 1]
4 3 1 1
After:  [0, 0, 2, 1]

Before: [3, 1, 3, 0]
0 1 2 2
After:  [3, 1, 0, 0]

Before: [2, 0, 2, 2]
3 0 2 0
After:  [4, 0, 2, 2]

Before: [0, 1, 3, 1]
0 1 2 0
After:  [0, 1, 3, 1]

Before: [2, 1, 2, 2]
5 1 2 1
After:  [2, 0, 2, 2]

Before: [1, 3, 0, 0]
2 0 2 3
After:  [1, 3, 0, 0]

Before: [1, 2, 2, 3]
3 2 2 0
After:  [4, 2, 2, 3]

Before: [0, 3, 2, 0]
6 0 0 2
After:  [0, 3, 0, 0]

Before: [0, 3, 2, 3]
13 0 3 3
After:  [0, 3, 2, 0]

Before: [3, 0, 0, 0]
4 0 2 3
After:  [3, 0, 0, 1]

Before: [2, 3, 2, 2]
4 3 2 1
After:  [2, 0, 2, 2]

Before: [2, 2, 2, 0]
3 0 2 2
After:  [2, 2, 4, 0]

Before: [3, 0, 3, 3]
9 3 0 3
After:  [3, 0, 3, 1]

Before: [0, 1, 2, 2]
13 0 1 3
After:  [0, 1, 2, 0]

Before: [1, 1, 1, 1]
1 2 3 3
After:  [1, 1, 1, 2]

Before: [2, 2, 3, 2]
4 2 0 0
After:  [1, 2, 3, 2]

Before: [3, 1, 3, 3]
15 1 3 0
After:  [0, 1, 3, 3]

Before: [0, 1, 1, 3]
6 0 0 1
After:  [0, 0, 1, 3]

Before: [0, 3, 1, 1]
6 0 0 1
After:  [0, 0, 1, 1]

Before: [3, 0, 0, 3]
9 3 0 2
After:  [3, 0, 1, 3]

Before: [0, 3, 3, 1]
9 2 3 2
After:  [0, 3, 0, 1]

Before: [2, 0, 1, 0]
14 0 3 3
After:  [2, 0, 1, 1]

Before: [0, 2, 3, 0]
6 0 0 3
After:  [0, 2, 3, 0]

Before: [1, 3, 0, 2]
2 0 2 1
After:  [1, 0, 0, 2]

Before: [1, 1, 2, 1]
11 3 2 0
After:  [1, 1, 2, 1]

Before: [1, 1, 2, 3]
8 0 2 0
After:  [0, 1, 2, 3]

Before: [3, 1, 2, 2]
5 1 2 3
After:  [3, 1, 2, 0]

Before: [1, 0, 0, 2]
2 0 2 2
After:  [1, 0, 0, 2]

Before: [2, 0, 3, 0]
12 3 2 3
After:  [2, 0, 3, 1]

Before: [1, 1, 2, 1]
8 0 2 3
After:  [1, 1, 2, 0]

Before: [1, 1, 2, 0]
8 0 2 1
After:  [1, 0, 2, 0]

Before: [3, 1, 2, 2]
5 1 2 0
After:  [0, 1, 2, 2]

Before: [1, 2, 2, 2]
8 0 2 3
After:  [1, 2, 2, 0]

Before: [3, 3, 0, 0]
10 2 0 2
After:  [3, 3, 0, 0]

Before: [0, 3, 2, 2]
6 0 0 2
After:  [0, 3, 0, 2]

Before: [3, 0, 3, 1]
10 1 0 0
After:  [0, 0, 3, 1]

Before: [2, 3, 1, 1]
7 2 3 1
After:  [2, 0, 1, 1]

Before: [3, 1, 2, 3]
3 2 2 1
After:  [3, 4, 2, 3]

Before: [0, 2, 2, 3]
15 1 3 1
After:  [0, 0, 2, 3]

Before: [0, 3, 2, 1]
6 0 0 0
After:  [0, 3, 2, 1]

Before: [0, 1, 3, 0]
6 0 0 3
After:  [0, 1, 3, 0]

Before: [1, 3, 2, 1]
7 3 3 1
After:  [1, 0, 2, 1]

Before: [1, 0, 0, 1]
2 0 2 1
After:  [1, 0, 0, 1]

Before: [3, 1, 0, 2]
12 2 3 3
After:  [3, 1, 0, 1]

Before: [3, 0, 2, 3]
10 1 0 0
After:  [0, 0, 2, 3]

Before: [3, 2, 2, 1]
11 3 2 0
After:  [1, 2, 2, 1]

Before: [1, 1, 2, 3]
5 1 2 1
After:  [1, 0, 2, 3]

Before: [1, 1, 0, 1]
2 0 2 3
After:  [1, 1, 0, 0]

Before: [2, 1, 2, 2]
5 1 2 0
After:  [0, 1, 2, 2]

Before: [0, 3, 3, 1]
6 0 0 3
After:  [0, 3, 3, 0]

Before: [3, 2, 2, 3]
9 2 1 3
After:  [3, 2, 2, 1]

Before: [1, 1, 2, 1]
8 0 2 0
After:  [0, 1, 2, 1]

Before: [2, 2, 1, 0]
14 0 3 2
After:  [2, 2, 1, 0]

Before: [3, 2, 2, 3]
9 2 1 2
After:  [3, 2, 1, 3]

Before: [1, 2, 2, 2]
8 0 2 2
After:  [1, 2, 0, 2]

Before: [1, 0, 2, 2]
8 0 2 0
After:  [0, 0, 2, 2]

Before: [1, 3, 0, 1]
2 0 2 1
After:  [1, 0, 0, 1]

Before: [3, 2, 3, 3]
15 1 3 3
After:  [3, 2, 3, 0]

Before: [0, 2, 0, 3]
6 0 0 1
After:  [0, 0, 0, 3]

Before: [2, 1, 3, 1]
0 1 2 1
After:  [2, 0, 3, 1]

Before: [1, 3, 0, 1]
2 0 2 0
After:  [0, 3, 0, 1]

Before: [2, 1, 2, 1]
5 1 2 1
After:  [2, 0, 2, 1]

Before: [1, 3, 0, 2]
12 2 3 2
After:  [1, 3, 1, 2]

Before: [0, 3, 0, 1]
13 0 1 2
After:  [0, 3, 0, 1]

Before: [3, 3, 0, 1]
7 3 3 3
After:  [3, 3, 0, 0]

Before: [0, 3, 0, 0]
13 0 1 2
After:  [0, 3, 0, 0]

Before: [2, 1, 1, 0]
4 2 1 3
After:  [2, 1, 1, 0]

Before: [3, 0, 0, 3]
10 1 0 1
After:  [3, 0, 0, 3]

Before: [2, 0, 2, 3]
15 2 3 0
After:  [0, 0, 2, 3]

Before: [1, 0, 0, 2]
2 0 2 3
After:  [1, 0, 0, 0]

Before: [1, 1, 0, 3]
2 0 2 0
After:  [0, 1, 0, 3]

Before: [3, 0, 0, 0]
10 2 0 3
After:  [3, 0, 0, 0]

Before: [3, 0, 2, 1]
11 3 2 0
After:  [1, 0, 2, 1]

Before: [3, 0, 0, 3]
10 2 0 2
After:  [3, 0, 0, 3]`

const prog = `
13 0 0 0
3 0 2 0
8 3 0 1
13 0 0 3
3 3 1 3
4 0 1 0
13 0 1 0
1 2 0 2
8 2 1 3
8 1 2 0
13 0 0 1
3 1 0 1
2 0 3 1
13 1 3 1
13 1 3 1
1 2 1 2
11 2 0 0
8 3 0 2
8 3 2 1
8 0 0 3
12 3 2 2
13 2 2 2
13 2 1 2
1 0 2 0
11 0 3 1
8 3 1 3
8 0 2 2
8 0 3 0
0 3 2 3
13 3 1 3
1 1 3 1
11 1 1 3
8 3 1 0
8 0 2 1
13 3 0 2
3 2 2 2
4 2 0 1
13 1 3 1
1 3 1 3
11 3 2 1
8 0 2 3
8 1 0 0
15 3 2 2
13 2 3 2
1 1 2 1
11 1 2 3
8 3 0 1
8 2 0 2
11 0 2 2
13 2 3 2
1 2 3 3
11 3 2 0
13 2 0 2
3 2 3 2
8 3 1 3
0 3 2 3
13 3 1 3
1 3 0 0
11 0 3 1
8 2 2 3
8 2 1 2
8 3 0 0
4 2 0 3
13 3 1 3
1 3 1 1
11 1 0 2
8 3 3 3
8 2 2 0
13 1 0 1
3 1 0 1
5 3 0 3
13 3 1 3
1 2 3 2
11 2 3 1
8 1 1 0
8 0 2 3
8 2 0 2
11 0 2 3
13 3 1 3
1 3 1 1
13 2 0 3
3 3 3 3
8 2 3 0
13 2 0 2
3 2 3 2
7 0 2 3
13 3 3 3
1 3 1 1
8 2 0 2
13 3 0 3
3 3 1 3
2 3 0 0
13 0 2 0
13 0 1 0
1 1 0 1
8 2 3 0
8 2 0 3
10 2 3 2
13 2 3 2
13 2 3 2
1 2 1 1
11 1 1 0
8 0 1 1
8 1 0 3
13 2 0 2
3 2 0 2
3 3 1 3
13 3 3 3
1 0 3 0
11 0 2 2
8 2 0 3
13 0 0 0
3 0 2 0
9 0 3 0
13 0 1 0
13 0 2 0
1 2 0 2
8 3 1 1
8 3 2 3
8 2 1 0
5 1 0 0
13 0 1 0
1 2 0 2
11 2 1 3
8 2 3 0
13 1 0 2
3 2 2 2
8 1 2 1
2 1 0 2
13 2 3 2
1 3 2 3
11 3 3 2
8 1 3 0
8 3 0 3
1 1 0 3
13 3 3 3
1 2 3 2
11 2 2 1
8 3 3 3
8 2 2 2
11 0 2 0
13 0 2 0
13 0 3 0
1 0 1 1
11 1 1 0
8 2 3 3
8 2 1 1
10 1 3 3
13 3 2 3
1 0 3 0
11 0 2 2
8 1 2 0
8 3 2 3
1 0 0 3
13 3 2 3
13 3 1 3
1 2 3 2
11 2 2 0
8 1 1 1
13 2 0 2
3 2 2 2
13 2 0 3
3 3 0 3
15 3 2 1
13 1 1 1
13 1 2 1
1 1 0 0
11 0 0 3
8 1 1 2
13 0 0 1
3 1 2 1
8 3 3 0
0 0 2 0
13 0 1 0
13 0 2 0
1 3 0 3
11 3 0 0
8 1 3 1
8 2 0 3
8 0 1 2
13 1 2 3
13 3 1 3
1 3 0 0
13 1 0 3
3 3 0 3
8 3 0 2
8 0 2 1
12 3 2 3
13 3 1 3
13 3 2 3
1 0 3 0
11 0 3 2
8 0 0 3
13 1 0 0
3 0 1 0
8 2 1 1
10 1 3 1
13 1 2 1
1 1 2 2
11 2 1 3
8 2 0 2
8 2 3 1
11 0 2 2
13 2 2 2
13 2 2 2
1 2 3 3
11 3 3 2
8 1 0 3
13 0 0 1
3 1 0 1
3 3 1 3
13 3 1 3
1 3 2 2
11 2 3 1
8 2 1 2
8 2 3 3
11 0 2 2
13 2 1 2
1 2 1 1
11 1 2 2
8 1 0 3
13 3 0 0
3 0 2 0
8 0 1 1
14 0 3 1
13 1 3 1
1 1 2 2
11 2 2 1
8 1 1 0
8 3 1 2
1 0 0 0
13 0 1 0
1 0 1 1
13 2 0 0
3 0 3 0
13 3 0 2
3 2 1 2
8 0 3 3
0 0 2 2
13 2 1 2
1 1 2 1
11 1 3 0
8 1 1 3
8 2 1 2
8 3 3 1
3 3 1 3
13 3 2 3
1 3 0 0
8 1 0 2
8 2 1 1
8 0 3 3
10 1 3 1
13 1 1 1
1 0 1 0
11 0 3 3
8 3 0 1
8 1 1 0
8 0 1 2
3 0 1 2
13 2 1 2
1 2 3 3
13 2 0 2
3 2 3 2
8 2 2 0
5 1 0 2
13 2 2 2
1 3 2 3
11 3 0 0
8 2 2 2
8 2 1 3
5 1 3 3
13 3 1 3
1 3 0 0
11 0 1 1
8 3 0 2
8 1 2 0
8 0 1 3
8 2 3 3
13 3 3 3
1 3 1 1
8 2 3 2
8 3 3 0
8 1 3 3
6 2 0 0
13 0 2 0
13 0 1 0
1 1 0 1
11 1 3 3
8 1 2 0
13 2 0 1
3 1 3 1
4 2 1 2
13 2 1 2
1 2 3 3
8 3 0 0
8 2 3 2
4 2 0 1
13 1 1 1
1 3 1 3
11 3 1 0
8 2 3 1
8 0 1 3
15 3 2 1
13 1 1 1
13 1 1 1
1 1 0 0
11 0 1 1
8 0 3 0
15 3 2 3
13 3 2 3
13 3 3 3
1 1 3 1
13 3 0 0
3 0 2 0
8 0 2 2
13 1 0 3
3 3 1 3
13 3 2 0
13 0 3 0
1 1 0 1
8 0 0 3
8 3 2 0
7 2 0 0
13 0 3 0
1 0 1 1
11 1 1 2
8 1 0 0
8 1 1 3
8 1 2 1
1 3 0 3
13 3 3 3
1 3 2 2
11 2 1 3
8 0 2 1
8 0 3 2
13 0 2 1
13 1 1 1
13 1 1 1
1 3 1 3
11 3 1 1
8 3 3 2
8 0 3 3
8 2 0 0
10 0 3 3
13 3 2 3
1 3 1 1
11 1 0 3
8 2 2 1
8 2 2 2
8 3 1 0
4 2 0 2
13 2 1 2
1 3 2 3
8 1 0 0
13 2 0 2
3 2 0 2
8 1 1 1
1 0 0 0
13 0 1 0
1 3 0 3
11 3 3 1
8 2 2 3
8 2 2 0
9 0 3 3
13 3 2 3
13 3 2 3
1 1 3 1
11 1 3 3
8 1 3 0
8 3 1 1
8 2 3 2
11 0 2 0
13 0 1 0
1 0 3 3
11 3 1 0
8 3 2 3
13 0 0 1
3 1 2 1
8 1 3 2
0 3 2 2
13 2 3 2
13 2 2 2
1 2 0 0
11 0 0 1
8 3 1 0
8 2 0 3
8 0 0 2
7 2 0 2
13 2 3 2
13 2 2 2
1 1 2 1
11 1 3 3
8 2 1 2
8 1 2 0
8 0 2 1
11 0 2 0
13 0 2 0
1 3 0 3
11 3 3 2
8 2 2 0
8 0 0 3
10 0 3 0
13 0 1 0
1 2 0 2
11 2 1 3
8 2 2 0
8 1 1 1
8 3 3 2
13 1 2 0
13 0 3 0
13 0 3 0
1 0 3 3
11 3 0 1
8 1 0 3
8 2 3 0
7 0 2 0
13 0 1 0
1 1 0 1
11 1 3 2
8 2 1 0
13 2 0 1
3 1 3 1
14 0 3 3
13 3 1 3
1 3 2 2
11 2 0 0
8 2 2 1
13 0 0 2
3 2 0 2
13 0 0 3
3 3 2 3
12 2 3 2
13 2 1 2
13 2 2 2
1 2 0 0
11 0 0 1
13 2 0 3
3 3 1 3
8 3 3 2
8 2 2 0
14 0 3 2
13 2 1 2
13 2 2 2
1 2 1 1
11 1 0 0
8 3 1 3
13 1 0 2
3 2 3 2
13 0 0 1
3 1 2 1
6 1 2 1
13 1 1 1
1 0 1 0
11 0 1 2
8 2 1 0
8 1 1 3
8 0 2 1
2 3 0 0
13 0 2 0
13 0 1 0
1 0 2 2
11 2 3 0
8 3 1 3
8 3 1 1
8 2 2 2
4 2 1 3
13 3 2 3
13 3 3 3
1 3 0 0
8 3 3 2
13 0 0 3
3 3 3 3
8 1 3 1
13 1 2 2
13 2 2 2
13 2 2 2
1 0 2 0
11 0 3 2
8 2 2 0
13 0 0 1
3 1 2 1
8 2 2 3
10 0 3 1
13 1 2 1
1 1 2 2
11 2 3 3
13 3 0 0
3 0 3 0
8 0 1 2
8 3 2 1
7 2 0 1
13 1 2 1
1 3 1 3
13 2 0 1
3 1 1 1
7 2 0 0
13 0 1 0
13 0 3 0
1 3 0 3
11 3 3 1
8 2 1 2
8 3 2 3
13 0 0 0
3 0 3 0
4 2 0 0
13 0 2 0
1 1 0 1
11 1 3 3
8 1 3 2
8 2 3 0
8 3 3 1
5 1 0 0
13 0 1 0
1 0 3 3
8 2 0 2
8 1 0 0
8 0 0 1
11 0 2 2
13 2 3 2
1 3 2 3
11 3 0 1
13 0 0 0
3 0 2 0
8 2 1 3
8 3 2 2
9 0 3 3
13 3 2 3
1 3 1 1
11 1 1 3
8 3 3 1
8 2 2 2
13 2 0 0
3 0 1 0
4 2 1 0
13 0 3 0
1 3 0 3
11 3 0 1
13 1 0 3
3 3 1 3
13 0 0 2
3 2 3 2
8 1 2 0
13 0 2 3
13 3 1 3
13 3 3 3
1 1 3 1
11 1 3 0
8 0 2 1
8 0 1 3
8 2 3 2
13 2 1 2
1 2 0 0
13 3 0 2
3 2 2 2
13 2 0 1
3 1 3 1
4 2 1 1
13 1 3 1
1 0 1 0
11 0 0 1
8 1 2 2
8 2 3 0
8 3 0 2
13 2 2 2
1 2 1 1
11 1 1 0
8 1 0 1
8 1 1 2
8 1 3 3
1 3 3 2
13 2 2 2
1 0 2 0
11 0 1 1
8 3 2 0
8 2 3 2
8 0 0 3
4 2 0 3
13 3 3 3
1 1 3 1
11 1 0 0
8 3 2 2
8 0 0 1
8 3 0 3
8 2 3 3
13 3 1 3
1 3 0 0
11 0 0 2
13 2 0 0
3 0 2 0
8 1 0 1
8 2 1 3
9 0 3 1
13 1 1 1
13 1 2 1
1 2 1 2
11 2 2 0
8 3 0 1
8 1 0 2
0 1 2 2
13 2 3 2
13 2 3 2
1 2 0 0
8 2 2 1
8 2 1 2
8 0 0 3
15 3 2 3
13 3 1 3
1 3 0 0
11 0 2 3
8 3 3 0
8 3 3 2
8 0 1 1
0 0 2 0
13 0 2 0
13 0 2 0
1 3 0 3
11 3 1 1
8 2 0 0
8 3 2 3
7 0 2 2
13 2 2 2
1 1 2 1
11 1 1 0
8 0 3 3
13 1 0 2
3 2 2 2
8 0 0 1
15 3 2 1
13 1 3 1
1 0 1 0
11 0 0 1
8 3 3 0
13 3 0 2
3 2 1 2
8 1 1 3
1 3 3 0
13 0 2 0
1 0 1 1
11 1 1 0
8 3 0 3
13 0 0 2
3 2 3 2
8 2 2 1
8 2 3 1
13 1 2 1
1 1 0 0
11 0 0 2
8 2 2 0
8 1 2 1
8 1 2 3
2 3 0 3
13 3 2 3
1 3 2 2
8 3 0 1
13 1 0 3
3 3 2 3
9 0 3 1
13 1 2 1
13 1 2 1
1 1 2 2
8 1 3 0
8 2 3 1
10 1 3 3
13 3 1 3
1 2 3 2
11 2 1 1
8 2 1 3
8 3 3 2
1 0 0 2
13 2 3 2
1 1 2 1
8 2 2 0
13 3 0 2
3 2 2 2
9 0 3 3
13 3 1 3
1 3 1 1
8 2 2 3
8 1 2 2
13 0 0 0
3 0 1 0
2 0 3 2
13 2 1 2
1 1 2 1
8 0 1 0
8 3 3 2
8 3 0 2
13 2 1 2
1 2 1 1
11 1 1 2
8 3 0 0
8 0 1 1
5 0 3 1
13 1 1 1
13 1 1 1
1 1 2 2
11 2 0 1
13 2 0 0
3 0 1 0
8 3 0 2
8 0 2 3
12 3 2 0
13 0 1 0
1 0 1 1
11 1 2 0
8 3 0 1
13 1 0 2
3 2 0 2
8 3 2 3
8 2 1 1
13 1 3 1
13 1 3 1
1 1 0 0
11 0 3 1
8 0 3 3
8 3 2 2
8 1 1 0
12 3 2 2
13 2 3 2
13 2 2 2
1 2 1 1
11 1 0 3
13 0 0 0
3 0 0 0
8 1 2 1
13 1 0 2
3 2 0 2
8 2 1 0
13 0 2 0
1 0 3 3
11 3 1 0
13 0 0 3
3 3 0 3
8 3 3 1
13 3 0 2
3 2 2 2
15 3 2 2
13 2 2 2
13 2 2 2
1 0 2 0
8 1 0 1
8 0 3 2
8 2 2 3
2 1 3 1
13 1 2 1
1 1 0 0
11 0 2 1
8 2 2 2
8 2 3 0
9 0 3 0
13 0 1 0
13 0 3 0
1 1 0 1
8 0 2 0
13 3 0 2
3 2 0 2
12 2 3 2
13 2 3 2
13 2 1 2
1 2 1 1
11 1 2 0
13 3 0 2
3 2 1 2
8 2 1 1
8 0 0 3
10 1 3 2
13 2 1 2
13 2 2 2
1 0 2 0
11 0 3 3
8 1 1 1
8 3 2 0
8 3 0 2
13 1 2 0
13 0 2 0
1 3 0 3
11 3 3 0
8 2 0 1
8 0 2 3
6 1 2 2
13 2 2 2
1 2 0 0
11 0 2 2
8 0 0 1
8 2 1 0
8 3 0 0
13 0 1 0
13 0 2 0
1 0 2 2
11 2 3 3
8 3 3 2
8 3 0 1
8 2 1 0
4 0 1 2
13 2 3 2
1 3 2 3
11 3 2 2
8 2 2 3
13 1 0 1
3 1 2 1
10 1 3 3
13 3 3 3
1 2 3 2
11 2 3 3
8 0 3 0
8 0 1 2
8 3 3 1
0 1 2 1
13 1 1 1
13 1 2 1
1 1 3 3
11 3 1 2
8 1 1 1
13 2 0 3
3 3 2 3
2 1 3 1
13 1 2 1
13 1 2 1
1 2 1 2
11 2 1 3
8 2 2 2
13 2 0 1
3 1 0 1
8 3 3 0
6 2 0 2
13 2 1 2
13 2 3 2
1 3 2 3
11 3 2 0
13 0 0 3
3 3 0 3
8 3 0 2
8 3 2 1
12 3 2 1
13 1 3 1
13 1 2 1
1 0 1 0
11 0 2 2
8 1 2 1
8 2 0 0
10 0 3 3
13 3 2 3
1 3 2 2
11 2 1 1
8 2 1 3
8 3 0 2
6 0 2 3
13 3 1 3
13 3 2 3
1 3 1 1
11 1 1 2
13 2 0 3
3 3 1 3
8 0 2 1
14 0 3 0
13 0 1 0
1 2 0 2
11 2 3 0
8 1 0 2
8 2 3 3
8 1 3 3
13 3 2 3
13 3 2 3
1 3 0 0
11 0 0 3
8 3 1 2
13 0 0 0
3 0 2 0
7 0 2 2
13 2 1 2
1 2 3 3
11 3 1 1
8 2 0 2
8 0 0 0
8 1 1 3
8 3 0 2
13 2 1 2
1 1 2 1
11 1 0 0
8 2 0 3
8 1 1 1
13 1 0 2
3 2 3 2
2 1 3 2
13 2 2 2
1 2 0 0
8 3 2 1
8 0 0 2
5 1 3 2
13 2 3 2
1 0 2 0
11 0 2 3
8 2 2 2
8 2 1 0
8 1 3 1
2 1 0 0
13 0 1 0
1 3 0 3
11 3 0 1
8 2 1 3
13 0 0 2
3 2 1 2
8 1 1 0
1 0 0 3
13 3 3 3
1 3 1 1
8 2 0 0
8 2 3 3
13 2 0 2
3 2 0 2
9 0 3 0
13 0 3 0
1 0 1 1
11 1 3 0
8 2 2 2
8 0 3 3
8 0 1 1
15 3 2 2
13 2 3 2
1 0 2 0
11 0 2 3
8 3 2 0
13 3 0 2
3 2 2 2
6 2 0 0
13 0 1 0
1 0 3 3
11 3 3 0
8 3 3 3
8 0 3 2
8 2 2 1
5 3 1 2
13 2 3 2
13 2 1 2
1 0 2 0
11 0 0 1
8 2 0 0
8 0 0 3
8 0 0 2
10 0 3 3
13 3 1 3
1 3 1 1
11 1 2 3
8 3 3 2
8 0 1 1
8 2 1 1
13 1 1 1
13 1 1 1
1 3 1 3
11 3 1 0`

var opcodes = map[string]func(regs [4]int, a, b, c int) [4]int{
	"addr": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = regs[a] + regs[b]
		return regs
	},
	"addi": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = regs[a] + b
		return regs
	},
	"mulr": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = regs[a] * regs[b]
		return regs
	},
	"muli": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = regs[a] * b
		return regs
	},
	"banr": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = regs[a] & regs[b]
		return regs
	},
	"bani": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = regs[a] & b
		return regs
	},
	"borr": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = regs[a] | regs[b]
		return regs
	},
	"bori": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = regs[a] | b
		return regs
	},
	"setr": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = regs[a]
		return regs
	},
	"seti": func(regs [4]int, a, b, c int) [4]int {
		regs[c] = a
		return regs
	},
	"gtir": func(regs [4]int, a, b, c int) [4]int {
		if a > regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"gtri": func(regs [4]int, a, b, c int) [4]int {
		if regs[a] > b {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"gtrr": func(regs [4]int, a, b, c int) [4]int {
		if regs[a] > regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"eqir": func(regs [4]int, a, b, c int) [4]int {
		if a == regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"eqri": func(regs [4]int, a, b, c int) [4]int {
		if regs[a] == b {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"eqrr": func(regs [4]int, a, b, c int) [4]int {
		if regs[a] == regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
}

func behavesLike(before, op, after [4]int) (like []string) {
	for name, fn := range opcodes {
		if fn(before, op[1], op[2], op[3]) == after {
			like = append(like, name)
		}
	}
	return
}

func notBehavesLike(before, op, after [4]int) (not []string) {
	for name, fn := range opcodes {
		if fn(before, op[1], op[2], op[3]) != after {
			not = append(not, name)
		}
	}
	return
}

func main() {
	// part 1
	var threeOrMore int
	lines := utils.Lines(input)
	for i := 0; i < len(lines); i += 4 {
		var before, op, after [4]int
		utils.Sscanf(lines[i], "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])
		utils.Sscanf(lines[i+1], "%d %d %d %d", &op[0], &op[1], &op[2], &op[3])
		utils.Sscanf(lines[i+2], "After: [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])
		if n := len(behavesLike(before, op, after)); n >= 3 {
			threeOrMore++
		}
	}
	utils.Println(threeOrMore)

	// part 2
	var opmap [16]string
	for i := 0; i < len(lines); i += 4 {
		var before, op, after [4]int
		utils.Sscanf(lines[i], "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3])
		utils.Sscanf(lines[i+1], "%d %d %d %d", &op[0], &op[1], &op[2], &op[3])
		utils.Sscanf(lines[i+2], "After: [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])
		like := behavesLike(before, op, after)
		m := make(map[string]struct{})
		for _, name := range like {
			m[name] = struct{}{}
		}
		for _, name := range opmap {
			delete(m, name)
		}
		if len(m) == 1 {
			for name := range m {
				opmap[op[0]] = name
			}
		}
	}

	var regs [4]int
	for _, line := range utils.Lines(prog) {
		var op [4]int
		utils.Sscanf(line, "%d %d %d %d", &op[0], &op[1], &op[2], &op[3])
		regs = opcodes[opmap[op[0]]](regs, op[1], op[2], op[3])
	}
	fmt.Println(regs[0])
}
