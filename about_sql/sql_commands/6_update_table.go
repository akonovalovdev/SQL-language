package main

/*
Запросы на обновление
Под обновлением данных подразумевается изменение значений в существующих записях таблицы.
При этом возможно как изменение значений полей в группе строк (даже всех строк таблицы),
так и правка значения поля отдельной строки.

Изменение записей в таблице реализуется с помощью запроса UPDATE. Простейший запрос на  обновление выглядит так:

UPDATE таблица SET поле = выражение
где
таблица – имя таблицы, в которой будут проводиться изменения;
поле – поле таблицы, в которое будет внесено изменение;
выражение – выражение,  значение которого будет занесено в поле.
*/

//Пример
//Уменьшить на 30% цену книг в таблице book.
//
//UPDATE book
//SET price = 0.7 * price;

/*
Запросом UPDATE можно обновлять значения нескольких столбцов одновременно. В этом случае простейший запрос будет выглядеть так:

UPDATE таблица SET поле1 = выражение1, поле2 = выражение2
На складе, кроме хранения и получения книг, выполняется их оптовая продажа.
Для реализации этого действия включим дополнительный столбец buy  в таблицу book:

В запросах на обновление можно использовать несколько таблиц, но тогда

для столбцов, имеющих одинаковые имена, необходимо указывать имя таблицы, к которой они относятся, например,
book.price – столбец price из таблицы book, supply.price – столбец price из таблицы supply;

все таблицы, используемые в запросе, нужно перечислить после ключевого слова UPDATE;
в запросе обязательно условие WHERE, в котором указывается условие при котором обновляются данные.
*/q

// Пример
// Если в таблице supply  есть те же книги, что и в таблице book, добавлять эти книги в таблицу book не имеет смысла.
// Необходимо увеличить их количество на значение столбца amount таблицы supply.
//
// UPDATE book, supply
// SET book.amount = book.amount + supply.amount
// WHERE book.title = supply.title AND book.author = supply.author;
