package pgx_storage

import (
	"context"
	"fmt"
	"os"
	"time"

	tm "github.com/buger/goterm"

	// "github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewStorage( /*username, password, host, dbname string*/ ) (*pgxpool.Pool, error) {
	// Preparing the statement
	// postgresStatment := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, dbname)
	postgresStatment := os.Getenv("POSTGRES_REMOTE_URI")
	conn, err := pgxpool.Connect(context.Background(), postgresStatment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	tm.Println(tm.Color(" pgx : DB Connected Succesfuly ... \n", tm.GREEN))

	// Continuously re-connects to the database
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ticker.C:
				{
					if conn != nil {
						err = conn.Ping(context.Background())
					}
					if err != nil || conn == nil {
						tm.Println(tm.Color("DB Connection: Trying to reconnect ...", tm.RED))
						i := 0
						for ; i <= 100; i++ {
							conn, err = pgxpool.Connect(context.Background(), postgresStatment)
							if err == nil {
								break
							}
						}
						if i == 100 {
							tm.Println(tm.Color("Database connection Failure ...", tm.RED))
							os.Exit(1)
						}
					}
				}
			}
		}
	}()
	// err = executeTableCreation(conn)
	// if err != nil {
	// 	return nil, err
	// }
	return conn, err
}

func executeTableCreation(conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), `CREATE TABLE address (
		address_id SERIAL PRIMARY KEY,
		kebele VARCHAR(100),
		woreda VARCHAR(100),
		city VARCHAR(100),
		region VARCHAR(100),
		unique_name VARCHAR(100),
		zone VARCHAR(20),
		latitude VARCHAR(20),
		longitude VARCHAR(20)
	);
	
	
	CREATE TABLE users (
		user_id SERIAL PRIMARY KEY,
		firstname VARCHAR(70) NOT NULL,
		lastname VARCHAR(70) NOT NULL,
		phone VARCHAR(13) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at INTEGER DEFAULT ROUND(EXTRACT(EPOCH FROM NOW())),
		role VARCHAR(20) NOT NULL CHECK (role IN ('customer', 'shop_owner', 'admin'))
	);
	
	
	CREATE TABLE shop (
		shop_id SERIAL PRIMARY KEY,
		shop_name VARCHAR(100) NOT NULL,
		owner_id INTEGER NOT NULL REFERENCES users(user_id),
		address_id INTEGER NOT NULL REFERENCES address(address_id),
		created_at INTEGER DEFAULT ROUND(EXTRACT(EPOCH FROM NOW()))
	);
	
	CREATE TABLE product (
		product_id SERIAL PRIMARY KEY,
		product_name VARCHAR(200) NOT NULL,
		description TEXT,
		unit_id INTEGER NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		shop_id INTEGER NOT NULL REFERENCES shop(shop_id),
		created_at INTEGER DEFAULT ROUND(EXTRACT(EPOCH FROM NOW()))
	);
	
	
	CREATE TABLE product_image (
		image_id SERIAL PRIMARY KEY,
		product_id INTEGER REFERENCES product(product_id),
		image_url VARCHAR(200) NOT NULL,
		created_at INTEGER DEFAULT ROUND(EXTRACT(EPOCH FROM NOW()))
	);
	
	
	CREATE TABLE orders (
		order_id SERIAL PRIMARY KEY,
		customer_id INTEGER NOT NULL REFERENCES users(user_id),
		shop_id INTEGER NOT NULL REFERENCES shop(shop_id),
		total_amount DECIMAL(10, 2) NOT NULL,
		status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'confirmed', 'shipped', 'delivered', 'canceled')),
		created_at INTEGER DEFAULT ROUND(EXTRACT(EPOCH FROM NOW()))
	);
	
	
	CREATE TABLE order_item (
		order_item_id SERIAL PRIMARY KEY,
		order_id INTEGER REFERENCES orders(order_id),
		product_id INTEGER REFERENCES product(product_id),
		quantity INTEGER NOT NULL,
		price DECIMAL(10, 2) NOT NULL
	);
	
	
	CREATE TABLE delivery (
		delivery_id SERIAL PRIMARY KEY,
		order_id INTEGER REFERENCES orders(order_id),
		delivery_address_id INTEGER REFERENCES address(address_id),
		delivery_status VARCHAR(20) NOT NULL CHECK (delivery_status IN ('pending', 'in_transit', 'delivered')),
		delivery_date INTEGER,
		created_at INTEGER DEFAULT ROUND(EXTRACT(EPOCH FROM NOW()))
	);`)
	return err
}
