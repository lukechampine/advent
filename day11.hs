import Data.List (group)

input :: String
input = "hxbxwxba"

increment :: String -> String
increment = reverse . inc . reverse where
	inc ('z':str) = 'a' : inc str
	inc (c:str)   = succ c : str

nextPassword :: String -> String
nextPassword = until valid increment where
	valid s = and [rule1 s, rule2 s, rule3 s]
	rule1 (x:y:z:s) = y == succ x && z == succ y || rule1 (y:z:s)
	rule1 _         = False
	rule2 s = not . or $ map (`elem` s) ['i','o','l']
	rule3 = (>= 2) . length . filter ((>= 2) . length) . group

part1 :: String -> String
part1 = nextPassword

part2 :: String -> String
part2 = nextPassword . increment . nextPassword

main :: IO ()
main = do
	print $ part1 input
	print $ part2 input
