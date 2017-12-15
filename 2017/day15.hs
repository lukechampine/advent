import Data.Bits

judge :: Int -> [(Int, Int)] -> Int
judge n = length . (filter match16) . (take n)
  where
    match16 (a, b) = (a .&. 0xFFFF) == (b .&. 0xFFFF)

genA :: [Int]
genA = iterate (\a -> (a * 16807) `mod` 2147483647) 634

genB :: [Int]
genB = iterate (\b -> (b * 48271) `mod` 2147483647) 301

part1 :: Int
part1 = judge 40000000 (zip genA genB)

part2 :: Int
part2 = judge 5000000 (zip (onlyMultiples 4 genA) (onlyMultiples 8 genB))
  where
    onlyMultiples n = filter (\x -> x `mod` n == 0)

main :: IO ()
main = do
  print $ part1
  print $ part2