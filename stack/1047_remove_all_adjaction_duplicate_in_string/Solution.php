<?php

/*
		problem: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/


	   check each character (c) from s
	   stack, if stack empty => push to stack
	   if the last element == c => remove last element
	   return stack

		time complexity: O(n) because traversal all character of s
		space complexity: O(n) create stack to store node
*/

class Solution
{

    /**
     * @param String $s
     * @return String
     */
    function removeDuplicates($s)
    {
        $stack = array();

        echo strlen($s) . "\n";

        for ($i = 0; $i < strlen($s); $i++) {
            echo $i;
            if (count($stack) == 0) {
                array_push($stack, $s[$i]);
            } elseif ($stack[count($stack) - 1] === $s[$i]) {
                array_pop($stack);
            } else {
                array_push($stack, $s[$i]);
            }
        }

        return implode("", $stack);
    }
}

$solution = new Solution();
$input1 = "abbaca";
//$expected1 = "ca";
echo $solution->removeDuplicates($input1);