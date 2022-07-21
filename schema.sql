CREATE TABLE language (
    language_id INTEGER PRIMARY KEY
    ,name TEXT NOT NULL
    ,last_update TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE film (
    film_id INTEGER PRIMARY KEY
    ,title VARCHAR(255) NOT NULL
    ,description VARCHAR(5000)
    ,release_year INT
    ,language_id INT NOT NULL
    ,original_language_id INT
    ,rental_duration INT NOT NULL DEFAULT 3
    ,rental_rate NUMERIC(5,2) NOT NULL DEFAULT 4.99
    ,length INT
    ,replacement_cost NUMERIC(5,2) NOT NULL DEFAULT 19.99
    ,rating TEXT DEFAULT 'G'
    ,special_features VARCHAR(255)
    ,last_update TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

    ,CONSTRAINT film_language_id_fkey FOREIGN KEY (language_id) REFERENCES language (language_id) ON UPDATE CASCADE ON DELETE RESTRICT
    ,CONSTRAINT film_original_language_id_fkey FOREIGN KEY (original_language_id) REFERENCES language (language_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE INDEX film_title_idx ON film (title);

CREATE INDEX film_language_id_idx ON film (language_id);

CREATE INDEX film_original_language_id_idx ON film (original_language_id);
