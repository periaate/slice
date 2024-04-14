# slice

```
`0`		`<index>`		index access
`1:2`	`<from>:<to>`	slice
`5:3`	`<from>:<to>`	slice wrap around
`:3`	`:<to>`			implicit from (equal to 0)
`3:`	`<from>:`		implicit to (equal to length)
`3:-1`	`<from>:<to>`	length relative slice

`//5`	`//<by>`	shift right by n, if n is undefined, n=1
`\\2`	`\\<by>`	shift left by n, if n is undefined, n=1
`_4`	`_<dst>`	take LHS input array and shift it into dst index

`#4`	`#<seed>`	shuffle LHS input array with seed, if seed undefined, seed is randomly generated
`-:`				reverse LHS input array
`.`					resolve LHS expression and provide as input array to RHS expression


example value:
[ 0y, 1y, 2n, 3n, 4n, 5y, 6n, 7y ]

We want all `y` values from that array.

[-3_-1//2.:4]			shift left by 1, slice from 0 to 3




```