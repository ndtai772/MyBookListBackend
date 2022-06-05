CREATE VIEW rate_count AS
SELECT book_id,
       count(*)                                  as count,
       avg(rate_value)                           as rate_avg
    --    count(*) FILTER ( WHERE rate_value = 1 )  as rate_1,
    --    count(*) FILTER ( WHERE rate_value = 2 )  as rate_2,
    --    count(*) FILTER ( WHERE rate_value = 3 )  as rate_3,
    --    count(*) FILTER ( WHERE rate_value = 4 )  as rate_4,
    --    count(*) FILTER ( WHERE rate_value = 5 )  as rate_5,
    --    count(*) FILTER ( WHERE rate_value = 6 )  as rate_6,
    --    count(*) FILTER ( WHERE rate_value = 7 )  as rate_7,
    --    count(*) FILTER ( WHERE rate_value = 8 )  as rate_8,
    --    count(*) FILTER ( WHERE rate_value = 9 )  as rate_9,
    --    count(*) FILTER ( WHERE rate_value = 10 ) as rate_10
FROM rates
GROUP BY book_id;

CREATE VIEW comment_count AS
SELECT book_id,
       count(*) as count
FROM comments
GROUP BY book_id;

CREATE VIEW bookmark_count AS
SELECT book_id,
       count(*) as count
FROM bookmarks
GROUP BY book_id;

CREATE VIEW category_summary AS
SELECT book_id,
       string_agg(category_id::varchar, ',') as categories
FROM book_category bc
GROUP BY book_id;

CREATE VIEW book_detail AS
SELECT b.*,
       COALESCE(cs.categories, '')::varchar as categories,
       COALESCE(c.count, 0)                 as comment_count,
       COALESCE(bm.count, 0)                as bookmark_count,
       COALESCE(r.count, 0)                 as rate_count,
       COALESCE(r.rate_avg, 0)              as rate_avg
    --    COALESCE(r.rate_1, 0)                as rate_1,
    --    COALESCE(r.rate_2, 0)                as rate_2,
    --    COALESCE(r.rate_3, 0)                as rate_3,
    --    COALESCE(r.rate_4, 0)                as rate_4,
    --    COALESCE(r.rate_5, 0)                as rate_5,
    --    COALESCE(r.rate_6, 0)                as rate_6,
    --    COALESCE(r.rate_7, 0)                as rate_7,
    --    COALESCE(r.rate_8, 0)                as rate_8,
    --    COALESCE(r.rate_9, 0)                as rate_9,
    --    COALESCE(r.rate_10, 0)               as rate_10
FROM books b
         LEFT OUTER JOIN category_summary cs ON b.id = cs.book_id
         LEFT OUTER JOIN comment_count c ON b.id = c.book_id
         LEFT OUTER JOIN bookmark_count bm on bm.book_id = b.id
         LEFT OUTER JOIN rate_count r ON r.book_id = b.id;

