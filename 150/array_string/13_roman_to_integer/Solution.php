<?php

class Solution
{

    /**
     * @param String $s
     * @return Integer
     */
    function romanToInt($s)
    {
        $dict = $this->buildDict();
        $sum = 0;

        for ($i = 0; $i < strlen($s); $i++) {
            if ($i <= strlen($s) - 2 && $dict[$s[$i]] < $dict[$s[$i + 1]]) {
                $sum -= $dict[$s[$i]];
            } else {
                $sum += $dict[$s[$i]];
            }
        }

        return $sum;
    }

    function buildDict(): array
    {
        return [
            "I" => 1,
            "V" => 5,
            "X" => 10,
            "L" => 50,
            "C" => 100,
            "D" => 500,
            "M" => 1000,
        ];
    }
}