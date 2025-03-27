# Problem 1: create account and create password was separated

1. If the create account success and create account password failed, ...what happens
   => Tạo ra người dùng ko có password
   => Giải pháp: Bắt buộc create account và create password cần thành công cùng nhau
   => Tương tác transactions
   => Khi 1 hoạt động thất bại thì tất cả cần được rollback lại
   => 1 là tất cả cùng thành công, 2 là tất cả cùng thất bại
