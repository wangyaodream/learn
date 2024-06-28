import sqlite3


number_of_top_customers = int(
        input("How many top customers do you want to query?")
)

conn = sqlite3.connect("data.db")

cursor = conn.cursor()


raw_sql = """
SELECT
    c.id,
    c.first_name,
    SUM(i.total) AS total
FROM Invoice i
LEFT JOIN Customer c ON i.customer_id = c.id
GROUP BY c.id, c.first_name
ORDER BY total DESC
LIMIT ?;
"""


for row in cursor.execute(raw_sql, (number_of_top_customers,)):
    print(row)
