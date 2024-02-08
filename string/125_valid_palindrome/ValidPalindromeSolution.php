<?php

class ValidPalindromeSolution
{

    /**
     * @param $s
     *
     * @return Boolean
     */
    function IsPalindrome($s)
    {
        $dict = $this->newDict();
        $s = strtolower($s);
        $sArr = str_split($s);
        $l = 0;
        $r = count($sArr) - 1;

        while ($l < $r) {
            while (true) {
                if ($dict[$sArr[$l]] === true) {
                    break;
                }
                $l++;
            }

            while (true) {
                if ($dict[$sArr[$r]] === true) {
                    break;
                }
                $r--;
            }

            if ($sArr[$l] != $sArr[$r]) {
                return false;
            }

            $l++;
            $r--;
        }

        return true;
    }

    /**
     * @return array
     */
    function newDict()
    {
        return [
            'a' => true,
            'b' => true,
            'c' => true,
            'd' => true,
            'e' => true,
            'f' => true,
            'g' => true,
            'h' => true,
            'i' => true,
            'j' => true,
            'k' => true,
            'l' => true,
            'm' => true,
            'n' => true,
            'o' => true,
            'p' => true,
            'q' => true,
            'r' => true,
            's' => true,
            't' => true,
            'u' => true,
            'v' => true,
            'w' => true,
            'x' => true,
            'y' => true,
            'z' => true,
            '0' => true,
            '1' => true,
            '2' => true,
            '3' => true,
            '4' => true,
            '5' => true,
            '6' => true,
            '7' => true,
            '8' => true,
            '9' => true
        ];
    }
}

//$s = new ValidPalindromeSolution();
//echo $s->IsPalindrome('A man, a plan, a canal: Panama');
//echo $s->isPalindrome('race a car');
//echo $s->isPalindrome(' ');

