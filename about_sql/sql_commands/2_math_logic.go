package main

// Логические функции

/*
 В SQL реализована возможность заносить в поле значение в зависимости от условия.
 Для этого используется функция IF():

 mySQL
 IF(логическое_выражение, выражение_1, выражение_2)

 Функция вычисляет логическое_выражение, если оно истина – в поле заносится значение выражения_1,
 в противном случае –  значение выражения_2. Все три параметра IF() являются обязательными.

 Допускается использование вложенных функций, вместо выражения_1 или выражения_2 может стоять новая функция IF.
*/

// Пример:
//
// SELECT title, amount, price,
//    IF(amount<4, price*0.5, price*0.7) AS sale
// FROM book;

// Усложним вычисление скидки в зависимости от количества книг. Если количество книг меньше 4 – то скидка 50%,
// меньше 11 – 30%, в остальных случаях – 10%. И еще укажем какая именно скидка на каждую книгу.
//
// SELECT title, amount, price,
//    ROUND(IF(amount < 4, price * 0.5, IF(amount < 11, price * 0.7, price * 0.9)), 2) AS sale,
//    IF(amount < 4, 'скидка 50%', IF(amount < 11, 'скидка 30%', 'скидка 10%')) AS Ваша_скидка
// FROM book;

/*
 В PostgreSQL та же самая задача решается  иначе
 case when <условие> then <выражение_если_условие_истинно> else <выражение_если_условие_ложно> end
*/
// Пример:
//
// SELECT title, amount, price,
//    ROUND(CASE WHEN amount < 4  THEN price*0.5 ELSE price*0.7 END, 2) AS sale
// FROM book;

/*
Логическое выражение после ключевого слова WHERE кроме операторов сравнения и выражений может включать
логические операции (И «and», ИЛИ «or», НЕ «not») и круглые скобки, изменяющие приоритеты выполнения операций.

Приоритеты операций:

круглые скобки
умножение  (*),  деление (/)
сложение  (+), вычитание (-)
операторы сравнения (=, >, <, >=, <=, <>);  BETWEEN и IN
NOT
AND
OR
*/

//Пример
//SELECT title, author, price
//FROM book
//WHERE price > 600 AND author = 'Булгаков М.А.';

/*
Оператор BETWEEN позволяет отобрать данные, относящиеся к некоторому интервалу, включая его границы.
*/
//Пример
//
//Выбрать названия и количества тех книг, количество которых от 5 до 14 включительно.
//SELECT title, amount
//FROM book
//WHERE amount BETWEEN 5 AND 14;

/*
Оператор  IN  позволяет выбрать данные, соответствующие значениям из списка.
*/
//Пример
//Выбрать названия и цены книг, написанных Булгаковым или Достоевским.
//
//SELECT title, price
//FROM book
//WHERE author IN ('Булгаков М.А.', 'Достоевский Ф.М.');

/*
ПОДСЧЁТ кол-ва суммы, повторений, всех элементов
Выборка данных, групповые функции SUM и COUNT
При группировке над элементами столбца, входящими в группу можно выполнить различные действия,
например, просуммировать их или найти количество элементов в группе.

Из таблицы с результатами запроса видно, что функцию COUNT() можно применять к любому столбцу,
в том числе можно использовать и *, если таблица не содержит пустых значений.
Если же в столбцах есть значения Null, (для группы по автору Есенин в нашем примере), то
COUNT(*) —  подсчитывает  все записи, относящиеся к группе, в том числе и со значением NULL;
COUNT(имя_столбца) — возвращает количество записей конкретного столбца (только NOT NULL), относящихся к группе.
ВАЖНО.
Если столбец указан в SELECT  БЕЗ применения групповой функции,
то он обязательно должен быть указан и вGROUP BY.Иначе получим ошибку.
*/

//SELECT author, COUNT(author), COUNT(amount), COUNT(*)
//FROM book
//GROUP BY author;
