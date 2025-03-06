# Write your MySQL query statement below
SELECT s.user_id, SUM((p.price * s.quantity)) AS spending
FROM Sales s
LEFT JOIN Product p
ON s.product_id = p.product_id
GROUP BY s.user_id
ORDER BY spending DESC;