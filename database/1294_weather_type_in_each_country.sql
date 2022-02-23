# problem: https://leetcode.com/problems/weather-type-in-each-country

SELECT c.country_name, 
    CASE
        WHEN avg(w.weather_state) <= 15 THEN 'Cold'
        WHEN avg(w.weather_state) >= 25 THEN 'Hot'
        ELSE 'Warm'
    END AS weather_type
    FROM Countries c INNER JOIN Weather w ON c.country_id = w.country_id
    WHERE Month(w.day) = 11 AND Year(w.day) = 2019 
    GROUP BY c.country_id;
