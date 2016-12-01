codenum :: Int -> Int -> Int
codenum row col = col + sum (enumFromTo 2 (row+col-2))

nextcode :: Int -> Int
nextcode x = x * 252533 `mod` 33554393

code :: Int -> Int
code n = iterate' nextcode 20151125 !! n where
	iterate' f x = x `seq` x : iterate' f (f x)

main :: IO ()
main = print $ code (codenum 2981 3075)
