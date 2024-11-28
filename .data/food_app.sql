-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: mysql:3306
-- Generation Time: Nov 28, 2024 at 06:00 AM
-- Server version: 8.0.40
-- PHP Version: 8.2.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `food_app`
--

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` int NOT NULL,
  `name` text NOT NULL,
  `mobile` varchar(10) NOT NULL,
  `email` varchar(50) NOT NULL,
  `status` set('active','inactive') NOT NULL DEFAULT 'active',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `name`, `mobile`, `email`, `status`, `created_at`, `updated_at`) VALUES
(1, 'John Doe', '1234567890', 'john.doe@example.com', 'active', '2024-11-20 14:30:00', '2024-11-20 14:30:00'),
(2, 'Jane Smith', '9876543210', 'jane.smith@example.com', 'inactive', '2024-11-19 10:20:00', '2024-11-19 10:20:00'),
(3, 'Alice Brown', '5551234567', 'alice.brown@example.com', 'active', '2024-11-18 08:15:00', '2024-11-18 08:15:00'),
(4, 'Bob Johnson', '4449876543', 'bob.johnson@example.com', 'inactive', '2024-11-17 16:45:00', '2024-11-17 16:45:00'),
(5, 'Charlie Lee', '6665554321', 'charlie.lee@example.com', 'active', '2024-11-16 12:00:00', '2024-11-16 12:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `notifications`
--

CREATE TABLE `notifications` (
  `id` int NOT NULL,
  `order_id` int NOT NULL,
  `customer_id` int NOT NULL,
  `notification_template_id` int NOT NULL DEFAULT '1000',
  `notification_channel` varchar(5) NOT NULL,
  `notification_status` set('success','failed','unknown','prepared') DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` int NOT NULL,
  `customer_id` int NOT NULL,
  `order_status` set('placed','preparing','ready','pickup','cancel','undelivered') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `customer_id`, `order_status`, `created_at`, `updated_at`) VALUES
(1, 1, 'placed', '2024-11-01 10:00:00', '2024-11-01 10:00:00'),
(2, 2, 'preparing', '2024-11-01 10:30:00', '2024-11-26 18:15:24'),
(3, 3, 'placed', '2024-11-01 11:00:00', '2024-11-01 11:00:00'),
(4, 4, 'placed', '2024-11-01 11:30:00', '2024-11-01 11:30:00'),
(5, 5, 'placed', '2024-11-01 12:00:00', '2024-11-01 12:00:00'),
(6, 1, 'placed', '2024-11-01 12:30:00', '2024-11-01 12:30:00'),
(7, 2, 'placed', '2024-11-01 13:00:00', '2024-11-01 13:00:00'),
(8, 3, 'placed', '2024-11-01 13:30:00', '2024-11-01 13:30:00'),
(9, 4, 'placed', '2024-11-01 14:00:00', '2024-11-01 14:00:00'),
(10, 5, 'placed', '2024-11-01 14:30:00', '2024-11-01 14:30:00'),
(11, 1, 'placed', '2024-11-01 15:00:00', '2024-11-01 15:00:00'),
(12, 2, 'placed', '2024-11-01 15:30:00', '2024-11-01 15:30:00'),
(13, 3, 'placed', '2024-11-01 16:00:00', '2024-11-01 16:00:00'),
(14, 4, 'placed', '2024-11-01 16:30:00', '2024-11-01 16:30:00'),
(15, 5, 'placed', '2024-11-01 17:00:00', '2024-11-01 17:00:00'),
(16, 1, 'placed', '2024-11-01 17:30:00', '2024-11-01 17:30:00'),
(17, 2, 'placed', '2024-11-01 18:00:00', '2024-11-01 18:00:00'),
(18, 3, 'placed', '2024-11-01 18:30:00', '2024-11-01 18:30:00'),
(19, 4, 'placed', '2024-11-01 19:00:00', '2024-11-01 19:00:00'),
(20, 5, 'placed', '2024-11-01 19:30:00', '2024-11-01 19:30:00'),
(21, 1, 'placed', '2024-11-01 20:00:00', '2024-11-01 20:00:00'),
(22, 2, 'placed', '2024-11-01 20:30:00', '2024-11-01 20:30:00'),
(23, 3, 'placed', '2024-11-01 21:00:00', '2024-11-01 21:00:00'),
(24, 4, 'placed', '2024-11-01 21:30:00', '2024-11-01 21:30:00'),
(25, 5, 'placed', '2024-11-01 22:00:00', '2024-11-01 22:00:00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD UNIQUE KEY `id` (`id`);

--
-- Indexes for table `notifications`
--
ALTER TABLE `notifications`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order_id` (`order_id`),
  ADD KEY `notification_status` (`notification_status`),
  ADD KEY `notification_channel` (`notification_channel`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order_status` (`order_status`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `notifications`
--
ALTER TABLE `notifications`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
