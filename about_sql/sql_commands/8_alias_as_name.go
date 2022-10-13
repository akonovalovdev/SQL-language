package main

/*
Чтобы не писать название таблицы каждый раз, удобно использовать алиасы.

Алиас, это псевдоним, который мы присваивали столбцам после ключевого слова AS(шаг).
Алиасы так же можно использовать и для таблиц. Это становится актуальным, при увеличении числа используемых таблиц,
их иногда может быть и 5 и 10 и более. Псевдонимы помогают сделать запрос чище и читабельнее.

Для присваивания псевдонима существует 2 варианта:

с использованием ключевого слова AS
FROM fine AS f, traffic_violation AS tv
а так же и без него
FROM fine f, traffic_violation tv
После присвоения таблице алиаса, он используется во всех разделах запроса, в котором алиас задан. Например:

WHERE f.violation = tv.violation
*/

// ПРИМЕР
//
//SELECT DISTINCT
//	id,
//	firstname,
//	u.lastname,
//	p.sku
//	FROM user as u join purchase as p on u.id = p.user_id
//	WHERE GETMONTH(date) = 2