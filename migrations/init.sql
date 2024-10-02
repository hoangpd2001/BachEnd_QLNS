-- +migrate Up
CREATE TABLE people (
    id INT AUTO_INCREMENT PRIMARY KEY
);

-- +migrate StatementBegin
DELIMITER $$

CREATE PROCEDURE do_something()
BEGIN
    DECLARE create_query TEXT;
    -- Do something here
END $$

DELIMITER ;
-- +migrate StatementEnd

-- +migrate Down
DROP PROCEDURE IF EXISTS do_something;
DROP TABLE IF EXISTS people;