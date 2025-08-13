-- ساخت دیتابیس
CREATE DATABASE IF NOT EXISTS chita
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

-- استفاده از دیتابیس
USE chita;

-- جدول تعرفه‌های VPN
CREATE TABLE IF NOT EXISTS plans (
    id INT AUTO_INCREMENT PRIMARY KEY,
    plan_name VARCHAR(100) NOT NULL,         -- نام تعرفه
    duration_days INT NOT NULL,               -- مدت زمان به روز
    price DECIMAL(10,2) NOT NULL,             -- قیمت
    data_limit BIGINT DEFAULT NULL,           -- حجم به بایت (NULL = نامحدود)
    description TEXT DEFAULT NULL,             -- توضیحات اضافی
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- جدول تراکنش‌ها
CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,           -- نام کاربر VPN
    amount DECIMAL(10,2) NOT NULL,             -- مبلغ
    transaction_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    payment_method VARCHAR(50),                -- روش پرداخت (آنلاین، کارت به کارت...)
    status VARCHAR(20) DEFAULT 'pending'       -- وضعیت (pending, completed, failed)
);
