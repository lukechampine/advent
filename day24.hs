presents :: [Int]
presents = [1, 3, 5, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113]

groups :: Int -> Int -> [Int] -> [[Int]]
groups 0 _ _  = []
groups m 0 [] = [[]]
groups m _ [] = []
groups m w (x:xs) = map (x:) (groups (m-1) (w-x) xs) ++ groups m w xs

bestGroup :: Int -> Int -> [Int] -> Int
bestGroup maxlen n xs = minimum . map product $ groups maxlen weight xs where
	weight = sum xs `div` n

main :: IO ()
main = do
	print $ bestGroup 7 3 presents
	print $ bestGroup 6 4 presents