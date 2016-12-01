import Data.List (group)

input :: String
input = "1113222113"

lookAndSay :: String -> [String]
lookAndSay = iterate next where
	next = concatMap say . group
	say str = (show (length str)) ++ [head str]

part1 :: String -> Int
part1 = length . (!! 40) . lookAndSay

part2 :: String -> Int
part2 = length . (!! 50) . lookAndSay

main :: IO ()
main = do
	print $ part1 input
	print $ part2 input
