# Kiến trúc


* `main.go`: entrypoint của chương trình

* `cmd`: nơi chứa các command để chạy chương trình theo chuẩn cloud native

* `configfile`: chứa các file config theo môi trường

* `Container`: tập hợp các transporter
    * `rest`: transporter http restful api
  
* `docs`: chứa tài liệu của chương trình
      
* `internal`: chứa những đặc trưng của project
    * `appconf`: khai báo và khởi tạo các config 
    * `appconst`: chứa các hằng số
    * `appctx`: khai báo và khởi tạo application context 
    * `apperr`: chứa error 
    * `apputils`: chứa các hàm utils 
  
* `services`: là tập hợp các domain
    * `company`: là domain company được aggregate và đóng gói
        * `models`: chứa các domain models mà repo cung cấp
        * `repository`: là nới cung cấp dữ liệu để đỏ vào models có thể là từ database, từ kafka...
        * `service`: là nơi triển khai các logic
        * `vo`: là các value obj sẽ tương tác với service