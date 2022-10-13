package main

/*
SELECT * FROM pg_indexes WHERE tablename = 'perf_test';

SELECT amname FROM pg_am;

CREATE TABLE perf_test(
    id int,
    reason text COLLATE "C",
    annotation text COLLATE "C"
);
DROP TABLE perf_test;

SELECT * FROM perf_test
ORDER BY id DESC
LIMIT 20
OFFSET 10;

INSERT INTO perf_test(id, reason, annotation)
SELECT s.id, md5(random()::text), null
FROM generate_series(1, 10000000) AS s(id)
ORDER BY random();

UPDATE perf_test
SET annotation = UPPER(md5(random()::text))
WHERE annotation IS NULL;

SELECT *
FROM perf_test
LIMIT 10;

EXPLAIN
SELECT *
FROM perf_test
WHERE id = 3700000;

CREATE INDEX idx_perf_test_id ON perf_test(id);

EXPLAIN
SELECT *
FROM perf_test
WHERE reason LIKE 'dd%' AND annotation  LIKE 'AC%';

CREATE INDEX idx_perf_test_reason_annotation ON perf_test(reason, annotation);

EXPLAIN analyse
SELECT *
FROM perf_test
WHERE LOWER(annotation) LIKE 'ab%';

CREATE INDEX ind_perf_test_annotation_lower ON perf_test(LOWER(annotation));

EXPLAIN
SELECT *
FROM perf_test
WHERE reason LIKE '%ddd%';

CREATE EXTENSION pg_trgm;

CREATE INDEX trgm_ind_perf_test_reason ON perf_test USING GIN(reason gin_trgm_ops);

ALTER TABLE perf_test
ADD COLUMN annotation_lower text;


EXPLAIN
UPDATE perf_test
SET annotation_lower = LOWER(annotation);

CREATE INDEX ind_perf_test_column_annotation_lower ON perf_test(annotation_lower);

EXPLAIN analyse
SELECT *
FROM perf_test
WHERE annotation_lower LIKE 'abcd%';

DROP INDEX ind_perf_test_annotation_lower;

DROP INDEX idx_perf_test_id;

DROP INDEX idx_perf_test_reason_annotation;

DROP INDEX trgm_ind_perf_test_reason;

*/
