/*
    problem: https://leetcode.com/problems/average-selling-price
*/
SELECT p.product_id, ROUND(SUM(price * units) / SUM(units), 2) as average_price 
    FROM Prices p INNER JOIN UnitsSold us ON p.product_id = us.product_id
    WHERE us.purchase_date BETWEEN p.start_date AND p.end_date GROUP BY p.product_id;
