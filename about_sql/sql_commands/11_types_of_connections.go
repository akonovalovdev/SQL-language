package main

/*  ОДИН КО МНОГИМ
Связь «один ко многим» имеет место, когда одной записи главной таблицы
соответствует несколько записей связанной таблицы,
а каждой записи связанной таблицы соответствует только одна запись главной таблицы.
*/

/*
 Связь «многие ко многим» имеет место когда каждой записи одной таблицы соответствует несколько записей во второй,
 и наоборот, каждой записи второй таблицы соответствует несколько записей в первой.
*/

/*
Создание таблицы с внешними ключами
При создании зависимой таблицы (таблицы, которая содержит внешние ключи) необходимо учитывать, что :

каждый внешний ключ должен иметь такой же тип данных, как связанное поле главной таблицы (в наших примерах это INT);
необходимо указать главную для нее таблицу и столбец, по которому осуществляется связь:

FOREIGN KEY (связанное_поле_зависимой_таблицы)
REFERENCES главная_таблица (связанное_поле_главной_таблицы)

По умолчанию любой столбец, кроме ключевого, может содержать значение NULL.
При создании таблицы это можно переопределить,  используя  ограничение NOT NULL для этого столбца:
*/

//CREATE TABLE таблица (
//    столбец_1 INT NOT NULL,
//    столбец_2 VARCHAR(10)
//);

// В созданной таблице в столбец_1 не может содержать пустое значение, а столбец_2 - может.
//
//Для внешних ключей рекомендуется устанавливать ограничение NOT NULL (если это совместимо с другими опциями

/*
Действия при удалении записи главной таблицы
С помощью выражения ON DELETE можно установить действия, которые выполняются для записей
подчиненной таблицы при удалении связанной строки из главной таблицы. При удалении можно установить следующие опции:

CASCADE: автоматически удаляет строки из зависимой таблицы при удалении  связанных строк в главной таблице.
SET NULL: при удалении  связанной строки из главной таблицы устанавливает для столбца внешнего ключа значение NULL.
(В этом случае столбец внешнего ключа должен поддерживать установку NULL).
SET DEFAULT похоже на SET NULL за тем исключением, что значение  внешнего ключа устанавливается не в NULL,
а в значение по умолчанию для данного столбца.
RESTRICT: отклоняет удаление строк в главной таблице при наличии связанных строк в зависимой таблице.

Важно! Если для столбца установлена опция SET NULL, то при его описании нельзя задать ограничение на пустое значение.

Будем считать, что при удалении автора из таблицы author, необходимо удалить все записи о книгах из таблицы book,
написанные этим автором. Данное действие необходимо прописать при создании таблицы.
*/

//CREATE TABLE book (
//    book_id INT PRIMARY KEY AUTO_INCREMENT,
//    title VARCHAR(50),
//    author_id INT NOT NULL,
//    price DECIMAL(8,2),
//    amount INT,
//    FOREIGN KEY (author_id)  REFERENCES author (author_id) ON DELETE CASCADE
//);
