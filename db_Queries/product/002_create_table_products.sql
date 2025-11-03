CREATE TABLE
    products (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        price REAL NOT NULL,
        ImgUrl TEXT,
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );
    -- type Product struct {
    -- 	ID          int    `json:"id"`
    -- 	Title       string `json:"title"`
    -- 	Description string `json:"description"`
    -- 	Price       string `json:"price"`
    -- 	ImgUrl      string `json:"imageUrl"`
    -- }