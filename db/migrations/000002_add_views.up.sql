CREATE VIEW rate_count AS
SELECT book_id,
       count(*)                                  as count,
       sum(rate_value)                           as rate_sum,
       count(*) FILTER ( WHERE rate_value = 1 )  as rate_1,
       count(*) FILTER ( WHERE rate_value = 2 )  as rate_2,
       count(*) FILTER ( WHERE rate_value = 3 )  as rate_3,
       count(*) FILTER ( WHERE rate_value = 4 )  as rate_4,
       count(*) FILTER ( WHERE rate_value = 5 )  as rate_5,
       count(*) FILTER ( WHERE rate_value = 6 )  as rate_6,
       count(*) FILTER ( WHERE rate_value = 7 )  as rate_7,
       count(*) FILTER ( WHERE rate_value = 8 )  as rate_8,
       count(*) FILTER ( WHERE rate_value = 9 )  as rate_9,
       count(*) FILTER ( WHERE rate_value = 10 ) as rate_10
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
       cs.categories,
       c.count                       as comment_count,
       bm.count                      as bookmark_count,
       r.count                       as rate_count,
       r.rate_sum::decimal / r.count as rate_avg,
       r.rate_1,
       r.rate_2,
       r.rate_3,
       r.rate_4,
       r.rate_5,
       r.rate_6,
       r.rate_7,
       r.rate_8,
       r.rate_9,
       r.rate_10
FROM books b
         LEFT OUTER JOIN category_summary cs ON b.id = cs.book_id
         LEFT OUTER JOIN comment_count c ON b.id = c.book_id
         LEFT OUTER JOIN bookmark_count bm on bm.book_id = b.id
         LEFT OUTER JOIN rate_count r ON r.book_id = b.id;
