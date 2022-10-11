package main

/*
Вложенный запрос, возвращающий одно значение, может использоваться в условии отбора записей WHERE как обычное значение совместно с операциями =, <>, >=, <=, >, <.

Для реализации этого запроса нам необходимо получить минимальную цену из столбца price таблицы book,
а затем вывести информацию о тех книгах, цена которых  равна минимальной.
Первая часть  – поиск  минимума – реализуется вложенным запросом.
*/

//Пример
//Вывести информацию о самых дешевых книгах, хранящихся на складе.
//
//SELECT title, author, price, amount
//FROM book
//WHERE price = (
//SELECT MIN(price)
//FROM book
//);

/*
Вложенный запрос, возвращающий одно значение, может использоваться в выражениях как обычный операнд,
например, к нему можно что-то прибавить, вычесть и пр.
*/

// Пример
// Вывести информацию о книгах, количество экземпляров которых отличается
// от среднего количества экземпляров книг на складе более чем на 3.
// То есть нужно вывести и те книги, количество экземпляров которых меньше среднего на 3, или больше среднего на 3.

//SELECT title, author, amount
//FROM book
//WHERE ABS(amount - (SELECT AVG(amount) FROM book)) >3;

/*
Вложенный запрос может возвращать несколько значений одного столбца.
Тогда его можно использовать в разделе WHERE совместно с оператором IN.

WHERE имя_столбца IN (вложенный запрос, возвращающий один столбец)
Оператор IN определяет, совпадает ли значение столбца с одним из значений, содержащихся во вложенном запросе.
При этом логическое выражение после WHERE получает значение истина.
Оператор NOT IN выполняет обратное действие – выражение истинно,
если значение столбца не содержится во вложенном запросе.
*/

//Пример
//Вывести информацию о книгах тех авторов, общее количество экземпляров книг которых не менее 12.
//
//SELECT title, author, amount, price
//FROM book
//WHERE author IN (
//        SELECT author
//        FROM book
//        GROUP BY author
//        HAVING SUM(amount) >= 12
//      );