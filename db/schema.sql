CREATE TABLE users
(
  id              INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
  account         VARCHAR(255) NOT NULL,
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY     (id)                                  # Make the id the primary key
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;

CREATE TABLE releases
(
  id                INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
  source            VARCHAR(32) NOT NULL,
  source_release_id BIGINT NOT NULL,
  title             VARCHAR(200) NOT NULL,                
  country           VARCHAR(5) NOT NULL,                
  genres            JSON NOT NULL,
  release_year      SMALLINT UNSIGNED NOT NULL,
  tracklist         JSON NOT NULL,
  images            VARCHAR(512),
  barcode           VARCHAR(20) NOT NULL,
  fetched_at        DATETIME NOT NULL,
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY     (id),
  UNIQUE KEY uniq_source_release (source, source_release_id)
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;

CREATE TABLE releases_users
(
  id              INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
  user_id         INT unsigned NOT NULL,
  release_id      INT unsigned NOT NULL,
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY     (id),
  UNIQUE KEY uniq_user_release (user_id, release_id)
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;
