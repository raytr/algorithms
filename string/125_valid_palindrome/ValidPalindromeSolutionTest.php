<?php

require './ValidPalindromeSolution.php';
class ValidPalindromeSolutionTest extends \PHPUnit\Framework\TestCase
{

    public function attachEstimatedDeliveryToOrderItemsDataProvider()
    {
        return [

            'case 1' => [
                's'    => 'A man, a plan, a canal: Panama',
                'want' => true,
            ],
            'case 2' => [
                's'    => 'race a car',
                'want' => false,
            ],
            'case 3' => [
                's'    => ' ',
                'want' => true,
            ],
        ];
    }

    /**
     * @dataProvider testcaseDataProvider
     * @param string $s
     * @param bool   $want
     * @return void
     */
    public function testIsPalindrome($s, $want)
    {
        $solution = new ValidPalindromeSolution();
        $this->assertEquals($want, $solution->IsPalindrome($s));
    }
}
