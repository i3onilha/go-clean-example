-- name: GetUserByID :one
SELECT user_id, name, email, location FROM users
WHERE user_id = ?;

-- name: GetOrdersByUserID :many
SELECT o.order_id, o.user_id, o.item, o.quantity, o.price,
       COALESCE(SUM(d.discount_percent), 0) as discount_percent
FROM orders o
LEFT JOIN discounts d ON o.order_id = d.order_id
WHERE o.user_id = ?
GROUP BY o.order_id, o.user_id, o.item, o.quantity, o.price
ORDER BY o.order_id
LIMIT ? OFFSET ?;
