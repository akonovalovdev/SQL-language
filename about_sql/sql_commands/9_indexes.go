package main

/*
Индексы — важнейший инструмент баз данных, ускоряющий поиск. Он не бесплатен,
создавать много индексов без лишней необходимости не стоит
— индексы занимают дополнительную память, и при любом обновлении проиндексированных
данных СУБД должна выполнять дополнительную работу по поддержанию индекса в актуальном состоянии
*/
// CREATE INDEX index_name
// ON table_name (column_name_1, column_name_2,....)

/*
В базах данных, таких как PostgreSQL, индекс формируется из значений одного или нескольких
столбцов таблицы и указателей на строки этой таблицы.
*/

// Команда выводящая все созданные индексы в кокнретной таблице
// SELECT * FROM pg_indexes WHERE tablename = 'perf_test';

// Команда выводящая все установленные конфигурации индексов
// SELECT amname FROM pg_am;
// по умолчанию:1 btree, 2 hash, 3 gist, 4 gin, 5 spgist, 6 brin

// Команда выводит все созданные внешние ключи в указанной таблице
// SELECT *
// FROM information_schema.KEY_COLUMN_USAGE
// WHERE TABLE_SCHEMA ='public' AND TABLE_NAME ='users' AND
//        CONSTRAINT_NAME <>'PRIMARY';

// Удаление внешнего ключа ALTER TABLE purchase DROP CONSTRAINT fkey_name;
