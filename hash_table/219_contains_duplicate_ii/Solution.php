<?php

class Solution
{

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Boolean
     */
    function containsNearbyDuplicate($nums, $k)
    {
        $appearsAt = [];

        for ($i = 0; $i < count($nums); $i++) {
            if (array_key_exists($nums[$i], $appearsAt) && ($i - $appearsAt[$nums[$i]]) <= $k) {
                return true;
            }
            $appearsAt[$nums[$i]] = $i;
        }

        return false;
    }
}

$solution = new Solution();
echo $solution->containsNearbyDuplicate([1, 2, 3, 1], 3) . "\n";
echo $solution->containsNearbyDuplicate([1, 0, 1, 1], 1) . "\n";
echo $solution->containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 2) . "\n";